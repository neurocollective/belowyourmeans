package structs

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CategoryItemPayload struct {
	CategoryId int `json:"categoryId"`
	DisplayName string `json:"displayName"`
}

type BudgetCategory struct {
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}
