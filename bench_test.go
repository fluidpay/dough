package dough

import "testing"

func BenchmarkGetISOFromAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := GetISOFromAlpha("USD"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetISOFromNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := GetISOFromNumeric("840"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFormatCurrency(b *testing.B) {
	iso, err := GetISOFromAlpha("USD")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatCurrency(123456789, iso)
	}
}

func BenchmarkDisplayFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := DisplayFull(123456789, "USD"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValidLuhn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !ValidLuhn("4111111111111111") {
			b.Fatal("expected valid")
		}
	}
}

func BenchmarkGetCardType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := GetCardType("4111111111111111"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMaskCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, _, _, err := MaskCard("4111111111111111"); err != nil {
			b.Fatal(err)
		}
	}
}
