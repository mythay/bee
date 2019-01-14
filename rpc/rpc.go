package rpc

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"bytes"
	"container/list"
	"encoding/gob"

	"github.com/golang/glog"

	"github.com/mythay/znp/codec"
)

type rpcCall struct {
	cmd     codec.Cmd     // The name of the service and method to call.
	Request interface{}   // The argument to the function (*struct).
	Reply   interface{}   // The reply from the function (*struct).
	Error   error         // After completion, the error status.
	Done    chan *rpcCall // Strobes when call is complete.
}

//ErrShutdown is returned if the connection is closed
var ErrShutdown = errors.New("connection is shutdown")

//ErrTimeout is returned if timeout
var ErrTimeout = errors.New("timeout")

//ErrUnexceptedResponse is returned if the response is not the one we want
var ErrUnexceptedResponse = errors.New("response doesn't match")

//NewClient create a new client to control ZNP device
func NewClient(conn io.ReadWriteCloser) *Client {
	client := &Client{
		enc:     codec.NewEncoder(conn),
		dec:     codec.NewDecoder(conn),
		pending: make(chan *rpcCall),
		timeout: time.Second * 2, // underlay timeout 2000ms
		indList: list.New(),
		conn:    conn,
	}

	go client.input()
	return client

}

func (call *rpcCall) done() {
	select {
	case call.Done <- call:
		// ok
	default:
		//		if debugLog {
		//			log.Println("rpc: discarding Call reply due to insufficient Done chan capacity")
		//		}
	}
}

type outstandingReq struct {
	req  *rpcCall
	done chan *rpcCall
	err  error
}

type notifyHanlder func(codec.Cmd, interface{})

// Client represent the  connection to a ZNP device
type Client struct {
	pending chan *rpcCall
	enc     *codec.Encoder
	dec     *codec.Decoder
	timeout time.Duration

	hanlder notifyHanlder

	mutex       sync.Mutex // protects following
	outstanding *outstandingReq
	closing     bool // user has called Close
	shutdown    bool // server has told us to stop
	indList     *list.List
	conn        io.ReadWriteCloser
}

// HandleFunc handle the notify coming from ZNP device
func (client *Client) HandleFunc(handler notifyHanlder) {
	client.hanlder = handler
}

type indication struct {
	cmd     codec.Cmd        // The name of the service and method to call.
	ind     interface{}      // The argument to the function (*struct).
	resp    interface{}      // The argument to the function (*struct).
	Done    chan *indication // Strobes when call is complete.
	Error   error
	handled bool
}

func (ind *indication) done() {
	select {
	case ind.Done <- ind:
		// ok
	default:

	}
}

// WaitAsync to wait the indication from ZNP
func (client *Client) WaitAsync(req interface{}, timeout time.Duration) (interface{}, error) {
	var ind *indication
	switch v := req.(type) {
	case codec.Cmd:

		ind = &indication{cmd: v, Done: make(chan *indication, 1)}

	default:
		c, ok := codec.GetCmdFromRespone(req)
		if ok {
			ind = &indication{cmd: c, ind: v, Done: make(chan *indication, 1)}
		} else {
			panic(fmt.Sprintf("invalid response, %v", v))
		}
	}
	client.mutex.Lock()
	e := client.indList.PushBack(ind)
	client.mutex.Unlock()
	select {
	case <-ind.Done:
		return ind.resp, ind.Error
	case <-time.After(timeout):
		client.mutex.Lock()
		client.indList.Remove(e)
		client.mutex.Unlock()
		return nil, ErrTimeout
	}

}

func (client *Client) send(call *rpcCall) {

	// Register this call.
	client.mutex.Lock()
	if client.shutdown || client.closing {
		call.Error = ErrShutdown
		client.mutex.Unlock()
		call.done()
		return
	}
	client.mutex.Unlock()
	if call.cmd.IsSyncReq() {
		client.mutex.Lock()
		client.outstanding = &outstandingReq{call, make(chan *rpcCall, 1), nil}
		client.mutex.Unlock()
	}

	glog.Infof("request %#v -%#v", call.cmd, call.Request)
	err := client.enc.Encode(call.cmd, call.Request)
	if err != nil {
		call.Error = err
		call.done()
		return
	}
	if call.cmd.IsSyncReq() {

		select {
		case <-client.outstanding.done:
			glog.Infof("reply %#v -%#v", call.cmd, call.Reply)
			call.done()

		case <-time.After(client.timeout):
			glog.Infof("reply timeout %#v ", call.cmd)
			call.Error = ErrTimeout
			call.done()

		}
		client.mutex.Lock()
		client.outstanding = nil
		client.mutex.Unlock()
	} else { // this is only for two reset command
		call.done()

	}
}

