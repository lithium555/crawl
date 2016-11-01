package aliExpress

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://www.aliexpress.com"

func main() {
	//newFile, err := os.Create("site.html")
	//if err != nil{
	//	log.Fatal(err)
	//}
	//defer newFile.Close()

	//get the response
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept-Language", "en-US")
	resp, err := http.DefaultClient.Do(req)

	//body
	body, err := ioutil.ReadAll(resp.Body)

	//header
	var header string
	for h, v := range resp.Header {
		for _, v := range v {
			header += fmt.Sprintf("%s %s \n", h, v)
		}
	}

	// append all to slice
	var write []byte
	write = append(write, body...)

	//write it to a file
	err = ioutil.WriteFile("site.html", write, 0644)
	if err != nil {
		panic(err)
	}
}
