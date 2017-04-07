package codec

import "reflect"

const (
	/* AREQ from host */
	MT_SYS_RESET_REQ = 0x00

	/* SREQ/SRSP */
	MT_SYS_PING              = 0x01
	MT_SYS_VERSION           = 0x02
	MT_SYS_SET_EXTADDR       = 0x03
	MT_SYS_GET_EXTADDR       = 0x04
	MT_SYS_RAM_READ          = 0x05
	MT_SYS_RAM_WRITE         = 0x06
	MT_SYS_OSAL_NV_ITEM_INIT = 0x07
	MT_SYS_OSAL_NV_READ      = 0x08
	MT_SYS_OSAL_NV_WRITE     = 0x09
	MT_SYS_OSAL_START_TIMER  = 0x0A
	MT_SYS_OSAL_STOP_TIMER   = 0x0B
	MT_SYS_RANDOM            = 0x0C
	MT_SYS_ADC_READ          = 0x0D
	MT_SYS_GPIO              = 0x0E
	MT_SYS_STACK_TUNE        = 0x0F
	MT_SYS_SET_TIME          = 0x10
	MT_SYS_GET_TIME          = 0x11
	MT_SYS_OSAL_NV_DELETE    = 0x12
	MT_SYS_OSAL_NV_LENGTH    = 0x13
	MT_SYS_SET_TX_POWER      = 0x14

	MT_SYS_GET_ANTENNA_MODE = 0X70
	/* AREQ to host */
	MT_SYS_RESET_IND          = 0x80
	MT_SYS_OSAL_TIMER_EXPIRED = 0x81

	/* Schneider Specific Indications */
	MT_SYS_SCHNEIDER_NV_CHANGE_IND = 0xF0
)

