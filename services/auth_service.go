package services

import "Url-Counter-Service/config"

func IsCredentialsValid(username, password string) bool {
	for _, user := range config.Data.Users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
