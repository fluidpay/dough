package dough

import (
	"regexp"
	"strings"
)

const (
	postalCodeGBString = "^GIR[ ]?0AA|((AB|AL|B|BA|BB|BD|BH|BL|BN|BR|BS|BT|CA|CB|CF|CH|CM|CO|CR|CT|CV|CW|DA|DD|DE|DG|DH|DL|DN|DT|DY|E|EC|EH|EN|EX|FK|FY|G|GL|GY|GU|HA|HD|HG|HP|HR|HS|HU|HX|IG|IM|IP|IV|JE|KA|KT|KW|KY|L|LA|LD|LE|LL|LN|LS|LU|M|ME|MK|ML|N|NE|NG|NN|NP|NR|NW|OL|OX|PA|PE|PH|PL|PO|PR|RG|RH|RM|S|SA|SE|SG|SK|SL|SM|SN|SO|SP|SR|SS|ST|SW|SY|TA|TD|TF|TN|TQ|TR|TS|TW|UB|W|WA|WC|WD|WF|WN|WR|WS|WV|YO|ZE)(\\d[\\dA-Z]?[ ]?\\d[ABD-HJLN-UW-Z]{2}))|BFPO[ ]?\\d{1,4}$"
	postalCodeJEString = "^JEd[\\dA-Z]?[ ]?\\d[ABD-HJLN-UW-Z]{2}$"
	postalCodeGGString = "^GYd[\\dA-Z]?[ ]?\\d[ABD-HJLN-UW-Z]{2}$"
	postalCodeIMString = "^IMd[\\dA-Z]?[ ]?\\d[ABD-HJLN-UW-Z]{2}$"
	postalCodeUSString = "^\\d{5}([ -]\\d{4})?$"
	postalCodeCAString = "^[ABCEGHJKLMNPRSTVXY]\\d[ABCEGHJ-NPRSTV-Z][ ]?\\d[ABCEGHJ-NPRSTV-Z]\\d$"
	postalCodeDEString = "^\\d{5}$"
	postalCodeJPString = "^\\d{3}-\\d{4"
	postalCodeFRString = "^\\d{2}[ ]?\\d{3}$"
	postalCodeAUString = "^\\d{4}$"
	postalCodeITString = "^\\d{5}$"
	postalCodeCHString = "^\\d{4}$"
	postalCodeATString = "^\\d{4}$"
	postalCodeESString = "^\\d{5}$"
	postalCodeNLString = "^\\d{4}[ ]?[A-Z]{2}$"
	postalCodeBEString = "^\\d{4}$"
	postalCodeDKString = "^\\d{4}$"
	postalCodeSEString = "^\\d{3}[ ]?\\d{2}$"
	postalCodeNOString = "^\\d{4}$"
	postalCodeBRString = "^\\d{5}[-]?\\d{3}$"
	postalCodePTString = "^\\d{4}([-]\\d{3})?$"
	postalCodeFIString = "^\\d{5}$"
	postalCodeAXString = "^22\\d{3}$"
	postalCodeKRString = "^\\d{3}[-]\\d{3}$"
	postalCodeCNString = "^\\d{6}$"
	postalCodeTWString = "^\\d{3}(\\d{2})?$"
	postalCodeSGString = "^\\d{6}$"
	postalCodeDZString = "^\\d{5}$"
	postalCodeADString = "^AD\\d{3}$"
	postalCodeARString = "^([A-HJ-NP-Z])?\\d{4}([A-Z]{3})?$"
	postalCodeAMString = "^(37)?\\d{4}$"
	postalCodeAZString = "^\\d{4}$"
	postalCodeBHString = "^((1[0-2]|[2-9])\\d{2})?$"
	postalCodeBDString = "^\\d{4}$"
	postalCodeBBString = "^(BB\\d{5})?$"
	postalCodeBYString = "^\\d{6}$"
	postalCodeBMString = "^[A-Z]{2}[ ]?[A-Z0-9]{2}$"
	postalCodeBAString = "^\\d{5}$"
	postalCodeIOString = "^BBND 1ZZ$"
	postalCodeBNString = "^[A-Z]{2}[ ]?\\d{4}$"
	postalCodeBGString = "^\\d{4}$"
	postalCodeKHString = "^\\d{5}$"
	postalCodeCVString = "^\\d{4}$"
	postalCodeCLString = "^\\d{7}$"
	postalCodeCRString = "^\\d{4,5}"
	postalCodeHRString = "^\\d{5}$"
	postalCodeCYString = "^\\d{4}$"
	postalCodeCZString = "^\\d{3}[ ]?d{2}$"
	postalCodeDOString = "^\\d{5}$"
	postalCodeECString = "^([A-Z]\\d{4}[A-Z]|(?:[A-Z]{2})?\\d{6})?$"
	postalCodeEGString = "^\\d{5}$"
	postalCodeEEString = "^\\d{5}$"
	postalCodeFOString = "^\\d{3}$"
	postalCodeGEString = "^\\d{4}$"
	postalCodeGRString = "^\\d{3}[ ]?\\d{2}$"
	postalCodeGLString = "^39\\d{2}$"
	postalCodeGTString = "^\\d{5}$"
	postalCodeHTString = "^\\d{4}$"
	postalCodeHNString = "^(?:\\d{5})?$"
	postalCodeHUString = "^\\d{4}$"
	postalCodeISString = "^\\d{3}$"
	postalCodeINString = "^\\d{6}$"
	postalCodeIDString = "^\\d{5}$"
	postalCodeILString = "^\\d{5}$"
	postalCodeJOString = "^\\d{5}$"
	postalCodeKZString = "^\\d{6}$"
	postalCodeKEString = "^\\d{5}$"
	postalCodeKWString = "^\\d{5}$"
	postalCodeLAString = "^\\d{5}$"
	postalCodeLVString = "^\\d{4}$"
	postalCodeLBString = "^(\\d{4}([ ]?\\d{4})?)?$"
	postalCodeLIString = "^(948[5-9])"
	postalCodeLTString = "^\\d{5}$"
	postalCodeLUString = "^\\d{4}$"
	postalCodeMKString = "^\\d{4}$"
	postalCodeMYString = "^\\d{5}$"
	postalCodeMVString = "^\\d{5}$"
	postalCodeMTString = "^[A-Z]{3}[ ]?\\d{2,4}$"
	postalCodeMUString = "^(\\d{3}[A-Z]{2}\\d{3})?$"
	postalCodeMXString = "^\\d{5}$"
	postalCodeMDString = "^\\d{4}$"
	postalCodeMCString = "^980d{2}$"
	postalCodeMAString = "^\\d{5}$"
	postalCodeNPString = "^\\d{5}$"
	postalCodeNZString = "^\\d{4}$"
	postalCodeNIString = "^((\\d{4}-)?\\d{3}-\\d{3}(-\\d{1})?)?$"
	postalCodeNGString = "^(\\d{6})?$"
	postalCodeOMString = "^(PC )?\\d{3}$"
	postalCodePKString = "^\\d{5}$"
	postalCodePYString = "^\\d{4}$"
	postalCodePHString = "^\\d{4}$"
	postalCodePLString = "^\\d{2}-\\d{3}$"
	postalCodePRString = "^00[679]\\d{2}([ -]\\d{4})?$"
	postalCodeROString = "^\\d{6}$"
	postalCodeRUString = "^\\d{6}$"
	postalCodeSMString = "^4789\\d$"
	postalCodeSAString = "^\\d{5}$"
	postalCodeSNString = "^\\d{5}$"
	postalCodeSKString = "^\\d{3}[ ]?\\d{2}$"
	postalCodeSIString = "^\\d{4}$"
	postalCodeZAString = "^\\d{4}$"
	postalCodeLKString = "^\\d{5}$"
	postalCodeTJString = "^\\d{6}$"
	postalCodeTHString = "^\\d{5}$"
	postalCodeTNString = "^\\d{4}$"
	postalCodeTRString = "^\\d{5}$"
	postalCodeTMString = "^\\d{6}$"
	postalCodeUAString = "^\\d{5}$"
	postalCodeUYString = "^\\d{5}$"
	postalCodeUZString = "^\\d{6}$"
	postalCodeVAString = "^00120$"
	postalCodeVEString = "^\\d{4}$"
	postalCodeZMString = "^\\d{5}$"
	postalCodeASString = "^96799$"
	postalCodeCCString = "^6799$"
	postalCodeCKString = "^\\d{4}$"
	postalCodeRSString = "^\\d{6}$"
	postalCodeMEString = "^8\\d{4}$"
	postalCodeCSString = "^\\d{5}$"
	postalCodeYUString = "^\\d{5}$"
	postalCodeCXString = "^6798$"
	postalCodeETString = "^\\d{4}$"
	postalCodeFKString = "^FIQQ 1ZZ$"
	postalCodeNFString = "^2899$"
	postalCodeFMString = "^(9694[1-4])([ -]\\d{4})?$"
	postalCodeGFString = "^9[78]3\\d{2}$"
	postalCodeGNString = "^\\d{3}$"
	postalCodeGPString = "^9[78][01]\\d{2}$"
	postalCodeGSString = "^SIQQ 1ZZ$"
	postalCodeGUString = "^969[123]\\d([ -]\\d{4})?$"
	postalCodeGWString = "^\\d{4}$"
	postalCodeHMString = "^\\d{4}$"
	postalCodeIQString = "^\\d{5}$"
	postalCodeKGString = "^\\d{6}$"
	postalCodeLRString = "^\\d{4}$"
	postalCodeLSString = "^\\d{3}$"
	postalCodeMGString = "^\\d{3}$"
	postalCodeMHString = "^969[67]\\d([ -]\\d{4})?$"
	postalCodeMNString = "^\\d{6}$"
	postalCodeMPString = "^9695[012]([ -]\\d{4})?$"
	postalCodeMQString = "^9[78]2\\d{2}$"
	postalCodeNCString = "^988\\d{2}$"
	postalCodeNEString = "^\\d{4}$"
	postalCodeVIString = "^008(([0-4]\\d)|(5[01]))([ \\-]\\d{4})?$"
	postalCodePFString = "^987\\d{2}$"
	postalCodePGString = "^\\d{3}$"
	postalCodePMString = "^9[78]5\\d{2}$"
	postalCodePNString = "^PCRN 1ZZ$"
	postalCodePWString = "^96940$"
	postalCodeREString = "^9[78]4\\d{2}$"
	postalCodeSHString = "^(ASCN|STHL) 1ZZ$"
	postalCodeSJString = "^\\d{4}$"
	postalCodeSOString = "^\\d{5}$"
	postalCodeSZString = "^[HLMS]\\d{3}$"
	postalCodeTCString = "^TKCA 1ZZ$"
	postalCodeWFString = "^986\\d{2}$"
	postalCodeXKString = "^\\d{5}$"
	postalCodeYTString = "^976\\d{2}$i"
)

