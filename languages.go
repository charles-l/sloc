package sloc

var CStyleComments = CommentStyle{
	"\\/\\/", // ughhh - go regex
	[]string{"\\/\\*", "\\*\\/"},
}

var Languages = []*Language{
	&Language{"C", []string{"c", "h"}, &CStyleComments},
	&Language{"C++", []string{"cpp", "cc", "h", "hh"}, &CStyleComments},
	&Language{"Java", []string{"java"}, &CStyleComments},
	&Language{"JavaScript", []string{"js"}, &CStyleComments},
	&Language{"Go", []string{"go"}, &CStyleComments},
	&Language{"Ruby", []string{"rb"}, &CommentStyle{"#", []string{"=begin", "^=end"}}},
	&Language{"Python", []string{"py"}, &CommentStyle{"#", []string{"=begin", "^=end"}}},
	&Language{"Lua", []string{"lua"}, &CommentStyle{"--", []string{"--[[", "]]--"}}},
	&Language{"Lisp", []string{"el", "lsp", "scm"}, &CommentStyle{";", []string{}}},
	&Language{"Shell", []string{"sh"}, &CommentStyle{"#", []string{}}},
}

func isAny(s string, a []string) bool {
	for _, e := range a {
		if e == s {
			return true
		}
	}
	return false
}

func GetLanguage(ext string) *Language {
	for _, l := range Languages {
		if isAny(ext, l.Ext) {
			return l
		}
	}
	return nil
}
