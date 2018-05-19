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
	log.Debug.Println("Left")
	log.Debug.Println(paras[0])
	// log.Debug.Print("Left: %d", paras[0])
}

func (c *Command) Right(paras []int, content *simple.Json) {
	log.Debug.Println("Right")
	log.Debug.Println(paras[0])
	// log.Debug.Printf("Right: %d", paras[0])
}

func (c *Command) Forward(paras []int, content *simple.Json) {
	log.Debug.Println("Forward")
	log.Debug.Println(paras[0])
	// log.Debug.Printf("Forward: %d", paras[0])
}

func (c *Command) Send(methodName string, args ...interface{}) interface{} {
	return metago.CallFunc(c, methodName, args...)
}
