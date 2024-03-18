package db

import "golang.org/x/crypto/bcrypt"

// This is a stub function, just to get the basic project
// working. This should be replaced with actual code to
// retrieve the user's hashed password from the datastore
func RetrieveUserHash(username string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte("a"), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