var (
	PostalCodeGBRegex = regexp.MustCompile(postalCodeGBString)
	PostalCodeJERegex = regexp.MustCompile(postalCodeJEString)
	PostalCodeGGRegex = regexp.MustCompile(postalCodeGGString)
	PostalCodeIMRegex = regexp.MustCompile(postalCodeIMString)
	PostalCodeUSRegex = regexp.MustCompile(postalCodeUSString)
	PostalCodeCARegex = regexp.MustCompile(postalCodeCAString)
	PostalCodeDERegex = regexp.MustCompile(postalCodeDEString)
	PostalCodeJPRegex = regexp.MustCompile(postalCodeJPString)
	PostalCodeFRRegex = regexp.MustCompile(postalCodeFRString)
	PostalCodeAURegex = regexp.MustCompile(postalCodeAUString)
	PostalCodeITRegex = regexp.MustCompile(postalCodeITString)
	PostalCodeCHRegex = regexp.MustCompile(postalCodeCHString)
	PostalCodeATRegex = regexp.MustCompile(postalCodeATString)
	PostalCodeESRegex = regexp.MustCompile(postalCodeESString)
	PostalCodeNLRegex = regexp.MustCompile(postalCodeNLString)
	PostalCodeBERegex = regexp.MustCompile(postalCodeBEString)
	PostalCodeDKRegex = regexp.MustCompile(postalCodeDKString)
	PostalCodeSERegex = regexp.MustCompile(postalCodeSEString)
	PostalCodeNORegex = regexp.MustCompile(postalCodeNOString)
	PostalCodeBRRegex = regexp.MustCompile(postalCodeBRString)
	PostalCodePTRegex = regexp.MustCompile(postalCodePTString)
	PostalCodeFIRegex = regexp.MustCompile(postalCodeFIString)
	PostalCodeAXRegex = regexp.MustCompile(postalCodeAXString)
	PostalCodeKRRegex = regexp.MustCompile(postalCodeKRString)
	PostalCodeCNRegex = regexp.MustCompile(postalCodeCNString)
	PostalCodeTWRegex = regexp.MustCompile(postalCodeTWString)
	PostalCodeSGRegex = regexp.MustCompile(postalCodeSGString)
	PostalCodeDZRegex = regexp.MustCompile(postalCodeDZString)
	PostalCodeADRegex = regexp.MustCompile(postalCodeADString)
	PostalCodeARRegex = regexp.MustCompile(postalCodeARString)
	PostalCodeAMRegex = regexp.MustCompile(postalCodeAMString)
	PostalCodeAZRegex = regexp.MustCompile(postalCodeAZString)
	PostalCodeBHRegex = regexp.MustCompile(postalCodeBHString)
	PostalCodeBDRegex = regexp.MustCompile(postalCodeBDString)
	PostalCodeBBRegex = regexp.MustCompile(postalCodeBBString)
	PostalCodeBYRegex = regexp.MustCompile(postalCodeBYString)
	PostalCodeBMRegex = regexp.MustCompile(postalCodeBMString)
	PostalCodeBARegex = regexp.MustCompile(postalCodeBAString)
	PostalCodeIORegex = regexp.MustCompile(postalCodeIOString)
	PostalCodeBNRegex = regexp.MustCompile(postalCodeBNString)
	PostalCodeBGRegex = regexp.MustCompile(postalCodeBGString)
	PostalCodeKHRegex = regexp.MustCompile(postalCodeKHString)
	PostalCodeCVRegex = regexp.MustCompile(postalCodeCVString)
	PostalCodeCLRegex = regexp.MustCompile(postalCodeCLString)
	PostalCodeCRRegex = regexp.MustCompile(postalCodeCRString)
	PostalCodeHRRegex = regexp.MustCompile(postalCodeHRString)
	PostalCodeCYRegex = regexp.MustCompile(postalCodeCYString)
	PostalCodeCZRegex = regexp.MustCompile(postalCodeCZString)
	PostalCodeDORegex = regexp.MustCompile(postalCodeDOString)
	PostalCodeECRegex = regexp.MustCompile(postalCodeECString)
	PostalCodeEGRegex = regexp.MustCompile(postalCodeEGString)
	PostalCodeEERegex = regexp.MustCompile(postalCodeEEString)
	PostalCodeFORegex = regexp.MustCompile(postalCodeFOString)
	PostalCodeGERegex = regexp.MustCompile(postalCodeGEString)
	PostalCodeGRRegex = regexp.MustCompile(postalCodeGRString)
	PostalCodeGLRegex = regexp.MustCompile(postalCodeGLString)
	PostalCodeGTRegex = regexp.MustCompile(postalCodeGTString)
	PostalCodeHTRegex = regexp.MustCompile(postalCodeHTString)
	PostalCodeHNRegex = regexp.MustCompile(postalCodeHNString)
	PostalCodeHURegex = regexp.MustCompile(postalCodeHUString)
	PostalCodeISRegex = regexp.MustCompile(postalCodeISString)
	PostalCodeINRegex = regexp.MustCompile(postalCodeINString)
	PostalCodeIDRegex = regexp.MustCompile(postalCodeIDString)
	PostalCodeILRegex = regexp.MustCompile(postalCodeILString)
	PostalCodeJORegex = regexp.MustCompile(postalCodeJOString)
	PostalCodeKZRegex = regexp.MustCompile(postalCodeKZString)
	PostalCodeKERegex = regexp.MustCompile(postalCodeKEString)
	PostalCodeKWRegex = regexp.MustCompile(postalCodeKWString)
	PostalCodeLARegex = regexp.MustCompile(postalCodeLAString)
	PostalCodeLVRegex = regexp.MustCompile(postalCodeLVString)
	PostalCodeLBRegex = regexp.MustCompile(postalCodeLBString)
	PostalCodeLIRegex = regexp.MustCompile(postalCodeLIString)
	PostalCodeLTRegex = regexp.MustCompile(postalCodeLTString)
	PostalCodeLURegex = regexp.MustCompile(postalCodeLUString)
	PostalCodeMKRegex = regexp.MustCompile(postalCodeMKString)
	PostalCodeMYRegex = regexp.MustCompile(postalCodeMYString)
	PostalCodeMVRegex = regexp.MustCompile(postalCodeMVString)
	PostalCodeMTRegex = regexp.MustCompile(postalCodeMTString)
	PostalCodeMURegex = regexp.MustCompile(postalCodeMUString)
	PostalCodeMXRegex = regexp.MustCompile(postalCodeMXString)
	PostalCodeMDRegex = regexp.MustCompile(postalCodeMDString)
	PostalCodeMCRegex = regexp.MustCompile(postalCodeMCString)
	PostalCodeMARegex = regexp.MustCompile(postalCodeMAString)
	PostalCodeNPRegex = regexp.MustCompile(postalCodeNPString)
	PostalCodeNZRegex = regexp.MustCompile(postalCodeNZString)
	PostalCodeNIRegex = regexp.MustCompile(postalCodeNIString)
	PostalCodeNGRegex = regexp.MustCompile(postalCodeNGString)
	PostalCodeOMRegex = regexp.MustCompile(postalCodeOMString)
	PostalCodePKRegex = regexp.MustCompile(postalCodePKString)
	PostalCodePYRegex = regexp.MustCompile(postalCodePYString)
	PostalCodePHRegex = regexp.MustCompile(postalCodePHString)
	PostalCodePLRegex = regexp.MustCompile(postalCodePLString)
	PostalCodePRRegex = regexp.MustCompile(postalCodePRString)
	PostalCodeRORegex = regexp.MustCompile(postalCodeROString)
	PostalCodeRURegex = regexp.MustCompile(postalCodeRUString)
	PostalCodeSMRegex = regexp.MustCompile(postalCodeSMString)
	PostalCodeSARegex = regexp.MustCompile(postalCodeSAString)
	PostalCodeSNRegex = regexp.MustCompile(postalCodeSNString)
	PostalCodeSKRegex = regexp.MustCompile(postalCodeSKString)
	PostalCodeSIRegex = regexp.MustCompile(postalCodeSIString)
	PostalCodeZARegex = regexp.MustCompile(postalCodeZAString)
	PostalCodeLKRegex = regexp.MustCompile(postalCodeLKString)
	PostalCodeTJRegex = regexp.MustCompile(postalCodeTJString)
	PostalCodeTHRegex = regexp.MustCompile(postalCodeTHString)
	PostalCodeTNRegex = regexp.MustCompile(postalCodeTNString)
	PostalCodeTRRegex = regexp.MustCompile(postalCodeTRString)
	PostalCodeTMRegex = regexp.MustCompile(postalCodeTMString)
	PostalCodeUARegex = regexp.MustCompile(postalCodeUAString)
	PostalCodeUYRegex = regexp.MustCompile(postalCodeUYString)
	PostalCodeUZRegex = regexp.MustCompile(postalCodeUZString)
	PostalCodeVARegex = regexp.MustCompile(postalCodeVAString)
	PostalCodeVERegex = regexp.MustCompile(postalCodeVEString)
	PostalCodeZMRegex = regexp.MustCompile(postalCodeZMString)
	PostalCodeASRegex = regexp.MustCompile(postalCodeASString)
	PostalCodeCCRegex = regexp.MustCompile(postalCodeCCString)
	PostalCodeCKRegex = regexp.MustCompile(postalCodeCKString)
	PostalCodeRSRegex = regexp.MustCompile(postalCodeRSString)
	PostalCodeMERegex = regexp.MustCompile(postalCodeMEString)
	PostalCodeCSRegex = regexp.MustCompile(postalCodeCSString)
	PostalCodeYURegex = regexp.MustCompile(postalCodeYUString)
	PostalCodeCXRegex = regexp.MustCompile(postalCodeCXString)
	PostalCodeETRegex = regexp.MustCompile(postalCodeETString)
	PostalCodeFKRegex = regexp.MustCompile(postalCodeFKString)
	PostalCodeNFRegex = regexp.MustCompile(postalCodeNFString)
	PostalCodeFMRegex = regexp.MustCompile(postalCodeFMString)
	PostalCodeGFRegex = regexp.MustCompile(postalCodeGFString)
	PostalCodeGNRegex = regexp.MustCompile(postalCodeGNString)
	PostalCodeGPRegex = regexp.MustCompile(postalCodeGPString)
	PostalCodeGSRegex = regexp.MustCompile(postalCodeGSString)
	PostalCodeGURegex = regexp.MustCompile(postalCodeGUString)
	PostalCodeGWRegex = regexp.MustCompile(postalCodeGWString)
	PostalCodeHMRegex = regexp.MustCompile(postalCodeHMString)
	PostalCodeIQRegex = regexp.MustCompile(postalCodeIQString)
	PostalCodeKGRegex = regexp.MustCompile(postalCodeKGString)
	PostalCodeLRRegex = regexp.MustCompile(postalCodeLRString)
	PostalCodeLSRegex = regexp.MustCompile(postalCodeLSString)
	PostalCodeMGRegex = regexp.MustCompile(postalCodeMGString)
	PostalCodeMHRegex = regexp.MustCompile(postalCodeMHString)
	PostalCodeMNRegex = regexp.MustCompile(postalCodeMNString)
	PostalCodeMPRegex = regexp.MustCompile(postalCodeMPString)
	PostalCodeMQRegex = regexp.MustCompile(postalCodeMQString)
	PostalCodeNCRegex = regexp.MustCompile(postalCodeNCString)
	PostalCodeNERegex = regexp.MustCompile(postalCodeNEString)
	PostalCodeVIRegex = regexp.MustCompile(postalCodeVIString)
	PostalCodePFRegex = regexp.MustCompile(postalCodePFString)
	PostalCodePGRegex = regexp.MustCompile(postalCodePGString)
	PostalCodePMRegex = regexp.MustCompile(postalCodePMString)
	PostalCodePNRegex = regexp.MustCompile(postalCodePNString)
	PostalCodePWRegex = regexp.MustCompile(postalCodePWString)
	PostalCodeRERegex = regexp.MustCompile(postalCodeREString)
	PostalCodeSHRegex = regexp.MustCompile(postalCodeSHString)
	PostalCodeSJRegex = regexp.MustCompile(postalCodeSJString)
	PostalCodeSORegex = regexp.MustCompile(postalCodeSOString)
	PostalCodeSZRegex = regexp.MustCompile(postalCodeSZString)
	PostalCodeTCRegex = regexp.MustCompile(postalCodeTCString)
	PostalCodeWFRegex = regexp.MustCompile(postalCodeWFString)
	PostalCodeXKRegex = regexp.MustCompile(postalCodeXKString)
	PostalCodeYTRegex = regexp.MustCompile(postalCodeYTString)
)

