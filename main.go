package layout

import "github.com/XPD-Kasun/layout/viewengine"

func New(layoutName string) viewengine.ViewEngine {
	return viewengine.NewLayoutViewEngine(layoutName)
}
