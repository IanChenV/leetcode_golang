package structure

type Point struct {
	X, Y int
}

func Intss2Points(points [][]int) []Point {
	res := make([]Point, len(points))
	for i, p := range points {
		res[i] = Point{
			X: p[0],
			Y: p[1],
		}
	}
	return res
}

func Point2Intss(points []Point) [][]int {
	res := make([][]int, len(points))
	for i, p := range points {
		res[i] = []int{p.X, p.Y}
	}
	return res
}
