package sloc

import (
	"bufio"
	"io"
	"regexp"
	"unicode"
)

type CommentStyle struct {
	SingleComment string
	MultiComment  []string
}

type Language struct {
	Name         string
	Ext          []string
	CommentStyle *CommentStyle
}

func isBlankLine(l string) bool {
	for _, r := range l {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func isSingleComment(l []byte, style *CommentStyle) bool {
	m, _ := regexp.Match("^\\s*"+style.SingleComment, l)
	return m
}

func isMultilineCommentStart(style *CommentStyle, l []byte) bool {
	m, _ := regexp.Match("^\\s*"+style.MultiComment[0], l)
	return m
}

func isMultilineCommentEnd(style *CommentStyle, l []byte) bool {
	m, _ := regexp.Match(style.MultiComment[1], l)
	return m
}

// return number of lines in multiline comment
// Note: this function expects s to have scanned the first line already
func skipMultilineComment(s *bufio.Scanner, style *CommentStyle) int {
	// bail if there aren't any multiline comment styles defined
	if len(style.MultiComment) < 2 {
		return 0
	}

	if isMultilineCommentStart(style, s.Bytes()) {
		i := 0
		for {
			i++
			if isMultilineCommentEnd(style, s.Bytes()) {
				return i
			}
			s.Scan()
		}
	} else {
		return 0
	}
}

func CountLines(i io.Reader, l *Language) int {
	s := bufio.NewScanner(i)
	sourceLines := 0

	var shouldSkipLine func(b string) bool
	if l == nil {
		shouldSkipLine = isBlankLine // skip blank lines
	} else {
		shouldSkipLine = func(b string) bool {
			return isBlankLine(b) ||
				isSingleComment([]byte(b), l.CommentStyle) ||
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
