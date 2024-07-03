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
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package bytes

import (
	"reflect"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/hex"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/json"
)

//nolint:gochecknoglobals // reflect.Type of Bytes set at runtime
var bytesT = reflect.TypeOf(Bytes(nil))

// Bytes marshals/unmarshals as a JSON string with 0x prefix.
// The empty slice marshals as "0x".
type Bytes []byte

// MarshalText implements encoding.TextMarshaler.
func (b Bytes) MarshalText() ([]byte, error) {
	return hex.EncodeBytes(b)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (b *Bytes) UnmarshalText(input []byte) error {
	dec, err := hex.UnmarshalByteText(input)
	if err != nil {
		return err
	}
	*b = Bytes(dec)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Bytes) UnmarshalJSON(input []byte) error {
	return json.UnmarshalJSONText(input, b, bytesT)
}

// String returns the hex encoding of b.
func (b Bytes) String() hex.String {
	return hex.FromBytes(b)
}