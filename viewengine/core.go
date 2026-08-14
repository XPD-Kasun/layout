package viewengine

import "io"

type ViewEngine interface {
	Render(writer io.Writer, viewName string, data any) error
}
