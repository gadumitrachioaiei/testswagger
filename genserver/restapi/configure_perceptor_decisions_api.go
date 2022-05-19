// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gadumitrachioaiei/testswagger/genserver/models"
	"github.com/gadumitrachioaiei/testswagger/genserver/restapi/operations"
	"github.com/gadumitrachioaiei/testswagger/genserver/restapi/operations/health"
	"github.com/gadumitrachioaiei/testswagger/genserver/restapi/operations/test"
)

//go:generate swagger generate server --target ../../genserver --name PerceptorDecisionsAPI --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.PerceptorDecisionsAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PerceptorDecisionsAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.HealthGetHealthHandler == nil {
		api.HealthGetHealthHandler = health.GetHealthHandlerFunc(func(params health.GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation health.GetHealth has not yet been implemented")
		})
	}
	api.TestTestrawjsonHandler = test.TestrawjsonHandlerFunc(func(params test.TestrawjsonParams) middleware.Responder {
		so := models.SomeObject{
			Data: struct{ A string }{A: "some data"},
			Kind: &models.SomeObjectKind{Data: "kind1"},
		}
		return &test.TestrawjsonOK{Payload: &so}
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
