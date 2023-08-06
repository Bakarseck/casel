package casel

import (
	"net/http"
	utils "casel/lib/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if utils.ValidateRequest(r, w, "/", http.MethodGet) {
		utils.RenderPage("index", nil, w)
	}
}
