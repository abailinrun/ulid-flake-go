<h1 align="left">
	<img width="240" src="https://raw.githubusercontent.com/ulid-flake/spec/main/logo.png" alt="ulid-flake">
</h1>

# Ulid-Flake, A 64-bit ULID variant featuring Snowflake - the go implementation

Ulid-Flake is a compact `64-bit` ULID (Universally Unique Lexicographically Sortable Identifier) variant inspired by ULID and Twitter's Snowflake. It features a 1-bit sign bit, a 43-bit timestamp, and a 20-bit randomness. Additionally, it offers a scalable version using the last 5 bits as a scalability identifier (e.g., machineID, podID, nodeID).

herein is proposed Ulid-Flake:

```go
ulidflake.New() // 00CMXB6TAK4SA
ulidflake.New().Int() // 14246757444195114
```

## Features

- **Compact and Efficient**: Uses only 64 bits, making it compatible with common integer types like `int64` and `bigint`.
- **Scalability**: Provides 32 configurations for scalability using a distributed system.
- **Lexicographically Sortable**: Ensures lexicographical order.
- **Canonical Encoding**: Encoded as a 13-character string using Crockford's Base32.
- **Monotonicity and Randomness**: Monotonic sort order within the same millisecond with enhanced randomness to prevent predictability.

## Installation

```sh
go get github.com/abailinrun/ulid-flake-go@latest
```

## Basic Usage

```go
package main

import (
    "fmt"
    "time"
	ulidflake "github.com/abailinrun/ulid-flake-go/ulidflake"
	ulidflakescalable "github.com/abailinrun/ulid-flake-go/ulidflakescalable"
)

func main() {
    // Configure settings for stand-alone version
    if err := ulidflake.SetConfig(
        ulidflake.WithEpochTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)), // Custom epoch time, default 2024-01-01
        ulidflake.WithEntropySize(1),                                         // Custom entropy size, 1, 2 or 3, default 1
    ); err != nil {
        log.Fatalf("Failed to set config: %v", err)
    }

    // Configure settings for scalable version
    if err := ulidflake.SetConfig(
        ulidflake.WithEpochTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)), // Custom epoch time, default 2024-01-01
        ulidflake.WithEntropySize(1),                                         // Custom entropy size, 1, 2 or 3, default 1
        ulidflake.WithScalabilityID(3),                                       // Custom scalability ID (e.g., machineID, podID, nodeID), 1~32, default 0
    ); err != nil {
        log.Fatalf("Failed to set config: %v", err)
    }

    // Generate a new Ulid-Flake instance
    flakeID, _ := ulidflake.New()
    fmt.Printf("Base32: %s\n", flakeID.String())
    fmt.Printf("Integer: %d\n", flakeID.Int())
    fmt.Printf("Timestamp: %d\n", flakeID.Timestamp())
    fmt.Printf("Randomness: %d\n", flakeID.Randomness())
    fmt.Printf("Hex: %s\n", flakeID.Hex())
    fmt.Printf("Binary: %s\n", flakeID.Bin())
    // Base32:     00F2N078MDT7J
    // Integer:    16981964897052914
    // Timestamp:  16195263764
    // Randomness: 452850
    // Hex:        0x3C5501D146E8F2
    // Bin:        0b111100010101010000000111010001010001101110100011110010
}
```

## Monotonicity Testing

Stand-alone version:

```go
package main

import (
    "fmt"
    "time"
	ulidflake "github.com/abailinrun/ulid-flake-go/ulidflake"
)

func main() {
	for range 5 {
		flakeID, err := ulidflake.New()
		if err != nil {
			log.Fatalf("Failed to generate Ulid-Flake: %v", err)
		}
		fmt.Printf("\nBase32: %v\n", flakeID.String())
		fmt.Printf("Int64: %v\n", flakeID.Int())
		fmt.Printf("Hex: %v\n", flakeID.Hex())
		fmt.Printf("Bin: %v\n\n", flakeID.Bin())
	}

    // Base32: 00F2T7EEA54NN
    // Int64: 16987710680109749
    // Hex: 0x3C5A3B9CA292B5
    // Bin: 0b111100010110100011101110011100101000101001001010110101


    // Base32: 00F2T7EEA54T4
    // Int64: 16987710680109892
    // Hex: 0x3C5A3B9CA29344
    // Bin: 0b111100010110100011101110011100101000101001001101000100


    // Base32: 00F2T7EEA54VF
    // Int64: 16987710680109935
    // Hex: 0x3C5A3B9CA2936F
    // Bin: 0b111100010110100011101110011100101000101001001101101111


    // Base32: 00F2T7EEA5510
    // Int64: 16987710680110112
    // Hex: 0x3C5A3B9CA29420
    // Bin: 0b111100010110100011101110011100101000101001010000100000


    // Base32: 00F2T7EEA5552
    // Int64: 16987710680110242
    // Hex: 0x3C5A3B9CA294A2
    // Bin: 0b111100010110100011101110011100101000101001010010100010
}
```

