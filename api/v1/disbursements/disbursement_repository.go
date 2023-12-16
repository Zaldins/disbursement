package disbursements

import (
	"github.com/paper.id.disbursement/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	DisbursementRepository struct {
		db *gorm.DB
	}

	IDisbursementRepository interface {
		CreateDisbursement(model *models.Disbursement) (err error)
		LastDisbursement() (lastDisbursement models.Disbursement)
		DisbursementInformation(id uint) (disbursement models.Disbursement)
		UpdateBalance(disbursement models.Disbursement) (tx *gorm.DB)
		UpdateDisbursement(disbursement models.Disbursement) (tx *gorm.DB)
	}
)

func NewDisbursementRepository(db *gorm.DB) *DisbursementRepository {
	return &DisbursementRepository{db: db}
}

func (r *DisbursementRepository) CreateDisbursement(model *models.Disbursement) (err error) {
	result := r.db.Create(model)
	if result.Error != nil {
		logrus.Error(result.Error)
		return result.Error
	}

	return nil
}

func (r *DisbursementRepository) LastDisbursement() (lastDisbursement models.Disbursement) {
	r.db.Last(&lastDisbursement)
	return lastDisbursement
}

func (r *DisbursementRepository) DisbursementInformation(id uint) (disbursement models.Disbursement) {
	result := r.db.Preload("Wallet").First(&disbursement, id)
	if result.Error != nil {
		logrus.Error(result.Error)
	}

	return disbursement
}

func (r *DisbursementRepository) UpdateBalance(disbursement models.Disbursement) (tx *gorm.DB) {
	return r.db.Save(&disbursement.Wallet)
}

func (r *DisbursementRepository) UpdateDisbursement(disbursement models.Disbursement) (tx *gorm.DB) {
	return r.db.Save(&disbursement)
}