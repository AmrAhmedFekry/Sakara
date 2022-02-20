package Application

import (
	"GoGin/Middlewares"
	"GoGin/Models"
	"GoGin/Utils/Token"
	"database/sql"

	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type ShareResources interface {
	Share()
}

type Request struct {
	Context         *gin.Context
	DB              *gorm.DB
	Connection      *sql.DB
	User            *Models.User
	IsAuth          bool
	IsSeller        bool
	IsBuyer         bool
	Lang            string
	ValidationError error
}

func (req *Request) Share() {}

// Handle request closure
func req() func(c *gin.Context) *Request {
	return func(c *gin.Context) *Request {
		var request Request
		request.Context = c
		connectToDatabase(&request)
		setLang(&request)
		return &request
	}
}

// Set Request language
func setLang(req *Request) {
	requestLang := gotrans.DetectLanguage(req.Context.GetHeader("Accept-language"))
	gotrans.SetDefaultLocale(requestLang)
	req.Lang = requestLang
}

// Init new request closure
func NewRequest(c *gin.Context) *Request {
	request := req()
	return request(c)
}

// Init new request closure with auth middleware
func NewRequestWithAuth(c *gin.Context) *Request {
	return NewRequest(c).Auth()
}

// return json response
func (req Request) Response(code int, data map[string]interface{}) {
	CloseConnection(&req)
	req.Context.JSON(code, data)
}

// Get Auth user and return true if user is auth
func RequestAuth(c *gin.Context) (*Request, bool) {
	r := NewRequest(c)
	if !r.IsAuth {
		r.NotAuth()
		return r, false
	}
	return r, true
}

// Find user by token and set user to request
func (req *Request) Auth() *Request {
	req.IsAuth = false
	req.IsSeller = false
	req.IsBuyer = false
	isTokenValid := Middlewares.IsAuth(req.Context)
	if isTokenValid {
		userId, _ := Token.ExtractTokenID(req.Context)
		req.DB.Where("id = ?", userId).First(&req.User)
		if req.User.ID != 0 {
			req.IsAuth = true
			if req.User.Role == "seller" {
				req.IsSeller = true
			}
			if req.User.Role == "buyer" {
				req.IsBuyer = true
			}
		}

	}
	return req
}

// Validate Request
func (req *Request) ValidateRequest(errors validation.Errors) *Request {
	req.ValidationError = errors.Filter()
	return req
}

// Return Bad Request response if there is a validation error in request
func (req *Request) FailsValidation() bool {
	if req.ValidationError != nil {
		req.BadRequest(req.ValidationError)
		return true
	}
	return false
}
