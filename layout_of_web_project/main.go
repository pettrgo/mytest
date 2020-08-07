package main

import (
	"net/http"
)

//thrift 请求结构体
type FeatureSetParams struct {
	DriverID  int64  `thrift:"driverID,1,required" json:"driver_id"`
	OrderID   int64  `thrift:"OrderID,2,required" json:"order_id"`
	UserID    int64  `thrift:"UserID,3,required" json:"user_id"`
	ProductID int    `thrift:"ProductID,4,required" json:"prod_id"`
	Addr      string `thrift:"Addr,5,required" json:"addr"`
}

//controller input struct
type CreateOrderParams struct {
	OrderID   int64
	UserID    int64
	ProductID int
	Addr      string
}

func HTTPCreateOrderHandler(wr http.ResponseWriter, r *http.Request) {

}

func main() {

}
