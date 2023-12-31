package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"net/http"
	"strconv"
	ncsql "github.com/neurocollective/go_utils/sql"
	"neurocollective.io/neurocollective/belowyourmeans/src/db"
)

func main() {

	log.Println("booting server...")
	
	router := gin.Default()

	router.LoadHTMLGlob("src/templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	client, getClientError := ncsql.BuildPostgresClient("user=postgres password=postgres dbname=postgres sslmode=disable")		

	// connect to the db to test if connection is valid

	if getClientError != nil {
		log.Fatal("error getting client")
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Below Your Means",
		})
	})

	router.GET("/user", func(c *gin.Context) {
		query := db.USER_QUERY

		// userInURLQuery := c.Query("id")
		args := []any{ 1 } // TODO: parse `userInURLQuery` from string to int, put into args instead of hardcoded 1

		users, parseError := ncsql.QueryForStructs[db.User](client, db.ScanForUser, query, args...)

		if parseError != nil {
			log.Fatal("error!", parseError.Error())
		}
		c.JSON(http.StatusOK, gin.H{ "data": users })
	})

	router.GET("/expenditure", func(c *gin.Context) {

		queryMap := c.Request.URL.Query()

		queryWhereClauses := ""
		argIndex := 2

		args := []any{ 1 } // user id

		for key, value := range queryMap {

			columnName := db.GetExpenditureColumnNameByQueryKey(key)

			if columnName == "" {
				continue
			}

			connector := " and "

			// ignore multiple query keys, only take first. Defies the spec but that's wacky, brah.

			if columnName == "category_id" && value[0] == "-1" {
				queryWhereClauses += connector + columnName + " is " + "null"		
			} else if columnName == "value" {
				impreciseFloat, parseError := strconv.ParseFloat(value[0], 32)
				if parseError != nil {
					log.Println(parseError.Error())
					c.JSON(http.StatusBadRequest, gin.H{ "error": "amount is not a valid float" })
					return
				}
				preciseFloat := fmt.Sprintf("%.2f", impreciseFloat) // TODO - should this be a float32 ?

				queryWhereClauses += connector + columnName + " = " + "$" + strconv.Itoa(argIndex)
				args = append(args, preciseFloat)
			} else {
				queryWhereClauses += connector + columnName + " = " + "$" + strconv.Itoa(argIndex)
				args = append(args, value[0])
			}

			argIndex++
		}

		queryStem := db.EXPENDITURE_QUERY_STEM
		fullQuery := queryStem + queryWhereClauses + ";"

		expenditures, parseError := ncsql.QueryForStructs[db.Expenditure](client, db.ScanForExpenditure, fullQuery, args...)

		if parseError != nil {
			log.Fatal("error!", parseError.Error())
		}
		c.JSON(http.StatusOK, gin.H{ "data": expenditures })
	})

	router.POST("/expenditure", func(c *gin.Context) {

		user := c.PostForm("user")
		category := c.PostForm("category")
		amount := c.PostForm("amount")
		description := c.PostForm("description")
		date := c.PostForm("date")

		fullQuery := db.CREATE_EXPENDITURE_QUERY_STEM

		args := []any{ user, category, amount, description, date }

		expenditures, parseError := ncsql.QueryForStructs[db.Expenditure](client, db.ScanForExpenditure, fullQuery, args...)
		
		if parseError != nil {
			log.Fatal("error!", parseError.Error())
		}
		c.JSON(http.StatusOK, gin.H{ "data": expenditures })
	})

	// router.GET("/bruh", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{ "message": "hello" })
	// })

	// router.GET("/rawhtml", func(c *gin.Context) {
	// 	c.Data(http.StatusOK, "text/html", []byte("<html><head><title>FROM GOLANG</title></head><body>yah brah</body></html>"))
	// })

	router.Run()
}