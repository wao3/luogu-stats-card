package handler

import (
	"net/http"

	"github.com/wao3/luogu-stats-card/common"
)

func RecoverHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				common.LogError("panic: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		handlerFunc(w, r)
	}
}
