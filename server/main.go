package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	ncsql "github.com/neurocollective/go_utils/sql"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"neurocollective.io/neurocollective/belowyourmeans/server/constants"
	"neurocollective.io/neurocollective/belowyourmeans/server/cookie"
	"neurocollective.io/neurocollective/belowyourmeans/server/db"
	"neurocollective.io/neurocollective/belowyourmeans/server/db/queries"
	// "neurocollective.io/neurocollective/belowyourmeans/server/parsing"
	// "bytes"
	"errors"
	"neurocollective.io/neurocollective/belowyourmeans/server/password"
	"neurocollective.io/neurocollective/belowyourmeans/server/structs"
	bymsql "neurocollective.io/neurocollective/belowyourmeans/server/structs/sql"
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

	router.LoadHTMLGlob("server/templates/*")

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

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	router.GET("/auth", authMiddleware, func(c *gin.Context) {
		userId := c.GetString(constants.USER_ID)
		c.JSON(http.StatusOK, gin.H{"status": "loggedIn", "userId": userId})
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

	// TODO - use `authMiddleware`
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

		query := queries.SelectExpendituresWithCategoryNameByUserAndMonth()
		executeQuery := db.ExecuteNodeQuery[structs.ExpenditureWithCategoryName]
		args := []any{userId, month}
		expenditures, err := executeQuery(query, args, "mapExpenditures")

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reach db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": expenditures})
		return
	})

	router.PUT("/expenditure", authMiddleware, func(c *gin.Context) {

		jsonBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		payload := new(structs.CategorizeExpenditurePayload)

		err = json.Unmarshal(jsonBytes, payload)

		expenditureId := payload.ExpenditureId
		categoryId := payload.CategoryId

		log.Println("expenditureId", expenditureId)
		log.Println("categoryId", categoryId)

		args := []any{categoryId, expenditureId}
		query := "update expenditure set category_id = $1 where id = $2;"
		_, err = db.ExecuteNodeQuery[any](query, args, "")

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "did not reach db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	})

	router.GET("/categories", func(c *gin.Context) {

		userIdString := c.Query("userId")

		userId, err := strconv.Atoi(userIdString)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "bad userId"})
			return
		}

		args := []any{userId}
		query := "select id, display_name, description, ignored from budget_category where user_id = $1;"
		categories, err := db.ExecuteNodeQuery[structs.BudgetCategory](query, args, "")

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "did not reach db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": categories})
		return
	})

	router.GET("/categories/expenditures/names", authMiddleware, func(c *gin.Context) {

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

		args := []any{userId}
		query := queries.SelectExpendituresWithCategoryNameByUserUnique()
		execute := db.ExecuteNodeQuery[structs.ExpenditureWithCategoryName]
		categories, err := execute(query, args, "mapExpenditures")

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "did not reach db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": categories})
		return
	})

	// not just creaeting a new row - creating it and then running
	// an update in expenditure table
	router.POST("/category-item/apply", authMiddleware, func(c *gin.Context) {

		userIdString, err := GetFromContext[string](c, constants.USER_ID)

		log.Println("userId:", userIdString)

		if err != nil {
			log.Println("userId not received from auth middleware for POST /category-item")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		userId, err := strconv.Atoi(userIdString)
		if err != nil {
			log.Println("userId un-convertable from string to int for POST /category-item")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		bodyBytes, err := io.ReadAll(c.Request.Body)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read request body"})
			return
		}

		payload := new(structs.CategoryItemPayload)

		err = json.Unmarshal(bodyBytes, payload)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not parse payload"})
			return
		}

		args := []any{payload.CategoryId, payload.Description}
		query := "insert into budget_category_items (category_id, description) values ($1, $2);"
		_, err = db.ExecuteNodeQuery[any](query, args, "")

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "did not reach db"})
			return
		}

		args = []any{userId, payload.CategoryId, payload.Description}
		err = db.ApplyBudgetCategoryItems(args)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "did not reach db"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{})
		return
	})

	router.PUT("/category-item", func(c *gin.Context) {

		bodyBytes, err := io.ReadAll(c.Request.Body)

		payload := new(structs.CategoryItemPayload)

		err = json.Unmarshal(bodyBytes, payload)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "did not reach db"})
			return
		}

		args := []any{payload.CategoryId, payload.Description}
		query := "update budget_category_items set category_id = $1 where display_name = $2;"
		_, err = db.ExecuteNodeQuery[any](query, args, "")

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "did not reach db"})
			return
		}

		c.JSON(http.StatusNoContent, gin.H{})
		return
	})

	router.DELETE("/category-item", func(c *gin.Context) {

		bodyBytes, err := io.ReadAll(c.Request.Body)

		payload := new(structs.CategoryItemPayload)

		err = json.Unmarshal(bodyBytes, payload)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "did not reach db"})
			return
		}

		args := []any{payload.CategoryId, payload.Description}
		query := "delete from budget_category_items where category_id = $1 and display_name = $2;"
		_, err = db.ExecuteNodeQuery[any](query, args, "")

		c.JSON(http.StatusNoContent, gin.H{})
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
