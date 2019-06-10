package genpass

import "testing"

func TestDecodePassword(t *testing.T) {
	err := DecodePassword()
	if err != nil {
		t.Errorf("unexpected error")
	}
}
