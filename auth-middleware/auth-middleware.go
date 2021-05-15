package authmiddleware

import (
	"log"
)

func IsAuthenticated(token string) bool {
	response := false
	log.Println("checking user auth")

	log.Println("auth check completed")
	return response
}
