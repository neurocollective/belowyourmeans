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
	UserId int `json:"userId"`
	Description string `json:"description"`
}

type BudgetCategory struct {
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}
