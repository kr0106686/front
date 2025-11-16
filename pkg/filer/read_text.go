package filer

import (
	"fmt"
	"text/template"
)

func (f *Filer) ReadTextTemplate(name string) (*template.Template, error) {
	filePath := fmt.Sprintf("%s/%s.%s", f.TmplDir, name, "tmpl")
	return template.ParseFiles(filePath)
}
