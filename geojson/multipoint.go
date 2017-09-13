/*
Copyright 2016, RadiantBlue Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package geojson

import (
	"encoding/json"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/core"
)

// MultiPoint is a geometry with multiple points
type MultiPoint struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
	BBox        BoundingBox `json:"bbox,omitempty"`
}

// ToMultiPoint returns a MultiPoint for the GeoJSON input
func ToMultiPoint(bytes []byte) (geom.MultiPoint, error) {
	var (
		result       geom.MultiPoint
		gjMultiPoint MultiPoint
		points       []geom.Point
		err          error
	)
	if err = json.Unmarshal(bytes, &gjMultiPoint); err == nil {

		for _, coords := range gjMultiPoint.Coordinates {
			points = append(points, core.Point{X: coords[0], Y: coords[1]})
		}
		result = core.MultiPoint{Points: points}
	}
	return result, err
}

// FromMultiPoint returns GeoJSON for the input point
func FromMultiPoint(input geom.MultiPoint, options FromOptions) (string, error) {
	var (
		result      string
		err         error
		multiPoint  MultiPoint
		coordinates [][]float64
		bytes       []byte
	)
	for _, point := range input.SubPoints() {
		x, y := point.XY()
		coordinates = append(coordinates, []float64{x, y})
	}
	multiPoint = MultiPoint{Type: MULTIPOINT, Coordinates: coordinates}
	if options.BBox {
		x1, y1, x2, y2 := input.BBox()
		multiPoint.BBox = BoundingBox{x1, y1, x2, y2}
	}
	if bytes, err = json.Marshal(multiPoint); err == nil {
		result = string(bytes)
	}
	return result, err
}
