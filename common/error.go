package common

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"net/http"
)

type ApiError struct {
	errCode int
	err     error
	realErr error
}

func (a ApiError) Error() string {
	return a.err.Error() + ", realErr: " + a.realErr.Error()
}

func (a *ApiError) RealErr(realErr error) *ApiError {
	a.realErr = realErr
	return a
}

func newApiError(errCode int, err error) *ApiError {
	return &ApiError{
		errCode: errCode,
		err:     err,
	}
}

// 定义所有错误
var (
	ErrSystem         = newApiError(ERROR, errors.New("system error, please contact administrator"))
	ErrInvalidAccount = newApiError(InvalidAccount, errors.New("invalid account"))
	ErrForbidden      = newApiError(Forbidden, errors.New("request forbidden"))
	ErrInvalidParam   = newApiError(InvalidParam, errors.New("invalid param"))
)

func ErrorResult(err error) (int, *Result) {
	switch err.(type) {
	case *ApiError:
		apiError := err.(*ApiError)
		glog.Errorf("ApiError: %v", apiError)
		return http.StatusOK, buildResultByApiError(apiError)
	default:
		glog.Errorf("not ApiError, maybe other error, log stack: %+v", err)
		return http.StatusOK, buildResultByApiError(ErrSystem)
	}
}

func buildResultByApiError(apiError *ApiError) *Result {
	return &Result{
		Code: apiError.errCode,
		Msg:  apiError.err.Error(),
		Data: nil,
	}
}
