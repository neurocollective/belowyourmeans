package main

import (
    "testing"
)

func TestGenerateCookie(t *testing.T) {
    cookieOne, _ := GenerateCookie()
    cookieTwo, _ := GenerateCookie()
    if cookieOne == cookieTwo {
        t.Fatalf("two cookies from `GenerateCookie` match! %v == %v", cookieOne, cookieTwo)
    }
}

