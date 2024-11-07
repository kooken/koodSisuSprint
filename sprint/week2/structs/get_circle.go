package sprint

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	circleRadius := r
	circleDiameter := (2 * r)
	circleArea := (3.14 * r * r)
	circlePerimeter := 2 * 3.14 * r
	result := Circle{Radius: circleRadius, Diameter: circleDiameter, Area: circleArea, Perimeter: circlePerimeter}
	return result
}
