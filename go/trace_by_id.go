package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/openzipkin/zipkin-api-example/go/client"
	"github.com/openzipkin/zipkin-api-example/go/client/operations"
)

func main() {
	var dev bool
	flag.BoolVar(&dev, "d", false, "Use the dev zipkin configurations")

	flag.Parse()
	traceID := flag.Arg(0)

	// http://goswagger.io/generate/client/
	apiHost := "prod-host.example.com"
	if dev {
		apiHost = "dev-host.example.com"
	}
	transport := httptransport.New(fmt.Sprintf("%s:9411", apiHost), "/api/v1", []string{"http"})
	zipkin := client.New(transport, strfmt.Default)

	resp, err := zipkin.Operations.GetTraceTraceID(operations.NewGetTraceTraceIDParams().WithTraceID(traceID))
	if err != nil {
		log.Fatal(err)
	}
	payload, _ := json.Marshal(resp.Payload)
	fmt.Println(string(payload))
}
