// Package plugin provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package plugin

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Post config to the scanner plugin and start scanner.
	// (POST /config)
	PostConfig(ctx echo.Context) error
	// Check the scanner plugin's health.
	// (GET /healthz)
	GetHealthz(ctx echo.Context) error
	// Get metadata from the scanner plugin.
	// (GET /metadata)
	GetMetadata(ctx echo.Context) error
	// Get the status of the scanner.
	// (GET /status)
	GetStatus(ctx echo.Context) error
	// Stop the scanner.
	// (POST /stop)
	PostStop(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostConfig converts echo context to params.
func (w *ServerInterfaceWrapper) PostConfig(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostConfig(ctx)
	return err
}

// GetHealthz converts echo context to params.
func (w *ServerInterfaceWrapper) GetHealthz(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetHealthz(ctx)
	return err
}

// GetMetadata converts echo context to params.
func (w *ServerInterfaceWrapper) GetMetadata(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMetadata(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// PostStop converts echo context to params.
func (w *ServerInterfaceWrapper) PostStop(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostStop(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/config", wrapper.PostConfig)
	router.GET(baseURL+"/healthz", wrapper.GetHealthz)
	router.GET(baseURL+"/metadata", wrapper.GetMetadata)
	router.GET(baseURL+"/status", wrapper.GetStatus)
	router.POST(baseURL+"/stop", wrapper.PostStop)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
