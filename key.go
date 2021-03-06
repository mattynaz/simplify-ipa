package simplifyipa

var ipaKey = map[string]string{
	".":    ".",
	"a":    "a",
	"aɪ":   "ai",
	"aʊ":   "ow",
	"aː":   "a",
	"b":    "b",
	"d":    "d",
	"dz":   "dz",
	"dʒ":   "j",
	"e":    "ey",
	"eː":   "ay",
	"f":    "f",
	"h":    "h",
	"i":    "ee",
	"iː":   "ee",
	"i̯":   "ee",
	"j":    "y",
	"k":    "k",
	"l":    "l",
	"l̩":   "l",
	"m":    "m",
	"m̩":   "m",
	"n":    "n",
	"n̩":   "n",
	"o":    "o",
	"oː":   "o",
	"p":    "p",
	"pf":   "pf",
	"r":    "r",
	"s":    "s",
	"t":    "t",
	"ts":   "ts",
	"tɕ":   "ch",
	"tʃ":   "ch",
	"u":    "oo",
	"uː":   "oo",
	"u̯":   "w",
	"v":    "v",
	"w":    "w",
	"x":    "kh",
	"y":    "u",
	"yː":   "ew",
	"z":    "z",
	"ã":    "ã",
	"ãː":   "aw",
	"æ":    "ah",
	"ç":    "hy",
	"ð":    "th",
	"õ":    "õ",
	"õː":   "o",
	"ø":    "u",
	"øː":   "ew",
	"øːr":  "er",
	"ŋ":    "ŋ",
	"œ":    "uh",
	"œːɐ̯": "er",
	"œ̃":   "uh",
	"œ̃ː":  "uh",
	"ɐ":    "u",
	"ɐ̯":   "ɐ̯",
	"ɑ":    "aw",
	"ɑ̃":   "ãw",
	"ɔ":    "o",
	"ɔʊ̯":  "ow",
	"ɔʏ":   "oy",
	"ɔ̃":   "aw",
	"ɕː":   "sh",
	"ə":    "uh",
	"ɛ":    "e",
	"ɛɪ̯":  "ay",
	"ɛː":   "ai",
	"ɛ̃":   "ũh",
	"ɛ̃ː":  "eh",
	"ɟʝ":   "j",
	"ɡ":    "g",
	"ɣ":    "ɣ",
	"ɥ":    "w",
	"ɨ":    "e",
	"ɪ":    "i",
	"ɫ":    "w",
	"ɱ":    "m",
	"ɲ":    "ny",
	"ɵ":    "oo",
	"ɹ":    "r",
	"ɾ":    "r",
	"ʁ":    "r",
	"ʂ":    "sh",
	"ʃ":    "sh",
	"ʉ":    "oo",
	"ʊ":    "w",
	"ʎ":    "yl",
	"ʏ":    "oo",
	"ʐ":    "zh",
	"ʑː":   "zh",
	"ʒ":    "zh",
	"ʔ":    "'",
	"ʝ":    "y",
	"ʲ":    "ʸ",
	"ˈ":    "ˈ",
	"ˌ":    "ˌ",
	"ː":    "ː",
	"͡":    "͡",
	"β":    "b",
	"θ":    "th",
}