scalable version:

```go
package main

import (
    "fmt"
    "time"
	ulidflakescalable "github.com/abailinrun/ulid-flake-go/ulidflakescalable"
)

func main() {

	for range 5 {
		flakeID, err := ulidflake.New()
		if err != nil {
			log.Fatalf("Failed to generate Ulid-Flake: %v", err)
		}
		fmt.Printf("\nBase32: %v\n", flakeID.String())
		fmt.Printf("Int64: %v\n", flakeID.Int())
		fmt.Printf("Hex: %v\n", flakeID.Hex())
		fmt.Printf("Bin: %v\n\n", flakeID.Bin())
	}

    // Base32: 00F2T8E54BKF0
    // Int64: 16987744731778528
    // Hex: 0x3C5A438A45CDE0
    // Bin: 0b111100010110100100001110001010010001011100110111100000


    // Base32: 00F2T8E54BP40
    // Int64: 16987744731781248
    // Hex: 0x3C5A438A45D880
    // Bin: 0b111100010110100100001110001010010001011101100010000000


    // Base32: 00F2T8E54BWZ0
    // Int64: 16987744731788256
    // Hex: 0x3C5A438A45F3E0
    // Bin: 0b111100010110100100001110001010010001011111001111100000


    // Base32: 00F2T8E54BXD0
    // Int64: 16987744731788704
    // Hex: 0x3C5A438A45F5A0
    // Bin: 0b111100010110100100001110001010010001011111010110100000


    // Base32: 00F2T8E54C2P0
    // Int64: 16987744731794112
    // Hex: 0x3C5A438A460AC0
    // Bin: 0b111100010110100100001110001010010001100000101011000000
}
```

## Creating Ulid-Flake Instances from other sources

### From Integer

```go
ulidFlake, _ := ulidflake.FromInt(1234567890123456789)
fmt.Printf("From Int: %s\n", ulidFlake.String())
```

### From Base32 String

```go
ulidFlake, _ := ulidflake.FromStr("01AN4Z07BY79K")
fmt.Printf("From String: %s\n", ulidFlake.String())
```

### From Unix Epoch Time

```go
ulidFlake, _ := ulidflake.FromUnixEpochTime(1672531200)
fmt.Printf("From Unix Time: %s\n", ulidFlake.String())
```

## Command Line Tool

This implementation also provides a tool to generate and parse Ulid-Flakes at the command line.

```sh
# stand-alone version:
go install github.com/abailinrun/ulid-flake-go/cmd/ulidflake@latest

# scalable version:
go install github.com/abailinrun/ulid-flake-go/cmd/ulidflakescalable@latest
```

Usage:

```
ulidflake
    -entropy int
        Set the custom entropy size (default: 1)
    -epoch string
        Set the custom epoch time (default "2024-01-01T00:00:00Z")
    -generate
        Generate a new Ulid-Flake
    -parse string
        Parse a Ulid-Flake string
```

```
ulidflakescalable
    -entropy int
        Set the custom entropy size (default: 1)
    -epoch string
        Set the custom epoch time (default "2024-01-01T00:00:00Z")
    -sid int
        Set the custom scalability ID (default: 0)
    -generate
        Generate a new Ulid-Flake
    -parse string
        Parse a Ulid-Flake string
```

Examples:

Stand-alone version:

```sh
./ulidflake -generate

    ██╗░░░██╗██╗░░░░░██╗██████╗░░░░░░░███████╗██╗░░░░░░█████╗░██╗░░██╗███████╗
    ██║░░░██║██║░░░░░██║██╔══██╗░░░░░░██╔════╝██║░░░░░██╔══██╗██║░██╔╝██╔════╝
    ██║░░░██║██║░░░░░██║██║░░██║█████╗█████╗░░██║░░░░░███████║█████═╝░█████╗░░
    ██║░░░██║██║░░░░░██║██║░░██║╚════╝██╔══╝░░██║░░░░░██╔══██║██╔═██╗░██╔══╝░░
    ╚██████╔╝███████╗██║██████╔╝░░░░░░██║░░░░░███████╗██║░░██║██║░╚██╗███████╗
    ░╚═════╝░╚══════╝╚═╝╚═════╝░░░░░░░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝╚══════╝

Generated Ulid-Flake:
  Base32:     00F2N6ZRB5HDG
  Integer:    16982197352449456
  Timestamp:  16195485451
  Randomness: 181680
  Hex:        0x3C5537F0B2C5B0
  Bin:        0b111100010101010011011111110000101100101100010110110000
```

