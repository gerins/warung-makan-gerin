package transaction

// Transaction Struct
type Transaction struct {
	ID             string      `json:"id"`
	SoldItems      []SoldItems `json:"solditems"`
	PeralatanMakan int         `json:"peralatanmakan"`
	Total          int         `json:"total"`
	Status         string      `json:"status"`
	Created        string      `json:"created"`
	Updated        string      `json:"updated"`
}

type SoldItems struct {
	ProductID   string `json:"productid" validate:"nonzero"`
	Category    string `json:"category"`
	Harga       int    `json:"harga"`
	EkstraPedas int    `json:"ekstrapedas"`
	Quantity    int    `json:"quantity" validate:"nonzero"`
	Total       int    `json:"total"`
}

type Laporan struct {
	TotalPendapatan int           `json:"totalpendapatan"`
	DetailTransaksi []Transaction `json:"detailtransaksi"`
}
