/*
 * example
 *
 * description example.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"strings"
)

// ServiceApiController binds http requests to an api service and writes the service results to the http response
type ServiceApiController struct {
	service ServiceApiServicer
	errorHandler ErrorHandler
}

// ServiceApiOption for how the controller is set up.
type ServiceApiOption func(*ServiceApiController)

// WithServiceApiErrorHandler inject ErrorHandler into controller
func WithServiceApiErrorHandler(h ErrorHandler) ServiceApiOption {
	return func(c *ServiceApiController) {
		c.errorHandler = h
	}
}

// NewServiceApiController creates a default api controller
func NewServiceApiController(s ServiceApiServicer, opts ...ServiceApiOption) Router {
	controller := &ServiceApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ServiceApiController
func (c *ServiceApiController) Routes() Routes {
	return Routes{ 
		{
			"Liveness",
			strings.ToUpper("Get"),
			"/v1/liveness",
			c.Liveness,
		},
		{
			"Ready",
			strings.ToUpper("Get"),
			"/v1/ready",
			c.Ready,
		},
	}
}

// Liveness - 
func (c *ServiceApiController) Liveness(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Liveness(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// Ready - 
func (c *ServiceApiController) Ready(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Ready(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
