// Llistat dels símbols i la seva correspondència en Range64.
// CreatedAt: 2024/12/10 dt. GPT

package core

var IdToSymbol = map[uint64]rune{}

var SymbolToID = map[rune]uint64{
	// Espais i control
	' ': 0x01, '\t': 0x02, '\n': 0x03, '\r': 0x04,

	// Lletres majúscules (inclou accents i diacrítics)
	'A': 0x10, 'Á': 0x11, 'À': 0x12, 'Ä': 0x13, 'Â': 0x14,
	'B': 0x20,
	'C': 0x30, 'Ç': 0x31,
	'D': 0x40,
	'E': 0x50, 'É': 0x51, 'È': 0x52, 'Ê': 0x53, 'Ë': 0x54,
	'F': 0x60,
	'G': 0x70,
	'H': 0x80,
	'I': 0x90, 'Í': 0x91, 'Ì': 0x92, 'Ï': 0x93,
	'J': 0xA0,
	'K': 0xB0,
	'L': 0xC0,
	'M': 0xD0,
	'N': 0xE0, 'Ñ': 0xE1,
	'O': 0xF0, 'Ó': 0xF1, 'Ò': 0xF2, 'Ö': 0xF3, 'Ô': 0xF4,
	'P': 0x100,
	'Q': 0x110,
	'R': 0x120,
	'S': 0x130,
	'T': 0x140,
	'U': 0x150, 'Ú': 0x151, 'Ù': 0x152, 'Ü': 0x153, 'Û': 0x154,
	'V': 0x160,
	'W': 0x170,
	'X': 0x180,
	'Y': 0x190,
	'Z': 0x1A0,

	// Lletres minúscules (inclou accents i diacrítics)
	'a': 0x1B0, 'á': 0x1B1, 'à': 0x1B2, 'ä': 0x1B3, 'â': 0x1B4,
	'b': 0x1C0,
	'c': 0x1D0, 'ç': 0x1D1,
	'd': 0x1E0,
	'e': 0x1F0, 'é': 0x1F1, 'è': 0x1F2, 'ê': 0x1F3, 'ë': 0x1F4,
	'f': 0x200,
	'g': 0x210,
	'h': 0x220,
	'i': 0x230, 'í': 0x231, 'ì': 0x232, 'ï': 0x233,
	'j': 0x240,
	'k': 0x250,
	'l': 0x260,
	'm': 0x270,
	'n': 0x280, 'ñ': 0x281,
	'o': 0x290, 'ó': 0x291, 'ò': 0x292, 'ö': 0x293, 'ô': 0x294,
	'p': 0x2A0,
	'q': 0x2B0,
	'r': 0x2C0,
	's': 0x2D0, 'ß': 0x2D1,
	't': 0x2E0,
	'u': 0x2F0, 'ú': 0x2F1, 'ù': 0x2F2, 'ü': 0x2F3, 'û': 0x2F4,
	'v': 0x300,
	'w': 0x310,
	'x': 0x320,
	'y': 0x330,
	'z': 0x340,

	// Números
	'0': 0x400, '1': 0x410, '2': 0x420, '3': 0x430, '4': 0x440,
	'5': 0x450, '6': 0x460, '7': 0x470, '8': 0x480, '9': 0x490,

	// Puntuació
	'.': 0x500, ',': 0x510, ';': 0x520, ':': 0x530,
	'!': 0x540, '?': 0x550,
	'(': 0x560, ')': 0x570,
	'[': 0x580, ']': 0x590,
	'{': 0x5A0, '}': 0x5B0,
	'-': 0x5C0, '_': 0x5D0,
	'\'': 0x5E0, '"': 0x5F0,

	// Nous símbols especials
	'º': 0x600, 'ª': 0x610, '·': 0x620, '|': 0x630,
	'#': 0x640, '$': 0x650, '%': 0x660, '&': 0x670,
	'¬': 0x680, '~': 0x690, '€': 0x6A0,

	// Elements Especials
	rune(0xFFFC): 0xFFF3, // Nul
	rune(0xFFFD): 0xFFF0, // Error
	rune(0xFFFE): 0xFFF1, // Desconegut
	rune(0xFFFF): 0xFFF2, // Qualsevol
}
