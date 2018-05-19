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
}

func (d *Hexa) Forward(paras []int, content *simple.Json) {
	log.Debug.Println("Forward")
	d.forward(paras[0])
}

func (d *Hexa) Dance(paras []int, content *simple.Json) {
	log.Debug.Println("Dance")
	d.dance(paras[0])
}

func (d *Hexa) Lightl(paras []int, content *simple.Json) {
	log.Debug.Println("Lightl")
	if d.lightL(float64(paras[0])) {
		d.ParseJson(content)
	}
}

func (d *Hexa) DistanceL(paras []int, content *simple.Json) {
	log.Debug.Println("DistanceL")
	if d.distanceL(float64(paras[0])) {
		d.ParseJson(content)
	}
}

func (d *Hexa) SelectGait(paras []int, content *simple.Json) {
	log.Debug.Println("Gait")
	d.ChangeGait(paras[0])
}

func (d *Hexa) Send(methodName string, args ...interface{}) interface{} {
	return metago.CallFunc(d, methodName, args...)
}
