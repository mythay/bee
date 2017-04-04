package codec

import (
	"bytes"
	//	"/*reflect*/"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncode(t *testing.T) {
	Convey("ZNP encode and decode", t, func() {
		Convey("only cmds", func() {
			buf := &bytes.Buffer{}
			enc := newEncoder(buf)
			enc.encodeRaw(cmd{}, nil)
			r := buf.Bytes()
			So(r[0], ShouldEqual, 254)
			So(buf.Len(), ShouldEqual, 5)

			dec := NewDecoder(buf)
			h, ebuf, err := dec.decodeHeader()
			So(err, ShouldBeNil)
			So(len(ebuf), ShouldEqual, 0)

			So(h.cmd0, ShouldEqual, 0)
			So(h.length, ShouldEqual, 0)
		})

		Convey("simply type", func() {
			buf := &bytes.Buffer{}
			enc := newEncoder(buf)
			err := enc.encodeRaw(cmd{}, uint8(2))
			So(err, ShouldBeNil)

			r := buf.Bytes()
			So(r[0], ShouldEqual, 254)
			So(r[buf.Len()-1], ShouldEqual, calcFcs(r[1:len(r)-1]))
			So(buf.Len(), ShouldEqual, 6)

			dec := NewDecoder(buf)
			h, ebuf, err := dec.decodeHeader()
			So(err, ShouldBeNil)
			So(len(ebuf), ShouldEqual, 1)

			So(h.cmd0, ShouldEqual, 0)
			So(h.length, ShouldEqual, 1)
		})

		Convey(" enable to decode two continuous bytes sequence", func() {
			buf := &bytes.Buffer{}
			enc := newEncoder(buf)
			enc.encodeRaw(cmd{}, uint8(2))
			enc.encodeRaw(cmd{1, 1}, nil)

			dec := NewDecoder(buf)
			_, _, err := dec.decodeHeader()
			So(err, ShouldBeNil)
			_, _, err = dec.decodeHeader()
			So(err, ShouldBeNil)

		})

		Convey("decode but no SOF", func() {
			buf := bytes.NewBuffer(make([]byte, 0))

			dec := NewDecoder(buf)
			_, _, err := dec.decodeHeader()
			So(err, ShouldNotBeNil)

		})

		Convey("check invalid data", func() {
			buf := &bytes.Buffer{}
			enc := newEncoder(buf)
			err := enc.encodeRaw(cmd{}, 2)
			So(err, ShouldNotBeNil)
		})

		Convey("data is struct", func() {
			type X struct {
				X uint8
				Y uint8
				Z uint32
			}
			buf := &bytes.Buffer{}
			enc := newEncoder(buf)
			err := enc.encodeRaw(cmd{}, &X{1, 2, 3})
			So(err, ShouldBeNil)

			So(buf.Len(), ShouldEqual, 11)

			dec := NewDecoder(buf)
			h, _, err := dec.decodeHeader()
			So(err, ShouldBeNil)

			So(h.cmd0, ShouldEqual, 0)
			So(h.length, ShouldEqual, 6)
		})

		Convey("data is struct, but has hidden field", func() {
			type XX struct {
				X uint8
				Y uint8
				z uint32
			}
			buf := &bytes.Buffer{}
			enc := newEncoder(buf)
			err := enc.encodeRaw(cmd{}, &XX{1, 2, 3})
			So(err, ShouldBeNil)

			So(buf.Len(), ShouldEqual, 11)

			dec := NewDecoder(buf)
			h, _, err := dec.decodeHeader()
			So(err, ShouldBeNil)

			So(h.cmd0, ShouldEqual, 0)
			So(h.length, ShouldEqual, 6)
		})
	})

}

