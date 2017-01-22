package sloc

import (
	"bufio"
	"bytes"
	"testing"
)

func TestBlankLine(t *testing.T) {
	var tests = []struct {
		input string
		out   bool
	}{
		{"", true},
		{"		", true},
		{"\n\t\n", true},
		{"   	", true},
		{" f", false},
		{" 		\nblah\n", false},
	}

	for _, test := range tests {
		if r := isBlankLine(test.input); r != test.out {
			t.Errorf("blankLine(%q) = %v; want %v", test.input, r, test.out)
		}
	}
}

func TestSingleCommentLine(t *testing.T) {
	var tests = []struct {
		input []byte
		out   bool
	}{
		{[]byte("a line without a comment"), false},
		{[]byte("// a comment line"), true},
		{[]byte("a line with a // comment at the end"), false},
		{[]byte("		// some whitespace"), true},
		{[]byte("		'//in a string'"), false},
		{[]byte("'//in a string'"), false},
	}
	for _, test := range tests {
		if r := rmatch(test.input, CStyleComments.SingleComment); r != test.out {
			t.Errorf("singleInputLine(%q) = %v; want %v", test.input, r, test.out)
		}
	}
}

func TestMultiCommentLine(t *testing.T) {
	var tests = []struct {
		input []byte
		out   int
	}{
		{[]byte("a\nprogram\nwith\nno\ncomments"), 0},
		{[]byte("/*single multiline comment*/"), 1},
		{[]byte("/*some\nmultlinline comment*/"), 2},
		{[]byte("		/* whitespace the start\n but not the end */"), 2},
		{[]byte("		'/* in a string nothing matters */'"), 0},
		{[]byte("'/* in a string there is no meaning */'"), 0},
	}
	for _, test := range tests {
		r := bytes.NewReader(test.input)
		s := bufio.NewScanner(r)
		s.Scan() // scan the first line in - this is expected for skipMultilineComment
		if r := skipMultilineComment(s, &CStyleComments); r != test.out {
			t.Errorf("skipMultilineComment(%q) = %v; want %v", test.input, r, test.out)
		}
	}
}

func TestCountLines(t *testing.T) {
	var tests = []struct {
		input []byte
		lines int
		out   int
	}{
		{[]byte("a\nprogram\nwith\nno\ncomments"), 5, 5},
		{[]byte("/*single multiline comment*/"), 1, 0},
		{[]byte("/*some\nmultlinline comment*/"), 2, 0},
		{[]byte("		/* whitespace the start\n but not the end */"), 2, 0},
		{[]byte("		'/* in a string nothing matters */'"), 1, 1},
		{[]byte("'/* in a string there is no meaning */'"), 1, 1},
		{[]byte("a nice\nprogram\n  /* with some\n long and\nfun comments */\n and\n\n // even more comments\nand done"), 8, 4},
	}
	for _, test := range tests {
		r := bytes.NewReader(test.input)
		if r := CountLines(r, Languages[0]); r != test.out {
			t.Errorf("CountLines(%q) = %v; want %v", test.input, r, test.out)
		}
		r = bytes.NewReader(test.input)
		if r := CountLines(r, nil); r != test.lines {
			t.Errorf("CountLines(%q, nil) = %v; want %v", test.input, r, test.lines)
		}
	}
}
