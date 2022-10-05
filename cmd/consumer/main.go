package main

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var validate *validator.Validate

type Payment struct {
	Transaction  string `validate:"required" json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `validate:"required" json:"currency"`
	Provider     string `validate:"required" json:"provider"`
	Amount       int    `validate:"required" json:"amount"`
	PaymentDt    int    `validate:"required" json:"payment_dt"`
	Bank         string `validate:"required" json:"bank"`
	DeliveryCost int    `validate:"required" json:"delivery_cost"`
	GoodsTotal   int    `validate:"required" json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Delivery struct {
	Name    string `validate:"required" json:"name"`
	Phone   string `validate:"required" json:"phone"`
	Zip     string `validate:"required" json:"zip"`
	City    string `validate:"required" json:"city"`
	Address string `validate:"required" json:"address"`
	Region  string `validate:"required" json:"region"`
	Email   string `validate:"required" json:"email"`
}

type Item struct {
	ChrtId      int    `validate:"required" json:"chrt_id"`
	TrackNumber string `validate:"required" json:"track_number"`
	Price       int    `validate:"required" json:"price"`
	Rid         string `validate:"required" json:"rid"`
	Name        string `validate:"required" json:"name"`
	Sale        int    `validate:"required" json:"sale"`
	Size        string `validate:"required" json:"size"`
	TotalPrice  int    `validate:"required" json:"total_price"`
	NmId        int    `validate:"required" json:"nm_id"`
	Brand       string `validate:"required" json:"brand"`
	Status      int    `validate:"required" json:"status"`
}

type Order struct {
	OrderUid          string    `validate:"required" json:"order_uid"`
	TrackNumber       string    `validate:"required" json:"track_number"`
	Entry             string    `validate:"required" json:"entry"`
	Delivery          Delivery  `validate:"required" json:"delivery"`
	Payment           Payment   `validate:"required" json:"payment"`
	Items             []Item    `validate:"required" json:"items"`
	Locale            string    `validate:"required" json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `validate:"required" json:"customer_id"`
	DeliveryService   string    `validate:"required" json:"delivery_service"`
	Shardkey          string    `validate:"required" json:"shardkey"`
	SmId              int       `validate:"required" json:"sm_id"`
	DateCreated       time.Time `validate:"required" json:"date_created"`
	OofShard          string    `validate:"required" json:"oof_shard"`
}

func main() {
	validate = validator.New()

	// коннектимся к базе
	connectToDb()
	defer db.Close()

	// заполняем кэш значениями из базы
	fillUpCache()

	// коннектимся к натс стриминг
	subscribeToNats()

	// запускаем сервак
	runServer()
}
