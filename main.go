package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup the webserver with routes to "health" and "query"
	router := gin.Default()
	router.GET("/health", health)
	router.GET("/query", query)

	// Start the server and listen on port 8080
	router.Run("localhost:8080")

}

func health(c *gin.Context) {
	// Respond to health check
	c.IndentedJSON(http.StatusOK, "OK")
}

func query(c *gin.Context) {
	// Create the response for the query endpoint

	// Define date variable and get the DAYS env variable in int form
	date := time.Now()
	days, _ := strconv.Atoi(os.Getenv("DAYS"))

	// Build URL query string
	query := fmt.Sprintf("apikey=%s&function=%s&symbol=%s", os.Getenv("KEY"), os.Getenv("FUNCTION"), os.Getenv("SYMBOL"))

	// Create URL to query alphavantage API endpoint
	u := &url.URL{
		Scheme:   "https",
		Host:     "alphavantage.co",
		Path:     "query",
		RawQuery: query,
	}

	// Build http client and send request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", u.String(), nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Save body in map object for parsing
	var marketData map[string]interface{}

	json.Unmarshal(body, &marketData)

	// Parse daily time series data
	dayMapInt := marketData["Time Series (Daily)"]
	dayMap := dayMapInt.(map[string]interface{})
	var closeAvg float64
	var count float64
	var finalJSON []byte

	// Iterate over data and get the last DAYS results
	for i := 1; i <= days; i++ {
		pastDay := date.AddDate(0, 0, -i)
		dateOnly := pastDay.Format(time.DateOnly)
		if dayMap[dateOnly] == nil {
			continue
		}

		// Get the days info. Grab the cloing price and convert to float
		count++
		closeQuery := dayMap[dateOnly].(map[string]interface{})["4. close"].(string)
		closePrice, _ := strconv.ParseFloat(closeQuery, 64)

		closeAvg = closeAvg + closePrice

		// Add date to JSON
		wrappedJSON := map[string]interface{}{
			dateOnly: dayMap[dateOnly],
		}
		jsonData, err := json.Marshal(wrappedJSON)
		if err != nil {
			fmt.Println(err)
		}

		// Append the new day's info
		finalJSON = append(finalJSON, jsonData...)

	}

	// Get closing average
	closeAvg = closeAvg / count
	finalAvg := fmt.Sprintf("%.2f", closeAvg)

	// Add closing average to JSON
	avgJSON := map[string]string{
		"average closing": finalAvg,
	}
	avgclose, err := json.Marshal(avgJSON)
	if err != nil {
		fmt.Println(err)
	}
	finalJSON = append(finalJSON, avgclose...)

	// Return JSON
	c.Data(http.StatusOK, "application/json", finalJSON)

}
