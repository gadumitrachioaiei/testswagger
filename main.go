package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/gadumitrachioaiei/testswagger/client"
	"github.com/gadumitrachioaiei/testswagger/client/health"
	"github.com/gadumitrachioaiei/testswagger/client/test"
)

func main() {
	host := "localhost:8080"
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	transport := httptransport.NewWithClient(host, client.DefaultBasePath, client.DefaultSchemes, c)
	apiClient := client.New(transport, strfmt.Default)
	testRawJson(apiClient)
	testRawJsonHttp()
}

func testGetHealth(apiClient *client.PerceptorDecisionsAPI) {
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
	log.Println("getHealth response is: ", response)
}

// testRawJson reads the data from the server, which is not deserialized yet, because we used the type json.RawMessage to store it.
// so this means that we can deserialize it ourselves,into whatever we think it is appropiate.
// we can do such things when the server returns a type that we don't like, but for which we have an equivalent type that we do like.
func testRawJson(apiClient *client.PerceptorDecisionsAPI) {
	testRawJsonResponse, err := apiClient.Test.Testrawjson(test.NewTestrawjsonParams())
	if err != nil {
		log.Fatalf("unexpected error calling the api: %v", err)
	}
	dataBytes := []byte(testRawJsonResponse.Payload.Data.RawMessage)
	type Data struct {
		A string
	}
	var data Data
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	log.Println("testRawJsonResponse", testRawJsonResponse.Payload.Kind, data)
}

func testRawJsonHttp() {
	resp, err := http.Get("http://localhost:8080/api/rawjson")
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode > http.StatusOK {
		log.Fatalf("unexpected status: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	var so interface{}
	if err := json.NewDecoder(resp.Body).Decode(&so); err != nil {
		log.Fatalf("unexpected err: %v", err)
	}
	log.Println("response is", so)
}
