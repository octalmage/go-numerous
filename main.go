package main

import "github.com/segmentio/go-log"
import "net/http"
import "net/url"
import "encoding/json"
import "io/ioutil"
import "strings"
import "os"
import "fmt"


func main() {

	username := os.Args[1]
	id := os.Args[2]
	value := os.Args[3]

	events := "https://api.numerousapp.com/v1/metrics/" + id + "/events"
	v := url.Values{"value": {value}}

	req, err := http.NewRequest("POST", events, strings.NewReader(v.Encode()))

	req.SetBasicAuth(username, "")

	cli := &http.Client{}
	res, err := cli.Do(req)

	if err != nil {
 		log.Error("Error")
 	}

	//var body string
	//json.NewDecoder(res.Body).Decode(&body)
	contents, err := ioutil.ReadAll(res.Body)
	fmt.Println(contents)
}
