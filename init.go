package simplifyipa

import "github.com/mattynaz/transliterate"

func init() {
	Initialize()
}

// Initialize builds
func Initialize() {
	converter = transliterate.NewConverter(ipaKey)
}
