// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Removes an exchange.
	// (DELETE /exchanges/{id})
	DeleteExchangesId(c *gin.Context, id string)
	// Returns one exchange by ID.
	// (GET /exchanges/{id})
	GetExchangesId(c *gin.Context, id string)
	// Update the status of an exchange
	// (PATCH /exchanges/{id})
	PatchExchangesId(c *gin.Context, id string)
	// Create a new exchange
	// (POST /exchanges/{trader_email}/{tradee_email})
	PostExchangesTraderEmailTradeeEmail(c *gin.Context, traderEmail string, tradeeEmail string)
	// Returns a list of all games
	// (GET /games)
	GetGames(c *gin.Context)
	// Adds a game to the list.
	// (POST /games)
	PostGames(c *gin.Context)
	// Removes a game from the exchange list.
	// (DELETE /games/{id})
	DeleteGamesId(c *gin.Context, id string)
	// Returns one game by ID.
	// (GET /games/{id})
	GetGamesId(c *gin.Context, id string)
	// Updates a property of a game.
	// (PATCH /games/{id})
	PatchGamesId(c *gin.Context, id string)
	// Returns a list of all users
	// (GET /users)
	GetUsers(c *gin.Context)
	// Adds a game to the list.
	// (POST /users)
	PostUsers(c *gin.Context)
	// Removes a user.
	// (DELETE /users/{id})
	DeleteUsersId(c *gin.Context, id string)
	// Returns one user by ID.
	// (GET /users/{id})
	GetUsersId(c *gin.Context, id string)
	// Updates a property of a user.
	// (PATCH /users/{id})
	PatchUsersId(c *gin.Context, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// DeleteExchangesId operation middleware
func (siw *ServerInterfaceWrapper) DeleteExchangesId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteExchangesId(c, id)
}

// GetExchangesId operation middleware
func (siw *ServerInterfaceWrapper) GetExchangesId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetExchangesId(c, id)
}

// PatchExchangesId operation middleware
func (siw *ServerInterfaceWrapper) PatchExchangesId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchExchangesId(c, id)
}

// PostExchangesTraderEmailTradeeEmail operation middleware
func (siw *ServerInterfaceWrapper) PostExchangesTraderEmailTradeeEmail(c *gin.Context) {

	var err error

	// ------------- Path parameter "trader_email" -------------
	var traderEmail string

	err = runtime.BindStyledParameter("simple", false, "trader_email", c.Param("trader_email"), &traderEmail)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter trader_email: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "tradee_email" -------------
	var tradeeEmail string

	err = runtime.BindStyledParameter("simple", false, "tradee_email", c.Param("tradee_email"), &tradeeEmail)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter tradee_email: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostExchangesTraderEmailTradeeEmail(c, traderEmail, tradeeEmail)
}

// GetGames operation middleware
func (siw *ServerInterfaceWrapper) GetGames(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetGames(c)
}

// PostGames operation middleware
func (siw *ServerInterfaceWrapper) PostGames(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostGames(c)
}

// DeleteGamesId operation middleware
func (siw *ServerInterfaceWrapper) DeleteGamesId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteGamesId(c, id)
}

// GetGamesId operation middleware
func (siw *ServerInterfaceWrapper) GetGamesId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetGamesId(c, id)
}

// PatchGamesId operation middleware
func (siw *ServerInterfaceWrapper) PatchGamesId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchGamesId(c, id)
}

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsers(c)
}

// PostUsers operation middleware
func (siw *ServerInterfaceWrapper) PostUsers(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUsers(c)
}

// DeleteUsersId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUsersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUsersId(c, id)
}

// GetUsersId operation middleware
func (siw *ServerInterfaceWrapper) GetUsersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersId(c, id)
}

