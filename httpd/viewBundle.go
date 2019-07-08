package main

import (
	"io/ioutil"
	"strings"
)

type viewBundle struct {
	views map[string]string
}

var vb viewBundle

func (vb viewBundle) hasView(name string) bool {
	_, ok := vb.views[name]
	return ok
}

func buildViews() error {
	vb.views = make(map[string]string)
	viewDir := "../views"

	files, err := ioutil.ReadDir(viewDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		name := strings.ReplaceAll(f.Name(), ".tpl", "")
		data, _ := ioutil.ReadFile(viewDir + "/" + f.Name())
		vb.views[name] = string(data)
	}
	return nil
}
