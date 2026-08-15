package layout

import "github.com/XPD-Kasun/layout/viewengine"

func New(opts ...viewengine.Opt) viewengine.LayoutRenderer {
	return viewengine.New()
}
