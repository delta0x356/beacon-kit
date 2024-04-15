// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

//nolint:lll // long hex strings.
package primitives_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/berachain/beacon-kit/mod/primitives"
)

func TestBytes4UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    primitives.Bytes4
		wantErr bool
	}{
		{
			name:  "valid input",
			input: `"0x01020304"`,
			want:  primitives.Bytes4{0x01, 0x02, 0x03, 0x04},
		},
		{
			name:    "invalid input - not hex",
			input:   `"01020304"`,
			wantErr: true,
		},
		{
			name:    "invalid input - wrong length",
			input:   `"0x010203"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got primitives.Bytes4
			err := got.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes4.UnmarshalJSON() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes4.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes4String(t *testing.T) {
	tests := []struct {
		name string
		h    primitives.Bytes4
		want string
	}{
		{
			name: "non-empty bytes",
			h:    primitives.Bytes4{0x01, 0x02, 0x03, 0x04},
			want: "0x01020304",
		},
		{
			name: "empty bytes",
			h:    primitives.Bytes4{},
			want: "0x00000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.String(); got != tt.want {
				t.Errorf("Bytes4.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes4MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		h       primitives.Bytes4
		want    string
		wantErr bool
	}{
		{
			name: "valid bytes",
			h:    primitives.Bytes4{0x01, 0x02, 0x03, 0x04},
			want: "0x01020304",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes4.MarshalText() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if string(got) != tt.want {
				t.Errorf(
					"Bytes4.MarshalText() = %v, want %v",
					string(got),
					tt.want,
				)
			}
		})
	}
}

func TestBytes4UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    primitives.Bytes4
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "0x01020304",
			want:  primitives.Bytes4{0x01, 0x02, 0x03, 0x04},
		},
		{
			name:    "invalid input - not hex",
			input:   "01020304",
			wantErr: true,
		},
		{
			name:    "invalid input - wrong length",
			input:   "0x010203",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got primitives.Bytes4
			err := got.UnmarshalText([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes4.UnmarshalText() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes4.UnmarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBytes32UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    primitives.Bytes32
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20",
			want: primitives.Bytes32{
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
				0x08,
				0x09,
				0x0a,
				0x0b,
				0x0c,
				0x0d,
				0x0e,
				0x0f,
				0x10,
				0x11,
				0x12,
				0x13,
				0x14,
				0x15,
				0x16,
				0x17,
				0x18,
				0x19,
				0x1a,
				0x1b,
				0x1c,
				0x1d,
				0x1e,
				0x1f,
				0x20,
			},
		},
		{
			name:    "invalid input - not hex",
			input:   "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20",
			wantErr: true,
		},
		{
			name:    "invalid input - wrong length",
			input:   "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got primitives.Bytes32
			err := got.UnmarshalText([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes32.UnmarshalText() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes32.UnmarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes32UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    primitives.Bytes32
		wantErr bool
	}{
		{
			name:  "valid input",
			input: `"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"`,
			want: primitives.Bytes32{
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
				0x08,
				0x09,
				0x0a,
				0x0b,
				0x0c,
				0x0d,
				0x0e,
				0x0f,
				0x10,
				0x11,
				0x12,
				0x13,
				0x14,
				0x15,
				0x16,
				0x17,
				0x18,
				0x19,
				0x1a,
				0x1b,
				0x1c,
				0x1d,
				0x1e,
				0x1f,
				0x20,
			},
		},
		{
			name:    "invalid input - not hex",
			input:   `"0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"`,
			wantErr: true,
		},
		{
			name:    "invalid input - wrong length",
			input:   `"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"`,
			wantErr: true,
		},
		{
			name:    "invalid input - extra characters",
			input:   `"0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got primitives.Bytes32
			err := json.Unmarshal([]byte(tt.input), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes32.UnmarshalJSON() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes32.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes32MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   primitives.Bytes32
		want    string
		wantErr bool
	}{
		{
			name: "valid input",
			input: primitives.Bytes32{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
				0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12,
				0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d,
				0x1e, 0x1f, 0x20},
			want: "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20",
		},
		{
			name:  "empty input",
			input: primitives.Bytes32{},
			want:  "0x0000000000000000000000000000000000000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes32.MarshalText() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if string(got) != tt.want {
				t.Errorf(
					"Bytes32.MarshalText() = %v, want %v",
					string(got),
					tt.want,
				)
			}
		})
	}
}

