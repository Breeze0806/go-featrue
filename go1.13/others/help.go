package others

var toValidUTF8Tests = []struct {
	input struct {
		s           string
		replacement string
	}
	want string
}{
	{
		input: struct {
			s           string
			replacement string
		}{
			s:           "a☺\xffb☺\xC0\xAFc☺\xff",
			replacement: "",
		},
		want: "a☺b☺c☺",
	},
	{
		input: struct {
			s           string
			replacement string
		}{
			s:           "a☺\xffb☺\xC0\xAFc☺\xff",
			replacement: "日本語",
		},
		want: "a☺日本語b☺日本語c☺日本語",
	},
	{
		input: struct {
			s           string
			replacement string
		}{
			s:           "\xed\xa0\x80",
			replacement: "abc",
		},
		want: "abc",
	},

	{
		input: struct {
			s           string
			replacement string
		}{
			s:           "abc",
			replacement: "\xed\xa0\x80",
		},
		want: "abc",
	},
}

func Equal(a, b []byte) bool {
	return string(a) == string(b)
}
