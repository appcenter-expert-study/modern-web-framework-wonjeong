package bootstrap

import (
	"reflect"

	"github.com/appcenter-expert-study/modern-web-framework-wonjeong/internal/context"
	"github.com/appcenter-expert-study/modern-web-framework-wonjeong/internal/dispatcher"
	"github.com/appcenter-expert-study/modern-web-framework-wonjeong/internal/http"
)

/*
bootstrap은 조립만 담당합니다.
요청 처리 로직은 여기에 포함하지 않습니다.
*/
func Run() {
	engine := http.NewEchoEngine()

	// ApplicationContext 생성 (IoC)
	ctx := context.NewApplicationContext()

	// ApplicationContext에 Dispatcher 등록
	ctx.RegisterBean(
		reflect.TypeFor[*dispatcher.Dispatcher](),
		func(ctx *context.ApplicationContext) any {
			return dispatcher.NewDispatcher()
		},
	)

	// ApplicationContext에서 Dispatcher 타입으로 Bean 가져오기
	dispatcher := ctx.GetBean(
		reflect.TypeFor[*dispatcher.Dispatcher](),
	).(*dispatcher.Dispatcher)

	// Dispatcher Bean을 Engine에 등록
	engine.RegisterDispatcher(dispatcher)
	engine.Start(":8080")
}
