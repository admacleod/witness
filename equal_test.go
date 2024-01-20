// Copyright (c) Alisdair MacLeod <copying@alisdairmacleod.co.uk>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
// OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

package witness_test

import (
	"testing"

	"github.com/admacleod/witness"
)

type mockT struct {
	helperCalled bool
	errorfCalled bool
	errorfFormat string
	errorfArgs   []any
}

func (mt *mockT) Helper() {
	mt.helperCalled = true
}

func (mt *mockT) Errorf(format string, args ...any) {
	mt.errorfCalled = true
	mt.errorfFormat = format
	mt.errorfArgs = args
}

func TestEqual(t *testing.T) {
	tests := map[string]struct {
		expect, actual bool
		callErrorf     bool
	}{
		"equal":     {expect: true, actual: true, callErrorf: false},
		"not equal": {expect: true, actual: false, callErrorf: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var mt mockT
			witness.Equal(&mt, test.expect, test.actual)
			if !mt.helperCalled {
				t.Errorf("test helper not registered")
			}
			switch {
			case !test.callErrorf:
			case !mt.errorfCalled:
				t.Errorf("expected errorf to be called but it was not")
			case mt.errorfFormat != witness.EqualFormat:
				t.Errorf("errorf called with incorrect format string: %q", mt.errorfFormat)
				fallthrough
			case len(mt.errorfArgs) < 2:
				t.Errorf("errorf not called with enough arguments: %#v", mt.errorfArgs)
			case mt.errorfArgs[0] != test.expect:
				t.Errorf("errorf not passed expect first: %#v", mt.errorfArgs)
				fallthrough
			case mt.errorfArgs[1] != test.actual:
				t.Errorf("errorf not passed actual second: %#v", mt.errorfArgs)
			}
		})
	}
}
