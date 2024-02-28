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
	// Returns API metrics
	// (GET /metrics)
	GetMetrics(c *gin.Context)
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

// GetMetrics operation middleware
func (siw *ServerInterfaceWrapper) GetMetrics(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetMetrics(c)
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
	router.GET(options.BaseURL+"/metrics", wrapper.GetMetrics)
	router.GET(options.BaseURL+"/users", wrapper.GetUsers)
	router.POST(options.BaseURL+"/users", wrapper.PostUsers)
	router.DELETE(options.BaseURL+"/users/:id", wrapper.DeleteUsersId)
	router.GET(options.BaseURL+"/users/:id", wrapper.GetUsersId)
	router.PATCH(options.BaseURL+"/users/:id", wrapper.PatchUsersId)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RZTW/jNhD9KwTbo/yx7RZY6JbuLgIXbRO0m16CIGDEsc2FRHJJyokR6L8XM/qwHUuO",
	"nMRxcpOV0XDmzePTE3PPE5NZo0EHz+N77pM5ZIIuv94lc6FngNfWGQsuKKC/zEQGSuJVWFrgMffBKT3j",
	"Eb8bzMygummdylRQCxie3XyHJEy+rAcMVGaNC5RchDmP+cwMM6NnRt4MjZuN6HognVqAG914o0dNQl5g",
	"IiOsGiRGwgz0AO6CE4MgZlQghmNGkcG1kpHJVIDMhiUvioi/h8q3q/ZBhJxCQOcZjy+5BS3L2kWSgA0g",
	"ecQlJKnSIPlV9KDFIuLBCQlwDZlQ6TYGdYDrDCianIZwwUdORdbCkMRoqYLCXlYFZ0oHHvGZMVjpVCjH",
	"I26Nca3Fvs8x6QqOrXbMrQb3HnqqCn3Ql81vUuXn4Fqbc5CC8HC9BOF68ubCl7k2edNNzfdJByu8vzVO",
	"trbkgwMI10JKB963huQeXAeltkElcdNTQ8EqpPi3/5QEw3CTslrP2cn5hA3YmQWNV78Ox8Mxj/gCnKcN",
	"yz/QHeSsBS2s4jGvgxA6qnQEVTY/uleywFsSUghUKU5U4O6fSB7zL3S/Xt1PJOVxIoMAzvP4Eh/1iVO2",
	"FAz+bQ6sTs8mX4Yc28IB49zqLYaEQOL9yJUDyePgcoiq11cbWlcY7K3RvmTaL+OPZdHrKzcQOcjMAiTz",
	"eZKA99M8TZdDhOTjzse0CWxqci2HpWbnWSbcksf8H8rnmdBNZ9hWyZ9LXt/jV0XEZxC2QTyF8PYQHFdS",
	"H0BTycLaVCVU9Oi7L8V/le9nB1Me859GK8cxquzGqPEaROLNZv749+xvZqYszKEEdw3C50wk5E57ZvQa",
	"VDfLCq3WwVgRkvn2aM7x9psYzo8cfPjdyOVec9lU4Of7jHZh2mylOBKXGl7kVorwghv8gvIRR0sEkbFr",
	"u72dU0W0IaTr9quoflZ2jRTWGt8iDefGr7ThG6X4io/QJdBlH0qWi3cTcsMb7kPNqHM1eGQ1eMJqL7YR",
	"uj5ynsjvD6/L78RBF7/H2/ye6IVIlWRK25zc2W/tUQHdSMo8uAU4Bs4Z92AffKZ1mWAabnuQH0Eu0e54",
	"6Z1SwAHVgr5fWpA8YfTeEc6JJe7lstLDA1i/lwRLlQ+kImlarb6CsfxNL6VOTVhB97Tt0A+11+N916TI",
	"3Aopj8D2EylxUjgNFgypP05t2DKphu09/TJNr5+ZoOVf1ScT5Pt5ZHrkcX9cNjN1JiM0G5/UBWu0Wzne",
	"BICHF6tti4wt7bEDnjSylYEmALfM85pI7XDOR53TYRxzdbRVHXRVx14dB13Hdss7VXU/l7ybL6VDxi1e",
	"wUevVVExtV0wMwhOJTsNwl9VSDtEmwVWsQyBEwl+yHSJfx0pza3u4P3J+YRlzeJ19fWdsv7cE5m7q7+g",
	"gAOOl47ZetibstJj2Zu8gqFGsfy9296soHt5e7NC7fXsTdek3ra9qSfVsL2nvaHp9VN9TPu69gaL29Pe",
	"0CM97A02M2xl+k6JeBNIHV6Vtn1MidfzfMxjs1n5GAJwy8esqdEOH3PUOb2Ej9nxv4pjm5QuvtBk9zMp",
	"+U4ydJmUrm1LT5OAlrPOXcpj/qdJRDo3PsSfxp/GHEdUPfiwms/oRNTG+5D8T8SElkwkGOZL7RWhEd+K",
	"KPWRRI+kVG+vpLUR6JG0Obnslbg5+WnPrYT2tZ9iU+NwABmEOeR+LUltr4qr4v8AAAD//8afo/S6IAAA",
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
