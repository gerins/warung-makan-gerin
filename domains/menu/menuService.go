package menu

import (
	"database/sql"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"warung_makan_gerin/utils/validation"

	"gopkg.in/validator.v2"
)

type MenuService struct {
	db       *sql.DB
	MenuRepo MenuRepository
}

type MenuServiceInterface interface {
	GetMenus(keyword, page, limit, status, orderBy, sort string) (*TotalMenu, error)
	GetMenuByID(id string) (*Menu, error)
	HandlePOSTMenu(d Menu, user string, uploadedFile multipart.File, handler *multipart.FileHeader) (*Menu, error)
	HandleUPDATEMenu(id string, data Menu) (*Menu, error)
	HandleDELETEMenu(id string) (*Menu, error)
}

func NewMenuService(db *sql.DB) MenuServiceInterface {
	return MenuService{db, NewMenuRepo(db)}
}

func (s MenuService) GetMenus(keyword, page, limit, status, orderBy, sort string) (*TotalMenu, error) {
	// var pageOffset = (page * limit) - limit

	Menu, err := s.MenuRepo.HandleGETAllMenu(keyword, page, limit, status, orderBy, sort)
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

func (s MenuService) HandlePOSTMenu(d Menu, user string, uploadedFile multipart.File, handler *multipart.FileHeader) (*Menu, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	min := 11111111111
	max := 99999999999
	fileLocation := filepath.Join(dir, "files", user+"-"+strconv.Itoa(rand.Intn(max-min+1)+min)+filepath.Ext(handler.Filename))

	log.Println(`FileLocation ->`, fileLocation)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		log.Println(`Error While Coping File to Local Storage`, err)
		return nil, err
	}

	if err := validator.Validate(d); err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := s.MenuRepo.HandlePOSTMenu(d)
	if err != nil {
		log.Println(err)
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
