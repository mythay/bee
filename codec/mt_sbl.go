package codec

const (
	MT_SBL_RESET_REQ           = 0x00
	MT_SBL_WRITE_FLASH_REQ     = 0x01
	MT_SBL_READ_FLASH_REQ      = 0x02
	MT_SBL_ENABLE_FLASH_REQ    = 0x03
	MT_SBL_HANDSHAKE_REQ       = 0x04
	MT_SBL_VERSION_SBL_REQ     = 0x05
	MT_SBL_VERSION_APP_REQ     = 0x06
	MT_SBL_DESCRIPTOR_APP_REQ  = 0x07
	MT_SBL_START_APP_REQ       = 0x08
	MT_SBL_RESET_RESP          = 0x80
	MT_SBL_WRITE_FLASH_RESP    = 0x81
	MT_SBL_READ_FLASH_RESP     = 0x82
	MT_SBL_ENABLE_FLASH_RESP   = 0x83
	MT_SBL_HANDSHAKE_RESP      = 0x84
	MT_SBL_VERSION_SBL_RESP    = 0x85
	MT_SBL_VERSION_APP_RESP    = 0x86
	MT_SBL_DESCRIPTOR_APP_RESP = 0x87
	MT_SBL_START_APP_RSP       = 0x88
)

type SblAppVersionIndFormat struct {
	SblStatus    uint8
	TransportRev uint8
	Product      uint8
	MajorRel     uint8
	MinorRel     uint8
	MaintRel     uint8
}

type SblStartAppIndFormat struct {
	SblStatus uint8
}

// func (client *Client) SblVersionApp() (ind *SblAppVersionIndFormat, err error) {
// 	ind = &SblAppVersionIndFormat{}
// 	_, err = client.Call(cmd{MT_RPC_SYS_SBL, MT_SBL_VERSION_APP_REQ})
// 	if err != nil {
// 		return
// 	}
// 	_, err = client.WaitAsync(ind, time.Second*3)
// 	return
// }

// func (client *Client) SblStartApp() (ind *ResetIndFormat, err error) {
// 	ind = &ResetIndFormat{}
// 	_, err = client.Call(cmd{MT_RPC_SYS_SBL, MT_SBL_START_APP_REQ})
// 	if err != nil {
// 		return
// 	}
// 	status := &SblStartAppIndFormat{}

// 	_, err = client.WaitAsync(status, time.Second*3)
// 	if err != nil {
// 		return
// 	}
// 	if status.SblStatus != 0 {
// 		err = errors.New("start fail")
// 		return
// 	}
// 	_, err = client.WaitAsync(ind, time.Second*3)
// 	return
// }
