// Generació d'embeddings per a símbols textuals amb criteri fonètic i funcional
// CreatedAt: 24-12-2024 dg. GPT(JIQ)

package main

import (
	"fmt"

	f64_2p "github.com/jibort/ld_mcac/internal/core/RF642Pi"
	f64_1 "github.com/jibort/ld_mcac/internal/core/RF64One"
)

// Tipus definit per RangF64 ja existent en un altre fitxer

type SymbolEmbedding struct {
	Symbol     rune
	ValueOne   f64_1.RangeF64One
	ValueTwoPi f64_2p.RangeF64TwoPi
	// Group  string // Categoria funcional o fonètica
}

func generateEmbeddings() []SymbolEmbedding {
	// Registre de valors usats per garantir que són únics
	usedValuesOne := map[float64]rune{}
	usedValuesTwoPi := map[float64]rune{}
	createRangeF64One := func(value float64, symbol rune) f64_1.RangeF64One {
		if existingSymbol, exists := usedValuesOne[value]; exists {
			panic(fmt.Sprintf("El valor %.4f ja està assignat al símbol '%s' (intentat assignar a '%s')", value, string(existingSymbol), string(symbol)))
		}
		usedValuesOne[value] = symbol
		return *f64_1.NewRangeF64One(value)
	}
	createRangeF64TwoPi := func(value float64, symbol rune) f64_2p.RangeF64TwoPi {
		if existingSymbol, exists := usedValuesTwoPi[value]; exists {
			panic(fmt.Sprintf("El valor %.4f ja està assignat al símbol '%s' (intentat assignar a '%s')", value, string(existingSymbol), string(symbol)))
		}
		usedValuesTwoPi[value] = symbol
		ret, _ := f64_2p.NewRangeF64TwoPi(value)
		return *ret
	}

	symbols := []SymbolEmbedding{
		{'9', createRangeF64One(1.0000000, '9'), createRangeF64TwoPi(6.2831853, '9')},
		{'á', createRangeF64One(0.9010200, 'á'), createRangeF64TwoPi(5.6548668, 'á')},
		{'à', createRangeF64One(0.9010100, 'à'), createRangeF64TwoPi(5.6548658, 'à')},
		{'a', createRangeF64One(0.9010000, 'a'), createRangeF64TwoPi(5.6548650, 'a')},
		{'<', createRangeF64One(0.8010400, '<'), createRangeF64TwoPi(5.0273234, '<')},
		{'k', createRangeF64One(0.8010301, 'k'), createRangeF64TwoPi(5.0273225, 'k')},
		{'[', createRangeF64One(0.8010300, '['), createRangeF64TwoPi(5.0273220, '[')},
		{'t', createRangeF64One(0.8010201, 't'), createRangeF64TwoPi(5.0273211, 't')},
		{'{', createRangeF64One(0.8010200, '{'), createRangeF64TwoPi(5.0273200, '{')},
		{'(', createRangeF64One(0.8010101, '('), createRangeF64TwoPi(5.0273191, '(')},
		{'p', createRangeF64One(0.8010100, 'p'), createRangeF64TwoPi(5.0273190, 'p')},
		{'8', createRangeF64One(0.7777778, '8'), createRangeF64TwoPi(4.8869219, '8')},
		{'😎', createRangeF64One(0.7000500, '😎'), createRangeF64TwoPi(4.3982297, '😎')},
		{'🤗', createRangeF64One(0.7000400, '🤗'), createRangeF64TwoPi(4.3982287, '🤗')},
		{'😂', createRangeF64One(0.7000300, '😂'), createRangeF64TwoPi(4.3982277, '😂')},
		{'😍', createRangeF64One(0.7000200, '😍'), createRangeF64TwoPi(4.3982267, '😍')},
		{'😊', createRangeF64One(0.7000100, '😊'), createRangeF64TwoPi(4.3982257, '😊')},
		{'g', createRangeF64One(0.6010300, 'g'), createRangeF64TwoPi(3.7763420, 'g')},
		{'d', createRangeF64One(0.6010200, 'd'), createRangeF64TwoPi(3.7763410, 'd')},
		{'b', createRangeF64One(0.6010100, 'b'), createRangeF64TwoPi(3.7763400, 'b')},
		{'✈', createRangeF64One(0.6000400, '✈'), createRangeF64TwoPi(3.7699112, '✈')},
		{'✅', createRangeF64One(0.6000300, '✅'), createRangeF64TwoPi(3.7699102, '✅')},
		{'🔥', createRangeF64One(0.6000200, '🔥'), createRangeF64TwoPi(3.7699092, '🔥')},
		{'😀', createRangeF64One(0.6000100, '😀'), createRangeF64TwoPi(3.7699082, '😀')},
		{'7', createRangeF64One(0.5555556, '7'), createRangeF64TwoPi(3.4906585, '7')},
		{'^', createRangeF64One(0.5011100, '^'), createRangeF64TwoPi(3.1492676, '^')},
		{'|', createRangeF64One(0.5011000, '|'), createRangeF64TwoPi(3.1492654, '|')},
		{'&', createRangeF64One(0.5010900, '&'), createRangeF64TwoPi(3.1492632, '&')},
		{'%', createRangeF64One(0.5010800, '%'), createRangeF64TwoPi(3.1492610, '%')},
		{'=', createRangeF64One(0.5010500, '='), createRangeF64TwoPi(3.1492545, '=')},
		{'/', createRangeF64One(0.5010400, '/'), createRangeF64TwoPi(3.1492523, '/')},
		{'*', createRangeF64One(0.5010300, '*'), createRangeF64TwoPi(3.1492501, '*')},
		{'-', createRangeF64One(0.5010201, '-'), createRangeF64TwoPi(3.1492479, '-')},
		{'é', createRangeF64One(0.5010200, 'é'), createRangeF64TwoPi(3.1492477, 'é')},
		{'è', createRangeF64One(0.5010101, 'è'), createRangeF64TwoPi(3.1492455, 'è')},
		{'+', createRangeF64One(0.5010100, '+'), createRangeF64TwoPi(3.1492453, '+')},
		{'e', createRangeF64One(0.5010000, 'e'), createRangeF64TwoPi(3.1492431, 'e')},
		{'🎉', createRangeF64One(0.5000400, '🎉'), createRangeF64TwoPi(3.1415926, '🎉')},
		{'🌟', createRangeF64One(0.5000300, '🌟'), createRangeF64TwoPi(3.1415904, '🌟')},
		{'💔', createRangeF64One(0.5000200, '💔'), createRangeF64TwoPi(3.1415882, '💔')},
		{'🧡', createRangeF64One(0.5000100, '🧡'), createRangeF64TwoPi(3.1415860, '🧡')},
		{'x', createRangeF64One(0.4010300, 'x'), createRangeF64TwoPi(2.5197115, 'x')},
		{'s', createRangeF64One(0.4010200, 's'), createRangeF64TwoPi(2.5197093, 's')},
		{'f', createRangeF64One(0.4010100, 'f'), createRangeF64TwoPi(2.5197071, 'f')},
		{'↕', createRangeF64One(0.4000600, '↕'), createRangeF64TwoPi(2.5132741, '↕')},
		{'↔', createRangeF64One(0.4000500, '↔'), createRangeF64TwoPi(2.5132731, '↔')},
		{'↓', createRangeF64One(0.4000400, '↓'), createRangeF64TwoPi(2.5132721, '↓')},
		{'↑', createRangeF64One(0.4000300, '↑'), createRangeF64TwoPi(2.5132711, '↑')},
		{'←', createRangeF64One(0.4000200, '←'), createRangeF64TwoPi(2.5132701, '←')},
		{'→', createRangeF64One(0.4000100, '→'), createRangeF64TwoPi(2.5132691, '→')},
		{'6', createRangeF64One(0.3333333, '6'), createRangeF64TwoPi(2.0943951, '6')},
		{'£', createRangeF64One(0.3010400, '£'), createRangeF64TwoPi(1.8901198, '£')},
		{'¥', createRangeF64One(0.3010300, '¥'), createRangeF64TwoPi(1.8901165, '¥')},
		{'$', createRangeF64One(0.3010200, '$'), createRangeF64TwoPi(1.8901132, '$')},
		{'€', createRangeF64One(0.3010100, '€'), createRangeF64TwoPi(1.8901099, '€')},
		{'j', createRangeF64One(0.2010300, 'j'), createRangeF64TwoPi(1.2630056, 'j')},
		{'z', createRangeF64One(0.2010200, 'z'), createRangeF64TwoPi(1.2630023, 'z')},
		{'v', createRangeF64One(0.2010100, 'v'), createRangeF64TwoPi(1.2629990, 'v')},
		{'∂', createRangeF64One(0.2000600, '∂'), createRangeF64TwoPi(1.2566371, '∂')},
		{'∞', createRangeF64One(0.2000500, '∞'), createRangeF64TwoPi(1.2566350, '∞')},
		{'π', createRangeF64One(0.2000400, 'π'), createRangeF64TwoPi(1.2566330, 'π')},
		{'∫', createRangeF64One(0.2000300, '∫'), createRangeF64TwoPi(1.2566310, '∫')},
		{'√', createRangeF64One(0.2000200, '√'), createRangeF64TwoPi(1.2566290, '√')},
		{'∑', createRangeF64One(0.2000100, '∑'), createRangeF64TwoPi(1.2566270, '∑')},
		{'5', createRangeF64One(0.1111111, '5'), createRangeF64TwoPi(0.6981317, '5')},
		{'¿', createRangeF64One(0.1010800, '¿'), createRangeF64TwoPi(0.6346652, '¿')},
		{'¡', createRangeF64One(0.1010700, '¡'), createRangeF64TwoPi(0.6346620, '¡')},
		{'?', createRangeF64One(0.1010600, '?'), createRangeF64TwoPi(0.6346588, '?')},
		{'!', createRangeF64One(0.1010500, '!'), createRangeF64TwoPi(0.6346556, '!')},
		{';', createRangeF64One(0.1010400, ';'), createRangeF64TwoPi(0.6346524, ';')},
		{':', createRangeF64One(0.1010300, ':'), createRangeF64TwoPi(0.6346492, ':')},
		{',', createRangeF64One(0.1010200, ','), createRangeF64TwoPi(0.6346460, ',')},
		{'.', createRangeF64One(0.1010100, '.'), createRangeF64TwoPi(0.6346428, '.')},
		{'ï', createRangeF64One(0.0010200, 'ï'), createRangeF64TwoPi(0.0064132, 'ï')},
		{'í', createRangeF64One(0.0010100, 'í'), createRangeF64TwoPi(0.0062821, 'í')},
		{'i', createRangeF64One(0.0010000, 'i'), createRangeF64TwoPi(0.0062832, 'i')},
		{'🤔', createRangeF64One(0.0000500, '🤔'), createRangeF64TwoPi(0.0003142, '🤔')},
		{'☆', createRangeF64One(0.0000400, '☆'), createRangeF64TwoPi(0.0002513, '☆')},
		{'😴', createRangeF64One(0.0000450, '😴'), createRangeF64TwoPi(0.0002827, '😴')},
		{'★', createRangeF64One(0.0000300, '★'), createRangeF64TwoPi(0.0001885, '★')},
		{'😲', createRangeF64One(0.0000350, '😲'), createRangeF64TwoPi(0.0002200, '😲')},
		{'✗', createRangeF64One(0.0000200, '✗'), createRangeF64TwoPi(0.0001257, '✗')},
		{'😕', createRangeF64One(0.0000250, '😕'), createRangeF64TwoPi(0.0001571, '😕')},
		{'✓', createRangeF64One(0.0000100, '✓'), createRangeF64TwoPi(0.0000628, '✓')},
		{'😐', createRangeF64One(0.0000150, '😐'), createRangeF64TwoPi(0.0000942, '😐')},
		{'¤', createRangeF64One(-0.1009100, '¤'), createRangeF64TwoPi(-0.6344532, '¤')},
		{'§', createRangeF64One(-0.1009200, '§'), createRangeF64TwoPi(-0.6344564, '§')},
		{'°', createRangeF64One(-0.1009300, '°'), createRangeF64TwoPi(-0.6344596, '°')},
		{'`', createRangeF64One(-0.1009400, '`'), createRangeF64TwoPi(-0.6344628, '`')},
		{'~', createRangeF64One(-0.1009500, '~'), createRangeF64TwoPi(-0.6344660, '~')},
		{'_', createRangeF64One(-0.1009600, '_'), createRangeF64TwoPi(-0.6344692, '_')},
		{'\'', createRangeF64One(-0.1009700, '\''), createRangeF64TwoPi(-0.6344724, '\'')},
		{'@', createRangeF64One(-0.1009800, '@'), createRangeF64TwoPi(-0.6344756, '@')},
		{'#', createRangeF64One(-0.1009900, '#'), createRangeF64TwoPi(-0.6344788, '#')},
		{'4', createRangeF64One(-0.1111111, '4'), createRangeF64TwoPi(-0.6981317, '4')},
		{'m', createRangeF64One(-0.2009900, 'm'), createRangeF64TwoPi(-1.2566369, 'm')},
		{'3', createRangeF64One(-0.3333333, '3'), createRangeF64TwoPi(-2.0943951, '3')},
		{'n', createRangeF64One(-0.4009900, 'n'), createRangeF64TwoPi(-2.5197072, 'n')},
		{0x0D, createRangeF64One(-0.5009600, 0x0D), createRangeF64TwoPi(-3.1415874, 0x0D)},
		{0x0A, createRangeF64One(-0.5009700, 0x0A), createRangeF64TwoPi(-3.1415906, 0x0A)},
		{'ó', createRangeF64One(-0.5009799, 'ó'), createRangeF64TwoPi(-3.1415947, 'ó')},
		{0x09, createRangeF64One(-0.5009800, 0x09), createRangeF64TwoPi(-3.1415953, 0x09)},
		{'ò', createRangeF64One(-0.5009899, 'ò'), createRangeF64TwoPi(-3.1415994, 'ò')},
		{0x20, createRangeF64One(-0.5009900, 0x20), createRangeF64TwoPi(-3.1416000, 0x20)},
		{'o', createRangeF64One(-0.5010000, 'o'), createRangeF64TwoPi(-3.1416022, 'o')},
		{'2', createRangeF64One(-0.5555556, '2'), createRangeF64TwoPi(-3.4906585, '2')},
		{'l', createRangeF64One(-0.6009900, 'l'), createRangeF64TwoPi(-3.7699068, 'l')},
		{'😔', createRangeF64One(-0.6999500, '😔'), createRangeF64TwoPi(-4.3982227, '😔')},
		{0x1F, createRangeF64One(-0.6999600, 0x1F), createRangeF64TwoPi(-4.3982259, 0x1F)},
		{'😨', createRangeF64One(-0.6999650, '😨'), createRangeF64TwoPi(-4.3982271, '😨')},
		{'😭', createRangeF64One(-0.6999750, '😭'), createRangeF64TwoPi(-4.3982281, '😭')},
		{0x1E, createRangeF64One(-0.6999700, 0x1E), createRangeF64TwoPi(-4.3982292, 0x1E)},
		{'😡', createRangeF64One(-0.6999850, '😡'), createRangeF64TwoPi(-4.3982303, '😡')},
		{'😢', createRangeF64One(-0.6999950, '😢'), createRangeF64TwoPi(-4.3982314, '😢')},
		{0x1C, createRangeF64One(-0.6999900, 0x1C), createRangeF64TwoPi(-4.3982325, 0x1C)},
		{'1', createRangeF64One(-0.7777778, '1'), createRangeF64TwoPi(-4.8869219, '1')},
		{'>', createRangeF64One(-0.8009600, '>'), createRangeF64TwoPi(-5.0273234, '>')},
		{']', createRangeF64One(-0.8009700, ']'), createRangeF64TwoPi(-5.0273245, ']')},
		{'}', createRangeF64One(-0.8009800, '}'), createRangeF64TwoPi(-5.0273256, '}')},
		{')', createRangeF64One(-0.8009899, ')'), createRangeF64TwoPi(-5.0273267, ')')},
		{'r', createRangeF64One(-0.8009900, 'r'), createRangeF64TwoPi(-5.0273278, 'r')},
		{'ü', createRangeF64One(-0.9009800, 'ü'), createRangeF64TwoPi(-5.6548658, 'ü')},
		{'ú', createRangeF64One(-0.9009900, 'ú'), createRangeF64TwoPi(-5.6548669, 'ú')},
		{'u', createRangeF64One(-0.9010000, 'u'), createRangeF64TwoPi(-5.6548679, 'u')},
		{'0', createRangeF64One(-1.0000000, '0'), createRangeF64TwoPi(-6.2831853, '0')},
	}

	return symbols
}

