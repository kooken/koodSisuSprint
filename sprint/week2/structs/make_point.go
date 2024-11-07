package sprint

// type Point struct {
// 	X    float32
// 	Y    float32
// 	Text string
// }

func MakePoint(x, y float32, text string) Point {
	result := Point{X: x, Y: y, Text: text}
	return result
}
