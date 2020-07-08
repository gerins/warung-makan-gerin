package kategorimenu

import (
	"database/sql"
	"warung_makan_gerin/utils/validation"

	"gopkg.in/validator.v2"
)

type KategoriMenuService struct {
	db               *sql.DB
	KategoriMenuRepo KategoriMenuRepository
}

type KategoriMenuServiceInterface interface {
	GetKategoriMenus() (*[]KategoriMenu, error)
	GetKategoriMenuByID(id string) (*KategoriMenu, error)
	HandlePOSTKategoriMenu(d KategoriMenu) (*KategoriMenu, error)
	HandleUPDATEKategoriMenu(id string, data KategoriMenu) (*KategoriMenu, error)
	HandleDELETEKategoriMenu(id string) (*KategoriMenu, error)
}

func NewKategoriMenuService(db *sql.DB) KategoriMenuServiceInterface {
	return KategoriMenuService{db, NewKategoriMenuRepo(db)}
}

func (s KategoriMenuService) GetKategoriMenus() (*[]KategoriMenu, error) {
	KategoriMenu, err := s.KategoriMenuRepo.HandleGETAllKategoriMenu()
	if err != nil {
		return nil, err
	}

	return KategoriMenu, nil
}

func (s KategoriMenuService) GetKategoriMenuByID(id string) (*KategoriMenu, error) {
	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	KategoriMenu, err := s.KategoriMenuRepo.HandleGETKategoriMenu(id, "A")
	if err != nil {
		return nil, err
	}
	return KategoriMenu, nil
}

func (s KategoriMenuService) HandlePOSTKategoriMenu(d KategoriMenu) (*KategoriMenu, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.KategoriMenuRepo.HandlePOSTKategoriMenu(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s KategoriMenuService) HandleUPDATEKategoriMenu(id string, data KategoriMenu) (*KategoriMenu, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	result, err := s.KategoriMenuRepo.HandleUPDATEKategoriMenu(id, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s KategoriMenuService) HandleDELETEKategoriMenu(id string) (*KategoriMenu, error) {
	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	if _, err := s.KategoriMenuRepo.HandleGETKategoriMenu(id, "A"); err != nil {
		return nil, err
	}

	result, err := s.KategoriMenuRepo.HandleDELETEKategoriMenu(id)
	if err != nil {
		return result, err
	}
	return result, nil
}
