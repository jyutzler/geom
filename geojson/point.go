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
)

// The Point object contains a single position
type Point struct {
	Type        string      `json:"type"`
	Coordinates []float64   `json:"coordinates"`
	BBox        BoundingBox `json:"bbox,omitempty"`
}

// FromPointOptions contains the various options for the FromPoint function
type FromPointOptions struct {
	BBox bool
}

// // ToPoint returns a Point for the GeoJSON input
// func ToPoint(input string) geom.Point {
// 	var (
// 		result
// 		point Point
// 		err error
// 	)
// 	if err = json.Unmarshal(bytes, &point); err == nil {
//
// 	}
// }

// FromPoint returns GeoJSON for the input point
func FromPoint(input geom.Point, options FromPointOptions) (string, error) {
	var (
		result string
		point  Point
		x, y   float64
		err    error
		bytes  []byte
	)
	x, y = input.XY()
	point = Point{Type: POINT, Coordinates: []float64{x, y}}
	if options.BBox {
		point.BBox = []float64{point.Coordinates[0], point.Coordinates[1], point.Coordinates[0], point.Coordinates[1]}
	}
	if bytes, err = json.Marshal(point); err == nil {
		result = string(bytes)
	}
	return result, err
}
