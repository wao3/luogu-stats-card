package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/wao3/luogu-stats-card/common"
	"github.com/wao3/luogu-stats-card/fetcher"
	"github.com/wao3/luogu-stats-card/model"
	"github.com/wao3/luogu-stats-card/tmpl"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	p, err := parseParamCommon(r)
	if err != nil {
		common.LogError("parseParamCommon, err: %v", err)
		handleError(w, common.ErrInvalidParam)
		return
	}
	userProfile, err := fetcher.FetchUserProfile(p.Id)
	if err != nil {
		common.LogError("fetcher.FetchUserProfile, err: %v", err)
		handleError(w, common.ErrInternal)
		return
	}
	stats := userProfile.ToStats()
	if stats.HidePractice {
		handleError(w, common.ErrPrivacy)
		return
	}
	color := p.GetColors()
	var nameTitle *tmpl.NameTitle
	if !p.IsHideTitle() {
		nameTitle = tmpl.NewNameTitleBoxFromStats(stats, common.GetContainerWidth(p.GetCardWidth()), color)
	}
	statsBox := tmpl.NewBarChartBoxFromStats(stats, common.GetContainerWidth(p.GetCardWidth()))
	card := tmpl.NewCard[tmpl.BarChartBox](nameTitle, statsBox, p.GetCardWidth(), color)
	respBuilder := &strings.Builder{}
	err = tmpl.Tmpl.ExecuteTemplate(respBuilder, tmpl.CardTmpl.Name, card)
	if err != nil {
		common.LogError("tmpl.Tmpl.ExecuteTemplate, err: %v", err)
		handleError(w, common.ErrInternal)
		return
	}
	_, err = w.Write([]byte(respBuilder.String()))
}

func handleError(w http.ResponseWriter, err error) {
	_ = tmpl.Tmpl.ExecuteTemplate(w, tmpl.ErrTmpl.Name, err.Error())
}

func parseParamCommon(r *http.Request) (model.CommonParam, error) {
	p := model.CommonParam{}
	if r == nil || r.URL == nil {
		return p, nil
	}
	var query = r.URL.Query()
	var err error
	p.Id, err = strconv.Atoi(query.Get("id"))
	if err != nil {
		return p, err
	}
	if query.Has("card_width") {
		p.CardWidth = new(int)
		*p.CardWidth, err = strconv.Atoi(query.Get("card_width"))
		if err != nil {
			return p, err
		}
	}
	if query.Has("dark_mode") {
		p.DarkMode = new(bool)
		*p.DarkMode, err = strconv.ParseBool(query.Get("dark_mode"))
		if err != nil {
			return p, err
		}
	}
	if query.Has("hide_title") {
		p.HideTitle = new(bool)
		*p.HideTitle, err = strconv.ParseBool(query.Get("hide_title"))
		if err != nil {
			return p, err
		}
	}
	return p, nil
}
