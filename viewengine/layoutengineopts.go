package viewengine

type Opt func(*layoutViewEngine)

func WithViewDir(viewDir string) Opt {
	return func(tve *layoutViewEngine) {
		tve.viewDir = viewDir
	}
}

func WithViews(views []*View) Opt {
	return func(lve *layoutViewEngine) {
		for _, view := range views {
			lve.fragmentViews[view.name] = view
		}
	}
}

func WithExt(ext string) Opt {
	return func(tve *layoutViewEngine) {
		tve.ext = ext
	}
}

func WithLayoutName(layoutName string) Opt {
	return func(tve *layoutViewEngine) {
		tve.layoutName = layoutName
	}
}

func WithLayout(layout *View) Opt {
	return func(lve *layoutViewEngine) {
		lve.layoutTemplate = layout
	}
}

func WithCaching(useCaching bool) Opt {
	return func(tve *layoutViewEngine) {
		tve.cacheEnabled = useCaching
	}
}
