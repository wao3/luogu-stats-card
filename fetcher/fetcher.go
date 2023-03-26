package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wao3/luogu-stats-card/common"
	"github.com/wao3/luogu-stats-card/model/fetch"
)

var HttpClient = http.Client{}

const BaseUrl = "https://www.luogu.com.cn"

func Fetch[Data fetch.DataType](path string, clientId *string) (*Data, error) {
	url := BaseUrl + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		common.LogError("http.NewRequest, err: %v", err)
		return nil, err
	}
	if !req.URL.Query().Has("_contentOnly") {
		q := req.URL.Query()
		q.Add("_contentOnly", "1")
		req.URL.RawQuery = q.Encode()
	}
	if clientId != nil {
		req.Header.Set("cookie", fmt.Sprintf("__client_id=%s", *clientId))
	}
	common.LogInfo("Fetch url: %s", req.URL)
	res, err := HttpClient.Do(req)
	if err != nil {
		common.LogError("HttpClient.Do, err: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		common.LogError("io.ReadAll, err: %v", err)
		return nil, err
	}
	var data fetch.Resp[Data]
	err = json.Unmarshal(body, &data)
	if err != nil {
		common.LogError("json.Unmarshal, err: %v", err)
		return nil, err
	}
	return &data.CurrentData, nil
}
