package model

type GuzhiType string

const (
	GuzhiTypeBasicCredit GuzhiType = "基础信用"
	GuzhiTypePractice    GuzhiType = "练习情况"
	GuzhiTypeContribute  GuzhiType = "社区贡献"
	GuzhiTypeCompetition GuzhiType = "比赛情况"
	GuzhiTypeAchievement GuzhiType = "获得成就"
)

type Guzhi struct {
	User  UserInfo
	Guzhi map[GuzhiType]int
}

var GuzhiOrder = []GuzhiType{GuzhiTypeBasicCredit, GuzhiTypePractice, GuzhiTypeContribute, GuzhiTypeCompetition, GuzhiTypeAchievement}

func NewGuzhi(user UserInfo, guzhi []int) *Guzhi {
	g := &Guzhi{
		User:  user,
		Guzhi: make(map[GuzhiType]int),
	}
	for i, v := range guzhi {
		g.Guzhi[GuzhiOrder[i]] = v
	}
	return g
}

func GetGuzhiColor(guzhi int, colors Colors) string {
	if guzhi >= 80 {
		return colors.GuzhiGreen
	}
	if guzhi >= 60 {
		return colors.GuzhiYellow
	}
	if guzhi >= 30 {
		return colors.GuzhiOrange
	}
	return colors.GuzhiRed
}
