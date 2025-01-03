// Interfícies de Range pel Grup B en float64.
// CreatedAt: 2025/01/03 dv. JIQ

package Intf

// Interfície per a valors de 64 bits (float64) del Grup B.
type RangeF64GroupBIntf interface {
	RangeF64Intf // Hereta les funcions generals per a float64
	GroupBIntf   // Hereta les funcions generals per a Grup B
}

// Interfície per a valors de 64 bits (float64) del Subgrup B.1 (símbols).
type RangeF64SymbolIntf interface {
	RangeF64GroupBIntf // Hereta les funcions del Grup B per a float64

	Symbol() rune // Retorna el símbol com a UTF-32.
}

// Interfície per a valors de 64 bits (float64) del Subgrup B.2.P (paddings).
type RangeF64PaddingIntf interface {
	RangeF64GroupBIntf // Hereta les funcions del Grup B per a float64

	IsStartPadding() bool  // Cert només si la instància és un valor de padding inicial.
	IsCommonPadding() bool // Cert només si la instància és un valor de padding comú.
	IsEndPadding() bool    // Cert només si la instància és un valor de padding final.
	PaddingRune() rune     // Retorna el símbol de padding
}

// Interfície per a valors de 64 bits (float64) del Subgrup B.2.N (nul).
type RangeF64NullIntf interface {
	RangeF64GroupBIntf // Hereta les funcions del Grup B per a float64
	IsNull() bool      // Cert només si la instància és un valor nul.
}

// Interfície per a valors de 64 bits (float64) del Subgrup B.2.C (coordenades).
type RangeF64CoordinatesIntf interface {
	RangeF64GroupBIntf // Hereta les funcions del Grup B per a float64

	Coordinates() (float64, float64) // Retorna les coordenades (x, y)
}

// Interfície per a valors de 64 bits (float64) del Subgrup B.2.R (rang de percentatges).
type RangeF64PercRangeIntf interface {
	RangeF64GroupBIntf // Hereta les funcions del Grup B per a float64

	Linits() (RangeF64PercentageIntf, RangeF64PercentageIntf)
}

// Interfície per a valors de 64 bits (float64) del Subgrup B.2.E (percentatge).
type RangeF64PercentageIntf interface {
	ComparableIntf
	RangeF64GroupBIntf

	Percentage() float64 // Retorna el percentatge
}

// Tecles i butons de mouse.
const (
	// Bytes:                        B7	      B6       B5       B4       B3       B2       B1       B0
	Key_F01               uint64 = 0b00000000_10000000_00000000_00000000_00000000_00000000_00000000_00000000
	Key_F02               uint64 = 0b00000000_01000000_00000000_00000000_00000000_00000000_00000000_00000000
	Key_F03               uint64 = 0b00000000_00100000_00000000_00000000_00000000_00000000_00000000_00000000
	Key_F04               uint64 = 0b00000000_00010000_00000000_00000000_00000000_00000000_00000000_00000000
	Key_F05               uint64 = 0b00000000_00001000_00000000_00000000_00000000_00000000_00000000_00000000
	Key_F06               uint64 = 0b00000000_00000100_00000000_00000000_00000000_00000000_00000000_00000000
	Key_F07               uint64 = 0b00000000_00000010_00000000_00000000_00000000_00000000_00000000_00000000
	Key_F08               uint64 = 0b00000000_00000001_00000000_00000000_00000000_00000000_00000000_00000000
	Key_F09               uint64 = 0b00000000_00000000_10000000_00000000_00000000_00000000_00000000_00000000
	Key_F10               uint64 = 0b00000000_00000000_01000000_00000000_00000000_00000000_00000000_00000000
	Key_F11               uint64 = 0b00000000_00000000_00100000_00000000_00000000_00000000_00000000_00000000
	Key_F12               uint64 = 0b00000000_00000000_00010000_00000000_00000000_00000000_00000000_00000000
	Key_Tab               uint64 = 0b00000000_00000000_00001000_00000000_00000000_00000000_00000000_00000000
	Key_CapsLock          uint64 = 0b00000000_00000000_00000100_00000000_00000000_00000000_00000000_00000000
	Key_LeftShift         uint64 = 0b00000000_00000000_00000010_00000000_00000000_00000000_00000000_00000000
	Key_RightShift        uint64 = 0b00000000_00000000_00000001_00000000_00000000_00000000_00000000_00000000
	Key_LeftControl       uint64 = 0b00000000_00000000_00000000_10000000_00000000_00000000_00000000_00000000
	Key_RightControl      uint64 = 0b00000000_00000000_00000000_01000000_00000000_00000000_00000000_00000000
	Key_Function          uint64 = 0b00000000_00000000_00000000_00100000_00000000_00000000_00000000_00000000
	Key_Windows           uint64 = 0b00000000_00000000_00000000_00010000_00000000_00000000_00000000_00000000
	Key_OSX               uint64 = 0b00000000_00000000_00000000_00001000_00000000_00000000_00000000_00000000
	Key_Alt               uint64 = 0b00000000_00000000_00000000_00000100_00000000_00000000_00000000_00000000
	Key_AltGraph          uint64 = 0b00000000_00000000_00000000_00000010_00000000_00000000_00000000_00000000
	Key_Delete            uint64 = 0b00000000_00000000_00000000_00000001_00000000_00000000_00000000_00000000
	Key_Supr              uint64 = 0b00000000_00000000_00000000_00000000_10000000_00000000_00000000_00000000
	Key_Escape            uint64 = 0b00000000_00000000_00000000_00000000_01000000_00000000_00000000_00000000
	Key_LeftMouseButton   uint64 = 0b00000000_00000000_00000000_00000000_00100000_00000000_00000000_00000000
	Key_MiddleMouseButton uint64 = 0b00000000_00000000_00000000_00000000_00010000_00000000_00000000_00000000
	Key_RightMouseButton  uint64 = 0b00000000_00000000_00000000_00000000_00001000_00000000_00000000_00000000
	Key_ExtendedBits      uint64 = 0b00000000_00000000_00000000_00000000_00000111_11111111_11111111_11111111
)

