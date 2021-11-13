package dlms

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
)

func TestCipherData(t *testing.T) {
	ciphering := Ciphering{
		Security:          SecurityAuthenticationEncryption,
		SystemTitle:       decodeHexString("4D4D4D0000BC614E"),
		BlockCipherKey:    decodeHexString("000102030405060708090A0B0C0D0E0F"),
		AuthenticationKey: decodeHexString("D0D1D2D3D4D5D6D7D8D9DADBDCDDDEDF"),
		InvocationCounter: 0x01234567,
	}
	data := decodeHexString("01011000112233445566778899AABBCCDDEEFF0000065F1F0400007E1F04B0")
	result := decodeHexString("21303001234567801302FF8A7874133D414CED25B42534D28DB0047720606B175BD52211BE6841DB204D39EE6FDB8E356855")

	out, err := CipherData(&ciphering, data, TagGloInitiateRequest, false)
	if err != nil {
		t.Errorf("Got an error when ciphering: %v", err)
	}

	res := bytes.Compare(out, result)
	if res != 0 {
		t.Errorf("Failed. get: %s, should: %s", encodeHexString(out), encodeHexString(result))
	}
}

func TestCipherError(t *testing.T) {
	ciphering := Ciphering{}
	data := decodeHexString("01011000112233445566778899AABBCCDDEEFF0000065F1F0400007E1F04B0")

	_, err := CipherData(&ciphering, data, TagGloInitiateRequest, false)
	if err == nil {
		t.Errorf("Should get an error when ciphering")
	}
}

func decodeHexString(s string) []byte {
	b, _ := hex.DecodeString(s)
	return b
}

func encodeHexString(b []byte) string {
	return strings.ToUpper(hex.EncodeToString(b))
}
