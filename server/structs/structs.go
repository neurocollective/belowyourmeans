package structs

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupPayload struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CategoryItemPayload struct {
	CategoryId  int    `json:"categoryId"`
	Description string `json:"description"`
}

type BudgetCategory struct {
	Id          int    `json:"id"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Ignored     bool   `json:"ignored"`
}

type ExpenditureWithCategoryName struct {
	Id           int64   `json:"id"`
	Value        float32 `json:"value"`
	Description  string  `json:"description"`
	DateOccurred string  `json:"date_occurred"`
	CategoryName string  `json:"category_name"`
}

type CategorizeExpenditurePayload struct {
	ExpenditureId int `json:"expenditureId"`
	CategoryId    int `json:"categoryId"`
}

type ApplyBudgetCategoryItemPayload struct {
	Description   string `json:"expenditureDescription"`
	CategoryName  string `json:"categoryName"`
	CategoryId    int    `json:"categoryId"`
	ExpenditureId int    `json:"expenditureId"`
}
