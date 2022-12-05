/**
 * Author: Rolly Maulana Awangga
 * File: gogis_test.go
 */

package gogis

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	expected := a + b

	if got := Add(a, b); got != expected {
		t.Errorf("Add(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestSubtract(t *testing.T) {
	a := 1
	b := 2
	expected := a - b

	if got := Subtract(a, b); got != expected {
		t.Errorf("Subtract(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestGetPublicIP(t *testing.T) {
	notfound := -1
	if got := strings.Index(GetPublicIP(), "{"); got <= notfound {
		t.Errorf("Response Body : %d, didn't return json", got)
	}

}
