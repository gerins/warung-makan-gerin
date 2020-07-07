package transaction

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"warung_makan_gerin/domains/menu"
)

type TransactionRepo struct {
	db *sql.DB
}

type TransactionRepository interface {
	HandleGETAllTransaction() (*[]Transaction, error)
	HandleGETTransaction(id, status string) (*Transaction, error)
	HandlePOSTTransaction(d Transaction) (*Transaction, error)
	HandleUPDATETransaction(id string, data Transaction) (*Transaction, error)
	HandleDELETETransaction(id string) (*Transaction, error)
	HandleGETAllTransactionDaily() (*Laporan, error)
}

func NewTransactionRepo(db *sql.DB) TransactionRepository {
	return TransactionRepo{db}
}

// HandleGETAllTransaction for GET all data from Transaction
func (p TransactionRepo) HandleGETAllTransaction() (*[]Transaction, error) {
	var d Transaction
	var AllTransaction []Transaction

	result, err := p.db.Query("SELECT * FROM daftar_transaksi WHERE status=?", "A")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.PeralatanMakan, &d.Total, &d.Status, &d.Created, &d.Updated)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		resultProduct, _ := p.db.Query(`SELECT nama,kategori,harga,extra_pedas,kuantiti,total FROM detail_transaksi_idx WHERE trans_id=?`, d.ID)

		var soldItem SoldItems
		var allSoldItem []SoldItems
		for resultProduct.Next() {
			err := resultProduct.Scan(&soldItem.ProductID, &soldItem.Category, &soldItem.Harga, &soldItem.EkstraPedas, &soldItem.Quantity, &soldItem.Total)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			allSoldItem = append(allSoldItem, soldItem)
		}
		d.SoldItems = allSoldItem
		AllTransaction = append(AllTransaction, d)
	}
	return &AllTransaction, nil
}

// HandleGETTransaction for GET single data from Transaction
func (p TransactionRepo) HandleGETTransaction(id, status string) (*Transaction, error) {
	results := p.db.QueryRow("SELECT * FROM daftar_transaksi WHERE id=? AND status=?", id, status)

	var d Transaction
	err := results.Scan(&d.ID, &d.PeralatanMakan, &d.Total, &d.Status, &d.Created, &d.Updated)
	if err != nil {
		return nil, errors.New("Transaction ID Not Found")
	}

	resultProduct, _ := p.db.Query(`SELECT nama,kategori,harga,extra_pedas,kuantiti,total FROM detail_transaksi_idx WHERE trans_id=?`, d.ID)
	var soldItem SoldItems
	for resultProduct.Next() {
		err := resultProduct.Scan(&soldItem.ProductID, &soldItem.Category, &soldItem.Harga, &soldItem.EkstraPedas, &soldItem.Quantity, &soldItem.Total)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		d.SoldItems = append(d.SoldItems, soldItem)
	}

	return &d, nil
}

// HandlePOSTTransaction will POST a new Transaction data
func (p TransactionRepo) HandlePOSTTransaction(d Transaction) (*Transaction, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := tx.Exec("INSERT INTO daftar_transaksi(peralatan_makan) VALUE(?)", d.PeralatanMakan)
	if err != nil {

		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()

	var Total int
	for _, value := range d.SoldItems {
		product, _ := menu.NewMenuRepo(p.db).HandleGETMenu(value.ProductID, "A")
		_, err = tx.Exec(`INSERT INTO detail_transaksi(daftar_transaksi_id, menu_id, kuantiti, extra_pedas, total) VALUE(?,?,?,?,?)`, lastInsertID, value.ProductID, value.Quantity, value.EkstraPedas, product.Harga*value.Quantity+value.EkstraPedas)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return nil, err
		}
		Total += product.Harga*value.Quantity + value.EkstraPedas

		_, err = tx.Exec(`UPDATE menu SET stock=stock-? WHERE id=?`, value.Quantity, value.ProductID)
	}

	_, err = tx.Exec(`UPDATE daftar_transaksi SET total=? WHERE id=?`, Total+d.PeralatanMakan, lastInsertID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return p.HandleGETTransaction(strconv.Itoa(int(lastInsertID)), "A")
}

// HandleUPDATETransaction is used for UPDATE data Transaction
func (p TransactionRepo) HandleUPDATETransaction(id string, data Transaction) (*Transaction, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(tx)

	// var Total int
	// for _, value := range data.SoldItems {
	// 	product, _ := product.NewProductRepo(p.db).HandleGETProduct(value.ProductID, "A")
	// 	hargaProduct, _ := strconv.Atoi(product.Harga)
	// 	quantitiProduct, _ := strconv.Atoi(value.Quantity)
	// 	_, err = tx.Exec(`UPDATE detail_ransaksi SET  kuantiti=?, total=? WHERE transaksi_penjualan_id=? AND produk_id=?`, value.Quantity, hargaProduct*quantitiProduct, id, value.ProductID)
	// 	if err != nil {
	// 		log.Println(err)
	// 		tx.Rollback()
	// 		return nil, err
	// 	}
	// 	Total += hargaProduct * quantitiProduct
	// }

	// _, err = tx.Exec(`UPDATE transaksi_penjualan SET total_penjualan=? WHERE id=?`, Total, id)
	// if err != nil {
	// 	log.Println(err)
	// 	tx.Rollback()
	// 	return nil, err
	// }

	// tx.Commit()
	checkAvaibility, err := p.HandleGETTransaction(id, "A")
	if err != nil {
		return nil, err
	}
	return checkAvaibility, nil
}

// HandleDELETETransaction for DELETE single data from Transaction
func (p TransactionRepo) HandleDELETETransaction(id string) (*Transaction, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("UPDATE daftar_transaksi SET status=? WHERE id=?", "NA", id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return p.HandleGETTransaction(id, "NA")
}

// HandleGETAllTransaction for GET all data this daily from Transaction
func (p TransactionRepo) HandleGETAllTransactionDaily() (*Laporan, error) {
	var d Transaction
	var AllTransaction []Transaction
	var todayReport Laporan
	var totalPendapatan int

	result, err := p.db.Query("SELECT * FROM daftar_transaksi WHERE status=? AND DATE(created_at) = CURDATE()", "A")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.PeralatanMakan, &d.Total, &d.Status, &d.Created, &d.Updated)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		totalPendapatan += d.Total
		resultProduct, _ := p.db.Query(`SELECT nama,kategori,harga,extra_pedas,kuantiti,total FROM detail_transaksi_idx WHERE trans_id=?`, d.ID)

		var soldItem SoldItems
		var allSoldItem []SoldItems
		for resultProduct.Next() {
			err := resultProduct.Scan(&soldItem.ProductID, &soldItem.Category, &soldItem.Harga, &soldItem.EkstraPedas, &soldItem.Quantity, &soldItem.Total)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			allSoldItem = append(allSoldItem, soldItem)
		}
		d.SoldItems = allSoldItem
		AllTransaction = append(AllTransaction, d)
	}
	todayReport.TotalPendapatan = totalPendapatan
	todayReport.DetailTransaksi = AllTransaction
	return &todayReport, nil
}
