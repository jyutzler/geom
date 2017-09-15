// Package core contains basic implementations of structs that implement geometry interfaces.
package core

import (
	"github.com/go-spatial/geom"
)

// BoundingBox contains x1, y1, x2, y2
type BoundingBox struct {
	coordinates [4]float64
}

// BBox returns x1, y1, x2, y2
func (bbox BoundingBox) BBox() (float64, float64, float64, float64) {
	return bbox.coordinates[0],
		bbox.coordinates[1],
		bbox.coordinates[2],
		bbox.coordinates[3]
}

// MakeBBox creates something implementing the geom.BoundingBox interface from the input
func MakeBBox(input [][2]float64) geom.BoundingBox {
	var (
		x1, x2, y1, y2 float64
	)

	for inx, coords := range input {
		if inx == 0 {
			x1 = coords[0]
			y1 = coords[1]
			x2 = coords[0]
			y2 = coords[1]
			continue
		}
		if coords[0] < x1 {
			x1 = coords[0]
		}
		if coords[0] > x2 {
			x2 = coords[0]
		}
		if coords[1] < y1 {
			y1 = coords[1]
		}
		if coords[1] > y2 {
			y2 = coords[1]
		}
	}
	return BoundingBox{coordinates: [4]float64{x1, y1, x2, y2}}
}
