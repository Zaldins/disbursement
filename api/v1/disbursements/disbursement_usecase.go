package disbursements

import (
	"github.com/gin-gonic/gin"
	"github.com/paper.id.disbursement/api/v1/disbursements/dto"
	"github.com/paper.id.disbursement/models"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type (
	DisbursementUseCase struct {
		repo IDisbursementRepository
	}
	IDisbursementUseCase interface {
		Disbursement(ctx *gin.Context, request *dto.DisbursementRequest)
		createDisbursement(recipientID uint, amount float64) (err error)
		processDisbursement(id uint)
	}
)

func NewDisbursementUseCase(repo IDisbursementRepository) *DisbursementUseCase {
	return &DisbursementUseCase{repo: repo}
}

func (s *DisbursementUseCase) Disbursement(ctx *gin.Context, request *dto.DisbursementRequest) {

	// create disbursement
	s.createDisbursement(request.WalletID, request.Amount)

	// retrieve the ID of the last created disbursement
	lastDisbursement := s.repo.LastDisbursement()

	// concurrently process the disbursement
	go s.processDisbursement(lastDisbursement.ID)
}

func (s *DisbursementUseCase) createDisbursement(walletID uint, amount float64) error {
	var dbMutex sync.Mutex

	dbMutex = sync.Mutex{}
	dbMutex.Lock()
	defer dbMutex.Unlock()

	disbursement := &models.Disbursement{
		WalletID: walletID,
		Amount:   amount,
		Status:   "pending",
	}

	return s.repo.CreateDisbursement(disbursement)
}

func (s *DisbursementUseCase) processDisbursement(id uint) {
	var dbMutex sync.Mutex
	dbMutex.Lock()
	defer dbMutex.Unlock()

	// get latest disbursement information
	disbursement := s.repo.DisbursementInformation(id)

	// check if the user has sufficient balance
	if disbursement.Wallet.Balance < disbursement.Amount {
		disbursement.Status = "failed"
	} else {
		// simulate some processing time
		time.Sleep(2 * time.Second)

		// update user balance
		disbursement.Wallet.Balance -= disbursement.Amount

		result := s.repo.UpdateBalance(disbursement)

		// simulate success or failure
		if result.Error != nil {
			logrus.Error(result.Error)
			disbursement.Status = "failed"
		} else {
			disbursement.Status = "completed"
		}
	}

	// update
	result := s.repo.UpdateDisbursement(disbursement)
	if result.Error != nil {
		logrus.Error(result.Error)
		return
	}
}