package getapis

import (
	"io/ioutil"
	"log"
	"net/http"
)

func get(url string) string {
	var username string = "inter"
	var passwd string = "com"
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://"+url, nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}
