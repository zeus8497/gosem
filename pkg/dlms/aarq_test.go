package dlms

import (
	"bytes"
	"testing"
)

func TestEncodeAARQWithoutAuthentication(t *testing.T) {
	settings := &Settings{
		Authentication: AuthenticationNone,
		MaxPduSize:     512,
	}

	t1, err := EncodeAARQ(settings)
	if err != nil {
		t.Errorf("t1 Encode Failed. err: %v", err)
	}
	result := []byte{0x60, 0x1D, 0xA1, 0x09, 0x06, 0x07, 0x60, 0x85, 0x74, 0x05, 0x08, 0x01, 0x01, 0xBE, 0x10, 0x04, 0x0E, 0x01, 0x00, 0x00, 0x00, 0x06, 0x5F, 0x1F, 0x04, 0x00, 0x00, 0x18, 0x1F, 0x02, 0x00}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}

func TestEncodeAARQWithLowAuthentication(t *testing.T) {
	settings := &Settings{
		Authentication: AuthenticationLow,
		MaxPduSize:     256,
		Password:       []byte("12345678"),
	}

	t1, err := EncodeAARQ(settings)
	if err != nil {
		t.Errorf("t1 Encode Failed. err: %v", err)
	}
	result := []byte{0x60, 0x36, 0xA1, 0x09, 0x06, 0x07, 0x60, 0x85, 0x74, 0x05, 0x08, 0x01, 0x01, 0x8A, 0x02, 0x07, 0x80, 0x8B, 0x07, 0x60, 0x85, 0x74, 0x05, 0x08, 0x02, 0x01, 0xAC, 0x0A, 0x80, 0x08, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0xBE, 0x10, 0x04, 0x0E, 0x01, 0x00, 0x00, 0x00, 0x06, 0x5F, 0x1F, 0x04, 0x00, 0x00, 0x18, 0x1F, 0x01, 0x00}
	res := bytes.Compare(t1, result)
	if res != 0 {
		t.Errorf("t1 Failed. get: %d, should:%v", t1, result)
	}
}
