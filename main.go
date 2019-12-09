package main

import (
	"image/png"
	"os"
)

// Simple rasterizer implementation, based on the tutorial found at
// https://www.scratchapixel.com/lessons/3d-basic-rendering/rasterization-practical-implementation/projection-stage
// (all chapters)

/*
Possible TODOs:
x add simple shading effect based on normals
- more efficient edgeFunction'ing (with constant steps, see tutorial)
- use textures
- add (non-)backface-culling
- take into account overlapping edges via top-left test

- add anti-aliasing by dividing up pixels
- remove apparent "cracks" between triangles in result
- add light sources

- create proper pipeline, sort of vertex and fragment shaders etc.
- create scene system to add and remove objects easily
- add interactive loop for animations and moving the camera
*/

var SHADING = true
var WIREFRAME = false

func main() {
    width := 800
    height := 600

    // comment out different scenes here
    //camera, triangles, background := createTestScene(width, height)
    camera, triangles, background := createModelScene(width, height, "elephant.stl")
    //camera, triangles, background := createModelScene(width, height, "tyranitar.stl")
    //camera, triangles, background := createGradientScene(width, height, false)

    pixels := rasterizeScene(width, height, camera, triangles, background)

    f, err := os.Create("output.png")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    if err := png.Encode(f, pixels); err != nil {
        panic(err)
    }
}
