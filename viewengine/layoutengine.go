package viewengine

import (
	"html/template"
	"io"
	"path"
)

type layoutViewEngine struct {
	layoutName string
}

// var s ViewEngine = &layoutViewEngine{}

// Render implements [ViewEngine].
func (t *layoutViewEngine) Render(writer io.Writer, viewName string, data any) error {

	layoutFile := path.Join("views", t.layoutName+".html")
	viewFile := path.Join("views", viewName+".html")
	tmplt, err := template.ParseFiles(layoutFile, viewFile)
	if err != nil {
		return err
	}

	return tmplt.ExecuteTemplate(writer, t.layoutName, data)

	//return fmt.Errorf("layout: Could not find the view '%s'", viewName)
}

func NewLayoutViewEngine(layoutName string) *layoutViewEngine {
	return &layoutViewEngine{layoutName: layoutName}
}
