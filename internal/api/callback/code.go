package callback

const (
	cErrForm uint8 = iota + 1
	cErrDBOperation
	cErrUnauthorized
	cErrUsernameOrPasswordNotCorrect
	cErrUnexpected
)

var (
	ErrForm = &Msg{
		Code:       cErrForm,
		Msg:        "参数错误",
		HttpStatus: 400,
	}
	ErrDBOperation = &Msg{
		Code:       cErrDBOperation,
		Msg:        "数据库操作失败",
		HttpStatus: 500,
	}
	ErrUnauthorized = &Msg{
		Code:       cErrUnauthorized,
		Msg:        "身份校验失败，令牌失效或过期",
		HttpStatus: 401,
	}
	ErrUsernameOrPasswordNotCorrect = &Msg{
		Code:       cErrUsernameOrPasswordNotCorrect,
		Msg:        "账号或密码错误",
		HttpStatus: 401,
	}
	ErrUnexpected = &Msg{
		Code:       cErrUnexpected,
		Msg:        "预期外的错误",
		HttpStatus: 500,
	}
)
