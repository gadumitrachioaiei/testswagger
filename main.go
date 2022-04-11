package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/gadumitrachioaiei/testswagger/client"
	"github.com/gadumitrachioaiei/testswagger/client/health"
)

func main() {
	host := "localhost:8080"
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	transport := httptransport.NewWithClient(host, client.DefaultBasePath, client.DefaultSchemes, c)
	apiClient := client.New(transport, strfmt.Default)
	response, err := apiClient.Health.GetHealth(health.NewGetHealthParams())
	if err != nil {
		switch apiErr := err.(type) {
		case *health.GetHealthBadRequest:
			log.Printf("received defined error: %v\n", apiErr)
		case *runtime.APIError:
			log.Printf("received runtime error: %v\n", apiErr.Code)
		default:
			log.Printf("no idea what error this is: %v\n", apiErr)
		}
		return
	}
	log.Print("response is: ", response)
}
