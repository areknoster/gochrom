package geom2d

type Mesh struct {
	Points PointsSet
	X, Y   int
}

func NewMesh(horizontal, vertical int) Mesh{
	m := Mesh{
		Points: make([]Point, horizontal * vertical),
		X:      horizontal,
		Y:      vertical,
	}
	for j := 0; j < vertical; j++{
		for i := 0; i < horizontal; i++{
			m.Points[m.getPointIndex(i,j)] = Point{
				X: float64(i)/float64(horizontal - 1),
				Y: float64(j)/float64(vertical - 1),
			}
		}
	}
	return m
}

func (m Mesh) GetVertex(x, y int) Point {
	return m.Points[m.getPointIndex(x,y)]
}

// returns -1 if does not belong to mesh
func (m Mesh) getPointIndex(x, y int) int {
	if x >= m.X || y >= m.Y {
		return -1
	}
	return x + y*m.X
}

func (m Mesh) GetSegments() EdgeSet {
	v, h := m.Y, m.X
	segments := make([]Edge, 3*v*h-2*(v+h)+1)
	s := 0
	for i := 0; i < m.X; i++ {
		for j := 0; j < m.Y; j++ {
			//add right
			if -1 != m.getPointIndex(i+1, j) {
				segments[s] = NewEdge(
					m.Points[m.getPointIndex(i, j)],
					m.Points[m.getPointIndex(i+1, j)])
				s++
			}
			//add up-right
			if -1 != m.getPointIndex(i+1, j+1) {
				segments[s] = NewEdge(
					m.Points[m.getPointIndex(i, j)],
					m.Points[m.getPointIndex(i+1, j+1)])
				s++
			}
			//add up
			if -1 != m.getPointIndex(i+1, j+1) {
				segments[s] = NewEdge(
					m.Points[m.getPointIndex(i, j)],
					m.Points[m.getPointIndex(i, j+1)])
				s++
			}
		}
	}
	return segments
}

func (m Mesh) GetTriangles() ShapeSet {
	triangles := make([]Shape, 2 * (m.X-1)*(m.Y -1))
	k:=0
	for j := 0 ; j < m.Y - 1; j++{
		for i := 0; i < m.X - 1; i++{
			triangles[k] = NewShape([]Point{
				m.GetVertex(i,j),
				m.GetVertex(i+1, j),
				m.GetVertex(i+1, j+1),
			})
			k++
			triangles[k] = NewShape([]Point{
				m.GetVertex(i,j),
				m.GetVertex(i, j+1),
				m.GetVertex(i+1, j+1),
			})
			k++
		}
	}
	return triangles
}


