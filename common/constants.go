// this package contain common elements
// all of the other packages can use all elements of this package
package common

const (
	RequestHeaderAccessToken = "Access-Token"
	ContextRequestId         = "RequestId"
	ContextIdentity          = "Identity"
)

// response code
const (
	SUCCESS = 0   //success
	ERROR   = 500 //error

	InvalidAccount = 400 //账号不合法
	Forbidden      = 403 //请求拒绝
	InvalidParam   = 405 //请求参数不合法
)
