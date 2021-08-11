package server

import "os"

func Init() {
	r := SetupRouter()

	port := getenv("PORT", "80")
	r.Run(":" + port)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}