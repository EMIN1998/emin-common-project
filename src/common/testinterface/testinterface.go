package testinterface

import "fmt"

type Animal interface {
	Roar()
}

func NewAnimal(name string) Animal {
	switch name {
	case "Cat":
		cat := newCat()
		return &cat
	default:
		return nil
	}
}

type cat struct {
}

func newCat() cat {
	return cat{}
}

func (c *cat) Roar() {
	fmt.Print("I am Cat ...")
}

type Dog struct {
	Name string
}

func (d *Dog) Roar() {
	fmt.Printf("I am Dog :%s...", d.Name)
}

//Shape 模型接口
type Shape interface {
	Draw()
}

//Circle 圆形类
type Circle struct{}

//NewCircle 实例化圆形
func NewCircle() *Circle {
	return &Circle{}
}

//Draw 输出方法，实现Shape接口
func (c *Circle) Draw() {
	fmt.Println("Circle Draw method.")
}

//RedShapeDecorator 红色装饰器
type RedShapeDecorator struct {
	DecoratedShape Shape
}

//NewRedShapeDecorator 实例化红色装饰器
func NewRedShapeDecorator(decShape Shape) *RedShapeDecorator {
	return &RedShapeDecorator{
		DecoratedShape: decShape,
	}
}

//SetRedBorder 装饰器方法
func (rs *RedShapeDecorator) SetRedBorder() {
	fmt.Println("Border Color: red")
}

//Draw 实现Shape接口的方法
func (rs *RedShapeDecorator) Draw() {
	rs.DecoratedShape.Draw()
	rs.SetRedBorder()
}
