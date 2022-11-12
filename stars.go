package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/jmcvetta/randutil"
)

type StarColor string

const (
	StarColorSol StarColor = "fce570"
	StarColorO   StarColor = "9bb0ff"
	StarColorB   StarColor = "aabfff"
	StarColorA   StarColor = "cad7ff"
	StarColorF   StarColor = "f8f7ff"
	StarColorG   StarColor = "fff4ea"
	StarColorK   StarColor = "ffd2a1"
	StarColorM   StarColor = "ffcc6f"
)

// Brighter colors, for testing
// const (
// 	StarColorSol StarColor = "F9EEB9"
// 	StarColorO   StarColor = "DAE1FC"
// 	StarColorB   StarColor = "E7ECFF"
// 	StarColorA   StarColor = "F9FBFF"
// 	StarColorF   StarColor = "FFFFFF"
// 	StarColorG   StarColor = "FFF9F3"
// 	StarColorK   StarColor = "FCE8D0"
// 	StarColorM   StarColor = "FCE1AB"
// )

func (c StarColor) String() string {
	return string(c)
}

func (c StarColor) Hex() uint {
	i, _ := strconv.ParseUint(string(c), 16, 32)
	return uint(i)
}

func addStars(amount int, scene *core.Node) {
	sphere := geometry.NewSphere(0.05, 10, 10)

	scene.Add(sol(sphere))
	for i := 0; i < amount; i++ {
		mat := material.NewStandard(color())
		s := graphic.NewMesh(sphere, mat)
		s.SetPositionVec(diskVector(200))
		scene.Add(s)
	}
}

// generate a Sol star
func sol(sphere *geometry.Geometry) *graphic.Mesh {
	mat := material.NewStandard(math32.NewColorHex(StarColorSol.Hex()))
	s := graphic.NewMesh(sphere, mat)
	s.SetPosition(0, 0, 0)
	return s
}

// a random vector on a disk of radius r
func diskVector(radius float32) *math32.Vector3 {
	theta := rand.Float32() * 2 * math32.Pi
	r := radius * math32.Sqrt(float32(rand.NormFloat64()))
	return math32.NewVector3(r*math32.Cos(theta), randomNormPosition(-5, 5), r*math32.Sin(theta))
}

// generate a random float32 between min and max
func randomNormPosition(min, max int) float32 {
	rand.Seed(time.Now().UnixNano())
	norm := rand.NormFloat64()
	return float32(norm*float64(max-min) + float64(min))
}

// generate a random color
func color() *math32.Color {
	colors := []randutil.Choice{
		{Weight: 1, Item: StarColorSol},
		{Weight: 1, Item: StarColorO},
		{Weight: 1, Item: StarColorB},
		{Weight: 3, Item: StarColorA},
		{Weight: 3, Item: StarColorF},
		{Weight: 1, Item: StarColorG},
		{Weight: 1, Item: StarColorK},
		{Weight: 1, Item: StarColorM},
	}
	color, _ := randutil.WeightedChoice(colors)
	return math32.NewColorHex(color.Item.(StarColor).Hex())
}
