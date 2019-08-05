package main

import "strings"

func buildView(content string) string {
	data, _ := Asset("views/_layout.tpl")

	return strings.ReplaceAll(string(data), "%%CONTENT%%", content)
}
