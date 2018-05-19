/*
* SensorWalkSkill implements a skill
* that makes the HEXA to walk forward and change direction
* when encountering obstacles
 */

// 最简单的走路，direction duration

package Hexa

import (

	// "time"

	"mind/core/framework/drivers/distance"
	"mind/core/framework/log"
)

func (d *Hexa) distance() float64 {
	distance, err := distance.Value()
	if err != nil {
		log.Error.Println(err)
	}
	return distance
}
