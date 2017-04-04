package codec

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
)

//ErrSOF invalid start of a ZNP message, because the first byte of the message should be 0xfe
var ErrSOF = errors.New("invalid SOF")

// ErrHDR the znp message should be at least 3 bytes
var ErrHDR = errors.New("invalid HDR")

// ErrPayload can not get the content of lenght in HDR
var ErrPayload = errors.New("unable to read enough payload")

// ErrFcs the checksum is wrong of ZNP message
var ErrFcs = errors.New("invalid FCS")

// ErrIO the underlayer connection read write error
var ErrIO = errors.New("underlayer IO error")

//
// RPC (Remote Procedure Call) definitions
//
const (
	// SOF (Start of Frame) indicator byte byte
	MT_RPC_SOF = (0xFE)

	// The 3 MSB's of the 1st command field byte (Cmd0) are for command type
	MT_RPC_CMD_TYPE_MASK = (0xE0)

	// The 5 LSB's of the 1st command field byte (Cmd0) are for the subsystem
	MT_RPC_SUBSYSTEM_MASK = (0x1F)

	// maximum length of RPC frame
	// (1 byte length + 2 bytes command + 0-250 bytes data)
	RPC_MAX_LEN = (256)

	// RPC Frame field lengths
	RPC_UART_SOF_LEN = (1)
	RPC_UART_FCS_LEN = (1)

	RPC_UART_FRAME_START_IDX = (1)

	RPC_LEN_FIELD_LEN  = (1)
	RPC_CMD0_FIELD_LEN = (1)
	RPC_CMD1_FIELD_LEN = (1)

	RPC_HDR_LEN = (RPC_LEN_FIELD_LEN + RPC_CMD0_FIELD_LEN + RPC_CMD1_FIELD_LEN)

	RPC_UART_HDR_LEN = (RPC_UART_SOF_LEN + RPC_HDR_LEN)
)

/***********************************************************************************
 * TYPEDEFS
 */

// Cmd0 Command Type
const (
	MT_RPC_CMD_POLL = 0x00 // POLL command
	MT_RPC_CMD_SREQ = 0x20 // SREQ (Synchronous Request) command
	MT_RPC_CMD_AREQ = 0x40 // AREQ (Acynchronous Request) command
	MT_RPC_CMD_SRSP = 0x60 // SRSP (Synchronous Response)
	MT_RPC_CMD_RES4 = 0x80 // Reserved
	MT_RPC_CMD_RES5 = 0xA0 // Reserved
	MT_RPC_CMD_RES6 = 0xC0 // Reserved
	MT_RPC_CMD_RES7 = 0xE0 // Reserved
)

// Cmd0 Command Subsystem
const (
	MT_RPC_SYS_RES0     = 0 // Reserved.
	MT_RPC_SYS_SYS      = 1 // SYS interface
	MT_RPC_SYS_MAC      = 2
	MT_RPC_SYS_NWK      = 3
	MT_RPC_SYS_AF       = 4 // AF interface
	MT_RPC_SYS_ZDO      = 5 // ZDO interface
	MT_RPC_SYS_SAPI     = 6 // Simple API interface
	MT_RPC_SYS_UTIL     = 7 // UTIL interface
	MT_RPC_SYS_DBG      = 8
	MT_RPC_SYS_APP      = 9
	MT_RPC_SYS_OTA      = 10
	MT_RPC_SYS_ZNP      = 11
	MT_RPC_SYS_SPARE_12 = 12
	MT_RPC_SYS_SBL      = 13 // 13 to be compatible with existing RemoTI - AKA MT_RPC_SYS_UBL
	MT_RPC_SYS_GP       = 0x15
	MT_RPC_SYS_RFT      = 0x1F // was 0x15, but this is now taken by TI GP

)

// Error codes in Attribute byte of SRSP packet
const (
	MT_RPC_SUCCESS        = 0 // success
	MT_RPC_ERR_SUBSYSTEM  = 1 // invalid subsystem
	MT_RPC_ERR_COMMAND_ID = 2 // invalid command ID
	MT_RPC_ERR_PARAMETER  = 3 // invalid parameter
	MT_RPC_ERR_LENGTH     = 4 // invalid length
)

type codecMap struct {
	tp       uint8
	reqType  reflect.Type
	respType reflect.Type
}

type marshaller interface {
	marshall() []byte
}

type unmarshaller interface {
	unmarshall([]byte) error
}

type decoder struct {
	r  io.Reader
	br *bufio.Reader
}

type cmd struct {
	cmd0 uint8
	cmd1 uint8
}

type header struct {
	length uint8
	cmd
}

func (c *cmd) getCmdType() uint8 {
	return c.cmd0 & MT_RPC_CMD_TYPE_MASK
}

func (c *cmd) getCmdID() uint16 {
	return uint16(c.cmd0&MT_RPC_SUBSYSTEM_MASK)<<8 | uint16(c.cmd1)
}

//func (c *cmd) isReqTypeSync() bool {
//	t, ok := c.getReqType()
//	return ok && (t&MT_RPC_CMD_SREQ) > 0
//}

func (c *cmd) isSyncResp() bool {
	return c.cmd0&MT_RPC_CMD_TYPE_MASK == MT_RPC_CMD_SRSP
}

func (c *cmd) isSyncReq() bool {
	return c.cmd0&MT_RPC_CMD_TYPE_MASK == MT_RPC_CMD_SREQ
}

type iCodec interface {
	// WriteRequest must be safe for concurrent use by multiple goroutines.
	encode(cmd cmd, v interface{}) error
	decode() (*header, interface{}, error)
	close() error
}

