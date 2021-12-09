package factory

import (
	"FullCycle/domain/repository"
)

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
