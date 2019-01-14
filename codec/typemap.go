package codec

import (
	"reflect"
)

type mapItem struct {
	cmd0     uint8
	cmd1     uint8
	respCmd  uint8
	reqT     uint8
	respT    uint8
	reqType  reflect.Type
	respType reflect.Type
	desc     string
}

var _reqCmdMap = make(map[uint16]*mapItem)  // request (cmd0,cmd1) to get item
var _respCmdMap = make(map[uint16]*mapItem) // response (cmd0,cmd1) to get item
var _reqTypeMap = make(map[reflect.Type]*mapItem)
var _respTypeMap = make(map[reflect.Type]*mapItem)

func addSubCommandMap(tbl []mapItem) {
	for i, v := range tbl {
		reqCmd := Cmd{v.cmd0, v.cmd1}
		respCmd := Cmd{v.cmd0, v.respCmd}
		_reqCmdMap[reqCmd.GetCmdID()] = &tbl[i]
		_respCmdMap[respCmd.GetCmdID()] = &tbl[i]
		if v.reqType != nil {
			_reqTypeMap[v.reqType] = &tbl[i]
		}
		if v.respType != nil {
			_respTypeMap[v.respType] = &tbl[i]
		}
	}
}

func getItemFromCmd(m map[uint16]*mapItem, cmd Cmd) (*mapItem, bool) {
	v, ok := m[cmd.GetCmdID()]
	return v, ok
}

func GetCmdFromRequest(req interface{}) (Cmd, bool) {
	st := reflect.TypeOf(req)
	if st.Kind() == reflect.Ptr { // the req may be struct address
		st = st.Elem()
	}
	v, ok := _reqTypeMap[st]
	if ok {
		return Cmd{v.cmd0 | v.reqT, v.cmd1}, true
	}
	return Cmd{}, false
}

func GetCmdFromRespone(req interface{}) (Cmd, bool) {
	st := reflect.TypeOf(req)
	if st.Kind() == reflect.Ptr { // the req may be struct address
		st = st.Elem()
	}
	v, ok := _respTypeMap[st]
	if ok {
		return Cmd{v.cmd0 | v.respT, v.respCmd}, true
	}
	return Cmd{}, false
}

func (c *Cmd) getReqType() (uint8, bool) {
	t, ok := getItemFromCmd(_reqCmdMap, *c)
	if !ok {
		return 0, false
	}
	return t.reqT, true
}

func (c *Cmd) AddReqType() bool {
	t, ok := c.getReqType()
	if !ok {
		return false
	}
	c.cmd0 |= t
	return true
}

func (c *Cmd) IsRespCmdMatch(respCmd Cmd) bool {
	reqItem, reqOk := getItemFromCmd(_reqCmdMap, *c)
	respItem, respOk := getItemFromCmd(_respCmdMap, respCmd)
	return reqOk && respOk && reqItem == respItem
}

//func getRespTypeFromCmd(cmd0, cmd1 uint8) uint8 {
//	t, ok := getItemFromCmd(_respCmdMap, cmd0, cmd1)
//	if !ok {
//		panic(errors.New(fmt.Sprintf("invalid response cmd(%d,%d)", cmd0, cmd1)))
//	}
//	return t.reqT
//}

// if this cmd contain CMD_TYPE try get the item and whether it is request
func guessMapItem(c Cmd) (*mapItem, bool, bool) {
	var item *mapItem
	var ok bool
	if c.cmd0&MT_RPC_CMD_TYPE_MASK == MT_RPC_CMD_SREQ || c.cmd0&MT_RPC_CMD_TYPE_MASK == MT_RPC_CMD_AREQ { //assume it is request
		item, ok = getItemFromCmd(_reqCmdMap, c)
	}
	if !ok {
		item, ok = getItemFromCmd(_respCmdMap, c) // try find it in response
		if !ok {
			return nil, false, false
		}
		return item, false, true

	}
	return item, true, true

}

func (c *Cmd) getDesc() string {
	item, _, ok := guessMapItem(*c)
	if ok {
		return item.desc
	}
	return ""

}

func (c *Cmd) newObject() (interface{}, bool) {
	var t reflect.Type
	item, isReq, ok := guessMapItem(*c)
	if ok {
		if isReq {
			t = item.reqType
		} else {
			t = item.respType
		}
		if t == nil { // we want to ingore the struct
			return nil, true
		}
		return reflect.New(t).Interface(), true
	}
	return nil, false

}
