package models

import (
	"encoding/json"
	"time"
)

type Sales struct {
	Tenant_name      string          `json:"tenant_name"`
	Create_dtm       string       	 `json:"create_dtm"`
	Total_trx		 int			 `json:"total_trx"`
	Total_amount	 int			 `json:"total_amount"`
}

type SalesDetail struct {
	ID               string          `json:"id"`
	Create_dtm       time.Time       `json:"create_dtm"`
	Sales_id         string          `json:"sales_id"`
	User_id          string          `json:"user_id"`
	Outlet_id        string          `json:"outlet_id"`
	Sales_type       string          `json:"sales_type"`
	Customer_id      string          `json:"customer_id"`
	Products         json.RawMessage `json:"products"`
	Subtotal         int             `json:"subtotal"`
	Total_diskon     int             `json:"total_diskon"`
	Total_tax        json.RawMessage `json:"total_tax"`
	Total_bill       int             `json:"total_bill"`
	Payment_method   string          `json:"payment_method"`
	Payment_due_date string          `json:"payment_due_date"`
	Total_payment    int             `json:"total_payment"`
	Exchange         int             `json:"exchange"`
	Notes            string          `json:"notes"`
	Total_buy_cost   int             `json:"total_buy_cost"`
	Payment_date     string          `json:"payment_date"`
	Reward_id        string          `json:"Reward_id"`
	Points_redeem    int             `json:"points_redeem"`
}

type QuerySales struct {
	User_id          string          `json:"user_id"`
	Outlet_id        string          `json:"outlet_id"`
	Start_date		 string			 `json:"start_date"`
	End_date		 string			 `json:"end_date"`
}

type QuerySalesRealTime struct {
	User_id          string          `json:"user_id"`
	Outlet_id        string          `json:"outlet_id"`
	Sales_type		 string			 `json:"sales_type"`
	Payment_due_date string			 `json:"payment_due_date"`
	Payment_method	 string			 `json:"payment_method"`
	Notes			 string			 `json:"notes"`
}

type SalesRT struct {
	Create_dtm       string       	 `json:"create_dtm"`
	Total_trx		 int			 `json:"total_trx"`
	Total_amount	 int			 `json:"total_amount"`
}