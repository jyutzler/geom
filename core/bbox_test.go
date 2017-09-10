package core

import (
	"testing"

	"github.com/go-spatial/geom"
)

func TestBBox(t *testing.T) {
	var (
		bbox geom.BoundingBox
	)
	bbox = BoundingBox{coordinates: [4]float64{10, 20, 30, 40}}
	bbox.BBox()
}
