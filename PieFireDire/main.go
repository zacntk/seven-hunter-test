package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/beef/summary", func(c *gin.Context) {
		url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to get data from %s", url)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read respond body")
		}

		bodyToString := string(body)

		bodyToString = strings.ReplaceAll(bodyToString, ",", "")
		bodyToString = strings.ReplaceAll(bodyToString, ".", "")

		beefs := strings.Fields(bodyToString)

		beefCount := make(map[string]int)

		for _, beef := range beefs {
			_, exists := beefCount[strings.ToLower(beef)]
			if exists {
				beefCount[strings.ToLower(beef)] += 1
			} else {
				beefCount[strings.ToLower(beef)] = 1
			}
		}

		c.JSON(200, gin.H{
			"beef": beefCount,
		})
	})

	router.Run()
}
