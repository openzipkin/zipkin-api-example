package main

import (
	"flag"
	"fmt"
	"log"
	"encoding/json"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/openzipkin/zipkin-api-example/go/client"
	"github.com/openzipkin/zipkin-api-example/go/client/operations"
	"github.com/go-openapi/strfmt"
	"time"
)

func main() {
	var dev bool
	flag.BoolVar(&dev, "d", false, "Use the dev zipkin configurations")

	flag.Parse()
	serviceName := flag.Arg(0)

	// http://goswagger.io/generate/client/
	apiHost := "prod-host.example.com"
	if dev {
		apiHost = "dev-host.example.com"
	}
	transport := httptransport.New(fmt.Sprintf("%s", apiHost), "/api/v1", []string{"https"})
	zipkin := client.New(transport, strfmt.Default)

	limit := int64(10)
	now := int64(time.Now().UnixNano() / int64(time.Millisecond))
	lookback := int64(36000000)
	resp, err := zipkin.Operations.GetTraces(operations.NewGetTracesParams().
		WithServiceName(&serviceName).
		WithLimit(&limit).
		WithEndTs(&now).
		WithLookback(&lookback))
	if err != nil {
		log.Fatal(err)
	}
	payload, _ := json.Marshal(resp.Payload)
	fmt.Println(string(payload))
}