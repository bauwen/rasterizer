package main

import "math"

type Camera struct {
    ScreenWidth, ScreenHeight float64
    Near, Far float64
    Matrix Mat4
}

func NewCamera(lookFrom, lookAt, up Vec3, vfov, aspect, near, far float64) *Camera {
    camera := new(Camera)

    z := lookFrom.Sub(lookAt).Normalize()
    x := up.CrossProduct(z).Normalize()
    y := z.CrossProduct(x).Normalize()

    m := NewMat4(
        x[0], x[1], x[2], 0,
        y[0], y[1], y[2], 0,
        z[0], z[1], z[2], 0,
        0, 0, 0, 1,
    )
    t := m.Multiply(lookFrom)

    camera.ScreenHeight = math.Tan(vfov * math.Pi / 180 / 2) * 2 * near
    camera.ScreenWidth = aspect * camera.ScreenHeight
    camera.Near = near
    camera.Far = far
    camera.Matrix = NewMat4(
        x[0], x[1], x[2], -t[0],
        y[0], y[1], y[2], -t[1],
        z[0], z[1], z[2], -t[2],
        0, 0, 0, 1,
    )

    return camera
}

func (c *Camera) FromWorldSpace(v Vec3) Vec3 {
    return c.Matrix.Multiply(v)
}