const (
	// OSAL NV item IDs
	ZCD_NV_EXTADDR        = 0x0001
	ZCD_NV_BOOTCOUNTER    = 0x0002
	ZCD_NV_STARTUP_OPTION = 0x0003
	ZCD_NV_START_DELAY    = 0x0004

	// NWK Layer NV item IDs
	ZCD_NV_NIB                      = 0x0021
	ZCD_NV_DEVICE_LIST              = 0x0022
	ZCD_NV_ADDRMGR                  = 0x0023
	ZCD_NV_POLL_RATE                = 0x0024
	ZCD_NV_QUEUED_POLL_RATE         = 0x0025
	ZCD_NV_RESPONSE_POLL_RATE       = 0x0026
	ZCD_NV_REJOIN_POLL_RATE         = 0x0027
	ZCD_NV_DATA_RETRIES             = 0x0028
	ZCD_NV_POLL_FAILURE_RETRIES     = 0x0029
	ZCD_NV_STACK_PROFILE            = 0x002A
	ZCD_NV_INDIRECT_MSG_TIMEOUT     = 0x002B
	ZCD_NV_ROUTE_EXPIRY_TIME        = 0x002C
	ZCD_NV_EXTENDED_PAN_ID          = 0x002D
	ZCD_NV_BCAST_RETRIES            = 0x002E
	ZCD_NV_PASSIVE_ACK_TIMEOUT      = 0x002F
	ZCD_NV_BCAST_DELIVERY_TIME      = 0x0030
	ZCD_NV_NWK_MODE                 = 0x0031
	ZCD_NV_CONCENTRATOR_ENABLE      = 0x0032
	ZCD_NV_CONCENTRATOR_DISCOVERY   = 0x0033
	ZCD_NV_CONCENTRATOR_RADIUS      = 0x0034
	ZCD_NV_CONCENTRATOR_RC          = 0x0036
	ZCD_NV_NWK_MGR_MODE             = 0x0037
	ZCD_NV_SRC_RTG_EXPIRY_TIME      = 0x0038
	ZCD_NV_ROUTE_DISCOVERY_TIME     = 0x0039
	ZCD_NV_NWK_ACTIVE_KEY_INFO      = 0x003A
	ZCD_NV_NWK_ALTERN_KEY_INFO      = 0x003B
	ZCD_NV_ROUTER_OFF_ASSOC_CLEANUP = 0x003C
	ZCD_NV_NWK_LEAVE_REQ_ALLOWED    = 0x003D
	ZCD_NV_NWK_CHILD_AGE_ENABLE     = 0x003E
	ZCD_NV_DEVICE_LIST_KA_TIMEOUT   = 0x003F

	// APS Layer NV item IDs
	ZCD_NV_BINDING_TABLE           = 0x0041
	ZCD_NV_GROUP_TABLE             = 0x0042
	ZCD_NV_APS_FRAME_RETRIES       = 0x0043
	ZCD_NV_APS_ACK_WAIT_DURATION   = 0x0044
	ZCD_NV_APS_ACK_WAIT_MULTIPLIER = 0x0045
	ZCD_NV_BINDING_TIME            = 0x0046
	ZCD_NV_APS_USE_EXT_PANID       = 0x0047
	ZCD_NV_APS_USE_INSECURE_JOIN   = 0x0048
	ZCD_NV_COMMISSIONED_NWK_ADDR   = 0x0049

	ZCD_NV_APS_NONMEMBER_RADIUS     = 0x004B // Multicast non_member radius
	ZCD_NV_APS_LINK_KEY_TABLE       = 0x004C
	ZCD_NV_APS_DUPREJ_TIMEOUT_INC   = 0x004D
	ZCD_NV_APS_DUPREJ_TIMEOUT_COUNT = 0x004E
	ZCD_NV_APS_DUPREJ_TABLE_SIZE    = 0x004F

	// Security NV Item IDs
	ZCD_NV_SECURITY_LEVEL         = 0x0061
	ZCD_NV_PRECFGKEY              = 0x0062
	ZCD_NV_PRECFGKEYS_ENABLE      = 0x0063
	ZCD_NV_SECURITY_MODE          = 0x0064
	ZCD_NV_SECURE_PERMIT_JOIN     = 0x0065
	ZCD_NV_APS_LINK_KEY_TYPE      = 0x0066
	ZCD_NV_APS_ALLOW_R19_SECURITY = 0x0067

	ZCD_NV_IMPLICIT_CERTIFICATE = 0x0069
	ZCD_NV_DEVICE_PRIVATE_KEY   = 0x006A
	ZCD_NV_CA_PUBLIC_KEY        = 0x006B

	ZCD_NV_USE_DEFAULT_TCLK = 0x006D
	ZCD_NV_TRUSTCENTER_ADDR = 0x006E
	ZCD_NV_RNG_COUNTER      = 0x006F
	ZCD_NV_RANDOM_SEED      = 0x0070

	// ZDO NV Item IDs
	ZCD_NV_USERDESC      = 0x0081
	ZCD_NV_NWKKEY        = 0x0082
	ZCD_NV_PANID         = 0x0083
	ZCD_NV_CHANLIST      = 0x0084
	ZCD_NV_LEAVE_CTRL    = 0x0085
	ZCD_NV_SCAN_DURATION = 0x0086
	ZCD_NV_LOGICAL_TYPE  = 0x0087
	ZCD_NV_NWKMGR_MIN_TX = 0x0088
	ZCD_NV_NWKMGR_ADDR   = 0x0089

	ZCD_NV_ZDO_DIRECT_CB = 0x008F

	// ZCL NV item IDs
	ZCD_NV_SCENE_TABLE       = 0x0091
	ZCD_NV_MIN_FREE_NWK_ADDR = 0x0092
	ZCD_NV_MAX_FREE_NWK_ADDR = 0x0093
	ZCD_NV_MIN_FREE_GRP_ID   = 0x0094
	ZCD_NV_MAX_FREE_GRP_ID   = 0x0095
	ZCD_NV_MIN_GRP_IDS       = 0x0096
	ZCD_NV_MAX_GRP_IDS       = 0x0097

	// Non-standard NV item IDs
	ZCD_NV_SAPI_ENDPOINT = 0x00A1

	// NV Items Reserved for Commissioning Cluster Startup Attribute Set (SAS):
	// 0x00B1 - 0x00BF: Parameters related to APS and NWK layers
	// 0x00C1 - 0x00CF: Parameters related to Security
	// 0x00D1 - 0x00DF: Current key parameters
	ZCD_NV_SAS_SHORT_ADDR    = 0x00B1
	ZCD_NV_SAS_EXT_PANID     = 0x00B2
	ZCD_NV_SAS_PANID         = 0x00B3
	ZCD_NV_SAS_CHANNEL_MASK  = 0x00B4
	ZCD_NV_SAS_PROTOCOL_VER  = 0x00B5
	ZCD_NV_SAS_STACK_PROFILE = 0x00B6
	ZCD_NV_SAS_STARTUP_CTRL  = 0x00B7

	ZCD_NV_SAS_TC_ADDR         = 0x00C1
	ZCD_NV_SAS_TC_MASTER_KEY   = 0x00C2
	ZCD_NV_SAS_NWK_KEY         = 0x00C3
	ZCD_NV_SAS_USE_INSEC_JOIN  = 0x00C4
	ZCD_NV_SAS_PRECFG_LINK_KEY = 0x00C5
	ZCD_NV_SAS_NWK_KEY_SEQ_NUM = 0x00C6
	ZCD_NV_SAS_NWK_KEY_TYPE    = 0x00C7
	ZCD_NV_SAS_NWK_MGR_ADDR    = 0x00C8

	ZCD_NV_SAS_CURR_TC_MASTER_KEY   = 0x00D1
	ZCD_NV_SAS_CURR_NWK_KEY         = 0x00D2
	ZCD_NV_SAS_CURR_PRECFG_LINK_KEY = 0x00D3

	// NV Items Reserved for Trust Center Link Key Table entries
	// 0x0101 - 0x01FF
	ZCD_NV_TCLK_TABLE_START = 0x0101
	ZCD_NV_TCLK_TABLE_END   = 0x01FF

	// NV Items Reserved for APS Link Key Table entries
	// 0x0201 - 0x02FF
	ZCD_NV_APS_LINK_KEY_DATA_START = 0x0201 // APS key data
	ZCD_NV_APS_LINK_KEY_DATA_END   = 0x02FF

	// NV Items Reserved for Master Key Table entries
	// 0x0301 - 0x03FF
	ZCD_NV_MASTER_KEY_DATA_START = 0x0301 // Master key data
	ZCD_NV_MASTER_KEY_DATA_END   = 0x03FF

	// NV Items Reserved for applications (user applications)
	// 0x0401 锟?0x0FFF

	MT_MAX_NV_CHANGE_NTF_ITEMS = 10
)

