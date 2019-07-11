package dough

import (
	"reflect"
	"testing"
)

func TestDesEncrypt(t *testing.T) {
	scenarios := []struct {
		src    string
		key    string
		result []byte
	}{
		{
			"random data",
			"12345678" + "12345678" + "12345678",
			[]byte{191, 8, 64, 227, 112, 217, 51, 234, 185, 64, 60, 119, 44, 171, 171, 161},
		},
		{
			"rand",
			"12345678" + "12345678" + "12345678",
			[]byte{70, 244, 153, 104, 131, 149, 75, 183},
		},
	}

	for i := 0; i < len(scenarios); i++ {
		result, err := DesEncrypt([]byte(scenarios[i].src), []byte(scenarios[i].key))
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(result, scenarios[i].result) {
			t.Errorf("Result should be %v, instead of %v in iteration %d", scenarios[i].result, result, i)
		}
	}

	_, err := DesEncrypt([]byte("test"), []byte("12345678"))
	if err == nil {
		t.Error("Error shouldn't be nil, because of length of key")
	}

	// NOTE: really hard to test ErrDesEncryptBlocksize because of PKCS5Padding
}

func TestDesDecrypt(t *testing.T) {
	scenarios := []struct {
		result    string
		key    string
		src []byte
	}{
		{
			"random data",
			"12345678" + "12345678" + "12345678",
			[]byte{191, 8, 64, 227, 112, 217, 51, 234, 185, 64, 60, 119, 44, 171, 171, 161},
		},
		{
			"rand",
			"12345678" + "12345678" + "12345678",
			[]byte{70, 244, 153, 104, 131, 149, 75, 183},
		},
	}

	for i := 0; i < len(scenarios); i++ {
		result, err := DesDecrypt(scenarios[i].src, []byte(scenarios[i].key))
		if err != nil {
			t.Error(err)
		}

		if string(result) != scenarios[i].result {
			t.Errorf("Result should be %s, instead of %s in iteration %d", scenarios[i].result, result, i)
		}
	}

	_, err := DesDecrypt([]byte{70, 244, 153, 104, 131, 149, 75, 183}, []byte("12345678"))
	if err == nil {
		t.Error("Error shouldn't be nil, because of length of key")
	}

	// NOTE: really hard to test ErrDesDecryptChiper
}

func TestZeroPadding(t *testing.T) {
	scenarios := []struct{
		chipertext string
		blockSize int
		result []byte
	} {
		{
			"test",
			16,
			[]byte{116, 101, 115, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
},
		{
			"testitwithotherdata",
			32,
			[]byte{116, 101, 115, 116, 105, 116, 119, 105, 116, 104, 111, 116, 104, 101, 114, 100, 97, 116, 97, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for i := 0; i < len(scenarios); i++ {
		result := ZeroPadding([]byte(scenarios[i].chipertext), scenarios[i].blockSize)

		if !reflect.DeepEqual(result, scenarios[i].result) {
			t.Errorf("Result should be %v, instead of %v in iteration %d", scenarios[i].result, result, i)
		}
	}
}

func TestZeroUnPadding(t *testing.T) {
	scenarios := []struct{
		result string
		origData []byte
	} {
		{
			"test",
			[]byte{116, 101, 115, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"testitwithotherdata",
			[]byte{116, 101, 115, 116, 105, 116, 119, 105, 116, 104, 111, 116, 104, 101, 114, 100, 97, 116, 97, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for i := 0; i < len(scenarios); i++ {
		result := ZeroUnPadding(scenarios[i].origData)

		if string(result) != scenarios[i].result {
			t.Errorf("Result should be %v, instead of %v in iteration %d", scenarios[i].result, result, i)
		}
	}
}