func (client *Client) dispatchSyncResponse(cmd codec.Cmd, resp interface{}) {
	client.mutex.Lock()
	outstanding := client.outstanding
	client.mutex.Unlock()
	if outstanding == nil { //no body wait for this synchonize response, discard
		// glog.V(4).Info("discard one synchonize response", hdr)
		return
	}
	if !outstanding.req.cmd.IsRespCmdMatch(cmd) { // don't match, error
		outstanding.req.Error = ErrUnexceptedResponse
		outstanding.done <- outstanding.req
	} else {
		outstanding.req.Reply = resp
		outstanding.done <- outstanding.req
	}
}

func (client *Client) dispatchAsyncResponse(cmd codec.Cmd, resp interface{}) {
	client.mutex.Lock()
	for e := client.indList.Front(); e != nil; e = e.Next() {
		ind := e.Value.(*indication)
		if cmd.GetCmdID() == ind.cmd.GetCmdID() { // item was found
			if resp != nil && ind.ind != nil {
				deepCopy(ind.ind, resp)
			}
			ind.resp = resp
			ind.handled = true
			client.indList.Remove(e)
			ind.done()
		}
	}
	client.mutex.Unlock()

	if client.hanlder != nil {
		client.hanlder(cmd, resp)
	} else {
		// glog.Errorf(" aysnchonize notify not handled:'%s' %v-%v", hdr.getDesc(), hdr.cmd, (resp))
	}

}

func (client *Client) clearPendingRequest() {
	var err error
	client.mutex.Lock()
	defer client.mutex.Unlock()
	client.shutdown = true
	closing := client.closing

	if closing {
		err = ErrShutdown
	} else {
		err = codec.ErrIO
	}

	if client.outstanding != nil { // synchronize
		client.outstanding.req.Error = err
		client.outstanding.req.done()
	}
	for e := client.indList.Front(); e != nil; e = e.Next() {
		ind := e.Value.(*indication)
		ind.Error = err
		ind.done()
	}
	client.indList.Init()

}

//input the reading function
func (client *Client) input() {
	var err error
	for err != codec.ErrIO { //never return until underlayer IO error happened
		var cmd codec.Cmd
		var resp interface{}
		cmd, resp, err = client.dec.Decode()
		if err != nil {
			continue
		}
		if cmd.IsSyncResp() {
			client.dispatchSyncResponse(cmd, resp)
		} else { // asynchonize indication, try to find if any function wait for it
			client.dispatchAsyncResponse(cmd, resp)
		}
	}
	// error happened, so terminate pending calls.

	client.clearPendingRequest()
	//	close(client.pending)
}

// Close the connection to chip
func (client *Client) Close() error {
	client.mutex.Lock()
	if client.closing {
		client.mutex.Unlock()
		return ErrShutdown
	}
	client.closing = true
	client.mutex.Unlock()
	close(client.pending)
	client.conn.Close()
	return nil
}

//gorequest for ansynchonized call
func (client *Client) gorequest(req interface{}) *rpcCall {
	var call *rpcCall
	switch v := req.(type) {
	case codec.Cmd:
		ok := v.AddReqType() // the request has no send type, add the type first
		if ok {
			call = &rpcCall{cmd: v, Request: nil, Done: make(chan *rpcCall, 1)}
		} else {
			panic(fmt.Sprintf("invalid cmd, %v", v))
		}

	default:
		c, ok := codec.GetCmdFromRequest(req)
		if ok {
			call = &rpcCall{cmd: c, Request: v, Done: make(chan *rpcCall, 1)}
		} else {
			panic(fmt.Sprintf("invalid request, %v", v))
		}
	}
	client.send(call)
	return call
}

// Call invokes the named function, waits for it to complete, and returns its error status.
func (client *Client) Call(req interface{}) (interface{}, error) {
	return client.CallWithTimeout(req, time.Millisecond*50)
}

// CallWithTimeout invokes the named function, waits for it to complete, and returns its error status.
func (client *Client) CallWithTimeout(req interface{}, timeout time.Duration) (interface{}, error) {
	select {
	case call := <-client.gorequest(req).Done:
		return call.Reply, call.Error
	case <-time.After(timeout):
		return nil, ErrTimeout
	}
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
