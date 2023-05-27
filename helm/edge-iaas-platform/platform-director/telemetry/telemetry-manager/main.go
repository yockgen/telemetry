package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
        "github.com/influxdata/influxdb-client-go/v2"
	//"github.com/influxdata/influxdb-client-go/v2/api/query"
        "context"
        "fmt"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Sales  int     `json:"sales"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, Sales: 0},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99, Sales: 0},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, Sales: 0},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// multiple params
type Sales struct {
	ID    string `uri:"id" binding:"required"`
	SALES int    `uri:"sales" binding:"required"`
}

func getMultiParams(c *gin.Context) {

	var sales Sales
	if err := c.ShouldBindUri(&sales); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "params incorrect"})
		return
	}

	var cnt = len(albums)
	for i := 0; i < cnt; i++ {
		if albums[i].ID == sales.ID {
			albums[i].Sales = sales.SALES
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

const (
	influxURL      = "http://192.168.1.107:32701"
	influxToken    = "X6zYQsXQdkC4K-WE7Uza_Z7yYWkENe3PAbNPIjryr4_KECA75QoLqALgsX9XQjWMFhdhZFz1TiLjxYUiM7B1zw=="
	influxOrg      = "intel"
	influxBucket   = "intel"
)

func getMetrices(c *gin.Context) {

        client := influxdb2.NewClient(influxURL, influxToken)
	defer client.Close()

        // Create a Flux query
	fluxQuery := fmt.Sprintf(`from(bucket: "%s")
		|> range(start: -1h)
		|> filter(fn: (r) => r._measurement == "tapo_p110_value")`, influxBucket)

        // Execute the query
	queryAPI := client.QueryAPI(influxOrg)
	result, err := queryAPI.Query(context.Background(), fluxQuery)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer result.Close()


        for result.Next() {
		if result.TableChanged() {
			// New table started, print the table name
			fmt.Printf("Table: %s\n", result.TableMetadata().String())
		}
		// Print the row values
		fmt.Printf("Row: %s\n", result.Record().String())
	}


	c.IndentedJSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums/:id/:sales", getMultiParams)


	router.GET("/telemetry", getMetrices)

	router.Run("0.0.0.0:8080")
}