```sh
./ulidflake -parse 7ZZZZZZZZZZZZ

    ██╗░░░██╗██╗░░░░░██╗██████╗░░░░░░░███████╗██╗░░░░░░█████╗░██╗░░██╗███████╗
    ██║░░░██║██║░░░░░██║██╔══██╗░░░░░░██╔════╝██║░░░░░██╔══██╗██║░██╔╝██╔════╝
    ██║░░░██║██║░░░░░██║██║░░██║█████╗█████╗░░██║░░░░░███████║█████═╝░█████╗░░
    ██║░░░██║██║░░░░░██║██║░░██║╚════╝██╔══╝░░██║░░░░░██╔══██║██╔═██╗░██╔══╝░░
    ╚██████╔╝███████╗██║██████╔╝░░░░░░██║░░░░░███████╗██║░░██║██║░╚██╗███████╗
    ░╚═════╝░╚══════╝╚═╝╚═════╝░░░░░░░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝╚══════╝

Parsed Ulid-Flake:
  Base32:     7ZZZZZZZZZZZZ
  Integer:    9223372036854775807
  Timestamp:  8796093022207
  Randomness: 1048575
  Hex:        0x7FFFFFFFFFFFFFFF
  Bin:        0b111111111111111111111111111111111111111111111111111111111111111
```

scalable version:

```sh
./ulidflake -generate -sid 31

    ██╗░░░██╗██╗░░░░░██╗██████╗░░░░░░░███████╗██╗░░░░░░█████╗░██╗░░██╗███████╗░░░░░░░██████╗
    ██║░░░██║██║░░░░░██║██╔══██╗░░░░░░██╔════╝██║░░░░░██╔══██╗██║░██╔╝██╔════╝░░░░░░██╔════╝
    ██║░░░██║██║░░░░░██║██║░░██║█████╗█████╗░░██║░░░░░███████║█████═╝░█████╗░░█████╗╚█████╗░
    ██║░░░██║██║░░░░░██║██║░░██║╚════╝██╔══╝░░██║░░░░░██╔══██║██╔═██╗░██╔══╝░░╚════╝░╚═══██╗
    ╚██████╔╝███████╗██║██████╔╝░░░░░░██║░░░░░███████╗██║░░██║██║░╚██╗███████╗░░░░░░██████╔╝
    ░╚═════╝░╚══════╝╚═╝╚═════╝░░░░░░░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝╚══════╝░░░░░░╚═════╝░

Generated Ulid-Flake:
  Base32:     00F2NC9TEXQ0Z
  Integer:    16982379959606303
  Timestamp:  16195659598
  Randomness: 30432
  SID:        31
  Hex:        0x3C556274EEDC1F
  Bin:        0b111100010101010110001001110100111011101101110000011111
```

```sh
./ulidflake -parse 7ZZZZZZZZZZZZ

    ██╗░░░██╗██╗░░░░░██╗██████╗░░░░░░░███████╗██╗░░░░░░█████╗░██╗░░██╗███████╗░░░░░░░██████╗
    ██║░░░██║██║░░░░░██║██╔══██╗░░░░░░██╔════╝██║░░░░░██╔══██╗██║░██╔╝██╔════╝░░░░░░██╔════╝
    ██║░░░██║██║░░░░░██║██║░░██║█████╗█████╗░░██║░░░░░███████║█████═╝░█████╗░░█████╗╚█████╗░
    ██║░░░██║██║░░░░░██║██║░░██║╚════╝██╔══╝░░██║░░░░░██╔══██║██╔═██╗░██╔══╝░░╚════╝░╚═══██╗
    ╚██████╔╝███████╗██║██████╔╝░░░░░░██║░░░░░███████╗██║░░██║██║░╚██╗███████╗░░░░░░██████╔╝
    ░╚═════╝░╚══════╝╚═╝╚═════╝░░░░░░░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝╚══════╝░░░░░░╚═════╝░

Parsed Ulid-Flake:
  Base32:     7ZZZZZZZZZZZZ
  Integer:    9223372036854775807
  Timestamp:  8796093022207
  Randomness: 32767
  SID:        31
  Hex:        0x7FFFFFFFFFFFFFFF
  Bin:        0b111111111111111111111111111111111111111111111111111111111111111
```

## Specification

Below is the default stand-alone version specification of Ulid-Flake.

<img width="600" alt="ulid-flake-stand-alone" src="https://github.com/ulid-flake/spec/assets/38312944/37d44c3f-1937-4c2e-b7ec-e7c0f0debe25">

