package codec

import "reflect"

const (
	/* SREQ/SRSP */
	MT_ZDO_NWK_ADDR_REQ        = 0x00
	MT_ZDO_IEEE_ADDR_REQ       = 0x01
	MT_ZDO_NODE_DESC_REQ       = 0x02
	MT_ZDO_POWER_DESC_REQ      = 0x03
	MT_ZDO_SIMPLE_DESC_REQ     = 0x04
	MT_ZDO_ACTIVE_EP_REQ       = 0x05
	MT_ZDO_MATCH_DESC_REQ      = 0x06
	MT_ZDO_COMPLEX_DESC_REQ    = 0x07
	MT_ZDO_USER_DESC_REQ       = 0x08
	MT_ZDO_DEVICE_ANNCE        = 0x0A
	MT_ZDO_USER_DESC_SET       = 0x0B
	MT_ZDO_SERVER_DISC_REQ     = 0x0C
	MT_ZDO_END_DEVICE_BIND_REQ = 0x20
	MT_ZDO_BIND_REQ            = 0x21
	MT_ZDO_UNBIND_REQ          = 0x22

	MT_ZDO_SET_LINK_KEY      = 0x23
	MT_ZDO_REMOVE_LINK_KEY   = 0x24
	MT_ZDO_GET_LINK_KEY      = 0x25
	MT_ZDO_NWK_DISCOVERY_REQ = 0x26
	MT_ZDO_JOIN_REQ          = 0x27

	MT_ZDO_MGMT_NWK_DISC_REQ    = 0x30
	MT_ZDO_MGMT_LQI_REQ         = 0x31
	MT_ZDO_MGMT_RTG_REQ         = 0x32
	MT_ZDO_MGMT_BIND_REQ        = 0x33
	MT_ZDO_MGMT_LEAVE_REQ       = 0x34
	MT_ZDO_MGMT_DIRECT_JOIN_REQ = 0x35
	MT_ZDO_MGMT_PERMIT_JOIN_REQ = 0x36
	MT_ZDO_MGMT_NWK_UPDATE_REQ  = 0x37

	/* AREQ optional, but no AREQ response. */
	MT_ZDO_MSG_CB_REGISTER  = 0x3E
	MT_ZDO_MSG_CB_REMOVE    = 0x3F
	MT_ZDO_STARTUP_FROM_APP = 0x40

	/* AREQ from host */
	MT_ZDO_AUTO_FIND_DESTINATION = 0x41

	/* AREQ to host */
	MT_ZDO_AREQ_TO_HOST    = 0x80 /* Mark the start of the ZDO CId AREQs to host. */
	MT_ZDO_NWK_ADDR_RSP    = 0x80
	MT_ZDO_IEEE_ADDR_RSP   = 0x81
	MT_ZDO_NODE_DESC_RSP   = 0x82
	MT_ZDO_POWER_DESC_RSP  = 0x83
	MT_ZDO_SIMPLE_DESC_RSP = 0x84
	MT_ZDO_ACTIVE_EP_RSP   = 0x85
	MT_ZDO_MATCH_DESC_RSP  = 0x86

	MT_ZDO_COMPLEX_DESC_RSP = 0x90
	MT_ZDO_USER_DESC_RSP    = 0x91
	//                                     =    0x92 */ ((uint8)Discovery_Cache_req | 0x80)
	MT_ZDO_USER_DESC_CONF  = 0x94
	MT_ZDO_SERVER_DISC_RSP = 0x95

	MT_ZDO_END_DEVICE_BIND_RSP = 0xA0
	MT_ZDO_BIND_RSP            = 0xA1
	MT_ZDO_UNBIND_RSP          = 0xA2

	MT_ZDO_MGMT_NWK_DISC_RSP    = 0xB0
	MT_ZDO_MGMT_LQI_RSP         = 0xB1
	MT_ZDO_MGMT_RTG_RSP         = 0xB2
	MT_ZDO_MGMT_BIND_RSP        = 0xB3
	MT_ZDO_MGMT_LEAVE_RSP       = 0xB4
	MT_ZDO_MGMT_DIRECT_JOIN_RSP = 0xB5
	MT_ZDO_MGMT_PERMIT_JOIN_RSP = 0xB6

	//                                        /* 0xB8 */ ((uint8)Mgmt_NWK_Update_req | 0x80)

	MT_ZDO_STATE_CHANGE_IND     = 0xC0
	MT_ZDO_END_DEVICE_ANNCE_IND = 0xC1
	MT_ZDO_MATCH_DESC_RSP_SENT  = 0xC2
	MT_ZDO_STATUS_ERROR_RSP     = 0xC3
	MT_ZDO_SRC_RTG_IND          = 0xC4
	MT_ZDO_BEACON_NOTIFY_IND    = 0xC5
	MT_ZDO_JOIN_CNF             = 0xC6
	MT_ZDO_NWK_DISCOVERY_CNF    = 0xC7
	MT_ZDO_CONCENTRATOR_IND_CB  = 0xC8
	MT_ZDO_LEAVE_IND            = 0xC9

	MT_ZDO_LEAVE_LOCAL_IND = 0xce

	MT_ZDO_MSG_CB_INCOMING = 0xFF

	// Some arbitrarily chosen value for a default error status msg.
	// MtZdoDef_rsp                         0x0040

	/*ZDO Status Responses Definitions for ZDO Startup from App*/
	RESTORED_NETWORK   = 0x00
	NEW_NETWORK        = 0x01
	LEAVEANDNOTSTARTED = 0x02
)

