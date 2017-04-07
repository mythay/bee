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
			var err error
			tag := t.Field(i).Tag
			var sliceLen int
			switch tag.Get("lentype") {
			case "uint8":
				var fl uint8
				err = binary.Read(r, order, &fl)
				sliceLen = int(fl)
			case "uint16":
				var fl uint16
				err = binary.Read(r, order, &fl)
				sliceLen = int(fl)
			case "uint32":
				var fl uint32
				err = binary.Read(r, order, &fl)
				sliceLen = int(fl)
			case "": // no tag
			default:
				panic(fmt.Sprintf("invalid lentype '%s'", tag.Get("lentype")))
			}
			if err != nil {
				return err
			}
			if sliceLen > 0 { // has tag
				if v.Field(i).Kind() != reflect.Slice {
					panic(fmt.Sprintf(" %v for field has tag'lentype' need to be slice ", v))
				}
				data := reflect.MakeSlice(v.Field(i).Type(), sliceLen, sliceLen)
				err = binary.Read(r, order, data.Interface())
				v.Field(i).Set(data)
			} else { // no tag, normal read
				err = read(r, order, v.Field(i))
			}
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return binary.Read(r, order, v.Addr().Interface())
	}

}

// Write similar function as binary.Write to decode struct from binary data,
// but to support struct element with array type and has a length element before this array.
// this is so common in ZNP
func Write(w io.Writer, order binary.ByteOrder, data interface{}) error {
	// Fallback to reflect-based decoding.
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Ptr: // may
		v = v.Elem()
	}
	return write(w, order, v)

}

func write(w io.Writer, order binary.ByteOrder, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			var err error
			tag := t.Field(i).Tag
			switch tag.Get("lentype") { //  write lenght filed first
			case "uint8":
				fl := uint8(v.Field(i).Len())
				err = binary.Write(w, order, &fl)
			case "uint16":
				fl := uint16(v.Field(i).Len())
				err = binary.Write(w, order, &fl)
			case "uint32":
				fl := uint32(v.Field(i).Len())
				err = binary.Write(w, order, &fl)
			case "": // no tag
			default:
				panic(fmt.Sprintf("invalid lentype '%s'", tag.Get("lentype")))
			}
			if err != nil {
				return err
			}
			if tag != "" { // it is a slice
				err = binary.Write(w, order, v.Field(i).Interface())
			} else {
				err = write(w, order, v.Field(i))
			}

			if err != nil {
				return err
			}
		}
		return nil
	default:
		return binary.Write(w, order, v.Addr().Interface())
	}

}
