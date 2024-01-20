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

package witness

// T is an interface that abstracts a `*testing.T` to simplify testing of this package.
type T interface {
	Helper()
	Errorf(string, ...any)
}

// EqualFormat is the format string used when raising errors.
const EqualFormat = "incorrect value:\nexpect=%#v\nactual=%#v\n"

// EqualFn compares the two passed values using the provided comparison function. If the result of the comparison
// function is false then a test error will be raised.
func EqualFn[Type any](t T, expect, actual Type, fn func(Type, Type) bool) {
	t.Helper()

	if !fn(expect, actual) {
		t.Errorf(EqualFormat, expect, actual)
	}
}

// Equal compares two comparable values using a simple equality function. If the values are not equal then a test error
// will be raised.
func Equal[Type comparable](t T, expect, actual Type) {
	t.Helper()

	EqualFn(t, expect, actual, func(e, a Type) bool { return e == a })
}
