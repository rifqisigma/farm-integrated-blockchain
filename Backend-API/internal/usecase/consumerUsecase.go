package usecase

type ConsumerUsecase interface {
}

type consumerUsecase struct {
}

func NewConsumerRepository() ConsumerUsecase {
	return &consumerUsecase{}
}
