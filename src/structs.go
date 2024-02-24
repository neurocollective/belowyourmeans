package main

type LoginPayload struct {
	Email    string
	Password string
}

type SignupPayload struct {
	LoginPayload
	FirstName string
	LastName  string
}