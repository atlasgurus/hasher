package hasher

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"hash"
	"reflect"
)

// hashStruct generates a hash for the given struct.
func ComputeHash(s interface{}) []byte {
	hasher := sha256.New()
	hashValue(reflect.ValueOf(s), hasher)
	return hasher.Sum(nil)
}

func hashValue(v reflect.Value, hasher hash.Hash) {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if !v.IsNil() {
			v = v.Elem()
		} else {
			return
		}
	}

	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if t.Field(i).Tag.Get("hash") == "-" {
				continue // Skip fields with the tag hash:"-"
			}
			if field.CanInterface() {
				hashValue(field, hasher)
			}
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			hashValue(v.Index(i), hasher)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var buf [8]byte
		binary.LittleEndian.PutUint64(buf[:], uint64(v.Int()))
		hasher.Write(buf[:])
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		var buf [8]byte
		binary.LittleEndian.PutUint64(buf[:], v.Uint())
		hasher.Write(buf[:])
	case reflect.String:
		hasher.Write([]byte(v.String()))
	case reflect.Bool:
		if v.Bool() {
			hasher.Write([]byte{1})
		} else {
			hasher.Write([]byte{0})
		}
	// Add cases for other types as needed
	default:
		// Fallback for types not explicitly handled
		hasher.Write([]byte(fmt.Sprintf("%v", v.Interface())))
	}
}
