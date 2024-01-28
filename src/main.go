package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	ncsql "github.com/neurocollective/go_utils/sql"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"neurocollective.io/neurocollective/belowyourmeans/src/db"
	"strconv"
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

	corsMiddleware := func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Vary", "origin")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Cookie, Set-Cookie, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}

	router.Use(corsMiddleware)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Below Your Means",
		})
	})

	type LoginPayload struct {
		Email    string
		Password string
	}

	type SignupPayload struct {
		LoginPayload
		FirstName string
		LastName  string
	}

	router.POST("/login", func(c *gin.Context) {

		jsonBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		payload := new(LoginPayload)

		err = json.Unmarshal(jsonBytes, payload)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		query := db.CHECK_LOGIN_QUERY

		args := []any{payload.Email}

		users, err := ncsql.QueryForStructs[db.User](client, db.ScanForUserLoginData, query, args...)

		if err != nil {
			log.Fatal("error!", err.Error())
		}

		userCount := len(users)

		if userCount == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(users[0].HashedPassword), []byte(payload.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		returnJson := map[string]int{"userId": users[0].Id}

		c.JSON(http.StatusOK, gin.H{"data": returnJson})
	})

	router.POST("/signup", func(c *gin.Context) {

		jsonBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		payload := new(SignupPayload)

		err = json.Unmarshal(jsonBytes, payload)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		query := db.CREATE_USER_QUERY

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		args := []any{payload.FirstName, payload.LastName, payload.Email, string(hashedPassword)}

		users, err := ncsql.QueryForStructs[db.User](client, db.ScanForUserLoginData, query, args...)

		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		userCount := len(users)

		if userCount == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists!"})
			return
		}

		returnJson := map[string]int{"userId": users[0].Id}
		c.JSON(http.StatusOK, gin.H{"data": returnJson})
	})

	router.GET("/user", func(c *gin.Context) {
		query := db.USER_QUERY

		// userInURLQuery := c.Query("id")
		args := []any{1} // TODO: parse `userInURLQuery` from string to int, put into args instead of hardcoded 1

		users, parseError := ncsql.QueryForStructs[db.User](client, db.ScanForUser, query, args...)

		if parseError != nil {
			log.Fatal("error!", parseError.Error())
		}
		c.JSON(http.StatusOK, gin.H{"data": users})
	})

	router.GET("/expenditure", func(c *gin.Context) {

		queryMap := c.Request.URL.Query()

		queryWhereClauses := ""
		argIndex := 2

		args := []any{1} // user id

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
					c.JSON(http.StatusBadRequest, gin.H{"error": "amount is not a valid float"})
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
		c.JSON(http.StatusOK, gin.H{"data": expenditures})
	})

	router.POST("/expenditure", func(c *gin.Context) {

		user := c.PostForm("user")
		category := c.PostForm("category")
		amount := c.PostForm("amount")
		description := c.PostForm("description")
		date := c.PostForm("date")

		fullQuery := db.CREATE_EXPENDITURE_QUERY_STEM

		args := []any{user, category, amount, description, date}

		expenditures, parseError := ncsql.QueryForStructs[db.Expenditure](client, db.ScanForExpenditure, fullQuery, args...)

		if parseError != nil {
			log.Fatal("error!", parseError.Error())
		}
		c.JSON(http.StatusOK, gin.H{"data": expenditures})
	})

	// router.GET("/bruh", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{ "message": "hello" })
	// })

	// router.GET("/rawhtml", func(c *gin.Context) {
	// 	c.Data(http.StatusOK, "text/html", []byte("<html><head><title>FROM GOLANG</title></head><body>yah brah</body></html>"))
	// })

	router.Run()
}
