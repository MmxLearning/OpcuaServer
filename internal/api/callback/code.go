package callback

const (
	cErrForm uint8 = iota + 1
	cErrDBOperation
	cErrUnauthorized
	cErrUsernameOrPasswordNotCorrect
	cErrUnexpected
	cErrNotExist
)

var (
	ErrForm = &Msg{
		Code:       cErrForm,
		Msg:        "error form",
		HttpStatus: 400,
	}
	ErrDBOperation = &Msg{
		Code:       cErrDBOperation,
		Msg:        "database operation failed",
		HttpStatus: 500,
	}
	ErrUnauthorized = &Msg{
		Code:       cErrUnauthorized,
		Msg:        "token not found or invalid",
		HttpStatus: 401,
	}
	ErrUsernameOrPasswordNotCorrect = &Msg{
		Code:       cErrUsernameOrPasswordNotCorrect,
		Msg:        "username or password not correct",
		HttpStatus: 401,
	}
	ErrUnexpected = &Msg{
		Code:       cErrUnexpected,
		Msg:        "unexpected error occurred",
		HttpStatus: 500,
	}
	ErrNotExist = &Msg{
		Code:       cErrNotExist,
		Msg:        "target not exist",
		HttpStatus: 404,
	}
)
