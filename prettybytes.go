package prettybytes

import (
	"math"
	"strconv"
)

// Options can be used to overwrite defaul values
type Options struct {
	// Format the number as bits instead of bytes.
	// This can be useful when, for example, referring to bit rate.
	// Defaults to false
	Bits bool

	// Include plus sign for positive numbers.
	// Defaults to false
	Signed bool
}

func getUnits(bits bool) []string {
	byteUnits := []string{
		"B",
		"kB",
		"MB",
		"GB",
		"TB",
		"PB",
		"EB",
		"ZB",
		"YB",
	}

	bitUnits := []string{
		"b",
		"kbit",
		"Mbit",
		"Gbit",
		"Tbit",
		"Pbit",
		"Ebit",
		"Zbit",
		"Ybit",
	}

	if bits {
		return bitUnits
	}

	return byteUnits
}

func getPrefix(isNegative, signed bool) string {
	if isNegative {
		return "-"
	}

	if signed {
		return "+"
	}

	return ""
}

// Convert transforms bytes to a human readable string
//
// Ex.: 1337 → 1.34 kB
func Convert(n float64, o ...Options) string {
	var options Options
	if len(o) > 0 {
		options = o[0]
	}

	units := getUnits(options.Bits)

	if options.Signed && n == 0 {
		return "0 " + units[0]
	}

	isNegative := n < 0
	prefix := getPrefix(isNegative, options.Signed)

	if isNegative {
		n = -n
	}

	if n < 1 {
		numberString := strconv.FormatFloat(n, 'f', -1, 64)
		return prefix + numberString + " " + units[0]
	}

	lastIndex := float64(len(units) - 1)
	exponent := math.Min(math.Floor(math.Log10(n)/3), lastIndex)

	exponentInt := int(exponent)

	n = n / math.Pow(1000, exponent)

	numberString := strconv.FormatFloat(n, 'f', 0, 64)

	return prefix + numberString + " " + units[exponentInt]
}
