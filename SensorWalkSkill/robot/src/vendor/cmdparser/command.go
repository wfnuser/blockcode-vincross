package cmdparser

import (
	"mind/core/framework/log"

	simple "github.com/bitly/go-simplejson"
	"github.com/st0012/metago"
)

type Command struct{}

func (c *Command) Repeat(paras []int, content *simple.Json) {
	for i := 0; i < paras[0]; i = i + 1 {
		ParseJson(content)
	}
}

func (c *Command) Left(paras []int, content *simple.Json) {
	log.Debug.Printf("Left: %d", paras[0])
}

func (c *Command) Right(paras []int, content *simple.Json) {
	log.Debug.Printf("Right: %d", paras[0])
}

func (c *Command) Forward(paras []int, content *simple.Json) {
	log.Debug.Printf("Forward: %d", paras[0])
}

func (c *Command) Send(methodName string, args ...interface{}) interface{} {
	return metago.CallFunc(c, methodName, args...)
}
