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

//go:build tools
// +build tools

// This is the canonical way to enforce dependency inclusion in go.mod for tools
// that are not directly involved in the build process.
// See
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package main

import (
	_ "github.com/bazelbuild/buildtools/buildifier"
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/cosmos/gosec/v2/cmd/gosec"
	_ "github.com/ethereum/go-ethereum/cmd/abigen"
	_ "github.com/ferranbt/fastssz/sszgen"
	_ "github.com/fjl/gencodec"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/addlicense"
	_ "github.com/segmentio/golines"
	_ "github.com/vektra/mockery/v2"
	_ "go.uber.org/nilaway/cmd/nilaway"
)
