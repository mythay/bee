package codec

const (
	// SAPI MT Command Identifiers
	/* AREQ from Host */
	MT_SAPI_SYS_RESET = 0x09

	/* SREQ/SRSP */
	MT_SAPI_START_REQ           = 0x00
	MT_SAPI_BIND_DEVICE         = 0x01
	MT_SAPI_ALLOW_BIND          = 0x02
	MT_SAPI_SEND_DATA_REQ       = 0x03
	MT_SAPI_READ_CONFIGURATION  = 0x04
	MT_SAPI_WRITE_CONFIGURATION = 0x05
	MT_SAPI_GET_DEVICE_INFO     = 0x06
	MT_SAPI_FIND_DEVICE_REQ     = 0x07
	MT_SAPI_PERMIT_JOINING_REQ  = 0x08
	MT_SAPI_APP_REGISTER_REQ    = 0x0a

	/* AREQ to host */
	MT_SAPI_START_CNF        = 0x80
	MT_SAPI_BIND_CNF         = 0x81
	MT_SAPI_ALLOW_BIND_CNF   = 0x82
	MT_SAPI_SEND_DATA_CNF    = 0x83
	MT_SAPI_FIND_DEVICE_CNF  = 0x85
	MT_SAPI_RECEIVE_DATA_IND = 0x87
)

const ( //configuration ID
	MT_SAPI_NV_CFG_ID_STARTUP_OPTION        = 0x0003 /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_LOGICAL_TYPE          = 0x0087 /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_DIRECT_CB             = 0x008F /* Size: 2byte */
	MT_SAPI_NV_CFG_ID_POLL_RATE             = 0x0024 /* Size: 2byte */
	MT_SAPI_NV_CFG_ID_QUEUED_POLL_RATE      = 0x0025 /* Size: 2bytes */
	MT_SAPI_NV_CFG_ID_RESPONSE_POLL_RATE    = 0x0026 /* Size: 2byte */
	MT_SAPI_NV_CFG_ID_POLL_FAILURE_RETRIES  = 0x0029 /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_INDIRECT_MSG_TIMEOUT  = 0x002B /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_EXTENDED_PAN_ID       = 0x002D /* Size: 8byte */
	MT_SAPI_NV_CFG_ID_APS_FRAME_RETRIES     = 0x0043 /* Size: 1bytes */
	MT_SAPI_NV_CFG_ID_APS_ACK_WAIT_DURAION  = 0x0044 /* Size: 2bytes */
	MT_SAPI_NV_CFG_ID_BINDING_TIME          = 0x0046 /* Size: 2bytes */
	MT_SAPI_NV_CFG_ID_APSF_WINDOW_SIZE      = 0x0049 /* Size: 3bytes */
	MT_SAPI_NV_CFG_ID_APSF_INTERFRAME_DELAY = 0x004A /* Size: 2bytes */
	MT_SAPI_NV_CFG_ID_USERDESC              = 0x0081 /* Size: 17bytes */
	MT_SAPI_NV_CFG_ID_NWKKEY                = 0x0082 /* Size: ?bytes */ // NEW MVDB CLONE TEST
	MT_SAPI_NV_CFG_ID_PANID                 = 0x0083 /* Size: 2bytes */
	MT_SAPI_NV_CFG_ID_CHANLIST              = 0x0084 /* Size: 4bytes */
	MT_SAPI_NV_CFG_ID_PRECFGKEY             = 0x0062 /* Size: 16bytes */
	MT_SAPI_NV_CFG_ID_PRECFGKEYS_ENABLE     = 0x0063 /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_SECURITY_MODE         = 0x0064 /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_USE_DEFAULT_TLCK      = 0x006D /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_BCAST_RETRIES         = 0x002E /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_PASSIVE_ACK_TIMEOUT   = 0x002F /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_BCAST_DELIVERY_TIME   = 0x0030 /* Size: 1byte */
	MT_SAPI_NV_CFG_ID_ROUTE_EXPIRY_TIME     = 0x002C /* Size: 1byte */

)
const (
	mt_ZCD_STARTOPT_DEFAULT_CONFIG_STATE  = 0x01
	mt_ZCD_STARTOPT_DEFAULT_NETWORK_STATE = 0x02
	mt_ZCD_STARTOPT_AUTO_START            = 0x04
	mt_ZCD_STARTOPT_CLEAR_CONFIG          = mt_ZCD_STARTOPT_DEFAULT_CONFIG_STATE | mt_ZCD_STARTOPT_DEFAULT_NETWORK_STATE
	mt_ZCD_STARTOPT_CLEAR_STATE           = mt_ZCD_STARTOPT_DEFAULT_NETWORK_STATE
)
const (
	mt_SAPI_DEV_INFO_TYPE_DEV_STATE          = 0 /* 1 byte */
	mt_SAPI_DEV_INFO_TYPE_DEV_IEEE_ADDR      = 1 /* 8 bytes */
	mt_SAPI_DEV_INFO_TYPE_DEV_SHORT_ADDR     = 2 /* 2 bytes */
	mt_SAPI_DEV_INFO_TYPE_PARENT_SHORT_ADDR  = 3 /* 2 bytes */
	mt_SAPI_DEV_INFO_TYPE_PARENT_IEEE_ADDR   = 4 /* 8 bytes */
	mt_SAPI_DEV_INFO_TYPE_ZB_CHANNEL         = 5 /* 1 byte */
	mt_SAPI_DEV_INFO_TYPE_ZB_PAN_ID          = 6 /* 2 bytes */
	mt_SAPI_DEV_INFO_TYPE_ZB_EXTENDED_PAN_ID = 7 /* 8 bytes */

	/* Schneider Extension */
	mt_SAPI_DEV_INFO_TYPE_KEY_SEQUENCE_NUMBER = 0x80 /* 1 byte */
	mt_SAPI_DEV_INFO_TYPE_MEM_METRICS         = 0x81 /* 4 bytes: szl_uint16 Usage, szl_uint16 highWater */

)

