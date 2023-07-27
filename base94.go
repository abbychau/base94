// base94.go

package base94

import "strings"

const minChar = 33

// Encode encodes the given binary data using the custom character set.
func Encode(data []byte) string {
	var encoded strings.Builder

	var buffer int         // Temporary buffer to hold bits from the input data.
	var bitsRemaining uint // Number of bits remaining in the buffer to form the next character.

	for _, b := range data {
		// Append the current byte to the buffer and shift it to the left to make room for the next byte.
		buffer |= int(b) << bitsRemaining
		bitsRemaining += 8

		// Form characters from the buffer as long as we have enough bits (6 or more) to form one.
		for bitsRemaining >= 6 {
			// Extract the lower 6 bits of the buffer (63 in binary is 111111) to get the character index.
			charIndex := buffer & 63

			// Add the minimum character value to the character index to get the actual character value.
			encoded.WriteByte(byte(charIndex) + minChar)

			// Shift the buffer to the right to remove the 6 bits that were just used to form the character.
			buffer >>= 6
			// Reduce the count of bits remaining in the buffer.
			bitsRemaining -= 6
		}
	}

	// If there are remaining bits in the buffer (less than 6), form a character from them.
	if bitsRemaining > 0 {
		// Extract the remaining bits (less than 6) to get the character index.
		charIndex := buffer & 63

		// Add the minimum character value to the character index to get the actual character value.
		encoded.WriteByte(byte(charIndex) + minChar)
	}

	return encoded.String()
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
