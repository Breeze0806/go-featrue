package numberliteral

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func Test_Help(t *testing.T) {
	funcs := []func(){
		LiteralPrinter,
		GoScanner,
		TextScanner,
		MathBigPrinter,
		StrconvPrinter,
	}

	for _, f := range funcs {
		fmt.Printf("=========%v==========\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		f()
		fmt.Printf("=========%v==========\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	}
	/*
		=== RUN   Test_Help
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.LiteralPrinter==========
		10
		1000
		10000
		2
		10
		20
		2
		8
		10
		2
		2
		8
		16
		(1+2i)
		(-1+2i)
		1000000000
		(-5+0i)
		1073741824
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.LiteralPrinter==========
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.GoScanner==========
		1:1	IDENT	"cos"
		1:4	(	""
		1:5	IDENT	"x"
		1:6	)	""
		1:8	+	""
		1:10	INT	"0b2"
		1:13	*	""
		1:14	IDENT	"sin"
		1:17	(	""
		1:18	IDENT	"x"
		1:19	)	""
		1:21	+	""
		1:23	INT	"100_0000"
		1:31	;	"\n"
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.GoScanner==========
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.TextScanner==========
		example:3:1: if
		example:3:4: a
		example:3:6: >
		example:3:8: 10_0000
		example:3:16: {
		example:4:2: someParsable
		example:4:15: =
		example:4:17: text
		example:5:1: }
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.TextScanner==========
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.MathBigPrinter==========
		2 true
		0 strconv.ParseFloat: parsing "0b010": invalid syntax
		2.5 2 <nil>
		2.5 2 <nil>
		2.5 true
		5/2 true
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.MathBigPrinter==========
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.StrconvPrinter==========
		2 <nil>
		2 <nil>
		0 strconv.ParseFloat: parsing "0b010": invalid syntax
		=========github.com/Breeze0806/go-featrue/go1.13/numberliteral.StrconvPrinter==========
		--- PASS: Test_Help (0.00s)
		PASS
	*/
}
