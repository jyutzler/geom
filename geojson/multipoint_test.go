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
	MP2   = `{"type":"MultiPoint","coordinates":[[10,20],[30,40]]}`
	MP2BB = `{"type":"MultiPoint","coordinates":[[10,20],[30,40]],"bbox":[10,20,30,40]}`
)

func TestFromMultiPoint(t *testing.T) {
	var (
		multiPoint core.MultiPoint
		text       string
		err        error
		options    FromOptions
	)
	multiPoint.Points = append(multiPoint.Points, core.Point{X: 10, Y: 20}, core.Point{X: 30, Y: 40})
	if text, err = FromMultiPoint(multiPoint, options); err == nil {
		if text != MP2 {
			t.Errorf("Expected %v, received %v.", MP2, text)
		}
	} else {
		t.Error(err)
	}

	options.BBox = true
	if text, err = FromMultiPoint(multiPoint, options); err == nil {
		if text != MP2BB {
			t.Errorf("Expected %v, received %v.", MP2BB, text)
		}
	} else {
		t.Error(err)
	}
}

func TestToMultiPoint(t *testing.T) {
	var (
		err        error
		multiPoint geom.MultiPoint
		bytes      []byte
		x, y       float64
	)
	if bytes, err = ioutil.ReadFile("test/multipoint.geojson"); err != nil {
		t.Error(err)
	}
	if multiPoint, err = ToMultiPoint(bytes); err != nil {
		t.Error(err)
	}
	if len(multiPoint.SubPoints()) != 2 {
		t.Errorf("Expected 2 points, received %v.", len(multiPoint.SubPoints()))
	}
	point := multiPoint.SubPoints()[0]
	x, y = point.XY()
	if x != 100 {
		t.Errorf("Expected x=100, received %v", x)
	}
	if y != 0 {
		t.Errorf("Expected y=0, received %v", y)
	}
}
