package viewengine

import (
	"html/template"
	"io"
)

type Renderer interface {
	Render(writer io.Writer, viewName string, data any) error
}

type LayoutRenderer interface {
	Renderer
	// Render the view without applying any layout. i.e. as a standalone view.
	RenderStandalone(writer io.Writer, viewName string, model any) error
}

type View struct {
	name  string
	tmpl  *template.Template
	files []string
}

func NewLayout(name string, files ...string) *View {
	return &View{
		name,
		template.Must(template.ParseFiles(files...)),
		nil,
	}
}

func NewView(viewName string, files ...string) *View {
	return &View{
		viewName,
		nil,
		files,
	}
}

func getFileName(filename, ext string) string {
	if ext == "" {
		return filename
	}
	return filename + "." + ext
}
