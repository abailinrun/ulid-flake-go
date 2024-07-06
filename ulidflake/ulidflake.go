package ulidflake

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	DefaultEpochSec = 1704067200 // Default epoch time in seconds (2024-01-01 00:00:00 UTC)

	IntSize = 63                 // 64-bit signed integer size without sign bit
	MinInt  = 0                  // 64-bit signed integer possible minimum value for Ulid-Flake (0)
	MaxInt  = (1 << IntSize) - 1 // 64-bit signed integer possible maximum value for Ulid-Flake (9223372036854775807)

	TimestampSize = 43                       // 43-bit timestamp size
	MinTimestamp  = 0                        // 43-bit minimum value (0)
	MaxTimestamp  = (1 << TimestampSize) - 1 // 43-bit maximum value (8796093022207)

	RandomnessSize = 20                        // 20-bit randomness size
	MinRandomness  = 0                         // 20-bit minimum value (0)
	MaxRandomness  = (1 << RandomnessSize) - 1 // 20-bit maximum value (1048575)

	MinEntropySize = 1 // Minimum entropy size (1 byte)
	MaxEntropySize = 3 // Maximum entropy size (3 byte)

	UlidFlakeSize = 64              // Ulid-Flake size in bits
	UlidFlakeLen  = 13              // Length of Ulid-Flake string
	MinUlidFlake  = "0000000000000" // Minimum possible Ulid-Flake value (0)
	MaxUlidFlake  = "7ZZZZZZZZZZZZ" // Maximum possible Ulid-Flake value (9223372036854775807)

	encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ" // Ulid-Flake Crockford's Base32 encoding characters
)

var (
	ErrOverflow         = errors.New("overflow error")
	ErrInvalidTimestamp = errors.New("invalid timestamp")
	ErrInvalidULID      = errors.New("invalid ULID")
	ErrInvalidConfig    = errors.New("invalid configuration")
	ErrInvalidEntropy   = errors.New("entropy size must be between 1 and 3")
)

type UlidFlake struct {
	value int64
}

var (
	mutex              sync.Mutex
	previousTimestamp  int64
	previousRandomness int64
	epochTime          time.Time = time.Unix(DefaultEpochSec, 0).UTC() // Default epoch time (2024-01-01 00:00:00 UTC)
	entropySize        int       = MinEntropySize
)

// Option defines the type for functional options
type Option func(*config) error

type config struct {
	epochTime   time.Time
	entropySize int
}

// NewUlidFlake creates a new UlidFlake
func NewUlidFlake(value int64) (*UlidFlake, error) {
	if value < 0 || value > (1<<63-1) {
		return nil, ErrOverflow
	}
	return &UlidFlake{value: value}, nil
}

// String returns the Base32 string representation
func (u *UlidFlake) String() string {
	return encodeBase32(u.value, UlidFlakeLen)
}

// Int returns the integer representation
func (u *UlidFlake) Int() int64 {
	return u.value
}

// Hex returns the hexadecimal string representation
func (u *UlidFlake) Hex() string {
	return "0x" + fmt.Sprintf("%X", u.value)
}

// Bin returns the binary string representation
func (u *UlidFlake) Bin() string {
	return "0b" + fmt.Sprintf("%b", u.value)
}

// Timestamp returns the timestamp component
func (u *UlidFlake) Timestamp() int64 {
	return (u.value >> 20) & MaxTimestamp
}

// Randomness returns the randomness component
func (u *UlidFlake) Randomness() int64 {
	return u.value & MaxRandomness
}

// bytes converts the UlidFlake value to a byte slice
func (u *UlidFlake) bytes() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(u.value))
	return b[1:] // return 7 bytes for 56 bits
}

// Helper functions for encoding Base32
func encodeBase32(value int64, length int) string {
	encoded := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		encoded[i] = encoding[value&31]
		value >>= 5
	}
	return string(encoded)
}

// decodeBase32 decodes a Base32 string to a numeric value
func decodeBase32(encoded string) (int64, error) {
	var value int64
	for _, c := range encoded {
		idx := -1
		for i, encChar := range encoding {
			if c == rune(encChar) {
				idx = i
				break
			}
		}
		if idx == -1 {
			return 0, ErrInvalidULID
		}
		value = value*32 + int64(idx)
	}
	return value, nil
}

// GenerateTimestamp generates a 43-bit timestamp
func generateTimestamp(now time.Time) (int64, error) {
	timestamp := now.Sub(epochTime).Milliseconds()
	if timestamp < MinTimestamp || timestamp > MaxTimestamp {
		return 0, ErrOverflow
	}
	return timestamp, nil
}

