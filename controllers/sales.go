package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andy-gate/artaka-tenant/models"
	"github.com/gin-gonic/gin"
)

func SalesList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QuerySales
	c.BindJSON(&query)

	var sales []models.Sales
  
	if err := models.MPosGORM.Raw("select b.owner_name as tenant_name, a.create_dtm::date, count(a.id) as total_trx, sum(a.total_bill) as total_amount from sales a join subscribers b on a.user_id = b.user_id where a.user_id = ? AND a.create_dtm > ? AND a.create_dtm < ? group by b.owner_name, a.create_dtm::date", query.User_id, query.Start_date, query.End_date).Scan(&sales).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (sales != nil) {
	  c.JSON(http.StatusOK, sales)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}

func SalesListDetail(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QuerySales
	c.BindJSON(&query)

	var sales []models.SalesDetail
  
	if err := models.MPosGORM.Raw("SELECT * from sales where user_id = ? AND create_dtm > ? AND create_dtm < ?", query.User_id, query.Start_date, query.End_date).Scan(&sales).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (sales != nil) {
	  c.JSON(http.StatusOK, sales)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}

func SalesListRealTime(c *gin.Context){
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QuerySalesRealTime
	c.BindJSON(&query)

	b, err := json.Marshal(query)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

	var jsonStr = []byte(b)

	apiUrl := "https://artaka-api.com/api/sales/show"

	req, err := http.NewRequest("POST",apiUrl,bytes.NewBuffer(jsonStr))
	if err != nil {   
		fmt.Printf("Request Failed: %s", err)
		return
	}

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

	defer resp.Body.Close()

	var sales []models.Sales

	var salesData []models.SalesDetail
  	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &salesData)
	collections := make(map[string][]models.SalesDetail)
    for _, b := range salesData {
        collections[b.Create_dtm.Format("2006-01-02")] = append(collections[b.Create_dtm.Format("2006-01-02")], b)
    }
	i := 0
	for k := range collections {
		total_trx := len(collections[k])
		total_amount := 0
		for _, n := range collections[k]{
			total_amount += n.Total_bill
		}
        sales = append(sales, models.Sales{Create_dtm: k, Total_trx: total_trx, Total_amount: total_amount})
		i++
    }
	if (sales != nil) {
		c.JSON(http.StatusOK, sales)
	  } else {
		c.JSON(http.StatusOK, json.RawMessage(`[]`))
	  }
}

func SalesListDetailRealTime(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QuerySalesRealTime
	c.BindJSON(&query)

	var sales []models.SalesDetail
  
	b, err := json.Marshal(query)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

	var jsonStr = []byte(b)

	apiUrl := "https://artaka-api.com/api/sales/show"

	req, err := http.NewRequest("POST",apiUrl,bytes.NewBuffer(jsonStr))
	if err != nil {   
		fmt.Printf("Request Failed: %s", err)
		return
	}

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

	defer resp.Body.Close()

  	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &sales)
  
	if (sales != nil) {
	  c.JSON(http.StatusOK, sales)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}