const (
	MT_ZDO_HOLD              = iota //0        //Initialized - not started automatically
	MT_ZDO_INIT                     //1        //Initialized - not connected to anything
	MT_ZDO_NWK_DISC                 //2        //Discovering PAN's to join
	MT_ZDO_NWK_JOINING              //3        //Joining a PAN
	MT_ZDO_NWK_REJOIN               //4        //ReJoining a PAN, only for end devices
	MT_ZDO_END_DEVICE_UNAUTH        //5        //Joined but not yet authenticated by trust center
	MT_ZDO_END_DEVICE               //6        //Started as device after authentication
	MT_ZDO_ROUTER                   //7        //Device joined, authenticated and is a router
	MT_ZDO_COORD_STARTING           //8        //Starting as Zigbee Coordinator
	MT_ZDO_ZB_COORD                 //9        //Started as Zigbee Coordinator
	MT_ZDO_NWK_ORPHAN               //10       //Device has lost information about its parent

)

type NetworkListItemFormat struct {
	PanID             uint64
	LogicalChannel    uint8
	StackProfZigVer   uint8
	BeacOrdSupFramOrd uint8
	PermitJoin        uint8
}

type NeighborLqiListItemFormat struct {
	ExtendedPanID           uint64
	ExtendedAddress         uint64
	NetworkAddress          uint16
	DevTypRxOnWhenIdleRelat uint8
	PermitJoining           uint8
	Depth                   uint8
	LQI                     uint8
}

type RoutingTableListItemFormat struct {
	DstAddr uint16
	Status  uint8
	NextHop uint16
}

type BindingTableListItemFormat struct {
	SrcIEEEAddr uint64
	SrcEndpoint uint8
	ClusterID   uint8
	DstAddrMode uint8
	DstIEEEAddr uint64
	DstEndpoint uint8
}

type BeaconListItemFormat struct {
	SrcAddr        uint16
	PanID          uint16
	LogicalChannel uint8
	PermitJoining  uint8
	RouterCap      uint8
	DevCap         uint8
	ProtocolVer    uint8
	StackProf      uint8
	Lqi            uint8
	Depth          uint8
	UpdateID       uint8
	ExtendedPanID  uint64
}

type NwkAddrReqFormat struct {
	IEEEAddress [8]uint8
	ReqType     uint8
	StartIndex  uint8
}

type IeeeAddrReqFormat struct {
	ShortAddr  uint16
	ReqType    uint8
	StartIndex uint8
}

type NodeDescReqFormat struct {
	DstAddr           uint16
	NwkAddrOfInterest uint16
}

type PowerDescReqFormat struct {
	DstAddr           uint16
	NwkAddrOfInterest uint16
}

type SimpleDescReqFormat struct {
	DstAddr           uint16
	NwkAddrOfInterest uint16
	Endpoint          uint8
}

