package modbus

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// ConvertBytesToUint16 ..
func ConvertBytesToUint16(rawResults []byte, byteOrder binary.ByteOrder) (results []uint16, err error) {
	var data uint16

	results = []uint16{}
	err = nil

	count := int(rawResults[0])
	rawResults = rawResults[1:]

	if count != len(rawResults) {
		return nil, fmt.Errorf("ConvertBytesToUint16(): data size mismatch")
	}

	for i := 0; i < count; i += 2 {
		//data, decodeErr := DecodeHiLoUInt16(order, rawResults[sliceStart:sliceStop + 1])
		buf := bytes.NewReader(rawResults[i : i+2])
		err = binary.Read(buf, byteOrder, &data)
		if err != nil {
			return nil, fmt.Errorf("ConvertBytesToUint16(): binary.Read() failed")
		}

		results = append(results, uint16(data))
	}

	return
}

// ConvertBytesToUint32 ..
func ConvertBytesToUint32(rawResults []byte, byteOrder binary.ByteOrder) (results []uint32, err error) {
	var data uint32

	results = []uint32{}
	err = nil

	count := int(rawResults[0])
	rawResults = rawResults[1:]

	if count != len(rawResults) {
		return nil, fmt.Errorf("ConvertBytesToUint32(): data size mismatch")
	}

	for i := 0; i < count; i += 4 {
		//data, decodeErr := DecodeHiLoUInt16(order, rawResults[sliceStart:sliceStop + 1])
		buf := bytes.NewReader(rawResults[i : i+4])
		err = binary.Read(buf, byteOrder, &data)
		if err != nil {
			return nil, fmt.Errorf("ConvertBytesToUint32(): binary.Read() failed")
		}

		results = append(results, uint32(data))
	}

	return
}

// ConvertBytesToInt16 ..
func ConvertBytesToInt16(rawResults []byte, byteOrder binary.ByteOrder) (results []int16, err error) {
	var data int16

	results = []int16{}
	err = nil

	count := int(rawResults[0])
	rawResults = rawResults[1:]

	if count != len(rawResults) {
		return nil, fmt.Errorf("ConvertBytesToInt16(): data size mismatch")
	}

	for i := 0; i < count; i += 2 {
		//data, decodeErr := DecodeHiLoUInt16(order, rawResults[sliceStart:sliceStop + 1])
		buf := bytes.NewReader(rawResults[i : i+2])
		err = binary.Read(buf, byteOrder, &data)
		if err != nil {
			return nil, fmt.Errorf("ConvertBytesToInt16(): binary.Read() failed")
		}

		results = append(results, int16(data))
	}

	return
}

// ConvertBytesToInt32 ..
func ConvertBytesToInt32(rawResults []byte, byteOrder binary.ByteOrder) (results []int32, err error) {
	var data int32

	results = []int32{}
	err = nil

	count := int(rawResults[0])
	rawResults = rawResults[1:]

	if count != len(rawResults) {
		return nil, fmt.Errorf("ConvertBytesToInt32(): data size mismatch")
	}

	for i := 0; i < count; i += 4 {
		//data, decodeErr := DecodeHiLoUInt32(order, rawResults[sliceStart:sliceStop + 1])
		buf := bytes.NewReader(rawResults[i : i+4])
		err = binary.Read(buf, byteOrder, &data)
		if err != nil {
			return nil, fmt.Errorf("ConvertBytesToInt32(): binary.Read() failed")
		}

		results = append(results, int32(data))
	}

	return
}

// ConvertBytesToFloat32 ..
func ConvertBytesToFloat32(rawResults []byte, byteOrder binary.ByteOrder) (results []float32, err error) {
	var data float32

	results = []float32{}
	err = nil

	count := int(rawResults[0])
	rawResults = rawResults[1:]

	if count != len(rawResults) {
		return nil, fmt.Errorf("ConvertBytesToFloat32(): data size mismatch")
	}

	for i := 0; i < count; i += 4 {
		//data, decodeErr := DecodeHiLoUFloat32(order, rawResults[sliceStart:sliceStop + 1])
		buf := bytes.NewReader(rawResults[i : i+4])
		err = binary.Read(buf, byteOrder, &data)
		if err != nil {
			return nil, fmt.Errorf("ConvertBytesToFloat32(): binary.Read() failed")
		}

		results = append(results, float32(data))
	}

	return
}
