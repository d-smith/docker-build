package quote

import (
	. "github.com/lsegal/gucumber"
	"net/http"
	"fmt"
	"io/ioutil"
)

func init() {
	Given(`^a symbbol$`, func() {

	})

	And(`^a request for a quote for that symbol$`, func() {

	})

	Then(`^the current price is returned$`, func() {
		//T.Errorf("fail!")
		resp, err := http.Get("http://mb:2525/imposters")
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		fmt.Println("Read ->", string(body))
		resp.Body.Close()
	})

}
