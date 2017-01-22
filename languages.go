package sloc

var CStyleComments = CommentStyle{
	LinePre("\\/\\/"), // ughhh - go regex
	[]string{LinePre("\\/\\*"), "\\*\\/"},
}

var Languages = []*Language{
	&Language{"C", []string{"c", "h"}, &CStyleComments},
	&Language{"C++", []string{"cpp", "cc", "hh", "hpp"}, &CStyleComments},
	&Language{"Java", []string{"java"}, &CStyleComments},
	&Language{"JavaScript", []string{"js"}, &CStyleComments},
	&Language{"Go", []string{"go"}, &CStyleComments},
	&Language{"Ruby", []string{"rb"}, &CommentStyle{LinePre("#"), []string{"^=begin", "^=end"}}},
	&Language{"Python", []string{"py"}, &CommentStyle{LinePre("#"), []string{"^=begin", "^=end"}}},
	&Language{"Lua", []string{"lua"}, &CommentStyle{LinePre("--"), []string{"--[[", "]]--"}}},
	&Language{"Lisp", []string{"el", "lsp", "scm"}, &CommentStyle{LinePre(";"), nil}},
	&Language{"Shell", []string{"sh"}, &CommentStyle{LinePre("#"), nil}},
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
