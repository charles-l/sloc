// Package sloc provides function for counting source lines of code.
package sloc

import (
	"bufio"
	"io"
	"regexp"
	"unicode"
)

// A CommentStyle for a Language describing how to match a comment
// so it can be ignored in the sloc count.
type CommentStyle struct {
	SingleComment string
	MultiComment  []string
}

// A Language is used for matching an extension and describing
// how to name a languge and match a comment
type Language struct {
	Name         string
	Ext          []string
	CommentStyle *CommentStyle
}

// Define the line prefix (ignoring whitespace)
func LinePre(match string) string {
	return "^\\s*" + match
}

func isBlankLine(l string) bool {
	for _, r := range l {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// regex match (throw away the error)
func rmatch(l []byte, r string) bool {
	m, _ := regexp.Match(r, l)
	return m
}

func multiStart(style *CommentStyle) string {
	return style.MultiComment[0]
}

func multiEnd(style *CommentStyle) string {
	return style.MultiComment[1]
}

// return number of lines in multiline comment
// Note: this function expects s to have scanned the first line already
func skipMultilineComment(s *bufio.Scanner, style *CommentStyle) int {
	// bail if there aren't any multiline comment styles defined
	if len(style.MultiComment) < 2 {
		return 0
	}

	if rmatch(s.Bytes(), multiStart(style)) {
		i := 0
		for {
			i++
			if rmatch(s.Bytes(), multiEnd(style)) {
				return i
			}
			s.Scan()
		}
	} else {
		return 0
	}
}

// Counts the lines of code from a Reader stream.
func CountLines(i io.Reader, l *Language) int {
	s := bufio.NewScanner(i)
	sourceLines := 0

	var shouldSkipLine func(b string) bool
	if l == nil {
		shouldSkipLine = isBlankLine // skip blank lines
	} else {
		shouldSkipLine = func(b string) bool {
			return isBlankLine(b) ||
				rmatch([]byte(b), l.CommentStyle.SingleComment) ||
				skipMultilineComment(s, l.CommentStyle) != 0
		}
	}

	for s.Scan() {
		if !shouldSkipLine(s.Text()) {
			sourceLines++
		}
	}
	return sourceLines
}
