package codec

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

type Cmd struct {
	cmd0 uint8
	cmd1 uint8
}

func (c *Cmd) GetCmdType() uint8 {
	return c.cmd0 & MT_RPC_CMD_TYPE_MASK
}

func (c *Cmd) GetCmdID() uint16 {
	return uint16(c.cmd0&MT_RPC_SUBSYSTEM_MASK)<<8 | uint16(c.cmd1)
}
