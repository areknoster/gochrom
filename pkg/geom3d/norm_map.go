package geom3d

type VectorMap struct {
	Vecs          []Vector
	Width, Height int
}

func NewVectorMap(width, height int) *VectorMap {
	return &VectorMap{
		Vecs:   make([]Vector, width*height),
		Width:  width,
		Height: height,
	}
}

func (vm *VectorMap) At(x, y int) Vector {
	if x >= vm.Width || y >= vm.Height {
		return WrongVector
	}
	return vm.Vecs[x+y*vm.Width]
}

func (vm *VectorMap) SetAt(x, y int, v Vector) {
	if x >= vm.Width || y >= vm.Height {
		return
	}
	vm.Vecs[x+y*vm.Width] = v
}
