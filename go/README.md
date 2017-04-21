## Swagger

The examples in this directory assume that the zipkin API client has been generated using
[go-swagger](https://github.com/go-swagger/go-swagger) into the `client` directory.

```bash
swagger generate client -f ../zipkin-api/zipkin-api.yaml -A zipkin
```