const (
	// ZCD_NV_STARTUP_OPTION values
	//   These are bit weighted - you can OR these together.
	//   Setting one of these bits will set their associated NV items
	//   to code initialized values.
	ZCD_STARTOPT_DEFAULT_CONFIG_STATE  = 0x01
	ZCD_STARTOPT_DEFAULT_NETWORK_STATE = 0x02
	ZCD_STARTOPT_AUTO_START            = 0x04
	ZCD_STARTOPT_CLEAR_CONFIG          = ZCD_STARTOPT_DEFAULT_CONFIG_STATE
	ZCD_STARTOPT_CLEAR_STATE           = ZCD_STARTOPT_DEFAULT_NETWORK_STATE

	DEVICETYPE_COORDINATOR = 0x00
	DEVICETYPE_ROUTER      = 0x01
	DEVICETYPE_ENDDEVICE   = 0x02

	ZCL_KE_IMPLICIT_CERTIFICATE_LEN = 48
	ZCL_KE_CA_PUBLIC_KEY_LEN        = 22
	ZCL_KE_DEVICE_PRIVATE_KEY_LEN   = 21

	ANTENNA_MODE_1         = 1
	ANTENNA_MODE_2         = 2
	ANTENNA_MODE_DIVERSITY = 3

	ProductId_TI_Znp                         = 0x00
	ProductId_Schneider_Pro_Znp              = 0x10
	ProductId_Schneider_GP_15_4_Brick        = 0x14
	ProductId_Schneider_ZB_Emc_Brick         = 0x16
	ProductId_Schneider_GPD_15_4_Brick       = 0x17
	ProductId_Schneider_Pro_GP_Znp           = 0x18
	ProductId_Schneider_Reserved_Future_Znp1 = 0x19
	ProductId_Schneider_Reserved_Future_Znp2 = 0x1A
	ProductId_Schneider_Reserved_Future_Znp3 = 0x1B
	ProductId_Schneider_Reserved_Future_Znp4 = 0x1C
	ProductId_Schneider_Reserved_Future_Znp5 = 0x1D
	ProductId_Schneider_Reserved_Future_Znp6 = 0x1E
	ProductId_Schneider_Reserved_Future_Znp7 = 0x1F
	ProductId_Schneider_Sbl                  = 0xFE
	ProductId_Unknown                        = 0xFF

	HardwareType_Error                            = 0x00
	HardwareType_CC2530_TI_EB                     = 0x01
	HardwareType_CC2531_TI_USB_Dongle             = 0x02
	HardwareType_CC2531_Spectec_USB_Dongle_PA     = 0x03
	HardwareType_CC2530_OUREA_V2                  = 0x04
	HardwareType_CC2531_SE_Inventek_USB_Dongle_PA = 0x05
	HardwareType_CC2538_TI_EB                     = 0x06
	HardwareType_CC2530_Nova                      = 0x07
	HardwareType_CC2530_SmartlinkIpz              = 0x08
	HardwareType_CC2538_OUREA_UART                = 0x09
	HardwareType_CC2538_OUREA_USB                 = 0x0A
	HardwareType_Not_Init                         = 0xFF
)

