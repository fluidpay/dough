package dough

import "testing"

var testCards = map[string]string{
	"4111111111111111":    "visa",
	"4457010000000009":    "visa",
	"4457010100000008":    "visa",
	"4457010140000141":    "visa",
	"4457010200000247":    "visa",
	"4100200300011001":    "visa",
	"4100200300012009":    "visa",
	"4100200300013007":    "visa",
	"4100200310000002":    "visa",
	"4024720001231239":    "visa",
	"4457012400000001":    "visa",
	"4457013200000001":    "visa",
	"4457119922390123":    "visa",
	"4457000300000007":    "visa",
	"4457000100000009":    "visa",
	"4457003100000003":    "visa",
	"4457000400000006":    "visa",
	"4457000200000008":    "visa",
	"4457000800000002":    "visa",
	"4457000900000001":    "visa",
	"4457001000000008":    "visa",
	"4005550000081019":    "visa",
	"4000000000000002":    "visa",
	"2223000148400010":    "mastercard",
	"2223000048400011":    "mastercard",
	"2223280062080010":    "mastercard",
	"2222630061560019":    "mastercard",
	"2222470061880012":    "mastercard",
	"2222400061240016":    "mastercard",
	"2222400041240011":    "mastercard",
	"2223520063560019":    "mastercard",
	"2223520043560014":    "mastercard",
	"2222420040560011":    "mastercard",
	"2222410040360017":    "mastercard",
	"2223020040760014":    "mastercard",
	"5112000100000003":    "mastercard",
	"5112002100000009":    "mastercard",
	"5112002200000008":    "mastercard",
	"5112000200000002":    "mastercard",
	"5112000300000001":    "mastercard",
	"5112000400000000":    "mastercard",
	"5112010400000009":    "mastercard",
	"5112000600000008":    "mastercard",
	"5112010000000003":    "mastercard",
	"5112010100000002":    "mastercard",
	"5112010140000004":    "mastercard",
	"5154605300000121":    "mastercard",
	"5167001020236549":    "mastercard",
	"5500000254444445":    "mastercard",
	"5592106621450897":    "mastercard",
	"5590409551104142":    "mastercard",
	"5587755665222179":    "mastercard",
	"5445840176552850":    "mastercard",
	"5390016478904678":    "mastercard",
	"5112010201000109":    "mastercard",
	"5112010202000108":    "mastercard",
	"5194560012341234":    "mastercard",
	"5435101234510196":    "mastercard",
	"5407102010000018":    "mastercard",
	"5112000900000005":    "mastercard",
	"6759649826438453":    "mastercard",
	"6011010000000003":    "discover",
	"6011010100000002":    "discover",
	"6011010140000004":    "discover",
	"6011010000000003011": "discover",
	"6011010100000002011": "discover",
	"6011010140000000011": "discover",
	"375000026600004":     "amex",
	"375001000000005":     "amex",
	"375001010000003":     "amex",
	"375001014000009":     "amex",
	"341234567890127":     "amex",
	"378734493671000":     "amex",
	"3530111333300000":    "jcb",
	"3566002020360505":    "jcb",
	"3530111333300000332": "jcb",
	"3566002020360505005": "jcb",
	"30569309025904":      "diners",
	"38520000023237":      "diners",
	"3056930902590411014": "diners",
	"3852000002323711017": "diners",
}

func TestValidLuhn(t *testing.T) {
	for key := range testCards {
		if !ValidLuhn(key) {
			t.Errorf("Should be valid %s", key)
		}
	}
}

func TestMaskCard(t *testing.T) {
	var testMaskableCards = []struct {
		cardNumber string
		firstSix   string
		lastFour   string
		maskedCard string
	}{
		{
			"4111111111111111",
			"411111",
			"1111",
			"411111******1111",
		},
		{
			"4457010000000009",
			"445701",
			"0009",
			"445701******0009",
		},
		{
			"2223000148400010",
			"222300",
			"0010",
			"222300******0010",
		},
		{
			"6011010000000003",
			"601101",
			"0003",
			"601101******0003",
		},
	}

	for i := 0; i < len(testMaskableCards); i++ {
		firstSix, lastFour, maskedCard, err := MaskCard(testMaskableCards[i].cardNumber)
		if err != nil {
			t.Error(err)
		}

		if firstSix != testMaskableCards[i].firstSix {
			t.Errorf("First six should be %s, instead of %s in iteration %d", testMaskableCards[i].firstSix, firstSix, i)
		}

		if lastFour != testMaskableCards[i].lastFour {
			t.Errorf("Last four should be %s, instead of %s in iteration %d", testMaskableCards[i].lastFour, lastFour, i)
		}

		if maskedCard != testMaskableCards[i].maskedCard {
			t.Errorf("Mask card should be %s, instead of %s in iteration %d", testMaskableCards[i].maskedCard, maskedCard, i)
		}
	}

	_, _, _, err := MaskCard("12345")
	if err != ErrCardLength {
		t.Errorf("Error should be %s", ErrCardLength.Error())
	}
}

func TestMaskACHAccount(t *testing.T) {
	// 8114460248
	var testMaskableAchs = []struct {
		accountNumber       string
		maskedAccountNumber string
	}{
		{
			"8114460248",
			"81******48",
		},
		{
			"1113445648",
			"11******48",
		},
		{
			"4315460748",
			"43******48",
		},
	}

	for i := 0; i < len(testMaskableAchs); i++ {
		maskedAccountNumber, err := MaskACHAccount(testMaskableAchs[i].accountNumber)
		if err != nil {
			t.Error(err)
		}

		if maskedAccountNumber != testMaskableAchs[i].maskedAccountNumber {
			t.Errorf("Mask account number should be %s, instead of %s in iteration %d", testMaskableAchs[i].maskedAccountNumber, maskedAccountNumber, i)
		}
	}

	_, err := MaskACHAccount("123")
	if err != ErrACHLength {
		t.Errorf("Error should be %s", ErrACHLength.Error())
	}
}

//TestInvalidCountryCode
func TestValidCardTypes(t *testing.T) {
	for key, value := range testCards {
		val, err := GetCardType(key)
		if err != nil {
			t.Error(err)
		}
		if val != value {
			t.Errorf("%s Expected %s received %s", key, value, val)
		}
	}

	_, err := GetCardType("1111111111111111")
	if err != ErrUnknownCardType {
		t.Errorf("Error should be %s", ErrUnknownCardType.Error())
	}
}
