//

package codec

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

// Read similar function as binary.Read to decode struct from binary data,
// but to support struct element with array type and has a length element before this array.
// this is so common in ZNP
func Read(r io.Reader, order binary.ByteOrder, data interface{}) error {
	// Fallback to reflect-based decoding.
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Ptr: // may
		v = v.Elem()
	}
	return read(r, order, v)

}

func read(r io.Reader, order binary.ByteOrder, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			tag := t.Field(i).Tag
			var sliceLen int
			switch tag.Get("lentype") {
			case "uint8":
				var fl uint8
				binary.Read(r, order, &fl)
				sliceLen = int(fl)
			case "uint16":
				var fl uint16
				binary.Read(r, order, &fl)
				sliceLen = int(fl)
			case "uint32":
				var fl uint32
				binary.Read(r, order, &fl)
				sliceLen = int(fl)
			case "": // no tag
			default:
				panic(fmt.Sprintf("invalid lentype '%s'", tag.Get("lentype")))
			}
			if sliceLen > 0 { // has tag
				data := reflect.MakeSlice(v.Field(i).Type(), sliceLen, sliceLen)
				binary.Read(r, order, data.Interface())
				v.Field(i).Set(data)
			} else { // no tag, normal read
				read(r, order, v.Field(i))
			}
		}
		return nil
	default:
		return binary.Read(r, order, v.Addr().Interface())
	}

}
