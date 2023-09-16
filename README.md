# Content service

### Code generating:
```
contentService %  go run github.com/ogen-go/ogen/cmd/ogen@v0.74.0 --package genapi --target pkg/api/v1/gen --clean pkg/api/v1/openapi.yaml
```

```
go get github.com/labstack/echo/v4 &&
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.15.0 &&
oapi-codegen --generate spec --package genapi pkg/api/v1/openapi.yaml > pkg/api/v1/spec.gen.go &&
oapi-codegen --generate types --package genapi pkg/api/v1/openapi.yaml > pkg/api/v1/types.gen.go &&
oapi-codegen --generate client --package genapi pkg/api/v1/openapi.yaml > pkg/api/v1/client.gen.go &&
oapi-codegen --generate server --package genapi pkg/api/v1/openapi.yaml > pkg/api/v1/server.gen.go 
```