type AppRegisterReqFormat struct {
	AppEndpoint        uint8
	AppProfileId       uint16
	DeviceId           uint16
	DeviceVersion      uint8
	Unused             uint8
	InputCommandsList  []uint16 `lentype:"uint8" `
	OutputCommandsList []uint16 `lentype:"uint8" `
}

type PermitJoiningReqFormat struct {
	Destination uint16
	Timeout     uint8
}

type BindDeviceFormat struct {
	Create    uint8
	CommandId uint16
	DstIeee   [8]uint8
}

type AllowBindFormat struct {
	Timeout uint8
}

type SendDataReqFormat struct {
	Destination uint16
	CommandId   uint16
	Handle      uint8
	Ack         uint8
	Radius      uint8
	Data        []uint8 `lentype:"uint8" `
}

type FindDeviceReqFormat struct {
	SearchKey [8]uint8
}

type WriteConfigurationFormat struct {
	ConfigId uint8
	Value    []uint8 `lentype:"uint8" `
}

type GetDeviceInfoFormat struct {
	Param uint8
}

type ReadConfigurationFormat struct {
	ConfigId uint8
}

type ReadConfigurationSrspFormat struct {
	Status   uint8
	ConfigId uint8
	Len      uint8
	Value    [128]uint8
}

type GetDeviceInfoSrspFormat struct {
	Param uint8
	Value [8]uint8
}

type FindDeviceCnfFormat struct {
	SearchKey uint16
	Result    uint64
}

type SendDataCnfFormat struct {
	Handle uint8
	Status uint8
}

type ReceiveDataIndFormat struct {
	Source  uint16
	Command uint16
	Data    [84]uint8 `lentype:"uint16" `
}

type AllowBindCnfFormat struct {
	Source uint16
}

type BindCnfFormat struct {
	CommandId uint16
	Status    uint8
}

type StartCnfFormat struct {
	Status uint8
}

