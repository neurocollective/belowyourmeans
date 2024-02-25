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
	"neurocollective.io/neurocollective/belowyourmeans/src/constants"
	"strconv"
	"strings"
)

func main() {

	FAKE_REDIS := make(map[string]int)

	authMiddleware := func(c *gin.Context) {
		headers := c.Request.Header

		log.Println("current redis:", FAKE_REDIS)

		cookieValues, present := headers["Cookie"]

		if !present {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		firstCookie := cookieValues[0]

		log.Println("firstCookie", firstCookie)

		index := strings.Index(firstCookie, constants.COOKIE_KEY)

		if index == -1 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		characterSlice := strings.Split(firstCookie, "")

		log.Println("characterSlice", characterSlice)

		keyLastIndex := index + len(constants.COOKIE_KEY) + 1

		afterKey := strings.Join(characterSlice[keyLastIndex:], "")

		log.Println("afterKey", afterKey)

		userId, present := FAKE_REDIS[afterKey]

		if !present {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(constants.USER_ID, userId)
		c.Next()
	}

	log.Println("booting server...")

	router := gin.Default()

	router.LoadHTMLGlob("src/templates/*")

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

		userId := users[0].Id

		returnJson := map[string]int{"userId": userId}

		cookie, err := GenerateCookie()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.Header("Set-Cookie", GetSetCookieHeaderValue(cookie))

		FAKE_REDIS[cookie] = userId

		c.JSON(http.StatusOK, gin.H{"data": returnJson})
	})

	router.POST("/signup", func(c *gin.Context) {

		payload, err := GetSignupPayload(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		query := db.CREATE_USER_QUERY

		hashedPassword, err := HashPassword(payload.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		// TODO - query for existing email first

		args := []any{payload.FirstName, payload.LastName, payload.Email, string(hashedPassword)}

		_, err = ncsql.QueryForStructs[db.User](client, db.ScanForUserLoginData, query, args...)

		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"staus": "success"})
	})

	router.GET("/auth", authMiddleware, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "loggedIn"})
	})

	router.GET("/user", authMiddleware, func(c *gin.Context) {
		query := db.USER_QUERY

		userInURLQuery := c.Query("id")

		id, err := strconv.Atoi(userInURLQuery)

		if err != nil {
			log.Fatal("error!", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id sent"})
			return
		}

		args := []any{ id }

		users, err := ncsql.QueryForStructs[db.User](client, db.ScanForUser, query, args...)

		if err != nil {
			log.Fatal("error!", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		if len(users) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return			
		}

		c.JSON(http.StatusOK, gin.H{"data": users[0]})
	})

	router.POST("/password/hash", func(c *gin.Context) {

		payload, err := GetSignupPayload(c)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		hashedPassword, err := HashPassword(payload.Password)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"password": hashedPassword})
	})

	router.GET("/expenditure", authMiddleware, func(c *gin.Context) {

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

	router.POST("/expenditure", authMiddleware, func(c *gin.Context) {

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

	router.Run()
}
