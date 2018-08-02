package xereb

import (
	"mind/core/framework/log"

	simple "github.com/bitly/go-simplejson"
	"github.com/st0012/metago"
)

func (d *xereb) Repeat(paras []int, content *simple.Json) {
	for i := 0; i < paras[0]; i = i + 1 {
		d.ParseJson(content)
	}
}

func (d *xereb) Left(paras []int, content *simple.Json) {
	log.Debug.Println("Left")
	d.left(float64(paras[0]))
}

func (d *xereb) Right(paras []int, content *simple.Json) {
	log.Debug.Println("Right")
	d.right(float64(paras[0]))
}

func (d *xereb) Forward(paras []int, content *simple.Json) {
	log.Debug.Println("Forward")
	d.forward(paras[0])
}

func (d *xereb) Send(methodName string, args ...interface{}) interface{} {
	return metago.CallFunc(d, methodName, args...)
}
