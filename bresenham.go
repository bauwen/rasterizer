package main

import (
    "image"
    "image/color"
    "math"
)


// https://en.wikipedia.org/wiki/Bresenham's_line_algorithm
func drawLine(frameBuffer *image.RGBA, p1 Vec3, p2 Vec3, color color.RGBA) {
    x0 := int(p1[0] + 0.5)
    y0 := int(p1[1] + 0.5)
    x1 := int(p2[0] + 0.5)
    y1 := int(p2[1] + 0.5)

    dx := x1 - x0
    sx := 1
    if x0 > x1 {
        dx = x0 - x1
        sx = -1
    }
    dy := y1 - y0
    sy := 1
    if y0 > y1 {
        dy = y0 - y1
        sy = -1
    }

    err := dx - dy
    for {
        frameBuffer.SetRGBA(x0, y0, color)
        if x0 == x1 && y0 == y1 {
            break
        }
        err2 := 2*err
        if err2 > -dy {
            err -= dy
            x0 += sx
        }
        if err2 < dx {
            err += dx
            y0 += sy
        }
    }
}

// https://en.wikipedia.org/wiki/Bresenham's_line_algorithm
func drawLineFloat64(frameBuffer *image.RGBA, p1 Vec3, p2 Vec3, color color.RGBA) {
    x0 := p1[0]
    y0 := p1[1]
    x1 := p2[0]
    y1 := p2[1]

    dx := math.Abs(x1 - x0)
    dy := math.Abs(y1 - y0)
    sx := 1.0
    if x0 > x1 {
        sx = -1.0
    }
    sy := 1.0
    if y0 > y1 {
        sy = -1.0
    }
    err := dx - dy
    for {
        frameBuffer.SetRGBA(int(x0), int(y0), color)
        if math.Abs(x0 - x1) < 1 && math.Abs(y0 - y1) < 1 {
            break
        }
        err2 := err*2
        if err2 > -dy {
            err -= dy
            x0 += sx
        }
        if err2 < dx {
            err += dx
            y0 += sy
        }
    }
}