type ActiveEpReqFormat struct {
	DstAddr           uint16
	NwkAddrOfInterest uint16
}

type MatchDescReqFormat struct {
	DstAddr           uint16
	NwkAddrOfInterest uint16
	ProfileID         uint16
	InClusterList     []uint16 `lentype:"uint8" `
	OutClusterList    []uint16 `lentype:"uint8" `
}

type ComplexDescReqFormat struct {
	DstAddr           uint16
	NwkAddrOfInterest uint16
}

type UserDescReqFormat struct {
	DstAddr           uint16
	NwkAddrOfInterest uint16
}

type DeviceAnnceFormat struct {
	NWKAddr      uint16
	IEEEAddr     [8]uint8
	Capabilities uint8
}

type UserDescSetFormat struct {
	DstAddr           uint16
	NwkAddrOfInterest uint16
	UserDescriptor    []uint8 `lentype:"uint8" `
}

type ServerDiscReqFormat struct {
	ServerMask uint16
}

type EndDeviceBindReqFormat struct {
	DstAddr          uint16
	LocalCoordinator uint16
	CoordinatorIEEE  [8]uint8
	EndPoint         uint8
	ProfileID        uint16
	NumInClusters    uint8
	InClusterList    []uint16 `lentype:"uint8" `
	OutClusterList   []uint16 `lentype:"uint8" `
}

type BindReqFormat struct {
	DstAddr     uint16
	SrcAddress  [8]uint8
	SrcEndpoint uint8
	ClusterID   uint16
	DstAddrMode uint8
	DstAddress  [8]uint8
	DstEndpoint uint8
}

type UnbindReqFormat struct {
	DstAddr     uint16
	SrcAddress  [8]uint8
	SrcEndpoint uint8
	ClusterID   uint16
	DstAddrMode uint8
	DstAddress  [8]uint8
	DstEndpoint uint8
}

type MgmtNwkDiscReqFormat struct {
	DstAddr      uint16
	ScanChannels [4]uint8
	ScanDuration uint8
	StartIndex   uint8
}

type MgmtLqiReqFormat struct {
	DstAddr    uint16
	StartIndex uint8
}

type MgmtRtgReqFormat struct {
	DstAddr    uint16
	StartIndex uint8
}

type MgmtBindReqFormat struct {
	DstAddr    uint16
	StartIndex uint8
}

type MgmtLeaveReqFormat struct {
	DstAddr             uint16
	DeviceAddr          [8]uint8
	RemoveChildreRejoin uint8
}

type MgmtDirectJoinReqFormat struct {
	DstAddr    uint16
	DeviceAddr [8]uint8
	CapInfo    uint8
}

type MgmtPermitJoinReqFormat struct {
	AddrMode       uint8
	DstAddr        uint16
	Duration       uint8
	TCSignificance uint8
}

type MgmtNwkUpdateReqFormat struct {
	DstAddr        uint16
	DstAddrMode    uint8
	ChannelMask    [4]uint8
	ScanDuration   uint8
	ScanCount      uint8
	NwkManagerAddr uint16
}

type StartupFromAppFormat struct {
	StartDelay uint16
}

type AutoFindDestinationFormat struct {
	Endpoint uint8
}

type SetLinkKeyFormat struct {
	ShortAddr   uint16
	IEEEaddr    [8]uint8
	LinkKeyData [16]uint8
}

type RemoveLinkKeyFormat struct {
	IEEEaddr [8]uint8
}

type GetLinkKeyFormat struct {
	IEEEaddr [8]uint8
}

type GetLinkKeySrspFormat struct {
	Status      uint8
	IEEEAddr    uint64
	LinkKeyData [16]uint8
}

type NwkDiscoveryReqFormat struct {
	ScanChannels [4]uint8
	ScanDuration uint8
}

type JoinReqFormat struct {
	LogicalChannel uint8
	PanID          uint16
	ExtendedPanID  [8]uint8
	ChosenParent   uint16
	ParentDepth    uint8
	StackProfile   uint8
}

type MsgCbRegisterFormat struct {
	ClusterID uint16
}

type MsgCbRemoveFormat struct {
	ClusterID uint16
}

