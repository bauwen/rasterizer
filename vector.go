package main

import "math"

type Vec3 [3]float64

func NewVec3(x, y, z float64) Vec3 {
    return Vec3{x, y, z}
}

func (v Vec3) Length() float64 {
    return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v Vec3) Normalize() Vec3 {
    x := 1 / v.Length()
    return NewVec3(
        v[0] * x,
        v[1] * x,
        v[2] * x,
    )
}

func (v Vec3) Add(w Vec3) Vec3 {
    return NewVec3(v[0]+w[0], v[1]+w[1], v[2]+w[2])
}

func (v Vec3) Sub(w Vec3) Vec3 {
    return NewVec3(v[0]-w[0], v[1]-w[1], v[2]-w[2])
}

func (v Vec3) Scale(x float64) Vec3 {
    return NewVec3(v[0]*x, v[1]*x, v[2]*x)
}

func (v Vec3) Dot(w Vec3) float64 {
    return v[0]*w[0] + v[1]*w[1] + v[2]*w[2]
}

func (v Vec3) CrossProduct(w Vec3) Vec3 {
    return NewVec3(
        v[1]*w[2] - v[2]*w[1],
        v[2]*w[0] - v[0]*w[2],
        v[0]*w[1] - v[1]*w[0],
    )
}
