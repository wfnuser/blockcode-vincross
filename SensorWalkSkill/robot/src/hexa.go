/*
* SensorWalkSkill implements a skill
* that makes the HEXA to walk forward and change direction
* when encountering obstacles
 */

// 最简单的走路，direction duration

package Hexa

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"math"
	"os"
	// "time"
	"mind/core/framework"
	"mind/core/framework/drivers/distance"
	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/drivers/media"
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

const (
	TIME_TO_NEXT_REACTION = 2000 // milliseconds
	DISTANCE_TO_REACTION  = 250  // millimeters
	MOVE_HEAD_DURATION    = 500  // milliseconds
	ROTATE_DEGREES        = 130  // degrees out of 360
	WALK_SPEED            = 1.0  // cm per second   走路
	SENSE_INTERVAL        = 250  // four times per second

	FOREVER = -1
)

func newDirection(direction float64) float64 {
	return math.Mod(direction+ROTATE_DEGREES, 360) * -1
}

type Hexa struct {
	skill.Base
	stop       chan bool
	isRunnning bool
	dir        chan bool
	direction  float64
}

func NewSkill() skill.Interface {
	return &Hexa{
		stop: make(chan bool),
	}
}

func (d *Hexa) left(degree float64) {
	if d.direction+degree > 360 {
		d.direction = d.direction + degree - 360
	} else {
		d.direction = d.direction + degree
	}
	hexabody.MoveHead(d.direction, MOVE_HEAD_DURATION)
}

func (d *Hexa) right(degree float64) {
	if d.direction-degree < 0 {
		d.direction = d.direction - degree + 360
	} else {
		d.direction = d.direction - degree
	}
	hexabody.MoveHead(d.direction, MOVE_HEAD_DURATION)
}

// 奇怪，walk好像并不能指定速度
func (d *Hexa) forward(duration int) {
	if duration == FOREVER {
		hexabody.WalkContinuously(0, MOVE_HEAD_DURATION)
	} else {
		for i := 0; i < duration; i = i + 1 {
			hexabody.Walk(d.direction, 100)
		}
	}
}
func (d *Hexa) changeGait(gait hexabody.GaitType) {
	if d.isRunnning == true {
		hexabody.StopWalkingContinuously()
	}
	hexabody.SelectGait(gait)
	// hexabody.WalkContinuously(d.direction(), WALK_SPEED)
}

func (d *Hexa) ChangeGait(info int) {
	switch info {
	case 1:
		d.changeGait(hexabody.GaitWave)
	case 2:
		d.changeGait(hexabody.GaitRipple)
	case 3:
		d.changeGait(hexabody.GaitTripod)
	case 0:
		d.changeGait(hexabody.GaitOriginal)
	}
}

func (d *Hexa) circle(info int) {
	var leg *hexabody.LegPosition
	if info == 1 {
		leg = hexabody.NewLegPosition().SetCoordinates(-100, 50.0, 70.0)
	} else {
		leg = hexabody.NewLegPosition().SetCoordinates(100, 50.0, 70.0)		
	}
	for j := 0; j < 20; j++ {
		for i := 1; i < 6; i++ {
			hexabody.MoveLeg(i, leg, FAST_DURATION*3)
		}
		if (j % 2 == 0) {
			leg = hexabody.NewLegPosition().SetCoordinates(0, 0, 70.0)
		} else {
			if info == 1 {
				leg = hexabody.NewLegPosition().SetCoordinates(-100, 50.0, 70.0)			
			} else {
				leg = hexabody.NewLegPosition().SetCoordinates(100, 50.0, 70.0)
			}
		}

	}
}

func (d *Hexa) shouldChangeDirection() bool {
	return d.distance() < DISTANCE_TO_REACTION
}

func (d *Hexa) walk() {
	// hexabody.WalkContinuously(0, WALK_SPEED)
	for {
		select {
		case <-d.stop:
			return
		// case <- d.dir:
		// 	d.TurnLeft(90)
		default:
			// if d.shouldChangeDirection() {
			// 	d.changeDirection()
			// }
			// time.Sleep(SENSE_INTERVAL * time.Millisecond)
			// d.right(90.0, 1000)
		}
	}
}

func (d *Hexa) OnStart() {
	err := hexabody.Start()
	if err != nil {
		log.Error.Println("Hexabody start err:", err)
		return
	}
	err = distance.Start()
	if err != nil {
		log.Error.Println("Distance start err:", err)
		return
	}
	if !distance.Available() {
		log.Error.Println("Distance sensor is not available")
	}
	if !media.Available() {
		log.Error.Println("Media driver not available")
		return
	}
	if err = media.Start(); err != nil {
		log.Error.Println("Media driver could not start")
	}
}

func (d *Hexa) OnConnect() {
	err := media.Start()
	if err != nil {
		log.Error.Println("Media start err:", err)
		return
	}
	for {
		log.Info.Println("Connected")
		buf := new(bytes.Buffer)
		log.Info.Println("JPEG")
		jpeg.Encode(buf, media.SnapshotYCbCr(), nil)
		log.Info.Println("BASE64")
		str := base64.StdEncoding.EncodeToString(buf.Bytes())
		log.Info.Println("SENDING")
		framework.SendString(str)
		log.Info.Println("Sent:", str[:20], len(str))
	}
}

func (d *Hexa) OnClose() {
	hexabody.Close()
	distance.Close()
}

func (d *Hexa) OnDisconnect() {
	os.Exit(0) // Closes the process when remote disconnects
}

func (d *Hexa) OnRecvString(data string) {
	d.direction = hexabody.Direction()
	b := []byte(data)
	err := d.Parse(b)
	log.Debug.Println(err)
	// switch data {
	// case "start":
	// 	// d.dir <- true
	// 	go d.walk()
	// case "stop":
	// 	d.stop <- true
	// 	hexabody.StopWalkingContinuously()
	// 	hexabody.Relax()
	// }
	// for {
	// 	// log.Debug.Println(d.distance())
	// 	log.Debug.Println(d.getAverageRGB())
	// 	time.Sleep(SENSE_INTERVAL * time.Millisecond)
	// }
	log.Debug.Println(data)
}
