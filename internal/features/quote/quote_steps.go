package quote

import (
	. "github.com/lsegal/gucumber"
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)




func postConfig() {

	payloadReader,err := os.Open("imposter.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer payloadReader.Close()

	resp, err := http.Post("http://mb:2525/imposters", "application/json", payloadReader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Println("Read ->", string(body))
	resp.Body.Close()
}

func getAndPrint(endpoint string) string {
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Print(err.Error())
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return ""
	}

	fmt.Println("Read ->", string(body))
 	resp.Body.Close()

	return string(body)
}

func init() {
	Given(`^a symbbol$`, func() {

	})

	And(`^a request for a quote for that symbol$`, func() {

	})

	Then(`^the current price is returned$`, func() {
		//T.Errorf("fail!")
		getAndPrint("http://testimg:2525/imposters")

		postConfig()

		getAndPrint("http://mb:2525/imposters")
		getAndPrint("http://mb:2525/imposters/4545")
		quote := getAndPrint("http://testimg:8080/quote/xxx")
		if strings.Trim(quote, "\n") != "34.5" {
			T.Errorf("Expected 34.5 got '%s'", quote)
		}
	})

}
