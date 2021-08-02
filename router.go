package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/recommend", recommendSubjects)
	router.Run(":2000")
}

// adapting the user preferences to a UserPreferences struct
// calling based on this structure a recommendation function
// and sending the json back as response
func recommendSubjects(c *gin.Context) {

	var uP *UserPreferences
	if err := c.BindJSON(&uP); err != nil {
		return
	}

	reco := Recommender{}
	recommendations := reco.provideRecommendation(*uP)
	fmt.Println(recommendations)

	c.IndentedJSON(http.StatusOK, recommendations)
	return
}