type NwkAddrRspFormat struct {
	Status       uint8
	IEEEAddr     uint64
	NwkAddr      uint16
	StartIndex   uint8
	AssocDevList []uint16 `lentype:"uint8" `
}

type IeeeAddrRspFormat struct {
	Status       uint8
	IEEEAddr     uint64
	NwkAddr      uint16
	StartIndex   uint8
	AssocDevList []uint16 `lentype:"uint8" `
}

type NodeDescRspFormat struct {
	SrcAddr                uint16
	Status                 uint8
	NwkAddr                uint16
	LoTyComDescAvUsrDesAv  uint8
	APSFlgFrqBnd           uint8
	MACCapFlg              uint8
	ManufacturerCode       uint16
	MaxBufferSize          uint8
	MaxTransferSize        uint16
	ServerMask             uint16
	MaxOutTransferSize     uint16
	DescriptorCapabilities uint8
}

type PowerDescRspFormat struct {
	SrcAddr                     uint16
	Status                      uint8
	NwkAddr                     uint16
	CurrntPwrModeAvalPwrSrcs    uint8
	CurrntPwrSrcCurrntPwrSrcLvl uint8
}

type SimpleDescRspFormat struct {
	SrcAddr        uint16
	Status         uint8
	NwkAddr        uint16
	Len            uint8
	Endpoint       uint8
	ProfileID      uint16
	DeviceID       uint16
	DeviceVersion  uint8
	InClusterList  [16]uint16 `lentype:"uint8" `
	OutClusterList [16]uint16 `lentype:"uint8" `
}

type ActiveEpRspFormat struct {
	SrcAddr      uint16
	Status       uint8
	NwkAddr      uint16
	ActiveEPList []uint8 `lentype:"uint8" `
}

type MatchDescRspFormat struct {
	SrcAddr   uint16
	Status    uint8
	NwkAddr   uint16
	MatchList []uint8 `lentype:"uint8" `
}

type ComplexDescRspFormat struct {
	SrcAddr     uint16
	Status      uint8
	NwkAddr     uint16
	ComplexList []uint8 `lentype:"uint8" `
}

type UserDescRspFormat struct {
	SrcAddr         uint16
	Status          uint8
	NwkAddr         uint16
	CUserDescriptor []uint8 `lentype:"uint8" `
}

type UserDescConfFormat struct {
	SrcAddr uint16
	Status  uint8
	NwkAddr uint16
}

type ServerDiscRspFormat struct {
	SrcAddr    uint16
	Status     uint8
	ServerMask uint16
}

type EndDeviceBindRspFormat struct {
	SrcAddr uint16
	Status  uint8
}

type BindRspFormat struct {
	SrcAddr uint16
	Status  uint8
}

type UnbindRspFormat struct {
	SrcAddr uint16
	Status  uint8
}

type MgmtNwkDiscRspFormat struct {
	SrcAddr      uint16
	Status       uint8
	NetworkCount uint8
	StartIndex   uint8
	NetworkList  []NetworkListItemFormat `lentype:"uint8" `
}

type MgmtLqiRspFormat struct {
	SrcAddr              uint16
	Status               uint8
	NeighborTableEntries uint8
	StartIndex           uint8
	NeighborLqiList      []NeighborLqiListItemFormat `lentype:"uint8" `
}

type MgmtRtgRspFormat struct {
	SrcAddr             uint16
	Status              uint8
	RoutingTableEntries uint8
	StartIndex          uint8
	RoutingTableList    []RoutingTableListItemFormat `lentype:"uint8" `
}

type MgmtBindRspFormat struct {
	SrcAddr             uint16
	Status              uint8
	BindingTableEntries uint8
	StartIndex          uint8
	BindingTableList    []BindingTableListItemFormat `lentype:"uint8" `
}

type MgmtLeaveRspFormat struct {
	SrcAddr uint16
	Status  uint8
}

type MgmtDirectJoinRspFormat struct {
	SrcAddr uint16
	Status  uint8
}

type EndDeviceAnnceIndFormat struct {
	SrcAddr      uint16
	NwkAddr      uint16
	IEEEAddr     uint64
	Capabilities uint8
}

