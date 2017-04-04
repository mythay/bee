package codec

import (
	"errors"
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

// NewDecoder create a new decoder to decode znp message from io.Reader
// func NewDecoder(r io.Reader) *decoder {
// 	return &decoder{r: r, br: bufio.NewReader(r)}
// }

// func (dec *decoder) decodeHeader() (*header, []byte, error) {
// 	buf := make([]byte, RPC_HDR_LEN+RPC_MAX_LEN+RPC_UART_FCS_LEN)
// 	sof, err := dec.br.ReadByte()
// 	if err != nil {
// 		return nil, nil, ErrIO
// 	}

// 	if sof != MT_RPC_SOF { //minimum package size
// 		return nil, nil, ErrSOF
// 	}

// 	for offset := 0; offset < RPC_HDR_LEN; {
// 		n, err := dec.br.Read(buf[:RPC_HDR_LEN])
// 		if err != nil {
// 			return nil, nil, ErrHDR
// 		}
// 		offset += n
// 	}
// 	hdr := &header{}
// 	hdr.length = buf[0]
// 	hdr.cmd0 = buf[1]
// 	hdr.cmd1 = buf[2]

// 	for offset := uint8(RPC_HDR_LEN); offset < RPC_HDR_LEN+hdr.length+RPC_UART_FCS_LEN; {
// 		n, err := dec.br.Read(buf[offset : RPC_HDR_LEN+hdr.length+RPC_UART_FCS_LEN])
// 		if err != nil {
// 			return hdr, nil, ErrPayload
// 		}
// 		offset += uint8(n)
// 	}
// 	fcs := buf[hdr.length+RPC_HDR_LEN]
// 	if calcFcs(buf[:hdr.length+RPC_HDR_LEN]) != fcs {
// 		return nil, nil, ErrFcs
// 	}
// 	return hdr, buf[RPC_HDR_LEN : hdr.length+RPC_HDR_LEN], nil
// }

// func (dec *decoder) decode() (*header, interface{}, error) {
// 	hdr, payload, err := dec.decodeHeader()
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	rsp, ok := hdr.newObject()
// 	if !ok {
// 		return hdr, payload, nil
// 	}

// 	if rsp == nil {
// 		return hdr, nil, nil
// 	}

// 	obj, ok := rsp.(unmarshaller)
// 	if ok {
// 		err = obj.unmarshall(payload)
// 	} else {
// err = binary.Read(bytes.NewBuffer(payload), binary.LittleEndian, rsp)
// 	}
// 	if err != nil {
// 		err = errors.New("unable to decode payload")
// 	}

// 	return hdr, rsp, err
// }

// // Dump to parase the input to human readable text
// func (dec *decoder) Dump(out io.Writer) error {
// 	buf := &bytes.Buffer{}
// 	hdr, resp, err := dec.decode()
// 	if err == nil {
// 		if resp == nil {
// 			buf.WriteString(fmt.Sprintf("'%s' : %02X-%02X \r\n", hdr.getDesc(), hdr.cmd0, hdr.cmd1))
// 		} else {
// 			buf.WriteString(fmt.Sprintf("'%s' : %02X-%02X | %v\r\n", hdr.getDesc(), hdr.cmd0, hdr.cmd1, reflect.ValueOf(resp).Elem()))
// 		}

// 	} else if err != nil && err != ErrIO {
// 		if hdr == nil {
// 			buf.WriteString(fmt.Sprintf("ERROR:  %v\r\n", err))
// 		} else {
// 			buf.WriteString(fmt.Sprintf("ERROR: %02X-%02X | %v\r\n", hdr.cmd0, hdr.cmd1, err))
// 		}

// 	}
// 	buf.WriteTo(out)
// 	return err
// }

// func calcFcs(msg []byte) uint8 {
// 	var result uint8
// 	// calculate FCS by XORing all bytes
// 	for _, v := range msg {
// 		result ^= v
// 	}
// 	return result
// }

// func cloneValue(source interface{}, destin interface{}) {
// 	x := reflect.ValueOf(source)
// 	if x.Kind() == reflect.Ptr {
// 		starX := x.Elem()
// 		y := reflect.New(starX.Type())
// 		starY := y.Elem()
// 		starY.Set(starX)
// 		reflect.ValueOf(destin).Elem().Set(y.Elem())
// 	} else {
// 		destin = x.Interface()
// 	}
// }
