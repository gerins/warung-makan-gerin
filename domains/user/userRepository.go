package user

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type UserRepo struct {
	db *sql.DB
}

type UserRepository interface {
	HandleGETAllUser() (*[]User, error)
	HandleUserLogin(username, status string) (*User, error)
	HandlePOSTUser(d User) (*User, error)
	HandleUPDATEUser(id string, data User) (*User, error)
	HandleDELETEUser(id string) (*User, error)
}

func NewUserRepo(db *sql.DB) UserRepository {
	return UserRepo{db}
}

// HandleGETAllUser for GET all data from User
func (p UserRepo) HandleGETAllUser() (*[]User, error) {
	var d User
	var AllUser []User

	result, err := p.db.Query("SELECT * FROM user WHERE status=?", "A")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.Username, &d.Password, &d.Status, &d.Created, &d.Updated)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		AllUser = append(AllUser, d)
	}
	return &AllUser, nil
}

// HandleUserLogin for GET single data from User
func (p UserRepo) HandleUserLogin(username, status string) (*User, error) {
	results := p.db.QueryRow("SELECT * FROM user WHERE username=? AND status=?", username, status)

	var d User
	err := results.Scan(&d.ID, &d.Username, &d.Password, &d.Status, &d.Created, &d.Updated)
	if err != nil {
		return nil, errors.New("Username atau Password salah")
	}

	return &d, nil
}

// HandlePOSTUser will POST a new User data
func (p UserRepo) HandlePOSTUser(d User) (*User, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt, _ := tx.Prepare(`INSERT INTO user(username,password) VALUES (?,?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.Username, d.Password)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return p.HandleUserLogin(strconv.Itoa(int(lastInsertID)), "A")
}

// HandleUPDATEUser is used for UPDATE data User
func (p UserRepo) HandleUPDATEUser(id string, data User) (*User, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE User SET username=?, password=? WHERE id=?`,
		data.Username, data.Password)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return p.HandleUserLogin(id, "A")
}

// HandleDELETEUser for DELETE single data from User
func (p UserRepo) HandleDELETEUser(id string) (*User, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("UPDATE User SET status=? WHERE id=?", "NA", id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return p.HandleUserLogin(id, "NA")
}
