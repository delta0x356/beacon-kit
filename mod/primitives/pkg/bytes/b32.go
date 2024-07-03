// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.
//
//nolint:dupl // it's okay to have similar code for different types
package bytes

import (
	"github.com/berachain/beacon-kit/mod/primitives/pkg/hex"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/ssz/types"
)

const (
	// B32Size represents a 32-byte size.
	B32Size = 32
)

var _ types.MinimalSSZType = (*B32)(nil)

// B32 represents a 32-byte fixed-size byte array.
// For SSZ purposes it is serialized a `Vector[Byte, 32]`.
type B32 [32]byte

// ToBytes32 is a utility function that transforms a byte slice into a fixed
// 32-byte array. If the input exceeds 32 bytes, it gets truncated.
func ToBytes32(input []byte) B32 {
	return B32(ExtendToSize(input, B32Size))
}

/* -------------------------------------------------------------------------- */
/*                                TextMarshaler                               */
/* -------------------------------------------------------------------------- */

// MarshalText implements the encoding.TextMarshaler interface for B32.
func (h B32) MarshalText() ([]byte, error) {
	return []byte("0x" + hex.EncodeToString(h[:])), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for B32.
func (h *B32) UnmarshalText(text []byte) error {
	return UnmarshalTextHelper(h[:], text)
}

/* -------------------------------------------------------------------------- */
/*                                JSONMarshaler                               */
/* -------------------------------------------------------------------------- */

// MarshalJSON implements the json.Marshaler interface for B32.
func (h B32) MarshalJSON() ([]byte, error) {
	return hex.EncodeFixedJSON(h[:]), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for B32.
func (h *B32) UnmarshalJSON(input []byte) error {
	return hex.DecodeFixedJSON(input, h[:])
}

/* -------------------------------------------------------------------------- */
/*                                SSZMarshaler                                */
/* -------------------------------------------------------------------------- */

// SizeSSZ returns the size of its SSZ encoding in bytes.
func (h B32) SizeSSZ() int {
	return B32Size
}

// MarshalSSZ implements the SSZ marshaling for B32.
func (h B32) MarshalSSZ() ([]byte, error) {
	return h[:], nil
}

// IsFixed returns true if the length of the B32 is fixed.
func (h B32) IsFixed() bool {
	return true
}

// Type returns the type of the B32.
func (h B32) Type() types.Type {
	return types.Composite
}

// HashTreeRoot returns the hash tree root of the B32.
func (h B32) HashTreeRoot() ([32]byte, error) {
	var result [32]byte
	copy(result[:], h[:])
	return result, nil
}
