package context

import (
	"log"
	"reflect"
)

type ApplicationContext struct {
	// 등록된 Bean 정의
	definitions map[reflect.Type]BeanDefinition
	// 이미 생성된 Bean 캐시
	singletons map[reflect.Type]any
	// 순환 의존성 감지용 플래그
	creating map[reflect.Type]bool
}

func NewApplicationContext() *ApplicationContext {
	return &ApplicationContext{
		definitions: map[reflect.Type]BeanDefinition{},
		singletons:  map[reflect.Type]any{},
		creating:    map[reflect.Type]bool{},
	}
}

/*
Bean 등록 메서드
Bean의 '정의'만 등록
객체를 여기서 만들지 않습니다.
*/
func (ctx *ApplicationContext) RegisterBean(beanType reflect.Type, factory func(ctx *ApplicationContext) any) {
	/*
		Bean 정의에 Bean Type을 Key로 BeanDefinition을 삽입한다.

		내부에는 BeanType과 의존성 주입 함수 팩토리 함수를 넣는다.
	*/
	ctx.definitions[beanType] = BeanDefinition{
		BeanType: beanType,
		Factory:  factory,
	}
}

/*
Bean 획득 메서드
GetBean의 동작 규칙을 지켜서 구현해야 합니다.

1. 이미 생성된 Bean이면 -> 캐시를 반환합니다.
2. 지금 생성중이면 순환 의존성이므로, 즉시 실패합니다.
3. 정의가 없으면 실패합니다.
4. Factory를 실행하고 Bean을 반환받습니다.
5. Bean 결과를 Singleton 캐시에 저장합니다.
6. 생성된 Bean을 반환합니다.
*/
func (ctx *ApplicationContext) GetBean(beanType reflect.Type) any {
	// BeanType이 일치하면 Application Context에서 Bean을 반환합니다.
	if bean, ok := ctx.singletons[beanType]; ok {
		return bean
	}

	// Bean이 생성중이면 오류
	if ctx.creating[beanType] {
		panic("순환 의존성 오류" + beanType.String())
	}

	definition, ok := ctx.definitions[beanType]
	if !ok {
		panic("빈 정의가 없습니다. " + beanType.String())
	}

	// Application Context에 Bean이 생성중이라고 표시
	ctx.creating[beanType] = true
	// Factory를 실행하고 Bean을 반환 받음
	bean := definition.Factory(ctx)
	// Application Context에 Bean 생성이 끝났다고 표시
	ctx.creating[beanType] = false

	// Application Context에 Bean 최종 등록
	log.Printf("[Application Context] Bean 등록 %s", beanType.String())
	ctx.singletons[beanType] = bean
	return bean
}
