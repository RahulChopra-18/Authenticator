package database

import (
	"context"
	user2 "example.com/m/pkg/models/user"
)

// LoadUser loads a dummy user.
func LoadUser(u *user2.Check) (*user2.Check, error) {

	var user user2.Check
	client := DB
	usersCollection := client.Database("Login").Collection("credentials")

	searchAttributes := make(map[string]interface{})
	searchAttributes["email"] = u.Email
	dbResponse := usersCollection.FindOne(context.TODO(), searchAttributes)
	if dbResponse.Err() != nil {
		return nil, dbResponse.Err()
	}

	err := dbResponse.Decode(&user)
	if err != nil {

		return nil, err
	}
	return &user, nil
}
