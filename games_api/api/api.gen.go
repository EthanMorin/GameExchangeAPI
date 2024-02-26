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
	// Returns one exchange by ID.
	// (GET /exchanges/{id})
	GetExchangesId(c *gin.Context, id string)
	// Update the status of an exchange
	// (PATCH /exchanges/{id})
	PatchExchangesId(c *gin.Context, id string)
	// Create a new exchange
	// (POST /exchanges/{traderid}/{tradeeid})
	PostExchangesTraderidTradeeid(c *gin.Context, traderid string, tradeeid string)
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

// PostExchangesTraderidTradeeid operation middleware
func (siw *ServerInterfaceWrapper) PostExchangesTraderidTradeeid(c *gin.Context) {

	var err error

	// ------------- Path parameter "traderid" -------------
	var traderid string

	err = runtime.BindStyledParameter("simple", false, "traderid", c.Param("traderid"), &traderid)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter traderid: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "tradeeid" -------------
	var tradeeid string

	err = runtime.BindStyledParameter("simple", false, "tradeeid", c.Param("tradeeid"), &tradeeid)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter tradeeid: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostExchangesTraderidTradeeid(c, traderid, tradeeid)
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

	router.GET(options.BaseURL+"/exchanges/:id", wrapper.GetExchangesId)
	router.PATCH(options.BaseURL+"/exchanges/:id", wrapper.PatchExchangesId)
	router.POST(options.BaseURL+"/exchanges/:traderid/:tradeeid", wrapper.PostExchangesTraderidTradeeid)
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

	"H4sIAAAAAAAC/9RYTW/jNhD9K8S0R/lj2y2w0C3dFIGLognapJcgCBhxbHMhkVyScmIE/u8FRx+OY0kr",
	"J7Gd3GR5OBq+eXzzpEdIdGa0QuUdxI/gkjlmnC7/eEjmXM0wXBurDVovkf6Z8Yzu/mxxCjH8NFrnGJUJ",
	"RmchZhWBFCHSLw1CDM5bqWYQwcNgpgflTWNlJr1c4PD87hsmfnL6NGAgM6OtpyK4n0MMMz3MtJppcTfU",
	"djai64GwcoF2dOe0GtUJw/MfBpobOUi0wBmqAT54yweez2gjIRxiuJUi0pn0mBm/hNUqAue5zykEVZ5B",
	"fA0GlShq50mCxqOACAQmqVQo4CZ6tsVVBN5ygXiLGZfpNgZVgG0NWNU5NeESlpyVyG/2I9FKSC/DXtYF",
	"Z1J5iGCmdah0yqWFCIzWtrHYj9kmVcKxtR19r9B+hD2VhT7bl8nvUunmaBs3ZzFF7vB2idz25M2VK3Jt",
	"8qadmuGEU0ioyvU962Uabi1fflxSGe7cvbaiERjnLaK/5UJYdK4xJHdoW4i53ZpwS6qppmDp0/Dff1Kg",
	"ZgFUVmkwO7mYsAE7N6jC1a/D8XAMESzQOjr28InuBOYbVNxIiKEKCtBRpSMss7nRoxQrUnIkdAMpeBCQ",
	"iYAYztBXz3UTQRksz9CjdRBfP4JAl1hpCsGByzmyKjGbnA4hbCi0NnSsOqKBCoG433NpUUDsbY5ROWya",
	"cLoJwc5o5Qoa/jIel0LnUVHJ3JhUJlT06JsrpG+dr4ut9Vwj8Dc38+e/538zPWV+jmyqcyXqrQ0DuJ/H",
	"n0P2zTV1j5T2xaJhMUPyLON2CTH8gz63yjGtnkB1tyzRKgh5DdU/cEMk9Ml8uzUX4fa7aM73HJ3/XYvl",
	"Tn3Z1J/XT9nmA7W5ldWRuFTzIjeCexTM5UmCzk3zNF2+hk5XlI84WiAYGMtV3ehmTq2iDQEozIcUq/IS",
	"S0Uw2jVIwoV2a024LJdeluv6kLB4XDsFq3J2ImLU+iT8wZPwWJQPg7VpLL6QyZ8Oy+TEYhuTx9tMnqgF",
	"T6VgUpmcXMhvzVE+zMuUObQLtAyt1fYZ47/ScxlnCu970Lx2L23j7YwC9qgLhSPaRvKE0YQhhxRObVHp",
	"/gGsJhBnqXSe9CJNy6evYSx+0/hpVYE1dC87Dv1QOxzv2zpF9osLcQS2nwgROhW6wbwmnQ9dGzZ0qmZ7",
	"7egEpuhxu3endJ+618820OP35+caJh9BbjHTi97DkpZ0+K6Qq8ZyanVGaNaOqA3WqFs53gWA+xerbTMc",
	"trTDCXhRy9ZWmQDcsslPRKrDIx+1T29hFDbqbHjL3OkN83CGuFNOdzPC3UQpTHA42yVuNE95SdFmpQxv",
	"5p2+4IoC9ggPfYfp4QuKSo/lC/IShgrD4ne3L1hD9/a+YI3a4XxBW6fety+oOlWzvacvoO71k8uQ9rC+",
	"IBS3oy+gJT18QdjMsJHpnRLxLpDavyptG4ACr9cZgB/1Zm0ACMAtA/BEjToMwFH79BYGoPMz9M7fmA/n",
	"ANrIRG3fzQHknUxpcwBtZ5pWk7oWRMhtCjH8pROezrXz8ZfxlzGE/pULn1fzVSvP5cawJHMRMa4E40kI",
	"c4Uwc18rc8mi6kW/R1Kqt1fSyiX0SFp/+euVuP6esrpZ/R8AAP//fIGxaiIeAAA=",
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
