package account

import (
	"net/http"
	"strconv"
)

func get(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	var limit int
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	} else {
		limit = 0
	}
	_ = limit

}