func TestBytes32String(t *testing.T) {
	tests := []struct {
		name  string
		input primitives.Bytes32
		want  string
	}{
		{
			name: "valid input",
			input: primitives.Bytes32{0x01, 0x02, 0x03, 0x04, 0x05, 0x06,
				0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
				0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a,
				0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20},
			want: "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20",
		},
		{
			name:  "empty input",
			input: primitives.Bytes32{},
			want:  "0x0000000000000000000000000000000000000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.String(); got != tt.want {
				t.Errorf("Bytes32.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBytes48String(t *testing.T) {
	tests := []struct {
		name  string
		input primitives.Bytes48
		want  string
	}{
		{
			name: "valid input",
			input: primitives.Bytes48{
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
				0x08,
				0x09,
				0x0a,
				0x0b,
				0x0c,
				0x0d,
				0x0e,
				0x0f,
				0x10,
				0x11,
				0x12,
				0x13,
				0x14,
				0x15,
				0x16,
				0x17,
				0x18,
				0x19,
				0x1a,
				0x1b,
				0x1c,
				0x1d,
				0x1e,
				0x1f,
				0x20,
				0x21,
				0x22,
				0x23,
				0x24,
				0x25,
				0x26,
				0x27,
				0x28,
				0x29,
				0x2a,
				0x2b,
				0x2c,
				0x2d,
				0x2e,
				0x2f,
				0x30,
			},
			want: "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30",
		},
		{
			name:  "empty input",
			input: primitives.Bytes48{},
			want:  "0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.String(); got != tt.want {
				t.Errorf("Bytes48.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes48MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   primitives.Bytes48
		want    string
		wantErr bool
	}{
		{
			name: "valid input",
			input: primitives.Bytes48{
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
				0x08,
				0x09,
				0x0a,
				0x0b,
				0x0c,
				0x0d,
				0x0e,
				0x0f,
				0x10,
				0x11,
				0x12,
				0x13,
				0x14,
				0x15,
				0x16,
				0x17,
				0x18,
				0x19,
				0x1a,
				0x1b,
				0x1c,
				0x1d,
				0x1e,
				0x1f,
				0x20,
				0x21,
				0x22,
				0x23,
				0x24,
				0x25,
				0x26,
				0x27,
				0x28,
				0x29,
				0x2a,
				0x2b,
				0x2c,
				0x2d,
				0x2e,
				0x2f,
				0x30,
			},
			want: "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30",
		},
		{
			name:  "empty input",
			input: primitives.Bytes48{},
			want:  "0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes48.MarshalText() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if string(got) != tt.want {
				t.Errorf(
					"Bytes48.MarshalText() = %s, want %s",
					string(got),
					tt.want,
				)
			}
		})
	}
}

func TestBytes48UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    primitives.Bytes48
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30",
			want: primitives.Bytes48{
				0x01,
				0x02,
				0x03,
				0x04,
				0x05,
				0x06,
				0x07,
				0x08,
				0x09,
				0x0a,
				0x0b,
				0x0c,
				0x0d,
				0x0e,
				0x0f,
				0x10,
				0x11,
				0x12,
				0x13,
				0x14,
				0x15,
				0x16,
				0x17,
				0x18,
				0x19,
				0x1a,
				0x1b,
				0x1c,
				0x1d,
				0x1e,
				0x1f,
				0x20,
				0x21,
				0x22,
				0x23,
				0x24,
				0x25,
				0x26,
				0x27,
				0x28,
				0x29,
				0x2a,
				0x2b,
				0x2c,
				0x2d,
				0x2e,
				0x2f,
				0x30,
			},
		},
		{
			name:    "invalid input - not hex",
			input:   "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f30",
			wantErr: true,
		},
		{
			name:    "invalid input - wrong length",
			input:   "0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got primitives.Bytes48
			err := got.UnmarshalText([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes48.UnmarshalText() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes48.UnmarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBytes96UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    primitives.Bytes96
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "0x" + strings.Repeat("01", 96),
			want: func() primitives.Bytes96 {
				var b primitives.Bytes96
				for i := range b {
					b[i] = 0x01
				}
				return b
			}(),
		},
		{
			name:    "invalid input - not hex",
			input:   strings.Repeat("01", 96),
			wantErr: true,
		},
		{
			name:    "invalid input - wrong length",
			input:   "0x" + strings.Repeat("01", 95),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got primitives.Bytes96
			err := got.UnmarshalText([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes96.UnmarshalText() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes96.UnmarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes96UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    primitives.Bytes96
		wantErr bool
	}{
		{
			name:  "valid input",
			input: `"0x` + strings.Repeat("01", 96) + `"`,
			want: func() primitives.Bytes96 {
				var b primitives.Bytes96
				for i := range b {
					b[i] = 0x01
				}
				return b
			}(),
		},
		{
			name:    "invalid input - not hex",
			input:   `"` + strings.Repeat("01", 96) + `"`,
			wantErr: true,
		},
		{
			name:    "invalid input - wrong length",
			input:   `"0x` + strings.Repeat("01", 95) + `"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got primitives.Bytes96
			err := got.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes96.UnmarshalJSON() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes96.UnmarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBytes96MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		h       primitives.Bytes96
		want    string
		wantErr bool
	}{
		{
			name: "valid bytes",
			h: func() primitives.Bytes96 {
				var b primitives.Bytes96
				for i := range b {
					b[i] = 0x01
				}
				return b
			}(),
			want: "0x" + strings.Repeat("01", 96),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"Bytes96.MarshalText() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if string(got) != tt.want {
				t.Errorf(
					"Bytes96.MarshalText() = %v, want %v",
					string(got),
					tt.want,
				)
			}
		})
	}
}

func TestBytes96String(t *testing.T) {
	tests := []struct {
		name string
		h    primitives.Bytes96
		want string
	}{
		{
			name: "non-empty bytes",
			h: func() primitives.Bytes96 {
				var b primitives.Bytes96
				for i := range b {
					b[i] = 0x01
				}
				return b
			}(),
			want: "0x" + strings.Repeat("01", 96),
		},
		{
			name: "empty bytes",
			h:    primitives.Bytes96{},
			want: "0x" + strings.Repeat("00", 96),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.String(); got != tt.want {
				t.Errorf("Bytes96.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
