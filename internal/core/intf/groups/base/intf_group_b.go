// Interfície general de tots els Subgrups B (símbols, paddings, nuls, coords, perc_ranges, percs, KeyMouse, Saturated i Error).
// CreatedAt: 2025/01/04 ds. JIQ

package base

// Interfície general de tots els Subgrups B (símbols, paddings, nuls, coords, perc_ranges, percs, KeyMouse, Saturated i Error).
type GroupBIntf interface {
	IsSymbolType() bool          // Cert només si la instància és un símbol.
	IsPaddingType() bool         // Cert només si la instància és un valor de padding.
	IsNullType() bool            // Cert només si la instància és un valor nul.
	IsCoordinatesType() bool     // Cert només si la instància és un valor de coordenades.
	IsPercentageRangeType() bool // Cert només si la instància és un rang de percentatges.
	IsPercentageType() bool      // Cert només si la instància és un valor de percentatge.
	IsKeyOrMouseType() bool      // Cert només si la instància és un valor de teclat o ratolí.
	IsSaturatedType() bool       // Cert només si la instància és un valor saturat.
	IsErrorType() bool           // Cert només si la instància és un error.
}
