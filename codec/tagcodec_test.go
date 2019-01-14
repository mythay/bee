//

package codec

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
)

type s1 struct {
	A uint8
	B uint8
}

type s2 struct {
	A uint16
	B uint16
}

type ls1 struct {
	A []uint8 `array:"yes" lentype:"uint8" size:"248"`
}

func TestRead(t *testing.T) {

	t.Run("simple struct", func(t *testing.T) {
		d := &s1{}
		err := Read(bytes.NewReader([]byte{2, 2}), binary.BigEndian, d)
		if err != nil {
			t.Errorf("should not error, but get %v", err)
		}
		if !reflect.DeepEqual(d, &s1{2, 2}) {
			t.Errorf("got %v", d)
		}
	})

	t.Run("uint16 struct big endian", func(t *testing.T) {
		d := &s2{}
		err := Read(bytes.NewReader([]byte{1, 2, 3, 4}), binary.BigEndian, d)
		if err != nil {
			t.Errorf("should not error, but get %v", err)
		}
		if !reflect.DeepEqual(d, &s2{0x0102, 0x0304}) {
			t.Errorf("got %v", d)
		}

	})

	t.Run("uint16 struct little endian", func(t *testing.T) {
		d := &s2{}
		err := Read(bytes.NewReader([]byte{1, 2, 3, 4}), binary.LittleEndian, d)
		if err != nil {
			t.Errorf("should not error, but get %v", err)
		}
		if !reflect.DeepEqual(d, &s2{0x0201, 0x0403}) {
			t.Errorf("got %v", d)
		}

	})
	t.Run("buffer too small", func(t *testing.T) {
		d := &s2{}
		err := Read(bytes.NewReader([]byte{1, 2, 3}), binary.LittleEndian, d)
		if err == nil {
			t.Errorf("should got error, but get %v", err)
		}

	})

	t.Run("struct has length tag", func(t *testing.T) {
		d := &ls1{}
		err := Read(bytes.NewReader([]byte{3, 1, 2, 3, 4}), binary.BigEndian, d)
		if err != nil {
			t.Errorf("should not error, but get %v", err)
		}
		if !reflect.DeepEqual(d, &ls1{[]uint8{1, 2, 3}}) {
			t.Errorf("got %v", d)
		}
	})

}

func TestWrite(t *testing.T) {
	RegisterTestingT(t)
	t.Run("simple struct", func(t *testing.T) {
		b := &bytes.Buffer{}
		err := Write(b, binary.BigEndian, &s1{1, 2})
		Expect(err).Should(BeNil(), "must be no error")
		Expect(b.Bytes()).Should(Equal([]byte{1, 2}), " buffer should equal")

	})

	t.Run("uint16 struct big endian", func(t *testing.T) {
		b := &bytes.Buffer{}
		err := Write(b, binary.BigEndian, &s2{0x0102, 0x0304})
		Expect(err).Should(BeNil(), "must be no error")
		Expect(b.Bytes()).Should(Equal([]byte{1, 2, 3, 4}), " buffer should equal")

	})

	t.Run("uint16 struct little endian", func(t *testing.T) {
		b := &bytes.Buffer{}
		err := Write(b, binary.LittleEndian, &s2{0x0201, 0x0403})
		Expect(err).Should(BeNil(), "must be no error")
		Expect(b.Bytes()).Should(Equal([]byte{1, 2, 3, 4}), " buffer should equal")

	})

	t.Run("struct has length tag", func(t *testing.T) {

		b := &bytes.Buffer{}
		err := Write(b, binary.LittleEndian, &ls1{[]uint8{1, 2, 3}})
		Expect(err).Should(BeNil(), "must be no error")
		Expect(b.Bytes()).Should(Equal([]byte{3, 1, 2, 3}), " buffer should equal")

	})

	t.Run("ResetReqFormat", func(t *testing.T) {

		b := &bytes.Buffer{}
		err := Write(b, binary.LittleEndian, &ResetReqFormat{1})
		Expect(err).Should(BeNil(), "must be no error")
		Expect(b.Bytes()).Should(Equal([]byte{1}), " buffer should equal")

	})

}