*Note: a `1-bit` sign bit is included in the timestamp.*

```text
Stand-alone version (default):

 00CMXB6TA      K4SA

|---------|    |----|
 Timestamp   Randomness
   44-bit      20-bit
   9-char      4-char
```

Also, a scalable version is provided for distributed system using purpose.

<img width="600" alt="ulid-flake-scalable" src="https://github.com/ulid-flake/spec/assets/38312944/e306ebd9-9406-436f-b6cd-a1004745f1b0">

*Note: a `1-bit` sign bit is included in the timestamp.*

```
Scalable version (optional):

 00CMXB6TA      K4S       A

|---------|    |---|     |-|
 Timestamp   Randomness  Scalability
   44-bit      15-bit    5-bit
   9-char      3-char    1-char
```

### Components

Total `64-bit` size for compatibility with common integer (`long int`, `int64` or `bigint`) types.

**Timestamp**
- The first `1-bit` is a sign bit, always set to 0.
- Remaining `43-bit` timestamp in millisecond precision.
- Custom epoch for extended usage span, starting from `2024-01-01T00:00:00.000Z`.
- Usable until approximately `2302-09-27` AD.

**Randomness**
- `20-bit` randomness for stand-alone version. Provides a collision resistance with a p=0.5 expectation of 1,024 trials. (not much)
- `15-bit` randomness for scalable version.
- Initial random value at each millisecond precision unit.
- adopt a `+n` bits entropy incremental mechanism to ensure uniqueness without predictability.

**Scalability (Scalable version ony)**
- Provide a `5-bit` scalability for distributed system using purpose.
- total 32 configurations can be used.

### Sorting

The left-most character must be sorted first, and the right-most character sorted last, ensuring lexicographical order.
The default ASCII character set must be used.

When using the stand-alone version strictly in a stand-alone environment, or using the scalable version in both stand-alone or distributed environment, sort order is guaranteed within the same millisecond. however, when using the stand-alone version in a distributed system, sort order is not guaranteed within the same millisecond.

*Note: within the same millisecond, sort order is guaranteed in the context of an overflow error could occur.*

### Canonical String Representation

```text
Stand-alone version (default):

tttttttttrrrr

where
t is Timestamp (9 characters)
r is Randomness (4 characters)
```

```text
Scalable version (optional):

tttttttttrrrs

where
t is Timestamp (9 characters)
r is Randomness (3 characters)
s is Scalability (1 characters)
```

#### Encoding

Crockford's Base32 is used as shown. This alphabet excludes the letters I, L, O, and U to avoid confusion and abuse.

```
0123456789ABCDEFGHJKMNPQRSTVWXYZ
```

### Optional Long Int Representation

```text
1234567890123456789

(with a maximum 13-character length in string format)
```

### Monotonicity and Overflow Error Handling

#### Randomness

When generating a Ulid-Flake within the same millisecond, the `randomness` component is incremented by a `n-bit` entropy in the least significant bit position (with carrying).
Thus, comparing just incremented `1-bit` one time, the incremented `n-bit` mechanism cloud lead to an overflow error sooner.

when the generation is failed with overflow error, it should be properly handled in the application to wait and create a new one till the next millisecond is coming. The implementation of Ulid-Flake should just return the overflow error, and leave the rest to the application.

#### Timestamp and Over All

Technically, a `13-character` Base32 encoded string can contain 65 bits of information, whereas a Ulid-Flake must only contain 64 bits. Further more, there is a `1-bit` sign bit at the beginning, only 63 bits are actually carrying effective information. Therefore, the largest valid Ulid-Flake encoded in Base32 is `7ZZZZZZZZZZZZ`, which corresponds to an epoch time of `8,796,093,022,207` or `2^43 - 1`.

Any attempt to decode or encode a Ulid-Flake larger than this should be rejected by all implementations and return an overflow error, to prevent overflow bugs.

### Binary Layout and Byte Order

The components are encoded as 16 octets. Each component is encoded with the Most Significant Byte first (network byte order).

```
Stand-alone version (default):

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                      32_bit_int_time_high                     |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| 12_bit_uint_time_low  |          20_bit_uint_random           |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

```
Scalable version (optional):

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                      32_bit_int_time_high                     |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| 12_bit_uint_time_low  |      15_bit_uint_random     | 5_bit_s |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

# Contributing
We welcome contributions! Please see our CONTRIBUTING.md for guidelines on how to get involved.

# License
This project is licensed under the MIT License. See the LICENSE file for details.

# Acknowledgments

[ULID](https://github.com/ulid/spec)

[Twitter's Snowflake](https://blog.x.com/engineering/en_us/a/2010/announcing-snowflake)
