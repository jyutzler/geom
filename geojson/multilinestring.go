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

// MultiLineString is a geometry with multiple LineStrings.
type MultiLineString struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
	BBox        BoundingBox   `json:"bbox,omitempty"`
}

// ToMultiLineString returns a MultiLineString for the GeoJSON input
func ToMultiLineString(bytes []byte) (geom.MultiLineString, error) {
	var (
		result      geom.MultiLineString
		gjMLS       MultiLineString
		err         error
		points      []geom.Point
		lineStrings []geom.LineString
	)
	if err = json.Unmarshal(bytes, &gjMLS); err == nil {

		for _, coords2 := range gjMLS.Coordinates {
			points = nil
			for _, coords := range coords2 {
				points = append(points, core.Point{X: coords[0], Y: coords[1]})
			}
			lineStrings = append(lineStrings, core.LineString{Points: points})
		}
		result = core.MultiLineString{LineStrings: lineStrings}
	}
	return result, err
}

// FromMultiLineString returns GeoJSON for the input polygon
func FromMultiLineString(input geom.Polygon, options FromOptions) (string, error) {
	var (
		result       string
		err          error
		mls          MultiLineString
		coordinates  [][][]float64
		coordinates2 [][]float64
		bytes        []byte
	)
	for _, lineString := range input.SubLineStrings() {
		coordinates2 = nil
		for _, point := range lineString.SubPoints() {
			x, y := point.XY()
			coordinates2 = append(coordinates2, []float64{x, y})
		}
		coordinates = append(coordinates, coordinates2)
	}
	mls = MultiLineString{Type: MULTILINESTRING, Coordinates: coordinates}
	if options.BBox {
		x1, y1, x2, y2 := input.BBox()
		mls.BBox = BoundingBox{x1, y1, x2, y2}
	}
	if bytes, err = json.Marshal(mls); err == nil {
		result = string(bytes)
	}
	return result, err
}
