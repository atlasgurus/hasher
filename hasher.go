package hasher

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"reflect"
)

type Hash []byte

func (h Hash) ToUint64() uint64 {
	return binary.LittleEndian.Uint64(h[0:8])
}

func ComputeHash(s interface{}) Hash {
	return hashValue(reflect.ValueOf(s))
}

func hashValue(v reflect.Value) Hash {
	hasher := sha256.New()
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if !v.IsNil() {
			v = v.Elem()
		} else {
			return nil
		}
	}

	// Check if the struct contains a field of type Hash
	if v.Kind() == reflect.Struct {
		if hash, found := getHashFromStruct(v); found {
			return hash
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
			if t.Field(i).Type == reflect.TypeOf(Hash(nil)) {
				continue // Skip fields of Hash type
			}
			if field.CanInterface() {
				hasher.Write(hashValue(field))
			}
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			hasher.Write(hashValue(v.Index(i)))
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
		hasher.Write(Hash(v.String()))
	case reflect.Bool:
		if v.Bool() {
			hasher.Write(Hash{1})
		} else {
			hasher.Write(Hash{0})
		}
	// Add cases for other types as needed
	default:
		// Fallback for types not explicitly handled
		hasher.Write(Hash(fmt.Sprintf("%v", v.Interface())))
	}
	return hasher.Sum(nil)
}

func getHashFromStruct(v reflect.Value) ([]byte, bool) {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Type == reflect.TypeOf(Hash(nil)) {
			if hash, ok := v.Field(i).Interface().(Hash); ok && len(hash) > 0 {
				return hash, true
			}
		}
	}
	return nil, false
}
