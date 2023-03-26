package model

type UserInfo struct {
	Name      string
	ColorName string
	CcfLevel  int
}

func (u *UserInfo) GetColor(colors Colors) string {
	colorMap := map[string]string{
		"Gray":    colors.NameGray,
		"Blue":    colors.NameBlue,
		"Green":   colors.NameGreen,
		"Orange":  colors.NameOrange,
		"Red":     colors.NameRed,
		"Purple":  colors.NamePurple,
		"Cheater": colors.NameCheater,
	}
	if color, ok := colorMap[u.ColorName]; ok {
		return color
	}
	return colors.NameGray
}

func (u *UserInfo) GetCcfColor(colors Colors) *string {
	if u.CcfLevel >= 3 && u.CcfLevel <= 5 {
		return &colors.CcfGreen
	}
	if u.CcfLevel >= 6 && u.CcfLevel <= 7 {
		return &colors.CcfBlue
	}
	if u.CcfLevel >= 8 {
		return &colors.CcfYellow
	}
	return nil
}
