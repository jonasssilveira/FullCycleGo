package factory

import (
	"FullCycle/adapter/repository"
	_"FullCycle/adapter/repository"
	"database/sql"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{
		DB: db,
	}
}

func (r *RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repository.NewTransactionRepositoryDb(r.DB)
}
