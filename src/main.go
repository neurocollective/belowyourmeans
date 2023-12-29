package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
		args := []any{ 1 }

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
			queryWhereClauses += connector + columnName + " = " + "$" + strconv.Itoa(argIndex)	

			args = append(args, value[0])
			argIndex++
		}

		queryStem := db.EXPENDITURE_QUERY_STEM
		fullQuery := queryStem + queryWhereClauses

		log.Println("fullQuery:", fullQuery)
		log.Println("args:", args)

		expenditures, parseError := ncsql.QueryForStructs[db.Expenditure](client, db.ScanForExpenditure, fullQuery, args...)

		if parseError != nil {
			log.Fatal("error!", parseError.Error())
		}
		c.JSON(http.StatusOK, gin.H{ "data": expenditures })
	})

	router.GET("/bruh", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "message": "hello" })
	})

	router.GET("/rawhtml", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte("<html><head><title>FROM GOLANG</title></head><body>yah brah</body></html>"))
	})

	router.Run()
}