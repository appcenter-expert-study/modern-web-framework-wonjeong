package http

import "github.com/labstack/echo/v4"

// Echo Context를 감싼 우리 프레임워크의 요청 객체
type RequestContext struct {
	Method string
	Path   string
}

// Echo Context를 외부로 노출하지 말 것, 최소한의 Method, Path만 포함
func NewRequestContext(c echo.Context) *RequestContext {
	return &RequestContext{
		Method: c.Request().Method,
		Path:   c.Request().URL.Path,
	}
}
