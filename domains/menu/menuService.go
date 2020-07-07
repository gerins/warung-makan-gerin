package menu

import (
	"database/sql"
	"warung_makan_gerin/utils/validation"

	"gopkg.in/validator.v2"
)

type MenuService struct {
	MenuRepo MenuRepository
}

type MenuServiceInterface interface {
	GetMenus() (*[]Menu, error)
	GetMenuByID(id string) (*Menu, error)
	HandlePOSTMenu(d Menu) (*Menu, error)
	HandleUPDATEMenu(id string, data Menu) (*Menu, error)
	HandleDELETEMenu(id string) (*Menu, error)
}

func NewMenuService(db *sql.DB) MenuServiceInterface {
	return MenuService{NewMenuRepo(db)}
}

func (s MenuService) GetMenus() (*[]Menu, error) {
	Menu, err := s.MenuRepo.HandleGETAllMenu()
	if err != nil {
		return nil, err
	}

	return Menu, nil
}

func (s MenuService) GetMenuByID(id string) (*Menu, error) {
	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	Menu, err := s.MenuRepo.HandleGETMenu(id, "A")
	if err != nil {
		return nil, err
	}
	return Menu, nil
}

func (s MenuService) HandlePOSTMenu(d Menu) (*Menu, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	result, err := s.MenuRepo.HandlePOSTMenu(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s MenuService) HandleUPDATEMenu(id string, data Menu) (*Menu, error) {
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	result, err := s.MenuRepo.HandleUPDATEMenu(id, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s MenuService) HandleDELETEMenu(id string) (*Menu, error) {
	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	_, err := s.MenuRepo.HandleGETMenu(id, "A")
	if err != nil {
		return nil, err
	}

	result, err := s.MenuRepo.HandleDELETEMenu(id)
	if err != nil {
		return result, err
	}
	return result, nil
}
