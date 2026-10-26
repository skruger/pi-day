package main

import "testing"

func TestPiDigits(t *testing.T) {
	want := "3141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067"
	for _, n := range []int{1, 2, 10, 50, 99} {
		if got := PiDigits(n); got != want[:n] {
			t.Errorf("n=%d: got %s want %s", n, got, want[:n])
		}
	}
}
