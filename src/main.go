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
	"bytes"
	"errors"
	"neurocollective.io/neurocollective/belowyourmeans/src/password"
	"neurocollective.io/neurocollective/belowyourmeans/src/structs"
	bymsql "neurocollective.io/neurocollective/belowyourmeans/src/structs/sql"
	"strconv"
	"strings"
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

type RawEx struct {
	Id           int64   `json:"id"`
	UserId       int64   `json:"user_id"`
	CategoryId   int64   `json:"category_id"`
	Value        float32 `json:"value"`
	Description  string  `json:"description"`
	DateOccurred string  `json:"date_occurred"`
	CreateDate   string  `json:"create_date"`
	ModifiedDate string  `json:"modified_date"`
}

func ExecuteNodeQuery[T any](query string, queryParams []any) ([]T, error) {

	client := new(http.Client)

	body := map[string]any{
		"query":       query,
		"parameters": queryParams,
		"specialRule": "mapExpenditures",
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bodyBytes)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:3001/query", bodyReader)

	request.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	results := []T{}

	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &results)

	if err != nil {
		return nil, err
	}

	return results, nil
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

	connectionString := "user=postgres password=postgres dbname=postgres sslmode=disable"

	client, getClientError := ncsql.BuildPostgresClient(connectionString)

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

	router.GET("/cache", func(c *gin.Context) {
		c.JSON(http.StatusOK, FAKE_REDIS)
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

	// old version
	router.GET("/ex", fakeAuthMiddleware, func(c *gin.Context) {

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

		query := "select * from expenditure limit 1;"
		args := make([]any, 0)

		expenditures, err := ncsql.Select[ncsql.Expenditure](client, query, args)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reach db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": expenditures})
		return
	})

	router.GET("/expenditure", func(c *gin.Context) {

		userIdString := c.Query("userId")

		if userIdString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "userId required"})
			return
		}

		userId, err := strconv.Atoi(userIdString)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad userId"})
			return
		}

		monthString := c.Query("month")

		if monthString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "monthString required"})
			return
		}

		month, err := strconv.Atoi(monthString)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad month"})
			return
		}

		args := []any{ userId, month }
		query := "select * from expenditure where user_id = $1 and EXTRACT(MONTH FROM date_occurred) = $2;"
		expenditures, err := ExecuteNodeQuery[RawEx](query, args)

		if err != nil {
			log.Println(err)
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
