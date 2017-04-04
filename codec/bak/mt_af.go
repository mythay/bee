package codec

import (
	"bytes"
	"encoding/binary"
)

const (
	afAddrNotPresent = 0
	afAddr16Bit      = 1
	afAddr64Bit      = 2
	afAddrGroup      = 3
	afAddrBroadcast  = 15
)

/***************************************************************************************************
 * AF COMMANDS
 ***************************************************************************************************/
const (
	/* SREQ/SRSP */
	MT_AF_REGISTER             = 0x00
	MT_AF_DATA_REQUEST         = 0x01 /* AREQ optional, but no AREQ response. */
	MT_AF_DATA_REQUEST_EXT     = 0x02 /* AREQ optional, but no AREQ response. */
	MT_AF_DATA_REQUEST_SRC_RTG = 0x03

	MT_AF_INTER_PAN_CTL   = 0x10
	MT_AF_DATA_STORE      = 0x11
	MT_AF_DATA_RETRIEVE   = 0x12
	MT_AF_APSF_CONFIG_SET = 0x13

	/* AREQ to host */
	MT_AF_DATA_CONFIRM     = 0x80
	MT_AF_INCOMING_MSG     = 0x81
	MT_AF_INCOMING_MSG_EXT = 0x82
	MT_AF_REFLECT_ERROR    = 0x83

	afStatus_SUCCESS           = 0x00
	afStatus_FAILED            = 0x01
	afStatus_INVALID_PARAMETER = 0x02
	afStatus_MEM_FAIL          = 0x10
	afStatus_NO_ROUTE          = 0xCD
	afStatus_DUPLICATE         = 0xB8
)

type RegisterFormat struct {
	EndPoint          uint8
	AppProfId         uint16
	AppDeviceId       uint16
	AppDevVer         uint8
	LatencyReq        uint8
	AppNumInClusters  uint8
	AppInClusterList  [16]uint16
	AppNumOutClusters uint8
	AppOutClusterList [16]uint16
}

func (register *RegisterFormat) marshall() []byte {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, register.EndPoint)
	binary.Write(buf, binary.LittleEndian, register.AppProfId)
	binary.Write(buf, binary.LittleEndian, register.AppDeviceId)
	binary.Write(buf, binary.LittleEndian, register.AppDevVer)
	binary.Write(buf, binary.LittleEndian, register.LatencyReq)
	binary.Write(buf, binary.LittleEndian, register.AppNumInClusters)
	for i := uint8(0); i < register.AppNumInClusters; i++ {
		binary.Write(buf, binary.LittleEndian, register.AppInClusterList[i])
	}

	binary.Write(buf, binary.LittleEndian, register.AppNumOutClusters)
	for i := uint8(0); i < register.AppNumOutClusters; i++ {
		binary.Write(buf, binary.LittleEndian, register.AppOutClusterList[i])
	}

	return buf.Bytes()
}

type DataRequestFormat struct {
	DstAddr     uint16
	DstEndpoint uint8
	SrcEndpoint uint8
	ClusterID   uint16
	TransID     uint8
	Options     uint8
	Radius      uint8
	Len         uint8
	Data        [128]uint8
}

type DataRequestExtFormat struct {
	DstAddrMode uint8
	DstAddr     [8]uint8
	DstEndpoint uint8
	DstPanID    uint16
	SrcEndpoint uint8
	ClusterId   uint16
	TransId     uint8
	Options     uint8
	Radius      uint8
	Len         uint16
	Data        [230]uint8
}

type DataRequestSrcRtgFormat struct {
	DstAddr     uint16
	DstEndpoint uint8
	SrcEndpoint uint8
	ClusterID   uint16
	TransID     uint8
	Options     uint8
	Radius      uint8
	RelayCount  uint8
	RelayList   [255]uint16
	Len         uint8
	Data        [128]uint8
}

type InterPanCtlFormat struct {
	Command uint8
	Data    [3]uint8
}

type DataStoreFormat struct {
	Index  uint16
	Length uint8
	Data   [247]uint8
}

type DataConfirmFormat struct {
	Status   uint8
	Endpoint uint8
	TransId  uint8
}

type IncomingMsgFormat struct {
	GroupId      uint16
	ClusterId    uint16
	SrcAddr      uint16
	SrcEndpoint  uint8
	DstEndpoint  uint8
	WasVroadcast uint8
	LinkQuality  uint8
	SecurityUse  uint8
	TimeStamp    uint32
	TransSeqNum  uint8
	Len          uint8
	Data         [99]uint8
}

type IncomingMsgExtFormat struct {
	GroupId      uint16
	ClusterId    uint16
	SrcAddrMode  uint8
	SrcAddr      uint64
	SrcEndpoint  uint8
	SrcPanId     uint16
	DstEndpoint  uint8
	WasVroadcast uint8
	LinkQuality  uint8
	SecurityUse  uint8
	TimeStamp    uint32
	TransSeqNum  uint8
	Len          uint8
	Data         [99]uint8
}

type DataRetrieveFormat struct {
	TimeStamp [4]uint8
	Index     uint16
	Length    uint8
}

type DataRetrieveSrspFormat struct {
	Status uint8
	Length uint8
	Data   [248]uint8
}

type ApsfConfigSetFormat struct {
	Endpoint   uint8
	FrameDelay uint8
	WindowSize uint8
}

type ReflectErrorFormat struct {
	Status      uint8
	Endpoint    uint8
	TransId     uint8
	DstAddrMode uint8
	DstAddr     uint16
}

func (self *Client) afRegister(req *RegisterFormat) error {
	_, err := self.Call(req)
	return err
}

//func (self *Client) DataRequest(req *DataRequestFormat) datare
//func (self *Client) DataRequestExt(DataRequestExtFormat_t *req)
//func (self *Client) DataRequestSrcRtg(DataRequestSrcRtgFormat_t *req)
//func (self *Client) InterPanCtl(InterPanCtlFormat_t *req)
//func (self *Client) DataStore(DataStoreFormat_t *req)
//func (self *Client) DataRetrieve(DataRetrieveFormat_t *req)
//func (self *Client) ApsfConfigSet(ApsfConfigSetFormat_t *req)
