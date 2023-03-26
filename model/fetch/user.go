package fetch

import (
	"github.com/wao3/luogu-stats-card/model"
)

type UserProfileData struct {
	User              UserInfoData   `json:"user"`
	PassedProblems    *[]ProblemData `json:"passedProblems,omitempty"`
	SubmittedProblems *[]ProblemData `json:"submittedProblems,omitempty"`
}

func (u *UserProfileData) ToUserInfoModel() model.UserInfo {
	return model.UserInfo{
		Name:      u.User.Name,
		ColorName: u.User.Color,
		CcfLevel:  u.User.CcfLevel,
	}
}

func (u *UserProfileData) ToStats() *model.Stats {
	var stats = &model.Stats{
		User: u.ToUserInfoModel(),
	}
	if u.PassedProblems == nil {
		stats.HidePractice = true
		return stats
	}
	passed := make(map[model.Level]int)
	for _, p := range *u.PassedProblems {
		passed[model.Level(p.Difficulty)]++
	}
	stats.Passed = passed
	return stats
}

type UserInfoData struct {
	RegisterTime          int           `json:"registerTime"`
	Introduction          string        `json:"introduction"`
	Prize                 []interface{} `json:"prize"`
	FollowingCount        int           `json:"followingCount"`
	FollowerCount         int           `json:"followerCount"`
	Ranking               int           `json:"ranking"`
	BlogAddress           string        `json:"blogAddress"`
	PassedProblemCount    int           `json:"passedProblemCount"`
	SubmittedProblemCount int           `json:"submittedProblemCount"`
	UID                   int           `json:"uid"`
	Name                  string        `json:"name"`
	Slogan                string        `json:"slogan"`
	Badge                 interface{}   `json:"badge"`
	IsAdmin               bool          `json:"isAdmin"`
	IsBanned              bool          `json:"isBanned"`
	Color                 string        `json:"color"`
	CcfLevel              int           `json:"ccfLevel"`
	Background            string        `json:"background"`
}

type ProblemData struct {
	Pid        string `json:"pid"`
	Title      string `json:"title"`
	Difficulty int    `json:"difficulty"`
	FullScore  int    `json:"fullScore"`
	Type       string `json:"type"`
}
