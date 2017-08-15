package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/openzipkin/zipkin-api-example/go/client"
	"github.com/openzipkin/zipkin-api-example/go/client/operations"
	"github.com/openzipkin/zipkin-api-example/go/models"
)

func main() {
	transport := httptransport.New("localhost:9411", "/api/v2", []string{"http"})
	zipkin := client.New(transport, strfmt.Default)

	id := uint64(rand.New(rand.NewSource(time.Now().UnixNano())).Int63())
	spanID := strconv.FormatUint(id, 16)
	timestamp := int64(time.Now().UnixNano() / int64(time.Microsecond))
	span := &models.Span{
		TraceID:   &spanID,
		ID:        &spanID,
		Kind:      models.SpanKindSERVER,
		Name:      "test",
		Timestamp: timestamp,
		Duration:  int64(200000),
		LocalEndpoint: &models.Endpoint{
			ServiceName: "testService",
			IPV4:        "192.228.233.62",
			Port:        80,
		},
		RemoteEndpoint: &models.Endpoint{
			IPV4: "192.228.233.62",
		},
		Tags: map[string]string{
			"http.path": "/api",
		},
	}

	resp, err := zipkin.Operations.PostSpans(
		operations.NewPostSpansParams().WithSpans(models.ListOfSpans{span}),
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp.Error()))
}
