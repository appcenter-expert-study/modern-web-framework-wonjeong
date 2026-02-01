package http

import "github.com/labstack/echo/v4"

// Engine과 Dispatcher 사이의 상호계약
type Handler interface {
	// Handler 타입과 호환되기 위해서는 Dispatch 메서드를 구현해야 함
	Dispatch(ctx *RequestContext) error
}

type EchoEngine struct {
	instance *echo.Echo
}

func NewEchoEngine() *EchoEngine {
	// 새로운 echo 인스턴스 생성
	e := echo.New()
	// EchoEngine 구조체로 반환 (한번 추상화)
	return &EchoEngine{instance: e}
}

// Echo에는 반드시 /* 단 하나만 등록, HTTP Method 구분 X, 전부 Dispatcher로 위임
func (engine *EchoEngine) RegisterDispatcher(dispatcher Handler) {
	engine.instance.Any("/*", func(context echo.Context) error {
		ctx := NewRequestContext(context)
		return dispatcher.Dispatch(ctx)
	})
}

func (engine *EchoEngine) Start(address string) error {
	return engine.instance.Start(address)
}
