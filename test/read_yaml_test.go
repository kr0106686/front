package test

import (
	"html/template"
	"os"
	"path"
	"testing"

	"github.com/kr0106686/front/pkg/filer"
	"github.com/stretchr/testify/require"
)

func Test_Read(t *testing.T) {
	err := os.Setenv("ENV_PATH", "/home/my/.env")
	require.NoError(t, err)

	filer, err := filer.NewFiler("front")
	require.NoError(t, err)

	layout := template.New("layout")
	layout.ParseFiles(path.Join(filer.TmplDir, "layout.html"))
	layout.ParseFiles(path.Join(filer.TmplDir, "head.html"))
	layout.ParseFiles(path.Join(filer.TmplDir, "header.html"))
	layout.ParseFiles(path.Join(filer.TmplDir, "main.html"))

	home, err := filer.ReadYaml("home")
	require.NoError(t, err)

	err = layout.ExecuteTemplate(os.Stdout, "layout", home)
	require.NoError(t, err)
	// t.Log(tmpl)
}
