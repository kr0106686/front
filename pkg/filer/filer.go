package filer

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert/yaml"
)

type Filer struct {
	TmplDir    string
	ProjectDir string
	StaticDir  string
}

func NewFiler(projectName string) (*Filer, error) {
	envPath := os.Getenv("ENV_PATH")
	if err := godotenv.Load(envPath); err != nil {
		return nil, err
	}

	staticDir := os.Getenv("STATIC_DIR")

	return &Filer{
		TmplDir:    path.Join(staticDir, "templates"),
		ProjectDir: path.Join(staticDir, projectName),
		StaticDir:  staticDir,
	}, nil
}

func (f *Filer) ReadYaml(name string) (map[string]interface{}, error) {
	filePath := fmt.Sprintf("%s/%s.yaml", f.ProjectDir, name)
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var dd map[string]interface{}

	if err := yaml.Unmarshal(b, &dd); err != nil {
		return nil, err
	}
	return dd, nil
}

func (f *Filer) ReadHtmlTemplate(name string) (*template.Template, error) {
	filePath := fmt.Sprintf("%s/%s.%s", f.TmplDir, name, "html")
	return template.ParseFiles(filePath)
}
