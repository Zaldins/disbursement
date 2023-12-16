package disbursements

import (
	"github.com/gin-gonic/gin"
	"github.com/paper.id.disbursement/api/v1/disbursements/dto"
	"github.com/paper.id.disbursement/helpers"
	"net/http"
)

type (
	DisbursementHandler struct {
		useCase IDisbursementUseCase
	}

	IDisbursementHandler interface {
		Disbursement(ctx *gin.Context)
	}
)

func NewDisbursementHandler(useCase IDisbursementUseCase) *DisbursementHandler {
	return &DisbursementHandler{useCase: useCase}
}

func (c *DisbursementHandler) Disbursement(ctx *gin.Context) {
	var (
		request dto.DisbursementRequest
		errInfo []string
	)

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "Disbursement request received and processing...",
	}

	// bind
	if err := ctx.ShouldBindJSON(&request); err != nil {
		errInfo = helpers.ErrorWrapper(errInfo, "no body payload")
		helpers.SendBack(ctx, struct{}{}, errInfo, http.StatusBadRequest)
		return
	}

	// inform client that disbursement received and processing
	helpers.SendBack(ctx, resp, []string{}, http.StatusOK)

	c.useCase.Disbursement(ctx, &request)
	return
}