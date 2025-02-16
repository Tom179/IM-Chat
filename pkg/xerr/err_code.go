package xerr

//业务错误码定义
const ( //业务错误码，区分不同类型的错误，并且根据错误码获取错误信息：只需要在业务logic层做，而不需要所有错误都使用。
	SERVER_COMMON_ERROR = 100001
	REQUEST_PARAMERROR  = 100002
	DB_ERROR            = 100003
)
