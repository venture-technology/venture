package utils

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

func allDigit(doc string) bool {
	for _, r := range doc {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func IsCNH(doc string) bool {
	if len(doc) != 11 {
		return false
	}
	if !allDigit(doc) {
		return false
	}

	sum := 0
	acc := 9
	for _, r := range doc[:len(doc)-2] {
		sum += toInt(r) * acc
		acc--
	}

	base := 0
	digit1 := sum % 11
	if digit1 == 10 {
		base = -2
	}
	if digit1 > 9 {
		digit1 = 0
	}

	sum = 0
	acc = 1
	for _, r := range doc[:len(doc)-2] {
		sum += toInt(r) * acc
		acc++
	}

	var digit2 int
	if (sum%11)+base < 0 {
		digit2 = 11 + (sum % 11) + base
	}
	if (sum%11)+base >= 0 {
		digit2 = (sum % 11) + base
	}
	if digit2 > 9 {
		digit2 = 0
	}

	return toInt(rune(doc[len(doc)-2])) == digit1 &&
		toInt(rune(doc[len(doc)-1])) == digit2
}

var (
	CPFRegexp  = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
	CNPJRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)

func IsCPF(doc string) bool {
	const (
		size = 9
		pos  = 10
	)

	return isCPFOrCNPJ(doc, CPFRegexp, size, pos)
}

func IsCNPJ(doc string) bool {
	const (
		size = 12
		pos  = 5
	)

	return isCPFOrCNPJ(doc, CNPJRegexp, size, pos)
}

func isCPFOrCNPJ(doc string, pattern *regexp.Regexp, size int, position int) bool {
	if !pattern.MatchString(doc) {
		return false
	}

	cleanNonDigits(&doc)

	if allEq(doc) {
		return false
	}

	d := doc[:size]
	digit := calculateDigit(d, position)

	d = d + digit
	digit = calculateDigit(d, position+1)

	return doc == d+digit
}

func cleanNonDigits(doc *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}

func allEq(doc string) bool {
	base := doc[0]
	for i := 1; i < len(doc); i++ {
		if base != doc[i] {
			return false
		}
	}

	return true
}

func calculateDigit(doc string, position int) string {
	var sum int
	for _, r := range doc {

		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}

func toInt(r rune) int {
	return int(r - '0')
}
