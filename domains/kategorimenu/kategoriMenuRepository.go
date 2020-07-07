package kategorimenu

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type KategoriMenuRepo struct {
	db *sql.DB
}

type KategoriMenuRepository interface {
	HandleGETAllKategoriMenu() (*[]KategoriMenu, error)
	HandleGETKategoriMenu(id, status string) (*KategoriMenu, error)
	HandlePOSTKategoriMenu(d KategoriMenu) (*KategoriMenu, error)
	HandleUPDATEKategoriMenu(id string, data KategoriMenu) (*KategoriMenu, error)
	HandleDELETEKategoriMenu(id string) (*KategoriMenu, error)
}

func NewKategoriMenuRepo(db *sql.DB) KategoriMenuRepository {
	return KategoriMenuRepo{db}
}

// HandleGETAllKategoriMenu for GET all data from KategoriMenu
func (p KategoriMenuRepo) HandleGETAllKategoriMenu() (*[]KategoriMenu, error) {
	var d KategoriMenu
	var AllKategoriMenu []KategoriMenu

	result, err := p.db.Query("SELECT * FROM category_menu WHERE status=?", "A")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.CategoryName, &d.Status, &d.Created, &d.Updated)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		resultMenu, _ := p.db.Query("SELECT nama FROM menu WHERE status=? AND category_menu_id=?", "A", d.ID)

		var menuName string
		var allMenuName []string
		for resultMenu.Next() {
			err := resultMenu.Scan(&menuName)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			allMenuName = append(allMenuName, menuName)
		}
		d.ListMenu = allMenuName
		AllKategoriMenu = append(AllKategoriMenu, d)
	}
	return &AllKategoriMenu, nil
}

// HandleGETKategoriMenu for GET single data from KategoriMenu
func (p KategoriMenuRepo) HandleGETKategoriMenu(id, status string) (*KategoriMenu, error) {
	results := p.db.QueryRow("SELECT * FROM category_menu WHERE id=? AND status=?", id, status)

	var d KategoriMenu
	err := results.Scan(&d.ID, &d.CategoryName, &d.Status, &d.Created, &d.Updated)
	if err != nil {
		return nil, errors.New("Category Menu ID Not Found")
	}

	resultMenu, _ := p.db.Query("SELECT nama FROM menu WHERE status=? AND category_menu_id=?", "A", d.ID)
	for resultMenu.Next() {
		var menuName string
		err := resultMenu.Scan(&menuName)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		d.ListMenu = append(d.ListMenu, menuName)
	}

	return &d, nil
}

// HandlePOSTKategoriMenu will POST a new KategoriMenu data
func (p KategoriMenuRepo) HandlePOSTKategoriMenu(d KategoriMenu) (*KategoriMenu, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stmnt, _ := tx.Prepare(`INSERT INTO category_menu(nama) VALUES (?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.CategoryName)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return p.HandleGETKategoriMenu(strconv.Itoa(int(lastInsertID)), "A")
}

// HandleUPDATEKategoriMenu is used for UPDATE data KategoriMenu
func (p KategoriMenuRepo) HandleUPDATEKategoriMenu(id string, data KategoriMenu) (*KategoriMenu, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE category_menu SET nama=? WHERE id=?`, data.CategoryName, id)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	checkAvaibility, err := p.HandleGETKategoriMenu(id, "A")
	if err != nil {
		return nil, err
	}
	return checkAvaibility, nil
}

// HandleDELETEKategoriMenu for DELETE single data from KategoriMenu
func (p KategoriMenuRepo) HandleDELETEKategoriMenu(id string) (*KategoriMenu, error) {
	if _, err := p.HandleGETKategoriMenu(id, "A"); err != nil {
		return nil, err
	}

	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("UPDATE category_menu SET status=? WHERE id=?", "NA", id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return p.HandleGETKategoriMenu(id, "NA")
}
