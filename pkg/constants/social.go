package constants

//1.未处理 2已处理 3拒绝
type HandlerResult int

const (
	NoHandlerResult HandlerResult = iota + 1
	PassHandlerResult
	RefuseHandlerResult
	CancelHandlerResult
)
