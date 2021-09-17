package mongo

import (
    "context"
    "encoding/json"
    "fmt"
    httpTransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
    "log"
    "github.com/sgahlot/go-service-quickstart/pkg/common"
    "net/http"
)

func logRequest(h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Received request: %v (address:%v)\n", r.URL, r.RemoteAddr)
        h.ServeHTTP(w, r)
        fmt.Printf("Response: %+v\n", w)
    }
    return http.HandlerFunc(fn)
}

type ErrorHandler struct {
}

func (handler ErrorHandler) Handle(ctx context.Context, err error) {
    log.Printf("Handled error: %+v\n", err)
}

func CreateHandlers(_ context.Context, endpoint EndPoints) http.Handler {
    router := mux.NewRouter()
    options := []httpTransport.ServerOption{
        httpTransport.ServerErrorHandler(ErrorHandler{}),
        httpTransport.ServerErrorEncoder(encodeError),
    }

    router.Use(logRequest)

    router.Methods(common.POST, common.PUT).Path("/api/v1/fruits").
        Handler(httpTransport.NewServer(
            endpoint.InsertFruit,
            decodeInsertFruitRequest,
            encodeFruitResponse,
            options...,
    ))

    router.Methods(common.GET).Path("/api/v1/fruits").
        Handler(httpTransport.NewServer(
            endpoint.GetFruits,
            decodeGetFruitsRequest,
            encodeFruitResponse,
            options...,
    ))

    return router
}

func setContentType(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
    if err == nil {
        panic("encodeError with nil error")
    }

    setContentType(w)
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "error": err.Error(),
    })
}