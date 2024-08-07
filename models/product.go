package models

import (
	"encoding/json"
	"time"
)

type Product struct {
	User_id          string   `json:"user_id"`
	Tenant_name		 string	  `json:"tenant_name"`
	Name             string   `json:"name"`
	Units            string   `json:"units"`
	Quantity	     int   	  `json:"quantity"`
	Price			 int	  `json:"price"`
}

type ProductRealtime struct {
	ID                string    `json:"id"`
	CreateDtm         time.Time `json:"create_dtm"`
	SkuID             string    `json:"sku_id"`
	UserID            string    `json:"user_id"`
	OutletID          string    `json:"outlet_id"`
	Name              string    `json:"name"`
	Category          string    `json:"category"`
	Variant           string    `json:"variant"`
	Units             string    `json:"units"`
	Weight            int       `json:"weight"`
	Quantity          int       `json:"quantity"`
	MinimumQuantity   int       `json:"minimum_quantity"`
	Description       string    `json:"description"`
	BuyCost           int       `json:"buy_cost"`
	SellCost          int       `json:"sell_cost"`
	ModifiersID       string    `json:"modifiers_id"`
	Images            []string  `json:"images"`
	Rawmaterial       json.RawMessage     `json:"rawmaterial"`
	IsStockTracked    string    `json:"is_stock_tracked"`
	NumberSold        int       `json:"number_sold"`
	Outlets           []string  `json:"outlets"`
	BuyCostDiscounted int       `json:"buy_cost_discounted"`
	IsActive          string    `json:"is_active"`
	WholesalerCost    []struct {
		Min  int `json:"min"`
		Cost int `json:"cost"`
	} `json:"wholesaler_cost"`
}

type QueryProduct struct {
	User_id			string	  `json:"user_id"`
}

type QueryProductRealTime struct {
	User_id			string	  `json:"user_id"`
	Outlet_id		string	  `json:"outlet_id"`
	Category		string	  `json:"category"`
	Is_active		string	  `json:"is_active"`
}