// PatchUsersId operation middleware
func (siw *ServerInterfaceWrapper) PatchUsersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchUsersId(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.DELETE(options.BaseURL+"/exchanges/:id", wrapper.DeleteExchangesId)
	router.GET(options.BaseURL+"/exchanges/:id", wrapper.GetExchangesId)
	router.PATCH(options.BaseURL+"/exchanges/:id", wrapper.PatchExchangesId)
	router.POST(options.BaseURL+"/exchanges/:trader_email/:tradee_email", wrapper.PostExchangesTraderEmailTradeeEmail)
	router.GET(options.BaseURL+"/games", wrapper.GetGames)
	router.POST(options.BaseURL+"/games", wrapper.PostGames)
	router.DELETE(options.BaseURL+"/games/:id", wrapper.DeleteGamesId)
	router.GET(options.BaseURL+"/games/:id", wrapper.GetGamesId)
	router.PATCH(options.BaseURL+"/games/:id", wrapper.PatchGamesId)
	router.GET(options.BaseURL+"/users", wrapper.GetUsers)
	router.POST(options.BaseURL+"/users", wrapper.PostUsers)
	router.DELETE(options.BaseURL+"/users/:id", wrapper.DeleteUsersId)
	router.GET(options.BaseURL+"/users/:id", wrapper.GetUsersId)
	router.PATCH(options.BaseURL+"/users/:id", wrapper.PatchUsersId)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RY3W7jNhN9FWK+71L+2XYLLHSXbhaBi6IJ2k1vgiBgxLHNhURyScqJEejdC45+bMeS",
	"IydxnNzJMjmcOXN4eMQHSHRmtELlHcQP4JI5Zpwev90nc65mGJ6N1Qatl0j/zHiGUoQnvzQIMThvpZpB",
	"BPeDmR5UL42VmfRygcPz2x+Y+Mnp+oCBzIy2noJzP4cYZnqYaTXT4nao7WxEzwNh5QLt6NZpNWoCQhEC",
	"aW7kINECZ6gGeO8tH3g+owTD8BCRZ3gjRaQz6TEzfglFEcFHyHw7a+e5z2kIqjyD+AoMKlHmzpMEjUcB",
	"EQhMUqlQwHX0qMQiAm+5QLzBjMt0G4N6gO0cUDQxNeESppzxrIUhiVZCehlqWSWcSeUhgpnWIdMplxYi",
	"MFrb1mQ/ZptUBcdWOfpOof0INVWJPqrL5LepdHO0rcVZTJE7vFkitz15c+nKWJu86abmx6SD4c7daSta",
	"S3LeIvobLoRF51qH5A5tB6W2QSVxU1NNg6VPw3//SoGahU3Kaj1nJxcTNmDnBlV4+nU4Ho4hggVaRxsW",
	"PtGbwFmDihsJMdSDAnSU6QiraG70IEURXglM0VOmoaM87P6JgBhO6X29upsIimN5hh6tg/gqTHWJlaYU",
	"DPg+R1aHZ5PTIYSyQoND3+otFggRiPczlxYFxN7mGFXHVxta12GwM1q5kmm/jD+XSa+v3EBkMdMLFMzl",
	"SYLOTfM0XQ4DJJ93TlPas6nOlRiWmp1nGbdLiOFviucYV01loaySP1dQv4PrIoIZ+m0Qz9C/PwTHldR7",
	"VJQyNyaVCSU9+uFK8V/F+7/FKcTwv9HKcYwquzFqvAaReLOYP/45/4vpKfNzLMFdg/AlHfG5VY5ptQbV",
	"7bJCq7Uxhvtkvt2ai/D6XTTnZ47O/67Fcq++bCrwy31GuzBtllIciUsNL3IjuH/FDX5J8YijJYKBsWu7",
	"vZ1TRbQhpOv2q6h+VnaNFNZo1yINF9qttOE7hfgWptAj0mMfSpaLdxNywxvuQ82oczV8YjV8xmqvthG6",
	"PnKeye9Pb8vvxGIXv8fb/J6oBU+lYFKZnNzZb+2jfHAjKXNoF2gZWqvto33wldZlnCm860H+AHKJdseh",
	"d0YDDqgW9P3SguQJo3OHW8uXYS+XmR4ewPpc4iyVzpOKpGm1+grG8jcdSp2asILueduhH2pvx/uuTpG5",
	"5UIcge0nQoROhW4wr0n9Q9eGLZ1q2N7TL1P3+pkJWv5NfTJBvp9HpilP++OymKnVGaHZ+KQuWKPdyvEu",
	"ADy8WG1b5FDSHjvgWS1bGWgCcMs8r4nUDud81D4dxjFXV1vVRVd17dVx0XVst7xTVfdzybv5UjrksMUr",
	"+OhY5RVT2wUzd0SGbntwSQMOCA9dU/WwB2Wmx7IHeQVDjWH5e7c9WEH3+vZghdrb2YOuTr1ve1B3qmF7",
	"T3tA3eunmiHs29qDkNye9oCm9LAHoZhhK9N3SsS7QOrwqrTtA0q8XuYDnurNygcQgFs+YE2NdviAo/bp",
	"NXzAjrv+Yx/yXXyhzu53yOc7ydB1yHdtW5pNAlr2OrcpxPCnTng6187HX8ZfxhBaVE18nM1XrTyXG+ch",
	"+YeIcSUYT8IwV2ov9434VkSpP+l7BKV8ewWtjUCPoM3NX6/Azc1JcV38FwAA//+xVBpAth8AAA==",
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