type MatchDescRspSentFormat struct {
	NwkAddr        uint16
	InClusterList  []uint16 `lentype:"uint8" `
	OutClusterList []uint16 `lentype:"uint8" `
}

type StatusErrorRspFormat struct {
	SrcAddr uint16
	Status  uint8
}

type SrcRtgIndFormat struct {
	DstAddr   uint16
	RelayList []uint16 `lentype:"uint8" `
}

type BeaconNotifyIndFormat struct {
	BeaconList []BeaconListItemFormat `lentype:"uint8" `
}

type JoinCnfFormat struct {
	Status     uint8
	DevAddr    uint16
	ParentAddr uint16
}

type NwkDiscoveryCnfFormat struct {
	Status uint8
}

type LeaveIndFormat struct {
	SrcAddr uint16
	ExtAddr uint64
	Request uint8
	Remove  uint8
	Rejoin  uint8
}

type LeavelocalIndFormat struct {
	Rejoin uint8
}

type MsgCbIncomingFormat struct {
	SrcAddr      uint16
	WasBroadcast uint8
	ClusterID    uint16
	SecurityUse  uint8
	SeqNum       uint8
	MacDstAddr   uint16
	Status       uint8
	ExtAddr      uint64
	NwkAddr      uint16
	NotUsed      uint8
}

func init() {
	addSubCommandMap([]mapItem{
		{MT_RPC_SYS_ZDO, MT_ZDO_IEEE_ADDR_REQ, MT_ZDO_IEEE_ADDR_REQ, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, reflect.TypeOf(IeeeAddrReqFormat{}), reflect.TypeOf(IeeeAddrReqFormat{}),
			"ZDO_IEEE_ADDR_REQ"},

		{MT_RPC_SYS_ZDO, MT_ZDO_STARTUP_FROM_APP, MT_ZDO_STARTUP_FROM_APP, MT_RPC_CMD_SREQ, MT_RPC_CMD_SRSP, reflect.TypeOf(StartupFromAppFormat{}), reflect.TypeOf(uint8(0)),
			"ZDO_STARTUP_FROM_APP"},
		{MT_RPC_SYS_ZDO, 0xff, MT_ZDO_STATE_CHANGE_IND, 0, 0, nil, reflect.TypeOf(uint8(0)),
			"ZDO_STATE_CHANGE_IND"},
		{MT_RPC_SYS_ZDO, 0xff, MT_ZDO_LEAVE_LOCAL_IND, 0, 0, nil, reflect.TypeOf(LeavelocalIndFormat{}),
			"ZDO_LEAVE_LOCAL_IND"},
	})
}

//ZdoInit to start a network, 0 means resume, 1 means new network
// func (c *Client) ZdoInit() (uint8, error) {
// 	return c.zdoStartupFromApp(&StartupFromAppFormat{100})
// }

//func (self *Client) ZdoNwkAddrReq(req *NwkAddrReqFormat) (*NwkAddrRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_NWK_ADDR_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*NwkAddrRspFormat), err
//}
//func (self *Client) ZdoIeeeAddrReq(req *IeeeAddrReqFormat) (*IeeeAddrRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_IEEE_ADDR_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*IeeeAddrRspFormat), err
//}
//func (self *Client) ZdoNodeDescReq(req *NodeDescReqFormat) (*NodeDescRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_IEEE_ADDR_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*NodeDescRspFormat), err
//}
//func (self *Client) ZdoPowerDescReq(req *PowerDescReqFormat) (*PowerDescRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_POWER_DESC_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*PowerDescRspFormat), err
//}
//func (self *Client) ZdoSimpleDescReq(req *SimpleDescReqFormat) (*SimpleDescRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_SIMPLE_DESC_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*SimpleDescRspFormat), err
//}
//func (self *Client) ZdoActiveEpReq(req *ActiveEpReqFormat) (*ActiveEpRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_ACTIVE_EP_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*ActiveEpRspFormat), err
//}

