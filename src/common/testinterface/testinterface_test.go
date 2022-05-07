package testinterface

import (
	"testing"
)

func TestNewRedShapeDecorator(t *testing.T) {
	type args struct {
		decShape Shape
	}
	tests := []struct {
		name string
		args args
		want *RedShapeDecorator
	}{
		{name: "装饰器", args: args{NewCircle()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRedShapeDecorator(tt.args.decShape)
			got.Draw()
		})
	}
}
