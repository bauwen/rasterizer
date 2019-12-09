package main

import (
    "image"
    "image/color"
    "math"
)

func rasterizeScene(width, height int, camera *Camera, triangles []*Triangle, background Vec3) *image.RGBA {
    frameBuffer := image.NewRGBA(image.Rect(0, 0, width, height))
    for i := 0; i < width; i++ {
        for j := 0; j < height; j++ {
            frameBuffer.SetRGBA(i, j, color.RGBA{
                uint8(background[0] * 255),
                uint8(background[1] * 255),
                uint8(background[2] * 255),
                255,
            })
        }
    }

    depthBuffer := make([]float64, width * height)
    for i := 0; i < len(depthBuffer); i++ {
        depthBuffer[i] = camera.Far
    }

    for _, triangle := range triangles {
        // from world space to camera space
        a0 := camera.FromWorldSpace(triangle.V0)
        a1 := camera.FromWorldSpace(triangle.V1)
        a2 := camera.FromWorldSpace(triangle.V2)

        // from camera space to screen space
        b0 := fromCameraToScreenSpace(a0, camera)
        b1 := fromCameraToScreenSpace(a1, camera)
        b2 := fromCameraToScreenSpace(a2, camera)

        // from screen space to raster space
        v0 := fromScreenToRasterSpace(b0, width, height, camera)
        v1 := fromScreenToRasterSpace(b1, width, height, camera)
        v2 := fromScreenToRasterSpace(b2, width, height, camera)

        // wireframe rendering
        if WIREFRAME {
            drawLine(frameBuffer, v0, v1, color.RGBA{ 0, 0, 0, 255 })
            drawLine(frameBuffer, v0, v2, color.RGBA{ 0, 0, 0, 255 })
            drawLine(frameBuffer, v1, v2, color.RGBA{ 0, 0, 0, 255 })
            continue
        }

        // compute bounding box of triangle
        xMin := int(math.Min(v0[0], math.Min(v1[0], v2[0])))
        xMax := int(math.Max(v0[0], math.Max(v1[0], v2[0])))
        yMin := int(math.Min(v0[1], math.Min(v1[1], v2[1])))
        yMax := int(math.Max(v0[1], math.Max(v1[1], v2[1])))

        if xMax < 0 || yMax < 0 || width - 1 < xMin || height - 1 < yMin {
            continue
        }

        if xMin < 0 {
            xMin = 0
        }
        if xMax > width - 1 {
            xMax = width - 1
        }
        if yMin < 0 {
            yMin = 0
        }
        if yMax > height - 1 {
            yMax = height - 1
        }

        // double area of the triangle
        area := edgeFunction(v0, v1, v2)

        // precompute reciprocal of z-coordinates
        v0[2] = 1 / v0[2]
        v1[2] = 1 / v1[2]
        v2[2] = 1 / v2[2]

        // prepare vertex attributes (color)
        c0 := triangle.C0.Scale(v0[2])
        c1 := triangle.C1.Scale(v1[2])
        c2 := triangle.C2.Scale(v2[2])

        // process points within triangle
        for y := yMin; y < yMax; y++ {
            for x := xMin; x < xMax; x++ {
                pixel := NewVec3(float64(x) + 0.5, float64(y) + 0.5, 0)
                w0 := edgeFunction(v1, v2, pixel)
                w1 := edgeFunction(v2, v0, pixel)
                w2 := edgeFunction(v0, v1, pixel)

                // check if point is inside triangle
                if w0 < 0 || w1 < 0 || w2 < 0 {
                    continue
                }

                // barycentric coordinates
                w0 /= area
                w1 /= area
                w2 /= area

                // interpolate z-coordinate
                z := 1 / (w0*v0[2] + w1*v1[2] + w2*v2[2])

                // depth buffer test
                if depthBuffer[x + y*width] <= z {
                    continue
                }
                depthBuffer[x + y*width] = z

                // shading effect based on normals
                shading := 1.0
                if SHADING {
                    normal := a1.Sub(a0).CrossProduct(a2.Sub(a0))
                    viewDirection := NewVec3(
                        z * (w0*b0[0] + w1*b1[0] + w2*b2[0]),
                        z * (w0*b0[1] + w1*b1[1] + w2*b2[1]),
                        -z,
                    )
                    shading = math.Max(0, -normal.Normalize().Dot(viewDirection.Normalize()))
                }

                // interpolate vertex attributes (color)
                c := NewVec3(
                    z * (w0*c0[0] + w1*c1[0] + w2*c2[0]) * shading,
                    z * (w0*c0[1] + w1*c1[1] + w2*c2[1]) * shading,
                    z * (w0*c0[2] + w1*c1[2] + w2*c2[2]) * shading,
                )

                /*
                // comment out to see normal map
                normal = normal.Normalize().Add(NewVec3(1, 1, 1)).Scale(0.5)
                c = NewVec3(
                    normal[0],
                    normal[1],
                    normal[2],
                )//*/

                // store resulting color in framebuffer
                frameBuffer.SetRGBA(x, y, color.RGBA{
                    uint8(c[0] * 255),
                    uint8(c[1] * 255),
                    uint8(c[2] * 255),
                    255,
                })
            }
        }
    }

    return frameBuffer
}

func fromCameraToScreenSpace(v Vec3, camera *Camera) Vec3 {
    return NewVec3(
        v[0]/-v[2] * camera.Near,
        v[1]/-v[2] * camera.Near,
        -v[2],
    )
}

func fromScreenToRasterSpace(v Vec3, width, height int, camera *Camera) Vec3 {
    return NewVec3(
        (0.5 + v[0]/camera.ScreenWidth) * float64(width),
        (0.5 - v[1]/camera.ScreenHeight) * float64(height),
        v[2],
    )
}

// this is just the magnitude of the cross product for 2D vectors
func edgeFunction(o Vec3, v Vec3, p Vec3) float64 {
    return (p[0] - o[0])*(v[1] - o[1]) - (p[1] - o[1])*(v[0] - o[0])
}
