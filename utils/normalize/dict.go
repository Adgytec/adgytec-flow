package normalize

// replacements handles special characters that Unicode normalization alone doesn't decompose.
var replacements = map[rune]string{
	'ß': "ss",
	'Æ': "AE", 'æ': "ae",
	'Œ': "OE", 'œ': "oe",
	'Ø': "O", 'ø': "o",
	'Ł': "L", 'ł': "l",
	'Þ': "Th", 'þ': "th",
	'Đ': "D", 'đ': "d",
	'ð': "d",
	'Ħ': "H", 'ħ': "h",
	'Ŋ': "Ng", 'ŋ': "ng",
	'ƒ': "f",
	'ĸ': "k",
	'Ŧ': "T", 'ŧ': "t",
}