func TestDecode(t *testing.T) {
	Convey("ZNP decode", t, func() {

		Convey("wrong EOF header", func() {
			stream := &bytes.Buffer{}

			dec := NewDecoder(stream)
			hdr, payload, err := dec.decode()
			So(err, ShouldEqual, ErrIO)
			So(hdr, ShouldBeNil)
			So(payload, ShouldBeNil)
		})
		Convey("wrong SOF header, need to consume at least one byte to avoid dead loop", func() {
			stream := &bytes.Buffer{}
			stream.WriteByte(0xee)
			dec := NewDecoder(stream)
			hdr, payload, err := dec.decode()
			So(err.Error(), ShouldContainSubstring, "SOF")
			So(hdr, ShouldBeNil)
			So(payload, ShouldBeNil)

			hdr, payload, err = dec.decode()
			So(err, ShouldEqual, ErrIO)
			So(hdr, ShouldBeNil)
			So(payload, ShouldBeNil)
		})
		Convey("wrong FCS ", func() {
			stream := &bytes.Buffer{}
			p := []byte{0xfe, 00, 00, 00, 11, 00}
			stream.Write(p)
			dec := NewDecoder(stream)
			hdr, payload, err := dec.decode()
			So(err.Error(), ShouldContainSubstring, "FCS")
			So(hdr, ShouldBeNil)
			So(payload, ShouldBeNil)
		})

		Convey("too short, no HDR ", func() {
			stream := &bytes.Buffer{}
			p := []byte{0xfe, 00}
			stream.Write(p)
			dec := NewDecoder(stream)
			hdr, payload, err := dec.decode()
			So(err.Error(), ShouldContainSubstring, "HDR")
			So(hdr, ShouldBeNil)
			So(payload, ShouldBeNil)
		})

		Convey(" HDR length not match", func() {
			stream := &bytes.Buffer{}
			p := []byte{0xfe, 01, 1, 1}
			stream.Write(p)
			// n, err := stream.Read(x)
			// fmt.Println(n, err)
			dec := NewDecoder(stream)
			hdr, payload, err := dec.decode()
			So(err.Error(), ShouldContainSubstring, "payload")
			So(hdr, ShouldBeNil)
			So(payload, ShouldBeNil)
		})

		Convey("decode dump request", func() {
			stream := &bytes.Buffer{}
			out := &bytes.Buffer{}
			enc := newEncoder(stream)
			enc.encodeRaw(cmd{(MT_RPC_CMD_SREQ | MT_RPC_SYS_SYS), MT_SYS_PING}, nil)
			dec := NewDecoder(stream)
			dec.Dump(out)
			So(out.String(), ShouldContainSubstring, "SYS_PING")
		})

		Convey("decode dump response", func() {
			stream := &bytes.Buffer{}
			out := &bytes.Buffer{}
			enc := newEncoder(stream)
			enc.encodeRaw(cmd{(MT_RPC_CMD_SRSP | MT_RPC_SYS_SYS), MT_SYS_GET_EXTADDR}, &GetExtAddrSrspFormat{121})
			dec := NewDecoder(stream)
			dec.Dump(out)
			So(out.String(), ShouldContainSubstring, "SYS_GET_EXTADDR")
		})
	})
}

func TestMarshall(t *testing.T) {
	Convey("marshall of ZNP object", t, func() {
		Convey("AF ", func() {
			Convey("RegisterFormat", func() {
				reg := &RegisterFormat{
					EndPoint:          1,
					AppProfId:         0x0104,
					AppDeviceId:       0x0100,
					AppDevVer:         1,
					LatencyReq:        0,
					AppNumInClusters:  1,
					AppNumOutClusters: 0,
				}
				reg.AppInClusterList[0] = 0x0006

				buf := reg.marshall()
				So(len(buf), ShouldEqual, 11)
				So(buf[8], ShouldEqual, 6)
				So(buf[9], ShouldEqual, 0)
			})
		})
	})

	Convey("Unmarshall of ZNP object", t, func() {
		Convey("use sys ping response", func() {
			buf := &bytes.Buffer{}
			enc := newEncoder(buf)
			err := enc.encodeRaw(cmd{MT_RPC_CMD_SRSP | MT_RPC_SYS_SYS, MT_SYS_PING}, &PingSrspFormat{121})
			So(err, ShouldBeNil)
			dec := NewDecoder(buf)
			hdr, rsp, err := dec.decode()
			So(err, ShouldBeNil)
			So(rsp, ShouldNotBeNil)
			r, ok := rsp.(*PingSrspFormat)
			So(ok, ShouldBeTrue)
			So(r.Capabilities, ShouldEqual, 121)
			So(hdr.cmd1, ShouldEqual, MT_SYS_PING)
		})
		Convey("invalid response, so no error, the response msessage should be  raw bytes", func() {
			buf := &bytes.Buffer{}
			enc := newEncoder(buf)
			err := enc.encodeRaw(cmd{MT_RPC_CMD_SRSP | MT_RPC_SYS_SYS, 0xff}, nil)
			So(err, ShouldBeNil)
			dec := NewDecoder(buf)
			_, rsp, err := dec.decode()
			So(err, ShouldBeNil)
			So(rsp, ShouldHaveSameTypeAs, []byte{})

		})

		Convey("sys nv read without response", func() {
			buf := &bytes.Buffer{}
			enc := newEncoder(buf)
			err := enc.encodeRaw(cmd{MT_RPC_CMD_SRSP | MT_RPC_SYS_SYS, MT_SYS_SET_EXTADDR}, nil)
			So(err, ShouldBeNil)
			dec := NewDecoder(buf)
			_, rsp, err := dec.decode()
			So(err, ShouldBeNil)
			So(rsp, ShouldBeNil)

		})
	})
}
