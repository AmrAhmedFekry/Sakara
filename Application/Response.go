package Application

import (
	"github.com/bykovme/gotrans"
)

// return success json response
func (req Request) Success(body interface{}) {
	req.Response(200, buildResponse(body, gotrans.T("success"), 200, nil))
}

// return success created json response
func (req Request) Created(body interface{}) {
	req.Response(201, buildResponse(body, gotrans.T("created"), 201, nil))
}

// return not auth response
func (req Request) NotAuth() {
	req.Response(401, buildResponse(nil, gotrans.T("you_are_not_auth"), 401, nil))
}

// return validation error response
func (req Request) BadRequest(err interface{}) {
	req.Response(422, buildResponse(nil, gotrans.T("validation_error"), 401, err))
}

// return resource not response
func (req Request) ResourceNotFound(resourceName string) {
	req.Response(404, buildResponse(nil, gotrans.T(resourceName+"_not_found"), 404, nil))
}

// return validation error response
func (req Request) ProductDeleted() {
	req.Response(200, buildResponse(nil, gotrans.T("product_deleted"), 200, nil))
}

// return resource already exits
func (req Request) ResourceAlreadyExists(resourceName string) {
	req.Response(409, buildResponse(nil, gotrans.T(resourceName+"_is_already_exits"), 409, nil))
}

// standardization our response
func buildResponse(payload interface{}, message string, code int, err interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["payload"] = payload
	response["message"] = message
	response["code"] = code
	response["errors"] = err

	return response
}