type PingSrspFormat struct {
	Capabilities uint16
}

type SetExtAddrFormat struct {
	ExtAddr [8]uint8
}

type GetExtAddrSrspFormat struct {
	ExtAddr uint64
}

type RamReadFormat struct {
	Address uint16
	Len     uint8
}

type RamReadSrspFormat struct {
	Status uint8
	Value  []uint8 `lentype:"uint8"`
}

type RamWriteFormat struct {
	Address uint16
	Value   []uint8 `lentype:"uint8"`
}

type ResetReqFormat struct {
	Type uint8
}

type ResetIndFormat struct {
	Reason       uint8
	TransportRev uint8
	ProductId    uint8
	MajorRel     uint8
	MinorRel     uint8
	HwRev        uint8
}

type VersionSrspFormat struct {
	TransportRev uint8
	Product      uint8
	MajorRel     uint8
	MinorRel     uint8
	MaintRel     uint8
	HwRev        uint8
}

type OsalNvReadFormat struct {
	Id     uint16
	Offset uint8
}

type OsalNvReadSrspFormat struct {
	Status uint8
	Value  []uint8 `lentype:"uint8" `
}

type OsalNvWriteFormat struct {
	Id     uint16
	Offset uint8
	Value  []uint8 ` lentype:"uint8" `
}

type OsalNvItemInitFormat struct {
	Id       uint16
	ItemLen  uint16
	InitData []uint8 ` lentype:"uint8" `
}

type OsalNvDeleteFormat struct {
	Id      uint16
	ItemLen uint16
}

type OsalNvLengthFormat struct {
	Id uint16
}

type OsalNvLengthSrspFormat struct {
	ItemLen uint16
}

type OsalStartTimerFormat struct {
	Id      uint8
	Timeout uint16
}

type OsalStopTimerFormat struct {
	Id uint8
}

type OsalTimerExpiredFormat struct {
	Id uint8
}

type StackTuneFormat struct {
	Operation uint8
	Value     uint8
}

type StackTuneSrspFormat struct {
	Value uint8
}

type AdcReadFormat struct {
	Channel    uint8
	Resolution uint8
}

type AdcReadSrspFormat struct {
	Value uint16
}

type GpioFormat struct {
	Operation uint8
	Value     uint8
}

type GpioSrspFormat struct {
	Value uint8
}

type RandomSrspFormat struct {
	Value uint16
}

