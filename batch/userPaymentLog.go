package batch

import "github.com/miyamoto-jo/go-batch-sample/domain/payment"

type userPaymentLog struct {
	service *payment.UserPaymentLogService
}

func NewUserPaymentLog() *userPaymentLog {
	return &userPaymentLog{
		service: &payment.UserPaymentLogService{},
	}
}

func (m *userPaymentLog) Execute(_ []string) error {
	return nil
}
