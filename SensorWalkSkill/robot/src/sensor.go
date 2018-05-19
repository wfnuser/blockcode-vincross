/*
* SensorWalkSkill implements a skill
* that makes the HEXA to walk forward and change direction
* when encountering obstacles
 */

// 最简单的走路，direction duration

package Hexa

import (
	"image"

	"mind/core/framework/drivers/distance"
	"mind/core/framework/drivers/media"
	"mind/core/framework/log"
)

func (d *Hexa) distance() float64 {
	distance, err := distance.Value()
	if err != nil {
		log.Error.Println(err)
	}
	return distance
}

func calculateAvgRGB(img *image.RGBA) float64 {
	average := float64(0)
	for i := 0; i < len(img.Pix); i++ {
		average += float64(img.Pix[i])
	}
	average = average / float64(len(img.Pix))
	return float64(average)
}

func (d *Hexa) getAverageRGB() float64 {
	img := media.SnapshotRGBA()

	return calculateAvgRGB(img)
}
