package main

//color:"heart",
//num:"4",
//x:"50px",//x坐标
//y:"100px",//y坐标
//提供给前端渲染扑克牌所用的数据结构
type Poker struct {
	Color string `json:"color"`
	Num string `json:"num"`
	X string `json:"x"`
	Y string `json:"y"`
}
