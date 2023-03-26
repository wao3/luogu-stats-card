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

func GuzhiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	p, err := parseParamGuzhi(r)
	if err != nil {
		common.LogError("parseParamGuzhi, err: %v", err)
		handleError(w, common.ErrInvalidParam)
		return
	}
	userProfile, err := fetcher.FetchUserProfile(p.Id)
	if err != nil {
		common.LogError("fetcher.FetchUserProfile, err: %v", err)
		handleError(w, common.ErrInternal)
		return
	}
	color := p.GetColors()
	guzhi := model.NewGuzhi(userProfile.ToUserInfoModel(), p.Scores)
	var nameTitle *tmpl.NameTitle
	if !p.IsHideTitle() {
		nameTitle = tmpl.NewNameTitleBoxFromGuzhi(guzhi, common.GetContainerWidth(p.GetCardWidth()), color)
	}
	guzhiBox := tmpl.NewBarChartBoxFromGuzhi(guzhi, common.GetContainerWidth(p.GetCardWidth()), color)
	card := tmpl.NewCard[tmpl.BarChartBox](nameTitle, guzhiBox, p.GetCardWidth(), color)
	respBuilder := &strings.Builder{}
	err = tmpl.Tmpl.ExecuteTemplate(respBuilder, tmpl.CardTmpl.Name, card)
	if err != nil {
		common.LogError("tmpl.Tmpl.ExecuteTemplate, err: %v", err)
		handleError(w, common.ErrInternal)
		return
	}
	_, err = w.Write([]byte(respBuilder.String()))
}

func parseParamGuzhi(r *http.Request) (model.GuzhiParam, error) {
	guzhiParam := model.GuzhiParam{}
	commonParam, err := parseParamCommon(r)
	if err != nil {
		return guzhiParam, err
	}
	guzhiParam.CommonParam = commonParam
	if r != nil && r.URL != nil && r.URL.Query().Has("scores") {
		guzhi, err := parseGuzhi(r.URL.Query().Get("scores"))
		if err != nil {
			common.LogError("parseGuzhi, err: %v", err)
			return guzhiParam, err
		}
		guzhiParam.Scores = guzhi
	}
	return guzhiParam, nil
}

func parseGuzhi(raw string) ([]int, error) {
	if len(raw) == 0 {
		return nil, common.ErrGuzhiInvalid
	}
	strings.ReplaceAll(raw, "，", ",")
	guzhiStrList := strings.Split(raw, ",")
	if len(guzhiStrList) != len(model.GuzhiOrder) {
		return nil, common.ErrGuzhiInvalid
	}
	res := make([]int, len(model.GuzhiOrder))
	for i, guzhiStr := range guzhiStrList {
		guzhiStr = strings.TrimSpace(guzhiStr)
		guzhiValue, err := strconv.ParseInt(guzhiStr, 10, 64)
		if err != nil {
			return nil, common.ErrGuzhiInvalid
		}
		res[i] = int(guzhiValue)
	}
	return res, nil
}
