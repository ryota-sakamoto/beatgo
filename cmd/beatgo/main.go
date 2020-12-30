package main

import (
	"image/color"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/ryota-sakamoto/beatgo/pkg/systems"
)

type scene struct {
}

func (*scene) Type() string { return "myGame" }

func (*scene) Preload() {
	log.Println("scene.Preload")
}

func (*scene) Setup(u engo.Updater) {
	log.Println("scene.Setup")

	world, _ := u.(*ecs.World)
	engo.Input.RegisterButton("AddCity", engo.KeyF1)
	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(common.NewKeyboardScroller(400, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	// world.AddSystem(&common.EdgeScroller{400, 20})
	world.AddSystem(&common.MouseZoomer{-0.125})

	world.AddSystem(&systems.KyeboardSystem{})
}

func main() {
	opts := engo.RunOptions{
		Title:          "Hello World",
		Width:          400,
		Height:         400,
		StandardInputs: true,
		FPSLimit:       120,
	}
	engo.Run(opts, &scene{})
}
