package casel

import (
	"net/http"
	utils "casel/lib/utils"
	models "casel/lib/models"
)

var err models.Error

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if code, ok := utils.ValidateRequest(r, w, "/", http.MethodGet); ok {
		utils.RenderPage("index", nil, w)
	} else {
		err.MsgError = models.Errors[code]
		utils.RenderPage("error", err, w)
	}
}
