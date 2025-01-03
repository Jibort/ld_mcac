// Identificació i codificació dels errors que es poden produïr.
// CreatedAt: 2024/12/11 dc. JIQ[GPT]

package Errors

// Errors de percentatges
const (
	ERR_INVALID_PERCENTAGE = 1001 // "Percentage out of range %.6f. Must be between 0.0 and 1.0."
	ERR_NULL_VALUE         = 1002 // Error: Operació amb valor null
)

// Errors relacionats amb operacions
const (
	ERR_INVALID_ADD_OPERATION = 2001 // "Invalid addition operation between groups."
	ERR_DIVISION_BY_ZERO      = 2002 // "Division by zero is not allowed."
)
