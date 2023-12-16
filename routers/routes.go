package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func API(router *gin.RouterGroup, db *gorm.DB) {

	disbursementsDI := Disbursements(db)

	v1 := router.Group("/v1")
	{
		disbursement := v1.Group("disbursement")
		{
			disbursement.POST("", disbursementsDI.Disbursement)
		}
	}
}