// GenerateRandomBytes generates a byte slice of random bytes
func generateRandomBytes(size int) ([]byte, error) {
	rnd := make([]byte, size)
	_, err := rand.Read(rnd)
	if err != nil {
		return nil, err
	}
	return rnd, nil
}

// GenerateRandomness generates a 20-bit randomness value
func generateRandomness(randomFunc func(size int) ([]byte, error)) (int64, error) {
	rnd, err := randomFunc(MaxEntropySize)
	if err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint32(append([]byte{0}, rnd...)) & MaxRandomness), nil
}

// GenerateEntropy generates an entropy value to increment randomness
func generateEntropy(size int, randomFunc func(size int) ([]byte, error)) (int64, error) {
	if size < MinEntropySize || size > MaxEntropySize {
		return 0, ErrInvalidEntropy
	}
	rnd, err := randomFunc(size)
	if err != nil {
		return 0, err
	}
	var entropy int64
	for _, b := range rnd {
		entropy = (entropy << 8) | int64(b)
	}
	return entropy, nil
}

// New generates a new Ulid-Flake with the given entropy size
func New() (*UlidFlake, error) {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now().UTC()
	timestamp, err := generateTimestamp(now)
	if err != nil {
		return nil, err
	}

	var randomness int64
	if timestamp < previousTimestamp {
		return nil, ErrInvalidTimestamp
	}
	if timestamp == previousTimestamp {
		entropy := int64(0)
		for entropy <= 0 {
			entropy, err = generateEntropy(entropySize, generateRandomBytes)
			if err != nil {
				return nil, err
			}
		}
		newRandomness := previousRandomness + entropy
		if newRandomness > MaxRandomness {
			return nil, ErrOverflow
		}
		randomness = newRandomness
	} else {
		randomness, err = generateRandomness(generateRandomBytes)
		if err != nil {
			return nil, err
		}
	}
	previousTimestamp = timestamp
	previousRandomness = randomness

	signBit := int64(0)
	combined := (signBit << 63) | (timestamp << 20) | randomness

	if combined > (1<<63 - 1) {
		return nil, ErrOverflow
	}

	return NewUlidFlake(combined)
}

// Parse parses a Ulid-Flake string
func Parse(ulidFlakeString string) (*UlidFlake, error) {
	if len(ulidFlakeString) != UlidFlakeLen {
		return nil, ErrInvalidULID
	}
	value, err := decodeBase32(ulidFlakeString)
	if err != nil {
		return nil, err
	}
	return NewUlidFlake(value)
}

// FromInt creates a Ulid-Flake instance from an integer
func FromInt(value int64) (*UlidFlake, error) {
	if value < 0 || value > (1<<63-1) {
		return nil, ErrOverflow
	}
	return NewUlidFlake(value)
}

// FromStr creates a Ulid-Flake instance from a Base32 string
func FromStr(ulidFlakeString string) (*UlidFlake, error) {
	return Parse(ulidFlakeString)
}

// FromUnixEpochTime creates a Ulid-Flake instance from a Unix epoch time
func FromUnixEpochTime(unixTimeSec int64) (*UlidFlake, error) {
	timestamp := int64(unixTimeSec*1000) - int64(DefaultEpochSec*1000)
	if timestamp < MinTimestamp || timestamp > MaxTimestamp {
		return nil, ErrOverflow
	}
	randomness, err := generateRandomness(generateRandomBytes)
	if err != nil {
		return nil, err
	}
	signBit := int64(0)
	combined := (signBit << 63) | (timestamp << 20) | randomness
	if combined > (1<<63 - 1) {
		return nil, ErrOverflow
	}
	return NewUlidFlake(combined)
}

// SetConfig sets the configuration values with functional options
func SetConfig(opts ...Option) error {
	cfg := &config{
		epochTime:   time.Unix(DefaultEpochSec, 0).UTC(),
		entropySize: MinEntropySize,
	}

	for _, opt := range opts {
		err := opt(cfg)
		if err != nil {
			return err
		}
	}

	epochTime = cfg.epochTime
	entropySize = cfg.entropySize

	return nil
}

// WithEpochTime sets the custom epoch time
func WithEpochTime(epoch time.Time) Option {
	return func(cfg *config) error {
		cfg.epochTime = epoch
		return nil
	}
}

// WithEntropySize sets the custom entropy size
func WithEntropySize(entropy int) Option {
	return func(cfg *config) error {
		if entropy < MinEntropySize || entropy > MaxEntropySize {
			return ErrInvalidEntropy
		}
		cfg.entropySize = entropy
		return nil
	}
}
