package main

import (
	"fmt"
	"github.com/d-smith/docker-build/customctx"
	"github.com/d-smith/docker-build/customctx/timing"
	"github.com/d-smith/docker-build/services/quote"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"net/url"
)

	func TestQuote(t *testing.T) {
	soapServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<SOAP-ENV:Envelope
  xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
  SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
   <SOAP-ENV:Body>
       <m:GetLastTradePriceResponse xmlns:m="Some-URI">
           <Price>34.5</Price>
       </m:GetLastTradePriceResponse>
   </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`)
	}))

	defer soapServer.Close()

		soapURL,_ := url.Parse(soapServer.URL)

	wrapped := quote.Middleware(customctx.ContextHandlerFunc(quote.NewQuoteHandler(soapURL.Host)))
	wrapped = timing.TimerMiddleware(wrapped)

	h := &customctx.ContextAdapter{
		Ctx:     context.Background(),
		Handler: wrapped,
	}

	ts := httptest.NewServer(h)
	defer ts.Close()



	req, err := http.NewRequest("GET", ts.URL+"/quote/MSFT", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	println(string(body))

	if !strings.Contains(string(body), "34.5") {
		t.Fail()
	}
}
