package respcode

type RespCode struct {
	Code string
	Info string
}

var (
	RC_GENERAL_SUCC *RespCode = &RespCode{"RC00000", "Request completed"}
	RC_GENERAL_PARAM_ERR *RespCode = &RespCode{"RC50000", "Input parameters error."}
	RC_GENERAL_APP_ERR *RespCode = &RespCode{"RC80000", "Application error."}

	RC_GENERAL_SYS_ERR *RespCode = &RespCode{"RC90000", "System error."}
	RC_SERVICE_UNAVAILABLE *RespCode = &RespCode{"RC90001", "Service temperarily unavailable"}
	RC_UNSUPPORT_FUNCTION *RespCode = &RespCode{"RC90002", "Unsupport function"}
)