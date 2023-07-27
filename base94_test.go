// base94_test.go

package base94

import (
	"encoding/base64"
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	// Test data
	data := []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}

	// Encode the data
	encodedData := Encode(data)

	// Ensure encoding is reversible (decode and compare with original data)
	decodedData := Decode(encodedData)
	if !reflect.DeepEqual(data, decodedData) {
		t.Errorf("Decode(Encode(data)) = %v, want %v", decodedData, data)
	}
}

func TestEncodeDecodeEmpty(t *testing.T) {
	// Test case with an empty binary data
	emptyData := []byte{}
	encodedEmpty := Encode(emptyData)
	decodedEmpty := Decode(encodedEmpty)
	if len(encodedEmpty) != 0 || len(decodedEmpty) != 0 {
		t.Errorf("Encode/Decode for empty data failed")
	}

	// Test case with a single byte of data
	singleByteData := []byte{255} // 11111111 in binary
	encodedSingleByte := Encode(singleByteData)
	decodedSingleByte := Decode(encodedSingleByte)
	if !reflect.DeepEqual(singleByteData, decodedSingleByte) {
		t.Errorf("Encode/Decode for single byte data failed")
	}

	// Test case with maximum value of binary data (all bytes are 255)
	maxData := make([]byte, 100) // 100 bytes with value 255 (11111111 in binary)
	for i := range maxData {
		maxData[i] = 255
	}
	encodedMaxData := Encode(maxData)
	decodedMaxData := Decode(encodedMaxData)
	if !reflect.DeepEqual(maxData, decodedMaxData) {
		t.Errorf("Encode/Decode for maximum data value failed")
	}
}

func BenchmarkEncode(b *testing.B) {
	// Generate a large binary data of 1 MB size
	dataSize := 1024 * 1024
	data := make([]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(data)
	}
}

func BenchmarkDecode(b *testing.B) {
	// Generate a large binary data of 1 MB size
	dataSize := 1024 * 1024
	data := make([]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = byte(i % 256)
	}

	// Encode the data to get the encoded string
	encodedData := Encode(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Decode(encodedData)
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	// Generate a large binary data of 1 MB size
	dataSize := 1024 * 1024
	data := make([]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.EncodeToString(data)
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	// Generate a large binary data of 1 MB size
	dataSize := 1024 * 1024
	data := make([]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = byte(i % 256)
	}

	// Encode the data to get the base64 encoded string
	encodedData := base64.StdEncoding.EncodeToString(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.DecodeString(encodedData)
	}
}
