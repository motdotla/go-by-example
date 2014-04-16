package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{}
	res, err := http.PostForm("http://requestb.in/api/v1/bins", values)

	if err != nil {

	} else {
		defer res.Body.Close()

		var v map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&v)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(v["name"])
	}
}
