package errorLog

import "net/http"

type LevelErr string

const LevelInfo LevelErr = "info"
const LevelWarning LevelErr = "warning"
const LevelError LevelErr = "error"
const LevelFatal LevelErr = "fatal"

type LayerErr string

const Domain LayerErr = "domain"
const Infra LayerErr = "infrastructure"
const Interface LayerErr = "interface"
const UseCase LayerErr = "use_case"

const ErrUnknownCode = "ERR_UNKNOWN"

type CustomError struct {
	error
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Messages    []string `json:"messages"`
	Level       LevelErr `json:"level"`
	HttpStatus  int      `json:"httpStatus"`
	Raw         string   `json:"raw"`
	Layer       LayerErr `json:"layer"`
	Function    string   `json:"function"`
}

func NewErr() *CustomError {
	errTmp := ErrUnknown
	return &errTmp
}

var (
	ErrUnknown = CustomError{
		Code:        ErrUnknownCode,
		Description: "UNKNOWN ERROR",
		Level:       LevelError,
		HttpStatus:  http.StatusInternalServerError,
	}
)

func (e *CustomError) Clone() *CustomError {
	return &CustomError{
		Code:        e.Code,
		Description: e.Description,
		Messages:    e.Messages,
		Level:       e.Level,
		HttpStatus:  e.HttpStatus,
		Raw:         e.Raw,
		Layer:       e.Layer,
		Function:    e.Function,
	}
}

func (e *CustomError) CopyCodeDescription(source *CustomError) *CustomError {
	e.Code = source.Code
	e.Description = source.Description
	return e
}

func (e *CustomError) SetCode(code string) *CustomError {
	e.Code = code
	return e
}

func (e *CustomError) SetDescription(description string) *CustomError {
	e.Description = description
	return e
}

func (e *CustomError) SetMessages(messages []string) *CustomError {
	e.Messages = messages
	return e
}

func (e *CustomError) SetLayer(layer LayerErr) *CustomError {
	e.Layer = layer
	return e
}

func (e *CustomError) SetLevel(level LevelErr) *CustomError {
	e.Level = level
	return e
}

func (e *CustomError) SetHttpStatus(httpStatus int) *CustomError {
	e.HttpStatus = httpStatus
	return e
}

func (e *CustomError) SetFunction(function string) *CustomError {
	e.Function = function
	return e
}

func (e *CustomError) SetRaw(err error) *CustomError {
	raw := ""
	if err != nil {
		raw = err.Error()
	}
	e.Raw = raw
	return e
}