// Interfície per a valors de 64 bits (float64) del Subgrup B.2.K (teclat i mouse).
type RangeF64KeyboardMouseIntf interface {
	RangeF64GroupBIntf

	IsKeyboard() bool // Cert només si la instància és un valor de teclat.
	IsMouse() bool    // Cert només si la instància és un valor de ratolí.

	KeysMask() []uint32        // Retorna la màscara de les tecles o butons de mouse premuts.
	IsF1() bool                // Cert només si la instància conté la tecla F1.
	IsF2() bool                // Cert només si la instància conté la tecla F2.
	IsF3() bool                // Cert només si la instància conté la tecla F3.
	IsF4() bool                // Cert només si la instància conté la tecla F4.
	IsF5() bool                // Cert només si la instància conté la tecla F5.
	IsF6() bool                // Cert només si la instància conté la tecla F6.
	IsF7() bool                // Cert només si la instància conté la tecla F7.
	IsF8() bool                // Cert només si la instància conté la tecla F8.
	IsF9() bool                // Cert només si la instància conté la tecla F9.
	IsF10() bool               // Cert només si la instància conté la tecla F10.
	IsF11() bool               // Cert només si la instància conté la tecla F11.
	IsF12() bool               // Cert només si la instància conté la tecla F12.
	IsTab() bool               // Cert només si la instància conté la tecla Tab.
	IsCapsLock() bool          // Cert només si la instància conté la tecla CapsLock.
	IsLeftShift() bool         // Cert només si la instància conté la tecla LeftShift.
	IsRightShift() bool        // Cert només si la instància conté la tecla RightShift.
	IsLeftControl() bool       // Cert només si la instància conté la tecla LeftControl.
	IsRightControl() bool      // Cert només si la instància conté la tecla RightControl.
	IsFunction() bool          // Cert només si la instància conté la tecla Function.
	IsWindows() bool           // Cert només si la instància conté la tecla Windows.
	IsOSX() bool               // Cert només si la instància conté la tecla OSX.
	IsAlt() bool               // Cert només si la instància conté la tecla Alt.
	IsAltGraph() bool          // Cert només si la instància conté la tecla AltGraph.
	IsDelete() bool            // Cert només si la instància conté la tecla Delete.
	IsSupr() bool              // Cert només si la instància conté la tecla Supr.
	IsEscape() bool            // Cert només si la instància conté la tecla Escape.
	IsLeftMouseButton() bool   // Cert només si la instància conté el butó esquerre del ratolí.
	IsMiddleMouseButton() bool // Cert només si la instància conté el butó cetral del ratolí.
	IsRightMouseButton() bool  // Cert només si la instància conté el butó dret del ratolí.
	IsAnExtendedBit() int      // Cert només si la instància és un bit extès.
}

// Interfície per a valors de 64 bits (float64) del Subgrup B.3 (saturats).
type RangeF64SaturatedIntf interface {
	RangeF64GroupBIntf

	IsSaturatedPos() bool        // Cert només si la instància és un valor saturat positiu.
	IsSaturatedNeg() bool        // Cert només si la instància és un valor saturat negatiu.
	IEEE754() RangeF64GroupAIntf // Retorna la instància com a RangeF64IE754Intf
}

// Interfície per a valors de 64 bits (float64) del Subgrup B.4 (errors)
type ErrorF64Intf interface {
	Error64Intf
}
