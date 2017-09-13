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

// LineString is a line of two or more points
type LineString struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
	BBox        BoundingBox `json:"bbox,omitempty"`
}

// ToLineString returns a LineString for the GeoJSON input
func ToLineString(bytes []byte) (geom.LineString, error) {
	var (
		result       geom.LineString
		gjLineString LineString
		lineString   core.LineString
		err          error
	)
	if err = json.Unmarshal(bytes, &gjLineString); err == nil {

		for _, coords := range gjLineString.Coordinates {
			lineString.Points = append(lineString.Points, core.Point{X: coords[0], Y: coords[1]})
		}
		result = lineString
	}
	return result, err
}

// FromLineString returns GeoJSON for the input LineString
func FromLineString(input geom.LineString, options FromOptions) (string, error) {
	var (
		result      string
		err         error
		lineString  LineString
		coordinates [][]float64
		bytes       []byte
	)
	for _, point := range input.SubPoints() {
		x, y := point.XY()
		coordinates = append(coordinates, []float64{x, y})
	}
	lineString = LineString{Type: LINESTRING, Coordinates: coordinates}
	if options.BBox {
		x1, y1, x2, y2 := input.BBox()
		lineString.BBox = BoundingBox{x1, y1, x2, y2}
	}
	if bytes, err = json.Marshal(lineString); err == nil {
		result = string(bytes)
	}
	return result, err
}
