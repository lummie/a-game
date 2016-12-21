package scene

import (
	"math"
	"fmt"
)

var (
//	APPROX_LIMIT float64 = 	0.000000000000001 // minimum
	APPROX_LIMIT float64 = 	0.00000000001
)

// Vector represents a point in 3d space
type Vector struct {
	X, Y, Z float64
}

// Vertices holds and owns a list of Vertex
// The order of items in the Vertices array has no relevance
type Vectors []Vector
type VectorsRef []*Vector



// Length returns the length of the vector, similar to pythagoras theorem but in 3 dimensional space
func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// LengthSquared can be used to compare the lengths of vectors without the additional (expensive) computation of the square root
func (v *Vector) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

//CrossProduct returns the vector that is perpendicular to the plane created by the two vectors
func (v *Vector) CrossProduct(other Vector) Vector {
	x := v.Y*other.Z - v.Z*other.Y
	y := v.Z*other.X - v.X*other.Z
	z := v.X*other.Y - v.Y*other.X
	return Vector{x, y, z}
}

//DotProduct Directionally multiplies the two vectors. If can be used to establish how one vector influences another's direction.
// It returns the cosine of the angle between the two vectors
// A return value of zero indicates there is no change in direction.
// A positive value indicates they are pointing in the same direction.
// A negative value indicates they are pointing in opposite directions.
func (v *Vector) DotProduct(other Vector) float64 {
	return v.X * other.X + v.Y * other.Y + v.Z * other.Z
}

// Normalize returns a resized vector within a unit of 1, keeping the direction of the original vector
func (v Vector) Normalize() Vector {
	l := v.Length()
	if l == 0 {
		return Vector{}
	}
	if l <= 1 {
		// no need to normalise
		return v
	}
	il := 1 / l
	return Vector{v.X * il, v.Y * il, v.Z * il}
}

// Add returns the addition of the two vectors
func (v *Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y, v.Z + o.Z }
}

// Subtract returns the subtraction of the two vectors
func (v *Vector) Subtract(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y, v.Z - o.Z }
}

// Multiply returns the multiplication of the two vectors
func (v *Vector) Multiply(o Vector) Vector {
	return Vector{v.X * o.X, v.Y * o.Y, v.Z * o.Z }
}

// Divide returns the division of the two vectors
func (v *Vector) Divide(o Vector) Vector {
	return Vector{v.X / o.X, v.Y / o.Y, v.Z / o.Z }
}

// AddScalar adds the float value to each dimension or the vector
func (v *Vector) AddScalar(f float64) Vector {
	return Vector{v.X + f, v.Y + f, v.Z + f }
}

// SubtractScalar subtracts the float value to each dimension or the vector
func (v *Vector) SubtractScalar(f float64) Vector {
	return Vector{v.X - f, v.Y - f, v.Z - f }
}

// MultiplyScalar multiplies each dimension or the vector by the float value
func (v *Vector) MultiplyScalar(f float64) Vector {
	return Vector{v.X * f, v.Y * f, v.Z * f }
}

// DivideScalar divides each dimension or the vector by the float value
func (v *Vector) DivideScalar(f float64) Vector {
	r := 1 / f
	return Vector{v.X * r, v.Y * r, v.Z * r}
}

// Equals returns true if the vectors are equal
func (v *Vector) ApproxEqual(o Vector) bool {
	d := v.Subtract(o)
	return (math.Abs(d.X) < APPROX_LIMIT)
}

// MidPointTo returns the vector that represents the point half way between the two vectors
func (v Vector) MidPointTo(o Vector) Vector {
	relative := o.Subtract(v) // calculate o relative to v
	mp := relative.DivideScalar(2.0) // find half way point
	return mp.Add(v) // translate back again
}

func (v *Vector) String() string {
	return fmt.Sprintf("x:%v y:%v z:%v", v.X, v.Y, v.Z)
}

func (v Vector) Round(dp int) Vector {
	return Vector{RoundUp(v.X, dp), RoundUp(v.Y, dp), RoundUp(v.Z, dp)}
}

