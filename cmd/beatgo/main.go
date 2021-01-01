package main

import (
	"image/color"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/ryota-sakamoto/beatgo/pkg/systems"
)

type Beatgo struct{}

func (pong *Beatgo) Preload() {
	err := engo.Files.Load("note.png")
	if err != nil {
		log.Println(err)
	}
}

func (pong *Beatgo) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.CollisionSystem{Solids: 1})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&systems.SpeedSystem{})
	w.AddSystem(&systems.BounceSystem{})
	w.AddSystem(&systems.LaneSystem{})
}

func (*Beatgo) Type() string { return "beatgo" }

func main() {
	opts := engo.RunOptions{
		Title:         "beatgo",
		Width:         600,
		Height:        500,
		ScaleOnResize: true,
		FPSLimit:      1000,
	}
	engo.Run(opts, &Beatgo{})
}
