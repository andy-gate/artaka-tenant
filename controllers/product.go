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

func ProductList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QueryProduct
	c.BindJSON(&query)

	var products []models.Product
  
	if err := models.MPosGORM.Raw("SELECT a.user_id, b.owner_name as tenant_name, a.name, a.units, a.quantity, a.sell_cost as price from products a join subscribers b on a.user_id = b.user_id where a.user_id = ?", query.User_id).Scan(&products).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (products != nil) {
	  c.JSON(http.StatusOK, products)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}

func ProductListRealTime(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QueryProductRealTime
	c.BindJSON(&query)

	var products []models.ProductRealtime
  
	b, err := json.Marshal(query)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

	var jsonStr = []byte(b)

	apiUrl := "https://artaka-api.com/api/products/show"

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
	_ = json.Unmarshal(body, &products)
  
	if (products != nil) {
	  c.JSON(http.StatusOK, products)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}