type Country struct {
	Code            string
	PostalCodeRegex *regexp.Regexp
}

var Countries = []Country{
	{
		Code:            "US",
		PostalCodeRegex: PostalCodeUSRegex,
	},
	{
		Code:            "CA",
		PostalCodeRegex: PostalCodeCARegex,
	},
	{
		Code:            "GB",
		PostalCodeRegex: PostalCodeGBRegex,
	},
	{
		Code:            "CN",
		PostalCodeRegex: PostalCodeCNRegex,
	},
	{
		Code:            "JE",
		PostalCodeRegex: PostalCodeJERegex,
	},
	{
		Code:            "GG",
		PostalCodeRegex: PostalCodeGGRegex,
	},
	{
		Code:            "IM",
		PostalCodeRegex: PostalCodeIMRegex,
	},
	{
		Code:            "DE",
		PostalCodeRegex: PostalCodeDERegex,
	},
	{
		Code:            "JP",
		PostalCodeRegex: PostalCodeJPRegex,
	},
	{
		Code:            "FR",
		PostalCodeRegex: PostalCodeFRRegex,
	},
	{
		Code:            "AU",
		PostalCodeRegex: PostalCodeAURegex,
	},
	{
		Code:            "IT",
		PostalCodeRegex: PostalCodeITRegex,
	},
	{
		Code:            "CH",
		PostalCodeRegex: PostalCodeCHRegex,
	},
	{
		Code:            "AT",
		PostalCodeRegex: PostalCodeATRegex,
	},
	{
		Code:            "ES",
		PostalCodeRegex: PostalCodeESRegex,
	},
	{
		Code:            "NL",
		PostalCodeRegex: PostalCodeNLRegex,
	},
	{
		Code:            "BE",
		PostalCodeRegex: PostalCodeBERegex,
	},
	{
		Code:            "DK",
		PostalCodeRegex: PostalCodeDKRegex,
	},
	{
		Code:            "SE",
		PostalCodeRegex: PostalCodeSERegex,
	},
	{
		Code:            "NO",
		PostalCodeRegex: PostalCodeNORegex,
	},
	{
		Code:            "BR",
		PostalCodeRegex: PostalCodeBRRegex,
	},
	{
		Code:            "PT",
		PostalCodeRegex: PostalCodePTRegex,
	},
	{
		Code:            "FI",
		PostalCodeRegex: PostalCodeFIRegex,
	},
	{
		Code:            "AX",
		PostalCodeRegex: PostalCodeAXRegex,
	},
	{
		Code:            "KR",
		PostalCodeRegex: PostalCodeKRRegex,
	},
	{
		Code:            "TW",
		PostalCodeRegex: PostalCodeTWRegex,
	},
	{
		Code:            "SG",
		PostalCodeRegex: PostalCodeSGRegex,
	},
	{
		Code:            "DZ",
		PostalCodeRegex: PostalCodeDZRegex,
	},
	{
		Code:            "AD",
		PostalCodeRegex: PostalCodeADRegex,
	},
	{
		Code:            "AR",
		PostalCodeRegex: PostalCodeARRegex,
	},
	{
		Code:            "AM",
		PostalCodeRegex: PostalCodeAMRegex,
	},
	{
		Code:            "AZ",
		PostalCodeRegex: PostalCodeAZRegex,
	},
	{
		Code:            "BH",
		PostalCodeRegex: PostalCodeBHRegex,
	},
	{
		Code:            "BD",
		PostalCodeRegex: PostalCodeBDRegex,
	},
	{
		Code:            "BB",
		PostalCodeRegex: PostalCodeBBRegex,
	},
	{
		Code:            "BY",
		PostalCodeRegex: PostalCodeBYRegex,
	},
	{
		Code:            "BM",
		PostalCodeRegex: PostalCodeBMRegex,
	},
	{
		Code:            "BA",
		PostalCodeRegex: PostalCodeBARegex,
	},
	{
		Code:            "IO",
		PostalCodeRegex: PostalCodeIORegex,
	},
	{
		Code:            "BN",
		PostalCodeRegex: PostalCodeBNRegex,
	},
	{
		Code:            "BG",
		PostalCodeRegex: PostalCodeBGRegex,
	},
	{
		Code:            "KH",
		PostalCodeRegex: PostalCodeKHRegex,
	},
	{
		Code:            "CV",
		PostalCodeRegex: PostalCodeCVRegex,
	},
	{
		Code:            "CL",
		PostalCodeRegex: PostalCodeCLRegex,
	},
	{
		Code:            "CR",
		PostalCodeRegex: PostalCodeCRRegex,
	},
	{
		Code:            "HR",
		PostalCodeRegex: PostalCodeHRRegex,
	},
	{
		Code:            "CY",
		PostalCodeRegex: PostalCodeCYRegex,
	},
	{
		Code:            "CZ",
		PostalCodeRegex: PostalCodeCZRegex,
	},
	{
		Code:            "DO",
		PostalCodeRegex: PostalCodeDORegex,
	},
	{
		Code:            "EC",
		PostalCodeRegex: PostalCodeECRegex,
	},
	{
		Code:            "EG",
		PostalCodeRegex: PostalCodeEGRegex,
	},
	{
		Code:            "EE",
		PostalCodeRegex: PostalCodeEERegex,
	},
	{
		Code:            "FO",
		PostalCodeRegex: PostalCodeFORegex,
	},
	{
		Code:            "GE",
		PostalCodeRegex: PostalCodeGERegex,
	},
	{
		Code:            "GR",
		PostalCodeRegex: PostalCodeGRRegex,
	},
	{
		Code:            "GL",
		PostalCodeRegex: PostalCodeGLRegex,
	},
	{
		Code:            "GT",
		PostalCodeRegex: PostalCodeGTRegex,
	},
	{
		Code:            "HT",
		PostalCodeRegex: PostalCodeHTRegex,
	},
	{
		Code:            "HN",
		PostalCodeRegex: PostalCodeHNRegex,
	},
	{
		Code:            "HU",
		PostalCodeRegex: PostalCodeHURegex,
	},
	{
		Code:            "IS",
		PostalCodeRegex: PostalCodeISRegex,
	},
	{
		Code:            "IN",
		PostalCodeRegex: PostalCodeINRegex,
	},
	{
		Code:            "ID",
		PostalCodeRegex: PostalCodeIDRegex,
	},
	{
		Code:            "IL",
		PostalCodeRegex: PostalCodeILRegex,
	},
	{
		Code:            "JO",
		PostalCodeRegex: PostalCodeJORegex,
	},
	{
		Code:            "KZ",
		PostalCodeRegex: PostalCodeKZRegex,
	},
	{
		Code:            "KE",
		PostalCodeRegex: PostalCodeKERegex,
	},
	{
		Code:            "KW",
		PostalCodeRegex: PostalCodeKWRegex,
	},
	{
		Code:            "LA",
		PostalCodeRegex: PostalCodeLARegex,
	},
	{
		Code:            "LV",
		PostalCodeRegex: PostalCodeLVRegex,
	},
	{
		Code:            "LB",
		PostalCodeRegex: PostalCodeLBRegex,
	},
	{
		Code:            "LI",
		PostalCodeRegex: PostalCodeLIRegex,
	},
	{
		Code:            "LT",
		PostalCodeRegex: PostalCodeLTRegex,
	},
	{
		Code:            "LU",
		PostalCodeRegex: PostalCodeLURegex,
	},
	{
		Code:            "MK",
		PostalCodeRegex: PostalCodeMKRegex,
	},
	{
		Code:            "MY",
		PostalCodeRegex: PostalCodeMYRegex,
	},
	{
		Code:            "MV",
		PostalCodeRegex: PostalCodeMVRegex,
	},
	{
		Code:            "MT",
		PostalCodeRegex: PostalCodeMTRegex,
	},
	{
		Code:            "MU",
		PostalCodeRegex: PostalCodeMURegex,
	},
	{
		Code:            "MX",
		PostalCodeRegex: PostalCodeMXRegex,
	},
	{
		Code:            "MD",
		PostalCodeRegex: PostalCodeMDRegex,
	},
	{
		Code:            "MC",
		PostalCodeRegex: PostalCodeMCRegex,
	},
	{
		Code:            "MA",
		PostalCodeRegex: PostalCodeMARegex,
	},
	{
		Code:            "NP",
		PostalCodeRegex: PostalCodeNPRegex,
	},
	{
		Code:            "NZ",
		PostalCodeRegex: PostalCodeNZRegex,
	},
	{
		Code:            "NI",
		PostalCodeRegex: PostalCodeNIRegex,
	},
	{
		Code:            "NG",
		PostalCodeRegex: PostalCodeNGRegex,
	},
	{
		Code:            "OM",
		PostalCodeRegex: PostalCodeOMRegex,
	},
	{
		Code:            "PK",
		PostalCodeRegex: PostalCodePKRegex,
	},
	{
		Code:            "PY",
		PostalCodeRegex: PostalCodePYRegex,
	},
	{
		Code:            "PH",
		PostalCodeRegex: PostalCodePHRegex,
	},
	{
		Code:            "PL",
		PostalCodeRegex: PostalCodePLRegex,
	},
	{
		Code:            "PR",
		PostalCodeRegex: PostalCodePRRegex,
	},
	{
		Code:            "RO",
		PostalCodeRegex: PostalCodeRORegex,
	},
	{
		Code:            "RU",
		PostalCodeRegex: PostalCodeRURegex,
	},
	{
		Code:            "SM",
		PostalCodeRegex: PostalCodeSMRegex,
	},
	{
		Code:            "SA",
		PostalCodeRegex: PostalCodeSARegex,
	},
	{
		Code:            "SN",
		PostalCodeRegex: PostalCodeSNRegex,
	},
	{
		Code:            "SK",
		PostalCodeRegex: PostalCodeSKRegex,
	},
	{
		Code:            "SI",
		PostalCodeRegex: PostalCodeSIRegex,
	},
	{
		Code:            "ZA",
		PostalCodeRegex: PostalCodeZARegex,
	},
	{
		Code:            "LK",
		PostalCodeRegex: PostalCodeLKRegex,
	},
	{
		Code:            "TJ",
		PostalCodeRegex: PostalCodeTJRegex,
	},
	{
		Code:            "TH",
		PostalCodeRegex: PostalCodeTHRegex,
	},
	{
		Code:            "TN",
		PostalCodeRegex: PostalCodeTNRegex,
	},
	{
		Code:            "TR",
		PostalCodeRegex: PostalCodeTRRegex,
	},
	{
		Code:            "TM",
		PostalCodeRegex: PostalCodeTMRegex,
	},
	{
		Code:            "UA",
		PostalCodeRegex: PostalCodeUARegex,
	},
	{
		Code:            "UY",
		PostalCodeRegex: PostalCodeUYRegex,
	},
	{
		Code:            "UZ",
		PostalCodeRegex: PostalCodeUZRegex,
	},
	{
		Code:            "VA",
		PostalCodeRegex: PostalCodeVARegex,
	},
	{
		Code:            "VE",
		PostalCodeRegex: PostalCodeVERegex,
	},
	{
		Code:            "ZM",
		PostalCodeRegex: PostalCodeZMRegex,
	},
	{
		Code:            "AS",
		PostalCodeRegex: PostalCodeASRegex,
	},
	{
		Code:            "CC",
		PostalCodeRegex: PostalCodeCCRegex,
	},
	{
		Code:            "CK",
		PostalCodeRegex: PostalCodeCKRegex,
	},
	{
		Code:            "RS",
		PostalCodeRegex: PostalCodeRSRegex,
	},
	{
		Code:            "ME",
		PostalCodeRegex: PostalCodeMERegex,
	},
	{
		Code:            "CS",
		PostalCodeRegex: PostalCodeCSRegex,
	},
	{
		Code:            "YU",
		PostalCodeRegex: PostalCodeYURegex,
	},
	{
		Code:            "CX",
		PostalCodeRegex: PostalCodeCXRegex,
	},
	{
		Code:            "ET",
		PostalCodeRegex: PostalCodeETRegex,
	},
	{
		Code:            "FK",
		PostalCodeRegex: PostalCodeFKRegex,
	},
	{
		Code:            "NF",
		PostalCodeRegex: PostalCodeNFRegex,
	},
	{
		Code:            "FM",
		PostalCodeRegex: PostalCodeFMRegex,
	},
	{
		Code:            "GF",
		PostalCodeRegex: PostalCodeGFRegex,
	},
	{
		Code:            "GN",
		PostalCodeRegex: PostalCodeGNRegex,
	},
	{
		Code:            "GP",
		PostalCodeRegex: PostalCodeGPRegex,
	},
	{
		Code:            "GS",
		PostalCodeRegex: PostalCodeGSRegex,
	},
	{
		Code:            "GU",
		PostalCodeRegex: PostalCodeGURegex,
	},
	{
		Code:            "GW",
		PostalCodeRegex: PostalCodeGWRegex,
	},
	{
		Code:            "HM",
		PostalCodeRegex: PostalCodeHMRegex,
	},
	{
		Code:            "IQ",
		PostalCodeRegex: PostalCodeIQRegex,
	},
	{
		Code:            "KG",
		PostalCodeRegex: PostalCodeKGRegex,
	},
	{
		Code:            "LR",
		PostalCodeRegex: PostalCodeLRRegex,
	},
	{
		Code:            "LS",
		PostalCodeRegex: PostalCodeLSRegex,
	},
	{
		Code:            "MG",
		PostalCodeRegex: PostalCodeMGRegex,
	},
	{
		Code:            "MH",
		PostalCodeRegex: PostalCodeMHRegex,
	},
	{
		Code:            "MN",
		PostalCodeRegex: PostalCodeMNRegex,
	},
	{
		Code:            "MP",
		PostalCodeRegex: PostalCodeMPRegex,
	},
	{
		Code:            "MQ",
		PostalCodeRegex: PostalCodeMQRegex,
	},
	{
		Code:            "NC",
		PostalCodeRegex: PostalCodeNCRegex,
	},
	{
		Code:            "NE",
		PostalCodeRegex: PostalCodeNERegex,
	},
	{
		Code:            "VI",
		PostalCodeRegex: PostalCodeVIRegex,
	},
	{
		Code:            "PF",
		PostalCodeRegex: PostalCodePFRegex,
	},
	{
		Code:            "PG",
		PostalCodeRegex: PostalCodePGRegex,
	},
	{
		Code:            "PM",
		PostalCodeRegex: PostalCodePMRegex,
	},
	{
		Code:            "PN",
		PostalCodeRegex: PostalCodePNRegex,
	},
	{
		Code:            "PW",
		PostalCodeRegex: PostalCodePWRegex,
	},
	{
		Code:            "RE",
		PostalCodeRegex: PostalCodeRERegex,
	},
	{
		Code:            "SH",
		PostalCodeRegex: PostalCodeSHRegex,
	},
	{
		Code:            "SJ",
		PostalCodeRegex: PostalCodeSJRegex,
	},
	{
		Code:            "SO",
		PostalCodeRegex: PostalCodeSORegex,
	},
	{
		Code:            "SZ",
		PostalCodeRegex: PostalCodeSZRegex,
	},
	{
		Code:            "TC",
		PostalCodeRegex: PostalCodeTCRegex,
	},
	{
		Code:            "WF",
		PostalCodeRegex: PostalCodeWFRegex,
	},
	{
		Code:            "XK",
		PostalCodeRegex: PostalCodeXKRegex,
	},
	{
		Code:            "YT",
		PostalCodeRegex: PostalCodeYTRegex,
	},
}

func ValidatePostalCodeByCountyCode(countryCode string, postalCode string) bool {
	//check to make sure countryCode is valid
	for _, v := range Countries {
		if strings.ToUpper(v.Code) == strings.ToUpper(countryCode) {
			if v.PostalCodeRegex != nil {
				return v.PostalCodeRegex.Match([]byte(postalCode))
			}
			return false
		}
	}

	return false
}