////func (self *Client) ZdoMatchDescReq(MatchDescReqFormat_t *req)
////func (self *Client) ZdoComplexDescReq(ComplexDescReqFormat_t *req)
////func (self *Client) ZdoUserDescReq(UserDescReqFormat_t *req)
////func (self *Client) ZdoDeviceAnnce(DeviceAnnceFormat_t *req)
////func (self *Client) ZdoUserDescSet(UserDescSetFormat_t *req)
////func (self *Client) ZdoServerDiscReq(ServerDiscReqFormat_t *req)
////func (self *Client) ZdoEndDeviceBindReq(EndDeviceBindReqFormat_t *req)
////func (self *Client) ZdoBindReq(BindReqFormat_t *req)
////func (self *Client) ZdoUnbindReq(UnbindReqFormat_t *req)
//func (self *Client) ZdoMgmtNwkDiscReq(req *MgmtNwkDiscReqFormat) (*MgmtNwkDiscRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_MGMT_NWK_DISC_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*MgmtNwkDiscRspFormat), err
//}
//func (self *Client) ZdoMgmtLqiReq(req *MgmtLqiReqFormat) (*MgmtLqiRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_MGMT_LQI_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*MgmtLqiRspFormat), err
//}
//func (self *Client) ZdoMgmtRtgReq(req *MgmtRtgReqFormat) (*MgmtRtgRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_MGMT_RTG_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*MgmtRtgRspFormat), err
//}
//func (self *Client) ZdoMgmtBindReq(req *MgmtBindReqFormat) (*MgmtBindRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_MGMT_BIND_REQ, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*MgmtBindRspFormat), err
//}

//func (self *Client) ZdoMgmtLeaveReq(req *MgmtLeaveReqFormat) (*MgmtLeaveRspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_MGMT_LEAVE_RSP, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*MgmtLeaveRspFormat), err
//}

////func (self *Client) ZdoMgmtDirectJoinReq(req *MgmtDirectJoinReqFormat) (*MgmtDirectJoinRspFormat, error) {

////}
//func (self *Client) ZdoMgmtPermitJoinReq(req *MgmtPermitJoinReqFormat) error {
//	_, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_MGMT_PERMIT_JOIN_REQ, req)
//	return err
//}

////func (self *Client) ZdoMgmtNwkUpdateReq(req *MgmtNwkUpdateReqFormat) error {

////}
// func (c *Client) zdoStartupFromApp(req *StartupFromAppFormat) (uint8, error) {
// 	resp, err := c.Call(req)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return *(resp.(*uint8)), err
// }

// func (c *Client) zdoWaitNwkStatus(status uint8, timeout time.Duration) error {
// 	s := time.Now()
// 	e := s.Add(timeout)
// 	var err error
// 	var resp interface{}
// 	for time.Now().Before(e) && err == nil {

// 		resp, err = c.WaitAsync(cmd{MT_RPC_SYS_ZDO, MT_ZDO_STATE_CHANGE_IND}, e.Sub(time.Now()))
// 		if err != nil {
// 			return err
// 		}
// 		st := *(resp.(*uint8))
// 		if st == status {
// 			return nil
// 		}
// 	}
// 	return ErrTimeout
// }

////func (self *Client) ZdoAutoFindDestination(AutoFindDestinationFormat_t *req)
//func (self *Client) ZdoSetLinkKey(req *SetLinkKeyFormat) error {
//	_, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_SET_LINK_KEY, req)
//	return err
//}

//func (self *Client) ZdoRemoveLinkKey(req *RemoveLinkKeyFormat) error {
//	_, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_REMOVE_LINK_KEY, req)
//	return err
//}

//func (self *Client) ZdoGetLinkKey(req *GetLinkKeyFormat) (*GetLinkKeySrspFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_GET_LINK_KEY, req)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*GetLinkKeySrspFormat), err
//}

//func (self *Client) ZdoNwkDiscoveryReq(req *NwkDiscoveryReqFormat) (*NwkDiscoveryCnfFormat, error) {
//	resp, err := self.Call(MT_RPC_SYS_ZDO, MT_ZDO_GET_LINK_KEY, req)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*NwkDiscoveryCnfFormat), err
//}

//func (self *Client) ZdoJoinReq(JoinReqFormat_t *req)
//func (self *Client) ZdoMsgCbRegister(MsgCbRegisterFormat_t *req)
//func (self *Client) ZdoMsgCbRemove(MsgCbRemoveFormat_t *req)
