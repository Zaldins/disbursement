//go:build wireinject
// +build wireinject

package routers

import (
	"github.com/google/wire"
	"github.com/paper.id.disbursement/api/v1/disbursements"
	"gorm.io/gorm"
)

func Disbursements(db *gorm.DB) *disbursements.DisbursementHandler {
	panic(wire.Build(wire.NewSet(
		disbursements.NewDisbursementRepository,
		disbursements.NewDisbursementUseCase,
		disbursements.NewDisbursementHandler,
		wire.Bind(new(disbursements.IDisbursementRepository), new(*disbursements.DisbursementRepository)),
		wire.Bind(new(disbursements.IDisbursementUseCase), new(*disbursements.DisbursementUseCase)),
		wire.Bind(new(disbursements.IDisbursementHandler), new(*disbursements.DisbursementHandler)),
	)))
	return &disbursements.DisbursementHandler{}
}