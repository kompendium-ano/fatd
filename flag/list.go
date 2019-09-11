// MIT License
//
// Copyright 2018 Canonical Ledgers, LLC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package flag

import (
	"strings"

	"github.com/Factom-Asset-Tokens/factom"
)

type FAAddressList []factom.FAAddress

func (adrs FAAddressList) String() string {
	if len(adrs) == 0 {
		return ""
	}
	var s string
	for _, adr := range adrs {
		s += adr.String() + ","
	}
	return s[:len(s)-1]
}

// Set appends a comma seperated list of factom.FAAddresses.
func (adrs *FAAddressList) Set(s string) error {
	adrStrs := strings.Split(s, ",")
	newAdrs := make(FAAddressList, len(adrStrs))
	for i, adrStr := range adrStrs {
		if err := newAdrs[i].Set(adrStr); err != nil {
			return err
		}
	}
	*adrs = append(*adrs, newAdrs...)
	return nil
}

type Bytes32List []factom.Bytes32

func (b32s Bytes32List) String() string {
	if len(b32s) == 0 {
		return ""
	}
	var s string
	for _, b32 := range b32s {
		s += b32.String() + ","
	}
	return s[:len(s)-1]
}

// Set appends a comma seperated list of factom.FAAddresses.
func (b32s *Bytes32List) Set(s string) error {
	b32Strs := strings.Split(s, ",")
	newB32s := make(Bytes32List, len(b32Strs))
	for i, b32Str := range b32Strs {
		if err := newB32s[i].Set(b32Str); err != nil {
			return err
		}
	}
	*b32s = append(*b32s, newB32s...)
	return nil
}
