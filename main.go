package main

/*
	From https://joeshaw.org/net-context-and-http-handler/
*/

import (
	"fmt"
	"github.com/d-smith/docker-build/customctx"
	"github.com/d-smith/docker-build/customctx/reqid"
	"github.com/d-smith/docker-build/customctx/timing"
	"github.com/d-smith/docker-build/services/quote"
	"golang.org/x/net/context"
	"net/http"
)

func handler(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	reqID := reqid.RequestIDFromContext(ctx)
	fmt.Fprintf(rw, "Hello request ID %s\n", reqID)
}

func main() {
	wrapped := quote.Middleware(customctx.ContextHandlerFunc(quote.NewQuoteHandler("localhost:4545")))
	wrapped = timing.TimerMiddleware(wrapped)

	h := &customctx.ContextAdapter{
		Ctx:     context.Background(),
		Handler: wrapped,
	}
	http.ListenAndServe(":8080", h)
}
