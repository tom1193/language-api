package websearch

import (
	"github.com/tom1193/language-api/proto"
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"
	//"fmt"
)

const host = "https://api.cognitive.microsoft.com"
const path = "/bing/v7.0/images/search"
const bingCreds = "4d316685980946218b2377cc6eea2b54"
const count = "5"

func ImageQuery (text string) ([]byte){
	client := &http.Client{}

	//build request
	req, err := http.NewRequest("GET", host+path, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", bingCreds)
	q := req.URL.Query()
	q.Add("q", text)
	q.Add("count", count)
	req.URL.RawQuery = q.Encode()

	//get response
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func ParseImageQueryResponse (body []byte) ([]proto.Image) {
	var result []proto.Image

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	// Extract images
	if data["value"] != nil {
		images := data["value"].([]interface{})
		for i, image := range images {
			imageData := image.(map[string]interface{})
			result = append(result, proto.Image{strconv.Itoa(i), imageData["contentUrl"].(string)})
		}
	}
	return result
}