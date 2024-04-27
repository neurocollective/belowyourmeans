package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	ncsql "github.com/neurocollective/go_utils/sql"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"neurocollective.io/neurocollective/belowyourmeans/src/constants"
	"neurocollective.io/neurocollective/belowyourmeans/src/cookie"
	"neurocollective.io/neurocollective/belowyourmeans/src/db"
	// "neurocollective.io/neurocollective/belowyourmeans/src/parsing"
	"neurocollective.io/neurocollective/belowyourmeans/src/password"
	"neurocollective.io/neurocollective/belowyourmeans/src/structs"
	bymsql "neurocollective.io/neurocollective/belowyourmeans/src/structs/sql"
	"strconv"
	"strings"
	"errors"
)

func GetFromContext[T any](c *gin.Context, key string) (T, error) {
	var empty T

	value, exists := c.Get(key)

	if !exists {
		return empty, errors.New(key + " not in context")
	}

	assertedValue, ok := value.(T)

	if !ok {
		return empty, errors.New(key + " cannot be asserted to requested type")
	}

	return assertedValue, nil
}

func main() {

	FAKE_REDIS := make(map[string]string)

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

		keyLastIndex := index + len(constants.COOKIE_KEY) + 1

		afterKey := strings.Join(characterSlice[keyLastIndex:], "")

		userId, present := FAKE_REDIS[afterKey]

		if !present {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(constants.USER_ID, userId)
		c.Next()
	}

	fakeAuthMiddleware := func(c *gin.Context) {
		c.Set(constants.USER_ID, "1")
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

		payload := new(structs.LoginPayload)

		err = json.Unmarshal(jsonBytes, payload)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		query := db.CHECK_LOGIN_QUERY

		args := []any{payload.Email}

		users, err := ncsql.Select[bymsql.User](client, query, args)

		if err != nil {
			log.Fatal("error!", err.Error())
		}

		userCount := len(users)

		if userCount == 0 {
			log.Println("no users found for", payload.Email)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(*users[0].HashedPassword), []byte(payload.Password))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		userId := users[0].Id

		returnJson := map[string]int{"userId": *userId}

		cookieValue, err := cookie.GenerateCookie()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.Header("Set-Cookie", cookie.GetSetCookieHeaderValue(cookieValue))

		FAKE_REDIS[cookieValue] = strconv.Itoa(*userId)

		c.JSON(http.StatusOK, gin.H{"data": returnJson})
	})

	router.POST("/signup", func(c *gin.Context) {

		payload, err := password.GetSignupPayload(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		query := db.CREATE_USER_QUERY

		hashedPassword, err := password.HashPassword(payload.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		// TODO - query for existing email first

		args := []any{payload.FirstName, payload.LastName, payload.Email, string(hashedPassword)}

		_, err = ncsql.Select[bymsql.User](client, query, args)

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

		args := []any{id}

		users, err := ncsql.Select[bymsql.User](client, query, args)

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

		payload, err := password.GetSignupPayload(c)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		hashedPassword, err := password.HashPassword(payload.Password)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"password": hashedPassword})
	})

	router.GET("/expenditure", fakeAuthMiddleware, func(c *gin.Context) {

		userIdString, err := GetFromContext[string](c, constants.USER_ID)

		log.Println("userId:", userIdString)

		if err != nil {
			log.Println("userId not received from auth middleware")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		userId, err := strconv.Atoi(userIdString)
		if err != nil {
			log.Println("userId un-convertable from string to int")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		if userId == 0 {
			userId = 1
		}

		query := "select * from expenditure;"
		args := make([]any, 0)

		expenditures, err := ncsql.Select[ncsql.Expenditure](client, query, args)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reach db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": expenditures})
		return
	})

	router.POST("/expenditure", authMiddleware, func(c *gin.Context) {

		user := c.PostForm("user")
		category := c.PostForm("category")
		amount := c.PostForm("amount")
		description := c.PostForm("description")
		date := c.PostForm("date")

		fullQuery := db.CREATE_EXPENDITURE_QUERY_STEM

		args := []any{user, category, amount, description, date}

		expenditures, parseError := ncsql.Select[ncsql.Expenditure](client, fullQuery, args)

		if parseError != nil {
			log.Fatal("error!", parseError.Error())
		}
		c.JSON(http.StatusOK, gin.H{"data": expenditures})
	})

	router.Run()
}
