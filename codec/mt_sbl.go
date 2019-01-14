package codec

import "reflect"

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

func init() {
	addSubCommandMap([]mapItem{
		{MT_RPC_SYS_SBL, MT_SBL_START_APP_REQ, MT_SBL_START_APP_RSP, MT_RPC_CMD_AREQ, MT_RPC_CMD_SRSP, nil, reflect.TypeOf(SblStartAppIndFormat{}),
			"SBL_START_APP_REQ"},
		{MT_RPC_SYS_SBL, MT_SBL_VERSION_APP_REQ, MT_SBL_VERSION_APP_RESP, MT_RPC_CMD_AREQ, MT_RPC_CMD_SRSP, nil, reflect.TypeOf(SblAppVersionIndFormat{}),
			"SBL_VERSION_APP_REQ"},
		{MT_RPC_SYS_SBL, MT_SBL_VERSION_SBL_REQ, MT_SBL_VERSION_SBL_RESP, MT_RPC_CMD_AREQ, MT_RPC_CMD_SRSP, nil, nil,
			"SBL_VERSION_SBL_REQ"},
	})
}
