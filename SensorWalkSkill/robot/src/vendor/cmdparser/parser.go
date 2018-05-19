package cmdparser

import (
	simple "github.com/bitly/go-simplejson"
)

func Parse(b []byte) error {
	js, err := simple.NewJson(b)
	if err != nil {
		return err
	}
	ParseJson(js)
	return nil
}

func ParseJson(js *simple.Json) error {
	cmds, err := js.Array()
	if err != nil {
		return err
	}
	for i, _ := range cmds {
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
		c := &Command{}
		c.Send(name, intParas, content)
	}
	return nil
}