//func (self *Client) ZbSystemReset() error {
//	_, err := self.Call(cmd{MT_RPC_SYS_SAPI, MT_SAPI_SYS_RESET})
//	return err
//}
//func (self *Client) ZbAppRegisterReq(req *AppRegisterReqFormat) error {
//	_, err := self.Call(cmd{MT_RPC_SYS_SAPI, MT_SAPI_APP_REGISTER_REQ})
//	return err
//}
//func (self *Client) ZbStartReq() error {
//	_, err := self.Call(cmd{MT_RPC_SYS_SAPI, MT_SAPI_START_REQ})
//	return err
//}
//func (self *Client) ZbPermitJoiningReq(req *PermitJoiningReqFormat) error {

//	_, err := self.Call(MT_RPC_SYS_SAPI, MT_SAPI_PERMIT_JOINING_REQ, req)
//	return err
//}

//func (self *Client) ZbBindDevice(req *BindDeviceFormat) error {
//	_, err := self.Call(MT_RPC_SYS_SAPI, MT_SAPI_BIND_DEVICE, req)
//	return err
//}

//func (self *Client) ZbAllowBind(req *AllowBindFormat) error {
//	_, err := self.Call(MT_RPC_SYS_SAPI, MT_SAPI_ALLOW_BIND, req)
//	return err
//}

//func (self *Client) ZbSendDataReq(req *SendDataReqFormat) error {
//	_, err := self.Call(MT_RPC_SYS_SAPI, MT_SAPI_SEND_DATA_REQ, req)
//	return err
//}

//func (self *Client) ZbFindDeviceReq(req *FindDeviceReqFormat) error {
//	_, err := self.Call(MT_RPC_SYS_SAPI, MT_SAPI_FIND_DEVICE_REQ, req)
//	return err
//}

// func (self *Client) ZbWriteConfiguration(req *WriteConfigurationFormat) error {
// 	_, err := self.Call(req)
// 	return err
// }

// func (c *Client) zbWriteConfig(configId uint8, value ...interface{}) error {
// 	buf := &bytes.Buffer{}
// 	var err error
// 	for _, v := range value {
// 		err = binary.Write(buf, binary.LittleEndian, v)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	req := &WriteConfigurationFormat{ConfigId: configId}
// 	if buf.Len() > len(req.Value) {
// 		req.Len = uint8(len(req.Value))
// 	}
// 	copy(req.Value[:], buf.Bytes())
// 	return c.ZbWriteConfiguration(req)
// }

// func (c *Client) zbReadDeviceInfo(configID uint8, value ...interface{}) error {
// 	resp, err := c.ZbGetDeviceInfo(&GetDeviceInfoFormat{configID})
// 	if err != nil {
// 		return err
// 	}
// 	buf := bytes.NewBuffer(resp.Value[:])
// 	for _, v := range value {
// 		err = binary.Read(buf, binary.LittleEndian, v)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// type errorCheck struct {
// 	c   *Client
// 	err error
// }

// func (c *errorCheck) ecWriteConfig(configId uint8, value ...interface{}) {
// 	if c.err != nil {
// 		return
// 	}
// 	c.err = c.c.zbWriteConfig(configId, value...)
// }

// func (c *errorCheck) ecGetDeviceInfo(configId uint8, value ...interface{}) {
// 	if c.err != nil {
// 		return
// 	}
// 	c.err = c.c.zbReadDeviceInfo(configId, value...)
// }

// func (self *Client) ZbGetDeviceInfo(req *GetDeviceInfoFormat) (*GetDeviceInfoSrspFormat, error) {
// 	resp, err := self.Call(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp.(*GetDeviceInfoSrspFormat), err
// }

// //func (self *Client) ZbReadConfiguration(req *ReadConfigurationFormat) (*ReadConfigurationSrspFormat, error) {
// //	resp, err := self.Call(MT_RPC_SYS_SAPI, MT_SAPI_READ_CONFIGURATION, req)
// //	if err != nil {
// //		return nil, err
// //	}
// //	return resp.(*ReadConfigurationSrspFormat), err
// //}
