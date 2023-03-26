package model

type CommonParam struct {
	Id        int
	CardWidth *int
	DarkMode  *bool
	HideTitle *bool
}

func (p *CommonParam) GetCardWidth() int {
	if p.CardWidth == nil {
		return 500
	}
	if *p.CardWidth < 500 {
		return 500
	}
	if *p.CardWidth > 1920 {
		return 1920
	}
	return *p.CardWidth
}

func (p *CommonParam) GetColors() Colors {
	if p.DarkMode == nil || !*p.DarkMode {
		return ColorDefault
	}
	return ColorDark
}

func (p *CommonParam) IsHideTitle() bool {
	if p.HideTitle == nil {
		return false
	}
	return *p.HideTitle
}

type GuzhiParam struct {
	CommonParam
	Scores []int
}
