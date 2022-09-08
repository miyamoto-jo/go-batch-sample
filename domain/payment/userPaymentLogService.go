package payment

type UserPaymentLogService struct {
	DBRepository  IDBRepository
	APIRepository IAPIRepository
}
