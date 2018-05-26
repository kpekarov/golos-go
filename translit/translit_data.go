package translit

var encMap = map[string]string{
	"а": "a",
	"б": "b",
	"в": "v",
	"ґ": "g",
	"г": "g",
	"д": "d",
	"е": "e",
	"ё": "yo",
	"є": "ye",
	"ж": "zh",
	"з": "z",
	"и": "i",
	"ї": "yi",
	"й": "ij",
	"к": "k",
	"л": "l",
	"м": "m",
	"н": "n",
	"о": "o",
	"п": "p",
	"р": "r",
	"с": "s",
	"т": "t",
	"у": "u",
	"ф": "f",
	"х": "kh",
	"ц": "cz",
	"ч": "ch",
	"ш": "sh",
	"щ": "shch",
	"ъ": "xx",
	"ы": "y",
	"ь": "x",
	"э": "ye",
	"ю": "yu",
	"я": "ya",
	" ": "-",
	".": "",
	",": "",
	"(": "",
	")": "",
	"{": "",
	"}": "",
	"/": "",
}

var decMap = map[string]string{
	"a":    "а",
	"b":    "б",
	"v":    "в",
	"g":    "г",
	"d":    "д",
	"e":    "е",
	"yo":   "ё",
	"zh":   "ж",
	"z":    "з",
	"i":    "и",
	"ij":   "й",
	"k":    "к",
	"l":    "л",
	"m":    "м",
	"n":    "н",
	"o":    "о",
	"p":    "п",
	"r":    "р",
	"s":    "с",
	"t":    "т",
	"u":    "у",
	"f":    "ф",
	"kh":   "х",
	"cz":   "ц",
	"ch":   "ч",
	"sh":   "ш",
	"shch": "щ",
	"xx":   "ъ",
	"y":    "ы",
	"x":    "ь",
	"ye":   "э",
	"yu":   "ю",
	"ya":   "я",
	"-":    " ",
}

var decOrder = []string{
	"-",
	"shch",
	"sh",
	"ch",
	"cz",
	"ij",
	"yo",
	"ye",
	"yu",
	"ya",
	"kh",
	"zh",
	"a",
	"b",
	"v",
	"g",
	"d",
	"e",
	"z",
	"i",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"r",
	"s",
	"t",
	"u",
	"f",
	"xx",
	"y",
	"x",
}
