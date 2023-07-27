// base94.go

package base94

const minChar = 33

// Encode encodes the given binary data using the custom character set.
func Encode(data []byte) string {
	// Calculate the maximum length of the encoded string.
	maxEncodedLen := (len(data)*8 + 5) / 6 // Ceiling of (len(data) * 8 / 6)

	// Create a fixed-length byte slice to hold the encoded characters.
	encoded := make([]byte, maxEncodedLen)

	var buffer int         // Temporary buffer to hold bits from the input data.
	var bitsRemaining uint // Number of bits remaining in the buffer to form the next character.
	var index int          // Index to track the current position in the encoded slice.

	for _, b := range data {
		buffer |= int(b) << bitsRemaining
		bitsRemaining += 8

		for bitsRemaining >= 6 {
			charIndex := buffer & 63
			encoded[index] = byte(charIndex) + minChar

			index++
			buffer >>= 6
			bitsRemaining -= 6
		}
	}

	if bitsRemaining > 0 {
		charIndex := buffer & 63
		encoded[index] = byte(charIndex) + minChar
		index++
	}

	return string(encoded[:index])
}

// Decode decodes the given encoded data using the custom character set and returns binary data.
func Decode(encodedData string) []byte {
	var decoded []byte

	var buffer int         // Temporary buffer to hold bits from the encoded data.
	var bitsRemaining uint // Number of bits remaining in the buffer to form the next byte.

	for _, char := range encodedData {
		// Append the current character to the buffer and shift it to the left to make room for the next character.
		buffer |= (int(char) - minChar) << bitsRemaining
		bitsRemaining += 6

		// Form bytes from the buffer as long as we have enough bits (8 or more) to form one byte.
		for bitsRemaining >= 8 {
			// Extract the lower 8 bits of the buffer (255 in binary is 11111111) to get the byte value.
			decoded = append(decoded, byte(buffer&255))

			// Shift the buffer to the right to remove the 8 bits that were just used to form the byte.
			buffer >>= 8
			// Reduce the count of bits remaining in the buffer.
			bitsRemaining -= 8
		}
	}

	return decoded
}
