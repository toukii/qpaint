package common

import (
	"encoding/json"
	"fmt"

	"github.com/toukii/bezier"
	"github.com/toukii/bezier/svg"
	"github.com/toukii/bytes"
	"github.com/toukii/goutils"
)

/*
[
	{
		"close": false,
		"points": [
			{
				"x": 45,
				"y": 75
			},
			{
				"x": 38,
				"y": 266
			},
			{
				"x": 145,
				"y": 322
			},
			{
				"x": 390,
				"y": 368
			},
			{
				"x": 461,
				"y": 434
			},
			{
				"x": 485,
				"y": 430
			},
			{
				"x": 485,
				"y": 430
			}
		],
		"style": {
			"fillColor": "white",
			"lineColor": "black",
			"lineWidth": 1
		}
	}
]
*/

type Shapes []*Shape
type Shape struct {
	Points []*bezier.Point `json:points`
	Style  *Style          `json:style`
}

type Style struct {
	FillColor string `json:fillColor`
	LineColor string `json:lineColor`
	LineWidth int    `json:lineWidth`
}

func BuildShapes(i interface{}) string {
	bs, err := json.Marshal(i)
	if err != nil {
		fmt.Errorf("json.Marshal err:%+v", err)
		return ""
	}
	var v Shapes
	err = json.Unmarshal(bs, &v)
	if err != nil {
		fmt.Errorf("json.Unmarshal err:%+v", err)
		return ""
	}

	buf := bytes.NewWriter(make([]byte, 0, 2048))
	for _, sp := range v {
		if len(sp.Points) <= 0 {
			continue
		}
		bs := svg.ExcutePath(sp.Style.LineColor, fmt.Sprint(sp.Style.LineWidth), sp.Points...)
		buf.Write(bs)
	}
	ret := goutils.ToString(buf.Bytes())
	fmt.Println(ret)
	return ret
}

type Point []struct {
	X int64 `json:x`
	Y int64 `json:y`
}

// type Points []*Point
type Points []*bezier.Point

func BezierPath(i interface{}) string {
	bs, err := json.Marshal(i)
	if err != nil {
		fmt.Errorf("json.Marshal err:%+v", err)
		return ""
	}
	var v Points
	err = json.Unmarshal(bs, &v)
	if err != nil {
		fmt.Errorf("json.Unmarshal err:%+v", err)
		return ""
	}
	return goutils.ToString(svg.ExcutePath("red", "2", v...))
}
