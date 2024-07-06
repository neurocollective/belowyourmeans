package sql

import (
	"errors"
	ncsql "github.com/neurocollective/go_utils/sql"
)

type User struct {
	Id *int
	FirstName *string
	LastName *string
	Email *string
	HashedPassword *string
	CreateDate *string
	LastModifiedDate *string
}

func (u User) GetId() *int {
	return u.Id
}

func (u User) Zero() ncsql.SQLMetaStruct {

	new := User{}

	one := 0
	two := ""
	three := ""
	four := ""
	five := ""
	six := ""
	seven := ""

	new.Id = &one
	new.FirstName = &two
	new.LastName = &three
	new.Email = &four
	new.HashedPassword = &five
	new.CreateDate = &six
	new.LastModifiedDate = &seven

	return new
}

// this should be generated code, based on column names
func (u User) Keys() []string {
	return []string{
		"first_name", 
		"last_name",
		"email text",
		"hashed_password",
		"create_date",
		"last_modified_date",
	}
}

func (u User) KeysAll() []string {
	return []string{
		"id",
		"first_name", 
		"last_name",
		"email text",
		"hashed_password",
		"create_date",
		"last_modified_date",
	}
}

func (u User) TableName() string {
	return "budget_user"
}

func (u User) Values() []any {

	return []any{
		u.FirstName,
		u.LastName,
		u.Email,
		u.HashedPassword,
		u.CreateDate,
		u.LastModifiedDate,
	}
}

func (u User) ValuesAll() []any {

	return []any{
		u.Id,
		u.FirstName,
		u.LastName,
		u.Email,
		u.HashedPassword,
		u.CreateDate,
		u.LastModifiedDate,
	}
}

func (u User) Get(structKey string) (any, error) {
	values := u.Values()
	for index, key := range u.Keys() {
		value := values[index]
		if key == structKey {
			return value, nil
		}
	}
	return nil, errors.New(structKey + "not found")
}
