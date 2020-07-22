package menu

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type MenuRepo struct {
	db *sql.DB
}

type MenuRepository interface {
	HandleGETAllMenu(keyword, offset, limit, status, orderBy, sort string) (*[]Menu, error)
	HandleGETMenu(id, status string) (*Menu, error)
	HandlePOSTMenu(d Menu) (*Menu, error)
	HandleUPDATEMenu(id string, data Menu) (*Menu, error)
	HandleDELETEMenu(id string) (*Menu, error)
}

func NewMenuRepo(db *sql.DB) MenuRepository {
	return MenuRepo{db}
}

// HandleGETAllMenu for GET all data from Menu
func (p MenuRepo) HandleGETAllMenu(keyword, offset, limit, status, orderBy, sort string) (*[]Menu, error) {
	var d Menu
	var AllMenu []Menu

	queryInput := fmt.Sprintf("SELECT * FROM menu_idx WHERE status=? AND nama LIKE ? ORDER BY %s %s LIMIT %s,%s", orderBy, sort, offset, limit)
	result, err := p.db.Query(queryInput, status, "%"+keyword+"%")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.MenuName, &d.Harga, &d.Stock, &d.Category,
			&d.Status, &d.Created, &d.Updated)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		AllMenu = append(AllMenu, d)
	}
	return &AllMenu, nil
}

// HandleGETMenu for GET single data from Menu
func (p MenuRepo) HandleGETMenu(id, status string) (*Menu, error) {
	results := p.db.QueryRow("SELECT * FROM menu_idx WHERE id=? AND status=?", id, status)

	var d Menu
	err := results.Scan(&d.ID, &d.MenuName, &d.Harga, &d.Stock, &d.Category,
		&d.Status, &d.Created, &d.Updated)
	if err != nil {
		return nil, errors.New("Menu ID Not Found")
	}

	return &d, nil
}

// HandlePOSTMenu will POST a new Menu data
func (p MenuRepo) HandlePOSTMenu(d Menu) (*Menu, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt, _ := tx.Prepare(`INSERT INTO menu(nama,harga,stock,category_menu_id) VALUES (?,?,?,?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.MenuName, d.Harga, d.Stock, d.Category)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return p.HandleGETMenu(strconv.Itoa(int(lastInsertID)), "A")
}

// HandleUPDATEMenu is used for UPDATE data Menu
func (p MenuRepo) HandleUPDATEMenu(id string, data Menu) (*Menu, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE menu SET nama=?, harga=?,stock=? WHERE id=?`,
		data.MenuName, data.Harga, data.Stock, id)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return p.HandleGETMenu(id, "A")
}

// HandleDELETEMenu for DELETE single data from Menu
func (p MenuRepo) HandleDELETEMenu(id string) (*Menu, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("UPDATE menu SET status=? WHERE id=?", "NA", id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return p.HandleGETMenu(id, "NA")
}
