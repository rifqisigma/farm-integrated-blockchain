package repository

type ConsumerRepository interface {
}

type consumerRepository struct {
}

func NewConsumerRepository() ConsumerRepository {
	return &consumerRepository{}
}
