package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/paper.id.disbursement/helpers"
	"net/http"
)

func NoRoute(ctx *gin.Context) {
	resp := struct {
		Message string `json:"message,omitempty"`
	}{
		Message: "no routes or wrong method",
	}
	helpers.SendBack(ctx, resp, []string{}, http.StatusNotFound)
	return
}