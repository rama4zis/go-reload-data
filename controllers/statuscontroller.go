package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rama4zis/go-reload-data/models"
)

func WriteDataJSON(data interface{}) {
	// marshal data
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Write data to file
	err = ioutil.WriteFile("status.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data saved to file")
}

func Index(c *gin.Context) {
	var status models.Status

	// Read data from file
	jsonData, err := ioutil.ReadFile("status.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &status)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}
