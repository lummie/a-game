package scene

import (
	"math"
	"fmt"
)

var (
//	APPROX_LIMIT float64 = 	0.000000000000001 // minimum
	APPROX_LIMIT float64 = 	0.00000000001
)

// Vertex represents a point in 3d space
type Vertex struct {
	X, Y, Z float64
}

// Vertices holds and owns a list of Vertex
// The order of items in the Vertices array has no relevance
type Vertices []Vertex
type VerticesRef []*Vertex



// Length returns the length of the vector, similar to pythagoras theorem but in 3 dimensional space
func (v *Vertex) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// LengthSquared can be used to compare the lengths of vectors without the additional (expensive) computation of the square root
func (v *Vertex) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

//CrossProduct returns the vector that is perpendicular to the vector and the supplied other vector
func (v *Vertex) CrossProduct(other Vertex) Vertex{
	x := v.Y*other.Z - v.Z*other.Y
	y := v.Z*other.X - v.X*other.Z
	z := v.X*other.Y - v.Y*other.X
	return Vertex{x, y, z}
}

//DotProduct Directionally multiplies the two vectors. If can be used to establish how one vector influences another's direction.
// A return value of zero indicates there is no change in direction.
// A positive value indicates there is an increase in direction.
// A negative value indicates there is a decrease in direction.
func (v *Vertex) DotProduct(other Vertex) float64 {
	return v.X * other.X + v.Y * other.Y + v.Z * other.Z
}

// Normalize returns a resized vector within a unit of 1, keeping the direction of the original vector
func (v *Vertex) Normalize() Vertex{
	l := v.Length()
	if l == 0 {
		return Vertex{}
	}
	return Vertex{v.X / l, v.Y / l, v.Z / l}
}

// Add returns the addition of the two vectors
func (v *Vertex) Add(o Vertex) Vertex{
	return Vertex{v.X + o.X, v.Y + o.Y, v.Z + o.Z }
}

// Subtract returns the subtraction of the two vectors
func (v *Vertex) Subtract(o Vertex) Vertex{
	return Vertex{v.X - o.X, v.Y - o.Y, v.Z - o.Z }
}

// Multiply returns the multiplication of the two vectors
func (v *Vertex) Multiply(o Vertex) Vertex{
	return Vertex{v.X * o.X, v.Y * o.Y, v.Z * o.Z }
}

// Divide returns the division of the two vectors
func (v *Vertex) Divide(o Vertex) Vertex{
	return Vertex{v.X / o.X, v.Y / o.Y, v.Z / o.Z }
}

// AddScalar adds the float value to each dimension or the vector
func (v *Vertex) AddScalar(f float64) Vertex{
	return Vertex{v.X + f, v.Y + f, v.Z + f }
}

// SubtractScalar subtracts the float value to each dimension or the vector
func (v *Vertex) SubtractScalar(f float64) Vertex{
	return Vertex{v.X - f, v.Y - f, v.Z - f }
}

// MultiplyScalar multiplies each dimension or the vector by the float value
func (v *Vertex) MultiplyScalar(f float64) Vertex{
	return Vertex{v.X * f, v.Y * f, v.Z * f }
}

// DivideScalar divides each dimension or the vector by the float value
func (v *Vertex) DivideScalar(f float64) Vertex{
	return Vertex{v.X / f, v.Y / f, v.Z / f }
}

// Equals returns true if the vectors are equal
func (v *Vertex) ApproxEqual(o Vertex) bool {
	d := v.Subtract(o)
	return (math.Abs(d.X) < APPROX_LIMIT)
}

// MidPointTo returns the vector that represents the point half way between the two vectors
func (v* Vertex) MidPointTo(o Vertex) Vertex{
	relative := o.Subtract(*v) // calculate o relative to v
	mp := relative.DivideScalar(2.0) // find half way point
	return mp.Add(*v) // translate back again
}

func (v *Vertex) String() string {
	return fmt.Sprintf("x:%v y:%v z:%v", v.X, v.Y, v.Z)
}