type SetTimeFormat struct {
	UTCTime [4]uint8
	Hour    uint8
	Minute  uint8
	Second  uint8
	Month   uint8
	Day     uint8
	Year    uint16
}

type GetTimeSrspFormat struct {
	UTCTime uint32
	Hour    uint8
	Minute  uint8
	Second  uint8
	Month   uint8
	Day     uint8
	Year    uint16
}

type SetTxPowerFormat struct {
	TxPower uint8
}

type SetTxPowerSrspFormat struct {
	TxPower uint8
}

// Schneider specific command

type GetAntennaModeSrspFormat struct {
	Status                uint8
	Version               uint8
	SupportedAntennaModes uint8
	CurrentAntennaMode    uint8
}

type SENvChangeIndFormat struct {
	Value []uint16 `lentype:"uint8" `
}

func init() {
	addSubCommandMap([]mapItem{
		{MT_RPC_SYS_SYS, MT_SYS_PING, MT_SYS_PING, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, nil, reflect.TypeOf(PingSrspFormat{}),
			"SYS_PING"},
		{MT_RPC_SYS_SYS, MT_SYS_SET_EXTADDR, MT_SYS_SET_EXTADDR, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, reflect.TypeOf(SetExtAddrFormat{}), nil,
			"SYS_SET_EXTADDR"},
		{MT_RPC_SYS_SYS, MT_SYS_GET_EXTADDR, MT_SYS_GET_EXTADDR, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, nil, reflect.TypeOf(GetExtAddrSrspFormat{}),
			"SYS_GET_EXTADDR"},
		{MT_RPC_SYS_SYS, MT_SYS_RAM_READ, MT_SYS_RAM_READ, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, reflect.TypeOf(RamReadFormat{}), reflect.TypeOf(RamReadSrspFormat{}),
			"SYS_RAM_READ"},
		{MT_RPC_SYS_SYS, MT_SYS_RAM_WRITE, MT_SYS_RAM_WRITE, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, reflect.TypeOf(RamWriteFormat{}), nil,
			"SYS_RAM_WRITE"},
		{MT_RPC_SYS_SYS, MT_SYS_RESET_REQ, MT_SYS_RESET_IND, MT_RPC_CMD_AREQ, 0, reflect.TypeOf(ResetReqFormat{}), reflect.TypeOf(ResetIndFormat{}),
			"SYS_RESET_REQ"},
		{MT_RPC_SYS_SYS, MT_SYS_VERSION, MT_SYS_VERSION, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, nil, reflect.TypeOf(VersionSrspFormat{}),
			"SYS_VERSION"},
		{MT_RPC_SYS_SYS, MT_SYS_OSAL_NV_READ, MT_SYS_OSAL_NV_READ, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, reflect.TypeOf(OsalNvReadFormat{}), reflect.TypeOf(OsalNvReadSrspFormat{}),
			"SYS_OSAL_NV_READ"},
		{MT_RPC_SYS_SYS, MT_SYS_OSAL_NV_WRITE, MT_SYS_OSAL_NV_WRITE, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, reflect.TypeOf(OsalNvWriteFormat{}), nil,
			"SYS_OSAL_NV_WRITE"},

		{MT_RPC_SYS_SYS, MT_SYS_SET_TX_POWER, MT_SYS_SET_TX_POWER, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, reflect.TypeOf(SetTxPowerFormat{}), reflect.TypeOf(SetTxPowerSrspFormat{}),
			"SYS_SET_TX_POWER"},

		{MT_RPC_SYS_SYS, MT_SYS_GET_ANTENNA_MODE, MT_SYS_GET_ANTENNA_MODE, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, nil, reflect.TypeOf(GetAntennaModeSrspFormat{}),
			"SYS_GET_ANTENNA_MODE"},

		{MT_RPC_SYS_SYS, 0xff, MT_SYS_SCHNEIDER_NV_CHANGE_IND, 0, 0, nil, reflect.TypeOf(SENvChangeIndFormat{}),
			"SYS_SCHNEIDER_NV_CHANGE_IND"},
	})
}
