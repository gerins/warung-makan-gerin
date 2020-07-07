package transaction

import (
	"database/sql"
	"errors"
	"warung_makan_gerin/config/database"
	"warung_makan_gerin/domains/menu"
	"warung_makan_gerin/utils/validation"

	"gopkg.in/validator.v2"
)

type TransactionService struct {
	TransactionRepo TransactionRepository
}

type TransactionServiceInterface interface {
	GetTransactions() (*[]Transaction, error)
	GetTransactionByID(id string) (*Transaction, error)
	HandlePOSTTransaction(d Transaction) (*Transaction, error)
	HandleUPDATETransaction(id string, data Transaction) (*Transaction, error)
	HandleDELETETransaction(id string) (*Transaction, error)
	GetTodayTransaction() (*Laporan, error)
}

func NewTransactionService(db *sql.DB) TransactionServiceInterface {
	return TransactionService{NewTransactionRepo(db)}
}

func (s TransactionService) GetTransactions() (*[]Transaction, error) {
	Transaction, err := s.TransactionRepo.HandleGETAllTransaction()
	if err != nil {
		return nil, err
	}

	return Transaction, nil
}

func (s TransactionService) GetTransactionByID(id string) (*Transaction, error) {
	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	Transaction, err := s.TransactionRepo.HandleGETTransaction(id, "A")
	if err != nil {
		return nil, err
	}
	return Transaction, nil
}

func (s TransactionService) HandlePOSTTransaction(d Transaction) (*Transaction, error) {
	db := database.ConnectDB()
	defer db.Close()

	for _, value := range d.SoldItems {
		if err := validator.Validate(value); err != nil {
			return nil, err
		}

		result, err := menu.NewMenuRepo(db).HandleGETMenu(value.ProductID, "A")
		if err != nil {
			return nil, err
		}
		if result.Stock == 0 {
			return nil, errors.New(result.MenuName + " Sementara Kosong")
		}
		if value.Quantity > result.Stock {
			return nil, errors.New("Stock " + result.MenuName + " Tidak Mencukupi")
		}
	}

	result, err := s.TransactionRepo.HandlePOSTTransaction(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s TransactionService) HandleUPDATETransaction(id string, data Transaction) (*Transaction, error) {
	return nil, errors.New("Update Not Yet Implemented")
	result, err := s.TransactionRepo.HandleUPDATETransaction(id, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s TransactionService) HandleDELETETransaction(id string) (*Transaction, error) {
	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	if _, err := s.TransactionRepo.HandleGETTransaction(id, "A"); err != nil {
		return nil, err
	}

	result, err := s.TransactionRepo.HandleDELETETransaction(id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s TransactionService) GetTodayTransaction() (*Laporan, error) {
	Transaction, err := s.TransactionRepo.HandleGETAllTransactionDaily()
	if err != nil {
		return nil, err
	}

	return Transaction, nil
}
