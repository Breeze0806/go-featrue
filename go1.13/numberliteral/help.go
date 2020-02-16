//https://github.com/golang/proposal/blob/master/design/19308-number-literals.md
package numberliteral

import (
	"fmt"
	"go/scanner"
	"go/token"
	"math/big"
	"strconv"
	"strings"
	textscanner "text/scanner"
)

func LiteralPrinter() {
	var (
		Binary      = 0b010
		Binary1     = 0b010
		Octal       = 0o010
		Hexadecimal = 0x1.0p-1021
		Imaginary1  = 0o2i + 1
		Imaginary2  = 0x2i - 1
		Digit       = 1000_000_000
	)

	fmt.Printf("%b\n", Binary)
	fmt.Printf("%b\n", Octal)
	fmt.Printf("%b\n", Hexadecimal)
	fmt.Printf("%o\n", Binary)
	fmt.Printf("%o\n", Octal)
	fmt.Printf("%x\n", Binary)
	fmt.Printf("%x\n", Octal)
	fmt.Printf("%v\n", Binary)
	fmt.Printf("%v\n", Binary1)
	fmt.Printf("%v\n", Octal)
	fmt.Printf("%v\n", Hexadecimal)
	fmt.Printf("%v\n", Imaginary1)
	fmt.Printf("%v\n", Imaginary2)
	fmt.Printf("%v\n", Digit)
	fmt.Printf("%v\n", Imaginary1*Imaginary2)

	var a int = 1
	fmt.Println(a << 30)
}

func GoScanner() {

	src := []byte("cos(x) + 0b2*sin(x) + 100_0000")
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}

func TextScanner() {
	const src = `
// This is scanned code.
if a > 10_0000 {
	someParsable = text
}`

	var s textscanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != textscanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
}

func MathBigPrinter() {
	fmt.Println(new(big.Int).SetString("0b010", 0))
	fmt.Println(big.ParseFloat("0b010.1", 0, 100, big.ToZero))
	fmt.Println(new(big.Float).Parse("0b010.1", 0))
	fmt.Println(new(big.Float).SetString("0b010.1"))
	fmt.Println(new(big.Rat).SetString("0b010.1"))
}

func StrconvPrinter() {
	fmt.Println(strconv.ParseInt("0b010", 0, 64))
	fmt.Println(strconv.ParseUint("0b010", 0, 64))
	fmt.Println(strconv.ParseFloat("0x010", 64))
}
