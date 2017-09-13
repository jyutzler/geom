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
	"io/ioutil"
	"testing"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/core"
)

const POLYGON2 = `{"type":"Polygon","coordinates":[[[10,20],[30,40],[-10,-5],[10,20]]]}`
const POLYGON2B = `{"type":"Polygon","coordinates":[[[10,20],[30,40],[-10,-5],[10,20]]],"bbox":[-10,-5,30,40]}`

func TestFromPolygon(t *testing.T) {
	var (
		polygon geom.Polygon
		text    string
		err     error
		options FromOptions
	)
	polygon = core.Polygon{LineStrings: []geom.LineString{core.LineString{Points: []geom.Point{core.Point{X: 10, Y: 20},
		core.Point{X: 30, Y: 40},
		core.Point{X: -10, Y: -5},
		core.Point{X: 10, Y: 20}}}}}

	if text, err = FromPolygon(polygon, options); err == nil {
		if text != POLYGON2 {
			t.Errorf("Expected %v, received %v.", POLYGON2, text)
		}
	} else {
		t.Error(err)
	}

	options.BBox = true
	if text, err = FromPolygon(polygon, options); err == nil {
		if text != POLYGON2B {
			t.Errorf("Expected %v, received %v.", POLYGON2B, text)
		}
	} else {
		t.Error(err)
	}
}

func TestToPolygon(t *testing.T) {
	var (
		err     error
		polygon geom.Polygon
		bytes   []byte
		x, y    float64
	)
	if bytes, err = ioutil.ReadFile("test/polygon.geojson"); err != nil {
		t.Error(err)
	}
	if polygon, err = ToPolygon(bytes); err != nil {
		t.Error(err)
	}
	x, y = polygon.SubLineStrings()[0].SubPoints()[0].XY()
	if x != 100 {
		t.Errorf("Expected x=100, received %v", x)
	}
	if y != 0 {
		t.Errorf("Expected y=0, received %v", y)
	}
	if bytes, err = ioutil.ReadFile("test/polygon-hole.geojson"); err != nil {
		t.Error(err)
	}
	if polygon, err = ToPolygon(bytes); err != nil {
		t.Error(err)
	}
	x, y = polygon.SubLineStrings()[1].SubPoints()[4].XY()
	if x != 100.8 {
		t.Errorf("Expected x=100, received %v", x)
	}
	if y != 0.8 {
		t.Errorf("Expected y=0, received %v", y)
	}
}
