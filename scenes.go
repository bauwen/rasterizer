package main

func addSquad(list []*Triangle, color, v0, v1, v2, v3 Vec3) []*Triangle {
    list = append(list, &Triangle{
        v0, v1, v2,
        color, color, color,
    })
    list = append(list, &Triangle{
        v2, v3, v0,
        color, color, color,
    })
    return list
}

func createTestScene(width, height int) (*Camera, []*Triangle, Vec3) {
	// background color
	background := NewVec3(0.2, 0.8, 1.0)

	// camera settings
	lookFrom := NewVec3(2, 2, 4)
	lookAt := NewVec3(0, 0, 0)
	up := NewVec3(0, 1, 0)
	fov := 90.0
	aspect := float64(width) / float64(height)
    near := 1.0
    far := 1000.0
	camera := NewCamera(lookFrom, lookAt, up, fov, aspect, near, far)

	// triangle list
	var list []*Triangle

	/*list = append(list, &Triangle{
		NewVec3(1, 0.0, -10),
        NewVec3(-1, 0.5, -1),
        NewVec3(-1, -0.5, -1),

        NewVec3(1, 0, 0),
        NewVec3(0, 1, 0),
        NewVec3(0, 0, 1),
	})*/
    list = addSquad(list, NewVec3(0, 1, 0),
        NewVec3(-1, -1, 1),
        NewVec3(1, -1, 1),
        NewVec3(1, -1, -1),
        NewVec3(-1, -1, -1),
    )
    list = addSquad(list, NewVec3(1, 0, 0),
        NewVec3(-1, -1, -1),
        NewVec3(1, -1, -1),
        NewVec3(1, 1, -1),
        NewVec3(-1, 1, -1),
    )
    list = addSquad(list, NewVec3(0, 0, 1),
        NewVec3(-1, -1, 1),
        NewVec3(-1, -1, -1),
        NewVec3(-1, 1, -1),
        NewVec3(-1, 1, 1),
    )//*/

	return camera, list, background
}

func createGradientScene(width, height int, shading bool) (*Camera, []*Triangle, Vec3) {
    SHADING = shading

	// background color
	background := NewVec3(0.2, 0.2, 0.3)

	// camera settings
	lookFrom := NewVec3(0, 0, 0)
	lookAt := NewVec3(0, 0, -1)
	up := NewVec3(0, 1, 0)
	fov := 90.0
	aspect := float64(width) / float64(height)
    near := 1.0
    far := 1000.0
	camera := NewCamera(lookFrom, lookAt, up, fov, aspect, near, far)

	// triangle list
	var list []*Triangle

	list = append(list, &Triangle{
		NewVec3(1, 0.0, -3),
        NewVec3(-1, 0.5, -1),
        NewVec3(-1, -0.5, -1),

        NewVec3(1, 0, 0),
        NewVec3(0, 1, 0),
        NewVec3(0, 0, 1),
	})

    list = append(list, &Triangle{
		NewVec3(1, 0.0, -3),
        NewVec3(-1, -0.5, -1),
        NewVec3(1, -0.5, -1),

        NewVec3(1, 0, 0),
        NewVec3(0, 0, 1),
        NewVec3(1, 1, 0),
	})

    return camera, list, background
}

func createModelScene(width, height int, filename string) (*Camera, []*Triangle, Vec3) {
    // background color
	background := NewVec3(0.2, 0.8, 1.0)

	// camera settings
	lookFrom := NewVec3(-3, 1, -3)
	lookAt := NewVec3(0, 0, 0)
	up := NewVec3(0, 1, 0)
	fov := 90.0
	aspect := float64(width) / float64(height)
    near := 1.0
    far := 1000.0
	camera := NewCamera(lookFrom, lookAt, up, fov, aspect, near, far)

	// triangle list
    color := NewVec3(1.0, 0.2, 0.8)
    scale := 1.0
    translation := NewVec3(0, 0, 0)
	list, err := loadBinarySTLModel(filename, color, scale, translation)
    if err != nil {
        panic(err)
    }

	return camera, list, background
}
