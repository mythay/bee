package codec

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"reflect"
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

//Decoder decode a byte stream to ZNP struct
type Decoder struct {
	r  io.Reader
	br *bufio.Reader
}

// NewDecoder create a new decoder to decode znp message from io.Reader
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r, br: bufio.NewReader(r)}
}

func (dec *Decoder) decodeHeader() (*Cmd, []byte, error) {
	buf := make([]byte, RPC_HDR_LEN+RPC_MAX_LEN+RPC_UART_FCS_LEN)
	sof, err := dec.br.ReadByte()
	if err != nil {
		return nil, nil, ErrIO
	}

	if sof != MT_RPC_SOF { //minimum package size
		return nil, nil, ErrSOF
	}

	for offset := 0; offset < RPC_HDR_LEN; {
		n, err := dec.br.Read(buf[:RPC_HDR_LEN])
		if err != nil {
			return nil, nil, ErrHDR
		}
		offset += n
	}

	length := buf[0]
	c := &Cmd{buf[1], buf[2]}

	for offset := uint8(RPC_HDR_LEN); offset < RPC_HDR_LEN+length+RPC_UART_FCS_LEN; {
		n, err := dec.br.Read(buf[offset : RPC_HDR_LEN+length+RPC_UART_FCS_LEN])
		if err != nil {
			return c, nil, ErrPayload
		}
		offset += uint8(n)
	}
	fcs := buf[length+RPC_HDR_LEN]
	if calcFcs(buf[:length+RPC_HDR_LEN]) != fcs {
		return nil, nil, ErrFcs
	}
	return c, buf[RPC_HDR_LEN : length+RPC_HDR_LEN], nil
}

func (dec *Decoder) Decode() (*Cmd, interface{}, error) {
	c, payload, err := dec.decodeHeader()
	if err != nil {
		return nil, nil, err
	}

	rsp, ok := c.newObject()
	if !ok {
		return c, payload, nil
	}

	if rsp == nil {
		return c, nil, nil
	}

	err = Read(bytes.NewBuffer(payload), binary.LittleEndian, rsp)
	if err != nil {
		err = errors.New("unable to decode payload")
		return c, nil, err
	}

	return c, rsp, err
}

// Dump to parase the input to human readable text
// func (dec *Decoder) Dump(out io.Writer) error {
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

func calcFcs(msg []byte) uint8 {
	var result uint8
	// calculate FCS by XORing all bytes
	for _, v := range msg {
		result ^= v
	}
	return result
}

func makeRawData(msg []byte) []byte {
	pad := make([]byte, len(msg)+1+1+1)
	pad[0] = MT_RPC_SOF
	pad[1] = byte(len(msg) - 2)
	copy(pad[2:], msg)
	pad[len(pad)-1] = calcFcs(pad[1:])
	return pad
}

func cloneValue(source interface{}, destin interface{}) {
	x := reflect.ValueOf(source)
	if x.Kind() == reflect.Ptr {
		starX := x.Elem()
		y := reflect.New(starX.Type())
		starY := y.Elem()
		starY.Set(starX)
		reflect.ValueOf(destin).Elem().Set(y.Elem())
	} else {
		destin = x.Interface()
	}
}
