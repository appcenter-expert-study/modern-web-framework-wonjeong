package context

import "reflect"

type BeanDefinition struct {
	BeanType reflect.Type
	Factory  func(ctx *ApplicationContext) interface{}
}
