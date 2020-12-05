package geom2d

import (
	"fmt"
	"testing"
)

func TestMesh(t *testing.T){
	mesh := NewMesh(3,5)
	for _, point := range mesh.Points {
		fmt.Println(point)
	}
}
