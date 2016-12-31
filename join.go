package main

import (
	"fmt"
	"os"
	"strings"
)

func Join(left, right [][]string, leftCol, rightCol int, trimAndIgnoreCase bool) (result [][]string) {
	// Collect values from right
	rightRowVal := make(map[string]int)
	for rightRow, fields := range right {
		if len(fields) <= rightCol {
			// No enough values on this row
			fmt.Fprintf(os.Stderr, "Right file row %d does not have column %d\n", rightRow, rightCol)
			continue
		}
		rightField := fields[rightCol]
		if trimAndIgnoreCase {
			rightField = strings.TrimSpace(strings.ToLower(rightField))
		}
		if val, exists := rightRowVal[rightField]; exists {
			fmt.Fprintf(os.Stderr, "Value from right file row %d ('%s') overrides previous value defined in row %d\n", rightRow, rightField, val)
		}
		rightRowVal[rightField] = rightRow
	}
	// Concatenate columns from right to left
	result = make([][]string, len(left))
	for leftRow, fields := range left {
		if len(fields) <= leftCol {
			// No enough values on this row
			fmt.Fprintf(os.Stderr, "Left file row %d does not have column %d, it is written down as-is.\n", leftRow, leftCol)
			result[leftRow] = fields
			continue
		}
		leftField := fields[rightCol]
		if trimAndIgnoreCase {
			leftField = strings.TrimSpace(strings.ToLower(leftField))
		}
		// Look up and concatenate fields
		rightRow, exists := rightRowVal[leftField]
		if !exists {
			fmt.Fprintf(os.Stderr, "Value from left file row %d ('%s') does not join, it is written down as-is.\n", leftRow, leftField)
			result[leftRow] = fields
			continue
		}
		rightRowFields := right[rightRow]
		concatenated := make([]string, 0, len(fields)+len(rightRowFields)-1)
		concatenated = append(concatenated, fields...)
		concatenated = append(concatenated, rightRowFields[:rightCol]...)
		concatenated = append(concatenated, rightRowFields[rightCol+1:]...)
		result[leftRow] = concatenated
	}
	return
}
