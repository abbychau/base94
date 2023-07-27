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

func TestEncodePrintableChars(t *testing.T) {
	const maxChar = 126

	// Test different combinations of binary data inputs
	testCases := [][]byte{
		// Some random binary data
		{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0},
		// Maximum and minimum values
		{0x00, 0x7f, 0x80, 0xff},
		// All zeros
		{0x00, 0x00, 0x00, 0x00},
		// All ones
		{0xff, 0xff, 0xff, 0xff},
	}

	// Generate longer binary data for testing
	for i := 0; i < 100; i++ {
		data := make([]byte, 1000*(i+1))
		for j := range data {
			data[j] = byte((i * j) % maxChar)
		}
		testCases = append(testCases, data)
	}

	for _, testData := range testCases {
		encodedData := Encode(testData)

		// Check if all characters in the encoded output are printable
		for _, char := range encodedData {
			if char < minChar || char > maxChar {
				t.Errorf("Encode result contains non-printable character: %q", char)
			}
		}
	}
}
