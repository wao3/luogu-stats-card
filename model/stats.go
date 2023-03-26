package model

type Level uint8

const (
	Level0 Level = iota
	Level1
	Level2
	Level3
	Level4
	Level5
	Level6
	Level7
)

var LevelOrder = []Level{Level0, Level1, Level2, Level3, Level4, Level5, Level6, Level7}

type Stats struct {
	User         UserInfo
	HidePractice bool
	Passed       map[Level]int
}

func (l Level) GetLevelName() string {
	levelMap := map[Level]string{
		Level0: "未评定",
		Level1: "入门",
		Level2: "普及-",
		Level3: "普及/提高-",
		Level4: "普及+/提高",
		Level5: "提高+/省选-",
		Level6: "省选/NOI-",
		Level7: "NOI/NOI+/CTSC",
	}
	if name, ok := levelMap[l]; ok {
		return name
	}
	return "未评定"
}

func (l Level) GetLevelColor(colors Colors) string {
	levelMap := map[Level]string{
		Level0: colors.PracticeLevel0,
		Level1: colors.PracticeLevel1,
		Level2: colors.PracticeLevel2,
		Level3: colors.PracticeLevel3,
		Level4: colors.PracticeLevel4,
		Level5: colors.PracticeLevel5,
		Level6: colors.PracticeLevel6,
		Level7: colors.PracticeLevel7,
	}
	if color, ok := levelMap[l]; ok {
		return color
	}
	return colors.PracticeLevel0
}
