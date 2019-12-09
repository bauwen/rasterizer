package main

type Mat4 []float64

func NewMat4(a0, a1, a2, a3, b0, b1, b2, b3, c0, c1, c2, c3, d0, d1, d2, d3 float64) Mat4 {
    return Mat4{
        a0, a1, a2, a3,
        b0, b1, b2, b3,
        c0, c1, c2, c3,
        d0, d1, d2, d3,
    }
}

// not really a correct kind of multiplication, but convenient
func (m Mat4) Multiply(n Vec3) Vec3 {
    return NewVec3(
        m[0]*n[0] + m[1]*n[1] + m[2] *n[2] + m[3],
        m[4]*n[0] + m[5]*n[1] + m[6] *n[2] + m[7],
        m[8]*n[0] + m[9]*n[1] + m[10]*n[2] + m[11],
    )
}
