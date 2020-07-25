package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参失败")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "授权失败,AppKey和AppSecret不存在")
	UnauthorizedTokenError    = NewError(10000004, "授权失败,Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "授权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "授权失败,Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
)
