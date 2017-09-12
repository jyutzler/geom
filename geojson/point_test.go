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

const POINT2 = `{"type":"Point","coordinates":[10,20]}`
const POINT2B = `{"type":"Point","coordinates":[10,20],"bbox":[10,20,10,20]}`

func TestFromPoint(t *testing.T) {
	var (
		point geom.Point
		text  string
		err   error
		fpo   FromPointOptions
	)
	point = core.Point{X: 10, Y: 20}

	if text, err = FromPoint(point, fpo); err == nil {
		if text != POINT2 {
			t.Errorf("Expected %v, received %v.", POINT2, text)
		}
	} else {
		t.Error(err)
	}

	fpo.BBox = true
	if text, err = FromPoint(point, fpo); err == nil {
		if text != POINT2B {
			t.Errorf("Expected %v, received %v.", POINT2B, text)
		}
	} else {
		t.Error(err)
	}
}

func TestToPoint(t *testing.T) {
	var (
		err   error
		point geom.Point
		bytes []byte
		x, y  float64
	)
	if bytes, err = ioutil.ReadFile("test/point.geojson"); err != nil {
		t.Error(err)
	}
	if point, err = ToPoint(bytes); err != nil {
		t.Error(err)
	}
	x, y = point.XY()
	if x != 100 {
		t.Errorf("Expected x=100, received %v", x)
	}
	if y != 0 {
		t.Errorf("Expected y=0, received %v", y)
	}
	if bytes, err = ioutil.ReadFile("test/point3.geojson"); err != nil {
		t.Error(err)
	}
	if point, err = ToPoint(bytes); err == nil {
		t.Error("Expected failure due to too many coordinates, but it succeeded.")
	}
}