func main() {
	embeddings := generateEmbeddings()

	for idx, e := range embeddings {
		fmt.Printf("[%3d] Symbol: '%s'\tOne: %.9f\tTwoPi:%.9f\n", idx, string(e.Symbol), e.ValueOne.GetF64Value(), e.ValueTwoPi.GetF64Value()) // , e.Group)
	}
}

// [  0] Symbol: '9'       One: 1.000000000        TwoPi:6.283185300
// [  1] Symbol: 'á'       One: 0.901020000        TwoPi:5.654866800
// [  2] Symbol: 'à'       One: 0.901010000        TwoPi:5.654865800
// [  3] Symbol: 'a'       One: 0.901000000        TwoPi:5.654865000
// [  4] Symbol: '<'       One: 0.801040000        TwoPi:5.027323400
// [  5] Symbol: 'k'       One: 0.801030100        TwoPi:5.027322500
// [  6] Symbol: '['       One: 0.801030000        TwoPi:5.027322000
// [  7] Symbol: 't'       One: 0.801020100        TwoPi:5.027321100
// [  8] Symbol: '{'       One: 0.801020000        TwoPi:5.027320000
// [  9] Symbol: '('       One: 0.801010100        TwoPi:5.027319100
// [ 10] Symbol: 'p'       One: 0.801010000        TwoPi:5.027319000
// [ 11] Symbol: '8'       One: 0.777777800        TwoPi:4.886921900
// [ 12] Symbol: '😎'      One: 0.700050000        TwoPi:4.398229700
// [ 13] Symbol: '🤗'      One: 0.700040000        TwoPi:4.398228700
// [ 14] Symbol: '😂'      One: 0.700030000        TwoPi:4.398227700
// [ 15] Symbol: '😍'      One: 0.700020000        TwoPi:4.398226700
// [ 16] Symbol: '😊'      One: 0.700010000        TwoPi:4.398225700
// [ 17] Symbol: 'g'       One: 0.601030000        TwoPi:3.776342000
// [ 18] Symbol: 'd'       One: 0.601020000        TwoPi:3.776341000
// [ 19] Symbol: 'b'       One: 0.601010000        TwoPi:3.776340000
// [ 20] Symbol: '✈'       One: 0.600040000        TwoPi:3.769911200
// [ 21] Symbol: '✅'      One: 0.600030000        TwoPi:3.769910200
// [ 22] Symbol: '🔥'      One: 0.600020000        TwoPi:3.769909200
// [ 23] Symbol: '😀'      One: 0.600010000        TwoPi:3.769908200
// [ 24] Symbol: '7'       One: 0.555555600        TwoPi:3.490658500
// [ 25] Symbol: '^'       One: 0.501110000        TwoPi:3.149267600
// [ 26] Symbol: '|'       One: 0.501100000        TwoPi:3.149265400
// [ 27] Symbol: '&'       One: 0.501090000        TwoPi:3.149263200
// [ 28] Symbol: '%'       One: 0.501080000        TwoPi:3.149261000
// [ 29] Symbol: '='       One: 0.501050000        TwoPi:3.149254500
// [ 30] Symbol: '/'       One: 0.501040000        TwoPi:3.149252300
// [ 31] Symbol: '*'       One: 0.501030000        TwoPi:3.149250100
// [ 32] Symbol: '-'       One: 0.501020100        TwoPi:3.149247900
// [ 33] Symbol: 'é'       One: 0.501020000        TwoPi:3.149247700
// [ 34] Symbol: 'è'       One: 0.501010100        TwoPi:3.149245500
// [ 35] Symbol: '+'       One: 0.501010000        TwoPi:3.149245300
// [ 36] Symbol: 'e'       One: 0.501000000        TwoPi:3.149243100
// [ 37] Symbol: '🎉'      One: 0.500040000        TwoPi:3.141592600
// [ 38] Symbol: '🌟'      One: 0.500030000        TwoPi:3.141590400
// [ 39] Symbol: '💔'      One: 0.500020000        TwoPi:3.141588200
// [ 40] Symbol: '🧡'      One: 0.500010000        TwoPi:3.141586000
// [ 41] Symbol: 'x'       One: 0.401030000        TwoPi:2.519711500
// [ 42] Symbol: 's'       One: 0.401020000        TwoPi:2.519709300
// [ 43] Symbol: 'f'       One: 0.401010000        TwoPi:2.519707100
// [ 44] Symbol: '↕'       One: 0.400060000        TwoPi:2.513274100
// [ 45] Symbol: '↔'       One: 0.400050000        TwoPi:2.513273100
// [ 46] Symbol: '↓'       One: 0.400040000        TwoPi:2.513272100
// [ 47] Symbol: '↑'       One: 0.400030000        TwoPi:2.513271100
// [ 48] Symbol: '←'       One: 0.400020000        TwoPi:2.513270100
// [ 49] Symbol: '→'       One: 0.400010000        TwoPi:2.513269100
// [ 50] Symbol: '6'       One: 0.333333300        TwoPi:2.094395100
// [ 51] Symbol: '£'       One: 0.301040000        TwoPi:1.890119800
// [ 52] Symbol: '¥'       One: 0.301030000        TwoPi:1.890116500
// [ 53] Symbol: '$'       One: 0.301020000        TwoPi:1.890113200
// [ 54] Symbol: '€'       One: 0.301010000        TwoPi:1.890109900
// [ 55] Symbol: 'j'       One: 0.201030000        TwoPi:1.263005600
// [ 56] Symbol: 'z'       One: 0.201020000        TwoPi:1.263002300
// [ 57] Symbol: 'v'       One: 0.201010000        TwoPi:1.262999000
// [ 58] Symbol: '∂'       One: 0.200060000        TwoPi:1.256637100
// [ 59] Symbol: '∞'       One: 0.200050000        TwoPi:1.256635000
// [ 60] Symbol: 'π'       One: 0.200040000        TwoPi:1.256633000
// [ 61] Symbol: '∫'       One: 0.200030000        TwoPi:1.256631000
// [ 62] Symbol: '√'       One: 0.200020000        TwoPi:1.256629000
// [ 63] Symbol: '∑'       One: 0.200010000        TwoPi:1.256627000
// [ 64] Symbol: '5'       One: 0.111111100        TwoPi:0.698131700
// [ 65] Symbol: '¿'       One: 0.101080000        TwoPi:0.634665200
// [ 66] Symbol: '¡'       One: 0.101070000        TwoPi:0.634662000
// [ 67] Symbol: '?'       One: 0.101060000        TwoPi:0.634658800
// [ 68] Symbol: '!'       One: 0.101050000        TwoPi:0.634655600
// [ 69] Symbol: ';'       One: 0.101040000        TwoPi:0.634652400
// [ 70] Symbol: ':'       One: 0.101030000        TwoPi:0.634649200
// [ 71] Symbol: ','       One: 0.101020000        TwoPi:0.634646000
// [ 72] Symbol: '.'       One: 0.101010000        TwoPi:0.634642800
// [ 73] Symbol: 'ï'       One: 0.001020000        TwoPi:0.006413200
// [ 74] Symbol: 'í'       One: 0.001010000        TwoPi:0.006282100
// [ 75] Symbol: 'i'       One: 0.001000000        TwoPi:0.006283200
// [ 76] Symbol: '🤔'      One: 0.000050000        TwoPi:0.000314200
// [ 77] Symbol: '☆'       One: 0.000040000        TwoPi:0.000251300
// [ 78] Symbol: '😴'      One: 0.000045000        TwoPi:0.000282700
// [ 79] Symbol: '★'       One: 0.000030000        TwoPi:0.000188500
// [ 80] Symbol: '😲'      One: 0.000035000        TwoPi:0.000220000
// [ 81] Symbol: '✗'       One: 0.000020000        TwoPi:0.000125700
// [ 82] Symbol: '😕'      One: 0.000025000        TwoPi:0.000157100
// [ 83] Symbol: '✓'       One: 0.000010000        TwoPi:0.000062800
// [ 84] Symbol: '😐'      One: 0.000015000        TwoPi:0.000094200
// [ 85] Symbol: '¤'       One: -0.100910000       TwoPi:-0.634453200
// [ 86] Symbol: '§'       One: -0.100920000       TwoPi:-0.634456400
// [ 87] Symbol: '°'       One: -0.100930000       TwoPi:-0.634459600
// [ 88] Symbol: '`'       One: -0.100940000       TwoPi:-0.634462800
// [ 89] Symbol: '~'       One: -0.100950000       TwoPi:-0.634466000
// [ 90] Symbol: '_'       One: -0.100960000       TwoPi:-0.634469200
// [ 91] Symbol: '''       One: -0.100970000       TwoPi:-0.634472400
// [ 92] Symbol: '@'       One: -0.100980000       TwoPi:-0.634475600
// [ 93] Symbol: '#'       One: -0.100990000       TwoPi:-0.634478800
// [ 94] Symbol: '4'       One: -0.111111100       TwoPi:-0.698131700
// [ 95] Symbol: 'm'       One: -0.200990000       TwoPi:-1.256636900
// [ 96] Symbol: '3'       One: -0.333333300       TwoPi:-2.094395100
// [ 97] Symbol: 'n'       One: -0.400990000       TwoPi:-2.519707200
// ' 98] SyOne: -0.500960000       TwoPi:-3.141587400
// [ 99] Symbol: '
// '       One: -0.500970000       TwoPi:-3.141590600
// [100] Symbol: 'ó'       One: -0.500979900       TwoPi:-3.141594700
// [101] Symbol: ' '       One: -0.500980000       TwoPi:-3.141595300
// [102] Symbol: 'ò'       One: -0.500989900       TwoPi:-3.141599400
// [103] Symbol: ' '       One: -0.500990000       TwoPi:-3.141600000
// [104] Symbol: 'o'       One: -0.501000000       TwoPi:-3.141602200
// [105] Symbol: '2'       One: -0.555555600       TwoPi:-3.490658500
// [106] Symbol: 'l'       One: -0.600990000       TwoPi:-3.769906800
// [107] Symbol: '😔'      One: -0.699950000       TwoPi:-4.398222700
// [108] Symbol: '▼'       One: -0.699960000       TwoPi:-4.398225900
// [109] Symbol: '😨'      One: -0.699965000       TwoPi:-4.398227100
// [110] Symbol: '😭'      One: -0.699975000       TwoPi:-4.398228100
// [111] Symbol: '▲'       One: -0.699970000       TwoPi:-4.398229200
// [112] Symbol: '😡'      One: -0.699985000       TwoPi:-4.398230300
// [113] Symbol: '😢'      One: -0.699995000       TwoPi:-4.398231400
// [114] Symbol: '∟'       One: -0.699990000       TwoPi:-4.398232500
// [115] Symbol: '1'       One: -0.777777800       TwoPi:-4.886921900
// [116] Symbol: '>'       One: -0.800960000       TwoPi:-5.027323400
// [117] Symbol: ']'       One: -0.800970000       TwoPi:-5.027324500
// [118] Symbol: '}'       One: -0.800980000       TwoPi:-5.027325600
// [119] Symbol: ')'       One: -0.800989900       TwoPi:-5.027326700
// [120] Symbol: 'r'       One: -0.800990000       TwoPi:-5.027327800
// [121] Symbol: 'ü'       One: -0.900980000       TwoPi:-5.654865800
// [122] Symbol: 'ú'       One: -0.900990000       TwoPi:-5.654866900
// [123] Symbol: 'u'       One: -0.901000000       TwoPi:-5.654867900
// [124] Symbol: '0'       One: -1.000000000       TwoPi:-6.283185300
