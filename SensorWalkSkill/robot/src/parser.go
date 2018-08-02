package Hexa

import (
	"mind/core/framework/log"

	simple "github.com/bitly/go-simplejson"
)

func (d *Hexa) Parse(b []byte) error {
	js, err := simple.NewJson(b)
	log.Debug.Println("Parse")
	if err != nil {
		return err
	}
	return d.ParseJson(js)
}

func (d *Hexa) ParseJson(js *simple.Json) error {
	cmds, err := js.Array()
	if err != nil {
		return err
	}
	for i, _ := range cmds {
		if stop == false {
			stop = true
			break
		}
		cmd := js.GetIndex(i)
		name, err := cmd.Get("cmd").String()
		if err != nil {
			return err
		}
		paras, err := cmd.Get("params").Array()
		if err != nil {
			return err
		}
		var intParas []int
		for j, _ := range paras {
			para, err := cmd.Get("params").GetIndex(j).Int()
			if err != nil {
				return err
			}
			intParas = append(intParas, para)
		}
		content := cmd.Get("contents")
		d.Send(name, intParas, content)
	}
	return nil
}
