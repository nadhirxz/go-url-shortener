package utils

import "os"

func GenerateFullURL(shortURL string) string {
	return os.Getenv("DOMAIN") + "/r/" + shortURL
}
