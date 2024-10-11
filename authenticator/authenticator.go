package authenticator

import (
	"crypto/sha256"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type Authenticator struct {
	path string
}

func New() *Authenticator {
	var auth Authenticator
	return &auth
}

func (*Authenticator) IsLoginCorrect(username string, password string) (bool, error) {

	file, err := os.Open("../pwd_test_vault.csv")
	if err != nil {
		fmt.Println("File reading error ", err)
		return false, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("CSV reading error ", err)
		return false, err
	}

	for _, record := range records {
		if record[0] == username {
			h := sha256.New()
			h.Write([]byte(password))

			pwdHash := fmt.Sprintf("%x", h.Sum(nil))

			if record[1] == pwdHash {
				return true, nil
			} else {
				return false, errors.New("password hash and csv password hash" +
					" do not match: " + pwdHash + " and " +
					record[1])
			}
		}
	}

	return false, errors.New("username not found")
}
