package viewengine

import (
	"html/template"
	"io"
	"path"
)

type layoutViewEngine struct {
	cached map[string]*template.Template
	// Layout filename in viewDir. Mutually exclusive with layoutTemplate.
	layoutName string
	// Template for the given layout. If layoutName is used this should be set to nil.
	layoutTemplate *View
	// Support for logical view. A logical view is a view constructed
	// with multiple files and identified by a logical name.
	fragmentViews map[string]*View
	// Extension for the template files. Eg: For product.html, ext is html
	ext string
	// Directory containing the views. When we refer views, we dont include this path segment.
	// Eg: View name of `views/products/list.html` is `products/list`, when `views` folder
	// is the `viewDir`
	viewDir string
	// Enables caching of parsed templates. Useful for hot reloading in dev mode.
	// Recommended to turn on in production. Default is true
	cacheEnabled bool
}

// var s LayoutRenderer = &layoutViewEngine{}

// RenderStandalone implements [LayoutRenderer].
func (t *layoutViewEngine) RenderStandalone(writer io.Writer, viewName string, model any) error {

	if t.cacheEnabled {
		if tmpl, ok := t.cached[viewName+"_stand"]; ok {
			err := tmpl.ExecuteTemplate(writer, viewName, model)
			if err != nil {
				return err
			}
		}
	}

	var tmplt *template.Template
	var err error

	if v, ok := t.fragmentViews[viewName]; ok {

		var rawFiles []string
		for _, file := range v.files {
			rawFiles = append(rawFiles, path.Join(t.viewDir, getFileName(file, t.ext)))
		}
		tmplt, err = t.layoutTemplate.tmpl.ParseFiles(rawFiles...)
		if err != nil {
			return err
		}

	} else {

		viewFileName := path.Join(t.viewDir, getFileName(viewName, t.ext))
		tmplt, err = template.ParseFiles(viewFileName)
		if err != nil {
			return err
		}
	}

	// use _stand suffix to differentiate from same name collision
	// from Render() registration
	if t.cacheEnabled {
		t.cached[viewName+"_stand"] = tmplt
	}

	return tmplt.ExecuteTemplate(writer, viewName, model)

}

// Render implements [LayoutRenderer].
func (t *layoutViewEngine) Render(writer io.Writer, viewName string, data any) error {

	var tmplt *template.Template
	var err error

	if v, ok := t.fragmentViews[viewName]; ok {

		var rawFiles []string
		for _, file := range v.files {
			rawFiles = append(rawFiles, path.Join(t.viewDir, getFileName(file, t.ext)))
		}
		tmplt, err = t.layoutTemplate.tmpl.Clone()
		if err != nil {
			return err
		}
		tmplt, err = tmplt.ParseFiles(rawFiles...)
		if err != nil {
			return err
		}

	} else {
		viewFile := path.Join(t.viewDir, getFileName(viewName, t.ext))
		tmplt, err = t.layoutTemplate.tmpl.Clone()
		if err != nil {
			return err
		}
		tmplt, err = tmplt.ParseFiles(viewFile)
		if err != nil {
			return err
		}
	}
	return tmplt.ExecuteTemplate(writer, t.layoutName, data)

	//return fmt.Errorf("layout: Could not find the view '%s'", viewName)
}

func New(options ...Opt) LayoutRenderer {

	var lve = &layoutViewEngine{
		layoutName:     "layout",
		cached:         make(map[string]*template.Template),
		layoutTemplate: nil,
		fragmentViews:  map[string]*View{},
		ext:            "html",
		viewDir:        "views",
		cacheEnabled:   true,
	}

	for _, opt := range options {
		opt(lve)
	}

	if lve.layoutTemplate == nil {
		lve.layoutTemplate = NewLayout(
			lve.layoutName,
			path.Join(lve.viewDir, getFileName(lve.layoutName, lve.ext)),
		)
	}

	return lve
}
