package main

import (
	"encoding/binary"
	"bytes"
	"io/ioutil"
	"math"
)

func loadBinarySTLModel(filename string, color Vec3, scale float64, translation Vec3) ([]*Triangle, error) {
    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
	}

    var list []*Triangle

    r := bytes.NewReader(contents[80:])
    var amount uint32
    binary.Read(r, binary.LittleEndian, &amount)

    for i := 0; i < int(amount); i++ {
        // skip normal vector
        for j := 0; j < 3*4; j++ {
            r.ReadByte()
        }

        // read vertex data
        vertices := make([]float32, 3*3)
        for j := 0; j < 3*3; j++ {
            binary.Read(r, binary.LittleEndian, &vertices[j])
        }
        list = append(list, &Triangle{
            NewVec3(float64(vertices[0]), float64(vertices[1]), float64(vertices[2])),
            NewVec3(float64(vertices[3]), float64(vertices[4]), float64(vertices[5])),
            NewVec3(float64(vertices[6]), float64(vertices[7]), float64(vertices[8])),
            color, color, color,
		})

        // skip attribute byte count
        for j := 0; j < 2; j++ {
            r.ReadByte()
        }
	}

	// normalize vertices
	N := float64(amount * 3)
	avgX := 0.0
	avgY := 0.0
	avgZ := 0.0
	for _, tr := range list {
		avgX += (tr.V0[0] + tr.V1[0] + tr.V2[0]) / N
		avgY += (tr.V0[1] + tr.V1[1] + tr.V2[1]) / N
		avgZ += (tr.V0[2] + tr.V1[2] + tr.V2[2]) / N
	}
	varX := 0.0
	varY := 0.0
	varZ := 0.0
	sq := func (x float64) float64 { return x*x }
	for _, tr := range list {
		varX += (sq(tr.V0[0] - avgX) + sq(tr.V1[0] - avgX) + sq(tr.V2[0] - avgX)) / N
		varY += (sq(tr.V0[1] - avgY) + sq(tr.V1[1] - avgY) + sq(tr.V2[1] - avgY)) / N
		varZ += (sq(tr.V0[2] - avgZ) + sq(tr.V1[2] - avgZ) + sq(tr.V2[2] - avgZ)) / N
	}
	varX = math.Sqrt(varX)
	varY = math.Sqrt(varY)
	varZ = math.Sqrt(varZ)
	for i := 0; i < len(list); i++ {
		tr := list[i]

		tx := tr.V0[0]
		ty := tr.V0[1]
		tz := tr.V0[2]

		tr.V0[0] = (tr.V2[0] - avgX) / varX * scale + translation[0]
		tr.V0[1] = (tr.V2[1] - avgY) / varY * scale + translation[1]
		tr.V0[2] = (tr.V2[2] - avgZ) / varZ * -scale + translation[2]

		tr.V1[0] = (tr.V1[0] - avgX) / varX * scale + translation[0]
		tr.V1[1] = (tr.V1[1] - avgY) / varY * scale + translation[1]
		tr.V1[2] = (tr.V1[2] - avgZ) / varZ * -scale + translation[2]

		tr.V2[0] = (tx - avgX) / varX * scale + translation[0]
		tr.V2[1] = (ty - avgY) / varY * scale + translation[1]
		tr.V2[2] = (tz - avgZ) / varZ * -scale + translation[2]
	}
	/*
	for _, tr := range list {
		fmt.Println(tr)
	}
	fmt.Println(avgX, avgY, avgZ)
	fmt.Println(varX, varY, varZ)
    fmt.Println("list size:", len(list))
	*/
    return list, nil
}
