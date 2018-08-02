/*
* SensorWalkSkill implements a skill
* that makes the HEXA to walk forward and change direction
* when encountering obstacles
 */

// 最简单的走路，direction duration

package xereb

import (
	"os"
	// "time"
	"mind/core/framework/drivers/xerebody"
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

var stop bool

const (
	TIME_TO_NEXT_REACTION = 2000 // milliseconds
)

type xereb struct {
	skill.Base
}

func NewSkill() skill.Interface {
	return &xereb{}
}

func (d *xereb) left(degree float64) {
	xerebody.Move(0, 1)
}

func (d *xereb) right(degree float64) {
	xerebody.Move(0, -1)
}

func (d *xereb) forward(duration int) {
	xerebody.Move(1, 1)
}

func (d *xereb) OnStart() {
	log.Info.Println("xerebody boot!")
	stop = false
	err := xerebody.Start()
	if err != nil {
		log.Error.Println("xerebody start err:", err)
		return
	}
}

func (d *xereb) OnConnect() {
}

func (d *xereb) OnClose() {
	xerebody.Close()
}

func (d *xereb) OnDisconnect() {
	os.Exit(0) // Closes the process when remote disconnects
}

func (d *xereb) OnRecvString(data string) {
	if data == "stop" {
		log.Info.Println("Receive stop!")
		stop = true
		return
	} else {
		stop = false
	}
	b := []byte(data)
	err := d.Parse(b)
	log.Debug.Println(err)
}
