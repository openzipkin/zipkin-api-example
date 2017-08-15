## Swagger

The examples in this directory assume that the [Zipkin API](http://zipkin.io/zipkin-api/#/default/post_spans) client has been generated using
[go-swagger](https://github.com/go-swagger/go-swagger) into the `client` directory.

```bash
# once you've installed swagger, build the client
$ wget https://raw.githubusercontent.com/openzipkin/zipkin-api/master/zipkin2-api.yaml
$ swagger generate client -f zipkin2-api.yaml -A zipkin
# now, you can run one of the examples, such as POSTing a test span
$ DEBUG=1 go run post_sample_trace.go
```
