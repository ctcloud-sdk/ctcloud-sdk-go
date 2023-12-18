package sdkerr

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/request"
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/response"
)

type ErrorHandler interface {
	HandleError(req *request.DefaultHttpRequest, resp *response.DefaultHttpResponse) error
}

type DefaultErrorHandler struct {
}

func (h DefaultErrorHandler) HandleError(req *request.DefaultHttpRequest, resp *response.DefaultHttpResponse) error {
	if resp.GetStatusCode() < 400 {
		return nil
	}
	return NewServiceResponseError(resp.Response)
}
