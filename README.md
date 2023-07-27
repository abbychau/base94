# Base94

Base94 is a Go library for encoding and decoding binary data using a custom character set, which includes characters ranging from decimal 33 ('!') to decimal 126 ('~'). This library allows you to represent binary data using a set of printable ASCII characters, making it easy to copy and paste the encoded data.

# Why Base94?

Base94 is an encoding scheme that uses a larger character set compared to popular encodings like Base64. By using a character set that includes printable ASCII characters, Base94 can achieve a higher data density, resulting in shorter encoded strings. This can be beneficial when you need to represent binary data in contexts that only support printable characters, such as when transmitting data in environments that do not handle binary data well.

Base94 provides a simple and straightforward way to encode and decode binary data, making it ideal for use cases where you want to convert binary data to a printable format without relying on more complex encryption or compression algorithms.

# Performance

Base94 offers a significant performance advantage over the built-in Base64 encoding and decoding functions provided by Go's standard library. Through careful optimization and utilization of bitwise operations, Base94 achieves approximately three times faster encoding and decoding speeds.

# Sample Usage

```go
package main

import (
	"fmt"
	"github.com/abbychau/base94"
)

func main() {
	// Example usage:

	// Binary data to be encoded
	data := []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}

	// Encoding binary data to a printable string
	encodedData := base94.Encode(data)
	fmt.Println("Encoded data:", encodedData)

	// Decoding the encoded string back to binary data
	decodedData := base94.Decode(encodedData)
	fmt.Println("Decoded data:", decodedData)
}
```

Please note that Base94 is intended for non-critical or non-sensitive use cases. If you require more secure data encryption, consider using stronger encryption algorithms. Base94 is designed for situations where the main goal is to represent binary data in a printable format for ease of sharing and handling.

# Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/abbychau/base94
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkEncode-16                   274           4199166 ns/op
BenchmarkDecode-16                   295           4048855 ns/op
BenchmarkBase64Encode-16             856           1397131 ns/op
BenchmarkBase64Decode-16             780           1494461 ns/op
```

# License

Base94 is licensed under the MIT license. See the LICENSE file for more information.
