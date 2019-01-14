package codec

const (
	/* MT GP Commands */

	MT_GP_CLEAR_ALL         = 0x00
	MT_GP_SET_PARAM         = 0x01
	MT_GP_COMMISS_MODE      = 0x02
	MT_GP_COMMISS_REPLY_REQ = 0x03

	MT_GP_COMMISS_GPDF_IND  = 0x80
	MT_GP_COMMISS_FINAL_IND = 0x81
	MT_GP_EXPIRE_TXQ_IND    = 0x82

	/* GP Parameters */

	MT_GP_PARAM_ID_PRE_COM_GROUP_ID    = 0x00
	MT_GP_PARAM_ID_LOOPBACK_DST_EP     = 0x01
	MT_GP_PARAM_ID_TX_Q_ENTRY_LIFETIME = 0x02
	MT_GP_PARAM_ID_TX_Q_EXPIRY_IND     = 0x03

	/* GP Commissioning Mode */

	MT_GP_COMMISSIONING_MODE_EXIT        = 0x00
	MT_GP_COMMISSIONING_MODE_NORMAL      = 0x01
	MT_GP_COMMISSIONING_MODE_RESTRICTIVE = 0x02
)

type GpdCommissioningNtfParams struct {
	Address         uint64 /**< Address of the GPD */
	DeviceId        uint8  /**< GP Device Id. If 0xFE then ManufacturerId & ManufacturerModelId are valid*/
	Options         uint8  /**< Options related to the command/device */
	ExtendedOptions uint8  /**< Ignore. For future expansion */
}

type GpComReplyReqParams struct {
	Address uint64 /**< Address of the GPD */
	Options uint8  /**<  */
}

type GpCmdReqParams struct {
	SourceEndpoint uint8   /**< The Endpoint for the source (App) */
	Address        uint64  /**< The Addressing mode to specify how the device communicates with the destination. Must be a GPD address. */
	GpCommand      uint8   /**< The GP Command Id */
	Payload        []uint8 `lentype:"uint8" ` /**< The ZGP command formatted payload in LITTLE ENDIAN - (extends beyond bounds) */
}

type GpdCommissionedNtfParams struct {
	Event    uint8  /**< Commissioned Event Type */
	Address  uint64 /**< Address of the GPD */
	DeviceId uint8  /**< Device ID of the GPD */
}

type GpSetParamFormat struct {
	ParamId uint8
	Value   uint16
}

type GpSetCommissioningMode struct {
	Mode       uint8
	Externsion uint8
	ModeExit   uint8
	Window     uint16
}
type GpSetParamSrspFormat struct {
	Status uint8
}

// var ErrGpConfirm = errors.New("GP request error")

// func (gp *GpSetParamFormat) marshall() []byte {
// 	buf := &bytes.Buffer{}
// 	binary.Write(buf, binary.LittleEndian, gp.ParamId)
// 	switch gp.ParamId {
// 	case MT_GP_PARAM_ID_PRE_COM_GROUP_ID:
// 		fallthrough
// 	case MT_GP_PARAM_ID_LOOPBACK_DST_EP:
// 		binary.Write(buf, binary.LittleEndian, uint8(gp.Value))
// 	case MT_GP_PARAM_ID_TX_Q_ENTRY_LIFETIME:
// 		fallthrough
// 	case MT_GP_PARAM_ID_TX_Q_EXPIRY_IND:
// 		binary.Write(buf, binary.LittleEndian, gp.Value)
// 	}

// 	return buf.Bytes()
// }

// func (gp *GpSetParamFormat) unmarshall(p []byte) error {
// 	if len(p) >= 2 {
// 		gp.ParamId = p[0]
// 		gp.Value = uint16(p[1])
// 		return nil
// 	}
// 	return errors.New("unmarshall error")
// }

// func (c *Client) gpSetParam(req *GpSetParamFormat) (uint8, error) {
// 	resp, err := c.Call(req)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return *(resp.(*uint8)), err
// }

// func (c *Client) gpSetLoopback() error {
// 	status, err := c.gpSetParam(&GpSetParamFormat{MT_GP_PARAM_ID_LOOPBACK_DST_EP, 1})
// 	if err == nil && status == 0 {
// 		return nil
// 	}
// 	return ErrGpConfirm
// }
