package simplifyipa

import "github.com/mattynaz/transliterate"

// IpaSimplifier is a
var converter *transliterate.Converter

// SimplifyIPA is a
func SimplifyIPA(text string) (result string) {
	result = converter.Convert(text)
	return
}
