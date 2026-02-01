package dispatcher

import (
	"fmt"

	"github.com/appcenter-expert-study/modern-web-framework-wonjeong/internal/http"
)

type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

/*
현재는 무조건 고정 응답
예, "Dispatcher received {METHOD} {PATH}""
*/
func (d *Dispatcher) Dispatch(ctx *http.RequestContext) error {
	fmt.Printf(
		"Dispatcher received %s %s \n",
		ctx.Method,
		ctx.Path,
	)
	return nil
}