type clientCodec struct {
	rwc io.ReadWriteCloser
	dec *decoder
	enc *encoder
}

func (codec *clientCodec) encode(cmd cmd, v interface{}) error {
	return codec.enc.encode(cmd, v)
}

func (codec *clientCodec) decode() (*header, interface{}, error) {
	return codec.dec.decode()
}

func (codec *clientCodec) close() error {
	return codec.rwc.Close()
}

// NewDecoder create a new decoder to decode znp message from io.Reader
func NewDecoder(r io.Reader) *decoder {
	return &decoder{r: r, br: bufio.NewReader(r)}
}

type encoder struct {
	w io.Writer
}

func newEncoder(w io.Writer) *encoder {
	return &encoder{w: w}
}

func (enc *encoder) encode(cmd cmd, v interface{}) error {
	tp, ok := cmd.getReqType()
	if !ok {
		panic(fmt.Sprintf("error request cmd(%v)", cmd))
	}
	cmd.cmd0 |= tp
	return enc.encodeRaw(cmd, v)
}

func (enc *encoder) encodeRaw(cmd cmd, v interface{}) error {
	var err error
	var payloadBuf []byte
	if v != nil {
		marshall, ok := v.(marshaller)
		if ok {
			payloadBuf = marshall.marshall()
		} else {
			buf := &bytes.Buffer{}
			err = binary.Write(buf, binary.LittleEndian, v)
			if err != nil {
				return err
			}
			payloadBuf = buf.Bytes()
		}
	}
	hdr := &header{uint8(len(payloadBuf)), cmd}
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, uint8(MT_RPC_SOF))
	binary.Write(buf, binary.LittleEndian, hdr)
	buf.Write(payloadBuf)
	fcs := calcFcs(buf.Bytes()[1:])
	binary.Write(buf, binary.LittleEndian, fcs)
	_, err = enc.w.Write(buf.Bytes())
	if err != nil {
		return ErrIO
	}
	return nil
}

func (dec *decoder) decodeHeader() (*header, []byte, error) {
	buf := make([]byte, RPC_HDR_LEN+RPC_MAX_LEN+RPC_UART_FCS_LEN)
	sof, err := dec.br.ReadByte()
	if err != nil {
		return nil, nil, ErrIO
	}

	if sof != MT_RPC_SOF { //minimum package size
		return nil, nil, ErrSOF
	}

	for offset := 0; offset < RPC_HDR_LEN; {
		n, err := dec.br.Read(buf[:RPC_HDR_LEN])
		if err != nil {
			return nil, nil, ErrHDR
		}
		offset += n
	}
	hdr := &header{}
	hdr.length = buf[0]
	hdr.cmd0 = buf[1]
	hdr.cmd1 = buf[2]

	for offset := uint8(RPC_HDR_LEN); offset < RPC_HDR_LEN+hdr.length+RPC_UART_FCS_LEN; {
		n, err := dec.br.Read(buf[offset : RPC_HDR_LEN+hdr.length+RPC_UART_FCS_LEN])
		if err != nil {
			return hdr, nil, ErrPayload
		}
		offset += uint8(n)
	}
	fcs := buf[hdr.length+RPC_HDR_LEN]
	if calcFcs(buf[:hdr.length+RPC_HDR_LEN]) != fcs {
		return nil, nil, ErrFcs
	}
	return hdr, buf[RPC_HDR_LEN : hdr.length+RPC_HDR_LEN], nil
}

func (dec *decoder) decode() (*header, interface{}, error) {
	hdr, payload, err := dec.decodeHeader()
	if err != nil {
		return nil, nil, err
	}

	rsp, ok := hdr.newObject()
	if !ok {
		return hdr, payload, nil
	}
	glog.V(5).Info("got rsp", rsp)

	if rsp == nil {
		return hdr, nil, nil
	}

	obj, ok := rsp.(unmarshaller)
	if ok {
		err = obj.unmarshall(payload)
	} else {
		err = binary.Read(bytes.NewBuffer(payload), binary.LittleEndian, rsp)
	}
	if err != nil {
		err = errors.New("unable to decode payload")
	}

	return hdr, rsp, err
}

// Dump to parase the input to human readable text
func (dec *decoder) Dump(out io.Writer) error {
	buf := &bytes.Buffer{}
	hdr, resp, err := dec.decode()
	if err == nil {
		if resp == nil {
			buf.WriteString(fmt.Sprintf("'%s' : %02X-%02X \r\n", hdr.getDesc(), hdr.cmd0, hdr.cmd1))
		} else {
			buf.WriteString(fmt.Sprintf("'%s' : %02X-%02X | %v\r\n", hdr.getDesc(), hdr.cmd0, hdr.cmd1, reflect.ValueOf(resp).Elem()))
		}

	} else if err != nil && err != ErrIO {
		if hdr == nil {
			buf.WriteString(fmt.Sprintf("ERROR:  %v\r\n", err))
		} else {
			buf.WriteString(fmt.Sprintf("ERROR: %02X-%02X | %v\r\n", hdr.cmd0, hdr.cmd1, err))
		}

	}
	buf.WriteTo(out)
	return err
}

func calcFcs(msg []byte) uint8 {
	var result uint8
	// calculate FCS by XORing all bytes
	for _, v := range msg {
		result ^= v
	}
	return result
}

func cloneValue(source interface{}, destin interface{}) {
	x := reflect.ValueOf(source)
	if x.Kind() == reflect.Ptr {
		starX := x.Elem()
		y := reflect.New(starX.Type())
		starY := y.Elem()
		starY.Set(starX)
		reflect.ValueOf(destin).Elem().Set(y.Elem())
	} else {
		destin = x.Interface()
	}
}
