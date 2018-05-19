package Hexa

import (
	"mind/core/framework/log"

	simple "github.com/bitly/go-simplejson"
	"github.com/st0012/metago"
)

func (d *Hexa) Repeat(paras []int, content *simple.Json) {
	for i := 0; i < paras[0]; i = i + 1 {
		d.ParseJson(content)
	}
}

func (d *Hexa) Left(paras []int, content *simple.Json) {
	log.Debug.Println("Left")
	d.left(float64(paras[0]))
}

func (d *Hexa) Right(paras []int, content *simple.Json) {
	log.Debug.Println("Right")
	d.right(float64(paras[0]))
	// log.Debug.Printf("Right: %d", paras[0])
}

func (d *Hexa) Forward(paras []int, content *simple.Json) {
	log.Debug.Println("Forward")
	d.forward(paras[0])
	// log.Debug.Printf("Forward: %d", paras[0])
}

func (d *Hexa) Forward(paras []int, content *simple.Json) {
	log.Debug.Println("Forward")
	d.dance(paras[0])
	// log.Debug.Printf("Forward: %d", paras[0])
}

func (d *Hexa) Send(methodName string, args ...interface{}) interface{} {
	return metago.CallFunc(d, methodName, args...)
}
