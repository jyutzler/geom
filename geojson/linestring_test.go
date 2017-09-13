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

const (
	LS2   = `{"type":"LineString","coordinates":[[10,20],[30,40]]}`
	LS2BB = `{"type":"LineString","coordinates":[[10,20],[30,40]],"bbox":[10,20,30,40]}`
)

func TestFromLineString(t *testing.T) {
	var (
		lineString core.LineString
		text       string
		err        error
		options    FromOptions
	)
	lineString.Points = append(lineString.Points, core.Point{X: 10, Y: 20}, core.Point{X: 30, Y: 40})
	if text, err = FromLineString(lineString, options); err == nil {
		if text != LS2 {
			t.Errorf("Expected %v, received %v.", LS2, text)
		}
	} else {
		t.Error(err)
	}

	options.BBox = true
	if text, err = FromLineString(lineString, options); err == nil {
		if text != LS2BB {
			t.Errorf("Expected %v, received %v.", LS2BB, text)
		}
	} else {
		t.Error(err)
	}
}

func TestToLineString(t *testing.T) {
	var (
		err        error
		lineString geom.LineString
		bytes      []byte
		x, y       float64
	)
	if bytes, err = ioutil.ReadFile("test/linestring.geojson"); err != nil {
		t.Error(err)
	}
	if lineString, err = ToLineString(bytes); err != nil {
		t.Error(err)
	}
	if len(lineString.SubPoints()) != 2 {
		t.Errorf("Expected 2 points, received %v.", len(lineString.SubPoints()))
	}
	point := lineString.SubPoints()[0]
	x, y = point.XY()
	if x != 100 {
		t.Errorf("Expected x=100, received %v", x)
	}
	if y != 0 {
		t.Errorf("Expected y=0, received %v", y)
	}
}
