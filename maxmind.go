package main

import(
	"log" 
	"fmt"
	"net/http"
	"io"

	"github.com/gin-gonic/gin"

)

var Cached float64
var NotCached float64
var Total float64

func Maxmind_Query_IP( c *gin.Context) {

	var body []byte
	var results string

	ip_address := c.Params.ByName("ip_address")

	body = Redis_Check_Cache(ip_address)

	Total++

	if body == nil {

		results = ip_address + " was not cached"
		NotCached++

		client := &http.Client{}

		connect_url := fmt.Sprintf("%s/%s", Config.Maxmind_Url, ip_address)

		req, _ := http.NewRequest("GET", connect_url, nil)
		req.Header.Add("Authorization", "Basic "+basicAuth(Config.Maxmind_Username, Config.Maxmind_Password))
		resp, _ := client.Do(req)

		body, _ = io.ReadAll(resp.Body) // DEBUG : err check!

		Redis_Store_Cache(string(body), ip_address)

		} else { 

		results = ip_address + " pull from cache."
		Cached++

		}

		log.Printf("[ Total Queries: %v | Cached: %v [%v%%]| Not Cached: %v [%v%%] - %s\n", Total, Cached, ( Cached / Total ) * 100, NotCached, (NotCached / Total) * 100, results )


		c.String(http.StatusOK, string(body))



}
