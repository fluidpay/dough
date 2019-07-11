package dough

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrCardLength = errors.New("card length should be > 10")
	ErrACHLength = errors.New("account length should be > 4")
	ErrUnknownCardType = errors.New("unknown card type")
)

var (
	amexCardFormatString                = "^3[47][0-9]{13}$"
	visaCardFormatString                = "^4[0-9]{15}$"
	mastercardCardFormatString          = "^(?:5[1-5][0-9]{2}|222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}$"
	discoverCardFormatString            = "^65[0-9]{14}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$"
	jcbCardFormatString                 = "^(?:2131|1800|35[0-9]{3})[0-9]{11}$"
	dinersClubInternationalFormatString = "^3(?:0[0-5]|[68][0-9])[0-9]{11}$"

	AMEXCardFormatRegex     = regexp.MustCompile(amexCardFormatString)
	VISACardFormatRegex     = regexp.MustCompile(visaCardFormatString)
	MasterCardFormatRegex   = regexp.MustCompile(mastercardCardFormatString)
	DISCOVERCardFormatRegex = regexp.MustCompile(discoverCardFormatString)
	JCBCardFormatRegex      = regexp.MustCompile(jcbCardFormatString)
	DINERSCardFormatRegex   = regexp.MustCompile(dinersClubInternationalFormatString)
)

// MaskCard takes in a card number and returns firstsix, lastfour, masked
func MaskCard(cardnumber string) (string, string, string, error) {
	length := len(cardnumber)
	if length < 10 {
		return "", "", "", ErrCardLength
	}
	maskLen := length - 10
	firstSix := cardnumber[:6]
	lastFour := cardnumber[length-4:]
	maskedCard := firstSix + strings.Repeat("*", maskLen) + lastFour

	return firstSix, lastFour, maskedCard, nil
}

// MaskACHAccount takes in an account number and returns masked
func MaskACHAccount(accountNumber string) (string, error) {
	length := len(accountNumber)
	if length < 4 {
		return "", ErrACHLength
	}
	maskLen := length - 4
	firstTwo := accountNumber[:2]
	lastTwo := accountNumber[length-2:]
	maskedAccount := firstTwo + strings.Repeat("*", maskLen) + lastTwo

	return maskedAccount, nil
}

// Valid returns a boolean indicating if the argument was valid according to the Luhn algorithm.
func ValidLuhn(luhnString string) bool {
	checksumMod := calculateChecksum(luhnString, false) % 10
	return checksumMod == 0
}

func calculateChecksum(luhnString string, double bool) int {
	source := strings.Split(luhnString, "")
	checksum := 0

	for i := len(source) - 1; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}

		checksum += n
	}

	return checksum
}

// GetCardType Accepts a string containing a credit card number and validates it against some regexes to return the card type.
func GetCardType(cardnum string) (string, error) {
	if AMEXCardFormatRegex.MatchString(cardnum) {
		return "amex", nil
	}

	if VISACardFormatRegex.MatchString(cardnum) {
		return "visa", nil
	}

	if MasterCardFormatRegex.MatchString(cardnum) {
		return "mastercard", nil
	}

	if DISCOVERCardFormatRegex.MatchString(cardnum) {
		return "discover", nil
	}

	if JCBCardFormatRegex.MatchString(cardnum) {
		return "jcb", nil
	}

	if DINERSCardFormatRegex.MatchString(cardnum) {
		return "diners", nil
	}

	return "", ErrUnknownCardType
}
