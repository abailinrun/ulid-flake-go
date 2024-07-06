package ulidflakescalable

import (
	reflect "reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUlidFlake(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name    string
		args    args
		want    *UlidFlake
		wantErr bool
	}{
		{
			name: "minimal value",
			args: args{
				value: 0,
			},
			want: &UlidFlake{
				value: 0,
			},
			wantErr: false,
		},
		{
			name: "maximal value",
			args: args{
				value: MaxInt,
			},
			want: &UlidFlake{
				value: MaxInt,
			},
			wantErr: false,
		},
		{
			name: "negative value",
			args: args{
				value: -1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUlidFlake(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUlidFlake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUlidFlake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlidFlake_String(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "minimal value",
			fields: fields{
				value: 0,
			},
			want: "0000000000000",
		},
		{
			name: "maximal value",
			fields: fields{
				value: MaxInt,
			},
			want: "7ZZZZZZZZZZZZ",
		},
		{
			name: "negative value (this should never happen, but theoretically possible)",
			fields: fields{
				value: -1,
			},
			want: "ZZZZZZZZZZZZZ",
		},
		{
			name: "overflow value (this should never happen, but theoretically possible)",
			fields: fields{
				value: -1 << IntSize,
			},
			want: "R000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UlidFlake{
				value: tt.fields.value,
			}
			if got := u.String(); got != tt.want {
				t.Errorf("UlidFlake.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlidFlake_Int(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "minimal value",
			fields: fields{
				value: 0,
			},
			want: 0,
		},
		{
			name: "maximal value",
			fields: fields{
				value: MaxInt,
			},
			want: MaxInt,
		},
		{
			name: "negative value (this should never happen, but theoretically possible)",
			fields: fields{
				value: -1,
			},
			want: -1,
		},
		{
			name: "overflow value (this should never happen, but theoretically possible)",
			fields: fields{
				value: -1 << IntSize,
			},
			want: -1 << IntSize,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UlidFlake{
				value: tt.fields.value,
			}
			if got := u.Int(); got != tt.want {
				t.Errorf("UlidFlake.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlidFlake_Timestamp(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "minimal value",
			fields: fields{
				value: 0,
			},
			want: 0,
		},
		{
			name: "maximal value",
			fields: fields{
				value: MaxInt,
			},
			want: 1<<43 - 1,
		},
		{
			name: "negative value (overflow, this should never happen, but theoretically possible)",
			fields: fields{
				value: -1,
			},
			want: 1<<43 - 1,
		},
		{
			name: "overflow value (this should never happen, but theoretically possible)",
			fields: fields{
				value: -1 << IntSize,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UlidFlake{
				value: tt.fields.value,
			}
			if got := u.Timestamp(); got != tt.want {
				t.Errorf("UlidFlake.Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlidFlake_Randomness(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "minimal value",
			fields: fields{
				value: 0,
			},
			want: 0,
		},
		{
			name: "maximal value",
			fields: fields{
				value: MaxInt,
			},
			want: 1<<15 - 1,
		},
		{
			name: "negative value (overflow, this should never happen, but theoretically possible)",
			fields: fields{
				value: -1,
			},
			want: 1<<15 - 1,
		},
		{
			name: "overflow value (this should never happen, but theoretically possible)",
			fields: fields{
				value: -1 << IntSize,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UlidFlake{
				value: tt.fields.value,
			}
			if got := u.Randomness(); got != tt.want {
				t.Errorf("UlidFlake.Randomness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlidFlake_SID(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "minimal value",
			fields: fields{
				value: 0,
			},
			want: 0,
		},
		{
			name: "maximal value",
			fields: fields{
				value: MaxInt,
			},
			want: 1<<5 - 1,
		},
		{
			name: "negative value (overflow, this should never happen, but theoretically possible)",
			fields: fields{
				value: -1,
			},
			want: 1<<5 - 1,
		},
		{
			name: "overflow value (this should never happen, but theoretically possible)",
			fields: fields{
				value: -1 << IntSize,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UlidFlake{
				value: tt.fields.value,
			}
			if got := u.SID(); got != tt.want {
				t.Errorf("UlidFlake.SID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUlidFlake_bytes(t *testing.T) {
	type fields struct {
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "minimal value",
			fields: fields{
				value: 0,
			},
			want: []byte{0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "maximal value",
			fields: fields{
				value: MaxInt,
			},
			want: []byte{255, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "negative value (overflow, this should never happen, but theoretically possible)",
			fields: fields{
				value: -1,
			},
			want: []byte{255, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "overflow value (this should never happen, but theoretically possible)",
			fields: fields{
				value: -1 << IntSize,
			},
			want: []byte{0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UlidFlake{
				value: tt.fields.value,
			}
			if got := u.bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UlidFlake.bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBase32(t *testing.T) {
	type args struct {
		value  int64
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "minimal value",
			args: args{
				value:  0,
				length: UlidFlakeLen,
			},
			want: "0000000000000",
		},
		{
			name: "maximal value",
			args: args{
				value:  MaxInt,
				length: UlidFlakeLen,
			},
			want: "7ZZZZZZZZZZZZ",
		},
		{
			name: "negative value (this should never happen, but theoretically possible)",
			args: args{
				value:  -1,
				length: UlidFlakeLen,
			},
			want: "ZZZZZZZZZZZZZ",
		},
		{
			name: "overflow value (this should never happen, but theoretically possible)",
			args: args{
				value:  -1 << IntSize,
				length: UlidFlakeLen,
			},
			want: "R000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBase32(tt.args.value, tt.args.length); got != tt.want {
				t.Errorf("encodeBase32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateTimestamp(t *testing.T) {
	type args struct {
		now time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "minimal value",
			args: args{
				now: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "maximal value",
			args: args{
				now: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(1<<43-1) * time.Millisecond),
			},
			want:    1<<43 - 1,
			wantErr: false,
		},
		{
			name: "negative value (overflow, this should never happen, but theoretically possible)",
			args: args{
				now: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Duration(1<<43) * time.Millisecond),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "before epoch time",
			args: args{
				now: time.Date(2023, 12, 31, 23, 59, 59, 999000000, time.UTC),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateTimestamp(tt.args.now)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomness(t *testing.T) {
	type args struct {
		randomFunc func(size int) ([]byte, error)
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "minimal value",
			args: args{
				randomFunc: func(size int) ([]byte, error) {
					return []byte{0, 0, 0}, nil
				},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "maximal value",
			args: args{
				randomFunc: func(size int) ([]byte, error) {
					return []byte{255, 255, 255}, nil
				},
			},
			want:    1<<15 - 1,
			wantErr: false,
		},
		{
			name: "overflow, this should never happen, but theoretically possible",
			args: args{
				randomFunc: func(size int) ([]byte, error) {
					return []byte{255, 255, 255, 255}, nil
				},
			},
			want:    1<<15 - 1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandomness(tt.args.randomFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandomness() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateRandomness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateEntropy(t *testing.T) {
	type args struct {
		size       int
		randomFunc func(size int) ([]byte, error)
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "minimal value",
			args: args{
				size: 1,
				randomFunc: func(size int) ([]byte, error) {
					return []byte{0}, nil
				},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "maximal value",
			args: args{
				size: 2,
				randomFunc: func(size int) ([]byte, error) {
					return []byte{255, 255}, nil
				},
			},
			want:    1<<16 - 1,
			wantErr: false,
		},
		{
			name: "too big size",
			args: args{
				size: 3,
				randomFunc: func(size int) ([]byte, error) {
					return []byte{255, 255, 255}, nil
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "too small size",
			args: args{
				size: 0,
				randomFunc: func(size int) ([]byte, error) {
					return []byte{}, nil
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateEntropy(tt.args.size, tt.args.randomFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateEntropy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateEntropy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Len(t, got.String(), UlidFlakeLen)
			assert.LessOrEqual(t, got.Timestamp(), time.Now().UTC().UnixMilli()-DefaultEpochSec)
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		ulidFlakeString string
	}
	tests := []struct {
		name    string
		args    args
		want    *UlidFlake
		wantErr bool
	}{
		{
			name: "minimal value",
			args: args{
				ulidFlakeString: "0000000000000",
			},
			want: &UlidFlake{
				value: 0,
			},
			wantErr: false,
		},
		{
			name: "maximal value",
			args: args{
				ulidFlakeString: "7ZZZZZZZZZZZZ",
			},
			want: &UlidFlake{
				value: MaxInt,
			},
			wantErr: false,
		},
		{
			name: "negative value (overflow, this should never happen, but theoretically possible)",
			args: args{
				ulidFlakeString: "8000000000000",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative value (this should never happen, but theoretically possible)",
			args: args{
				ulidFlakeString: "R000000000000",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative value (this should never happen, but theoretically possible)",
			args: args{
				ulidFlakeString: "ZZZZZZZZZZZZZ",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid length 12",
			args: args{
				ulidFlakeString: "000000000000",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid length 0",
			args: args{
				ulidFlakeString: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.ulidFlakeString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name    string
		args    args
		want    *UlidFlake
		wantErr bool
	}{
		{
			name: "minimal value",
			args: args{
				value: 0,
			},
			want: &UlidFlake{
				value: 0,
			},
			wantErr: false,
		},
		{
			name: "maximal value",
			args: args{
				value: MaxInt,
			},
			want: &UlidFlake{
				value: MaxInt,
			},
			wantErr: false,
		},
		{
			name: "negative value (this should never happen, but theoretically possible)",
			args: args{
				value: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "overflow value (this should never happen, but theoretically possible)",
			args: args{
				value: -1 << IntSize,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromInt(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromStr(t *testing.T) {
	type args struct {
		ulidFlakeString string
	}
	tests := []struct {
		name    string
		args    args
		want    *UlidFlake
		wantErr bool
	}{
		{
			name: "minimal value",
			args: args{
				ulidFlakeString: "0000000000000",
			},
			want: &UlidFlake{
				value: 0,
			},
			wantErr: false,
		},
		{
			name: "maximal value",
			args: args{
				ulidFlakeString: "7ZZZZZZZZZZZZ",
			},
			want: &UlidFlake{
				value: MaxInt,
			},
			wantErr: false,
		},
		{
			name: "negative value (overflow, this should never happen, but theoretically possible)",
			args: args{
				ulidFlakeString: "8000000000000",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative value (this should never happen, but theoretically possible)",
			args: args{
				ulidFlakeString: "R000000000000",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative value (this should never happen, but theoretically possible)",
			args: args{
				ulidFlakeString: "ZZZZZZZZZZZZZ",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromStr(tt.args.ulidFlakeString)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUnixEpochTime(t *testing.T) {
	type args struct {
		unixTimeSec int64
	}
	tests := []struct {
		name    string
		args    args
		want    *UlidFlake
		wantErr bool
	}{
		{
			name: "minimal value of Unix epoch time (before the custom epoch time of 2024-01-01T00:00:00Z)",
			args: args{
				unixTimeSec: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "minimal value of Unix epoch time (after the custom epoch time of 2024-01-01T00:00:00Z)",
			args: args{
				unixTimeSec: DefaultEpochSec,
			},
			want: &UlidFlake{
				value: 0,
			},
			wantErr: false,
		},
		{
			name: "maximal value of Unix epoch time (starting from the custom epoch time of 2024-01-01T00:00:00Z)",
			args: args{
				unixTimeSec: DefaultEpochSec + (1<<43-1)/1000,
			},
			want: &UlidFlake{
				value: 1<<43 - 1,
			},
			wantErr: false,
		},
		{
			name: "negative value of Unix epoch time (overflow, this should never happen, but theoretically possible)",
			args: args{
				unixTimeSec: DefaultEpochSec + (1<<43-1)/1000 + 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromUnixEpochTime(tt.args.unixTimeSec)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUnixEpochTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.GreaterOrEqual(t, got.String(), tt.want.String())
			}
		})
	}
}

func TestSetConfig(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				opts: []Option{
					WithEpochTime(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
					WithEntropySize(1),
					WithSID(0),
				},
			},
			wantErr: false,
		},
		{
			name: "too small entropy size",
			args: args{
				opts: []Option{
					WithEntropySize(0),
				},
			},
			wantErr: true,
		},
		{
			name: "too big entropy size",
			args: args{
				opts: []Option{
					WithEntropySize(3),
				},
			},
			wantErr: true,
		},
		{
			name: "too small scalability ID",
			args: args{
				opts: []Option{
					WithSID(-1),
				},
			},
			wantErr: true,
		},
		{
			name: "too big scalability ID",
			args: args{
				opts: []Option{
					WithSID(32),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetConfig(tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("SetConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		assert.Equal(t, int64(DefaultEpochSec), epochTime.Unix())
		assert.Equal(t, 1, entropySize)
		assert.Equal(t, int64(0), sid)
	}
}

func TestMonotonicallyIncreasingUlidFlake(t *testing.T) {
	ulidFlakeID, err := New()
	assert.Nil(t, err)
	for i := 0; i < 100; i++ {
		newUlidFlake, err := New()
		assert.Nil(t, err)
		assert.Greater(t, newUlidFlake.Int(), ulidFlakeID.Int())
		ulidFlakeID = newUlidFlake
	}
}
