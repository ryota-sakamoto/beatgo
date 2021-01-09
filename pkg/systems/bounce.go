package systems

import (
	"log"
	"sync"
	"time"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type bounceEntity struct {
	*ecs.BasicEntity
	*SpeedComponent
	*common.SpaceComponent
	*common.AudioComponent
}

type BounceSystem struct {
	entities    []bounceEntity
	playing     []bounceEntity
	audioSystem *common.AudioSystem
	mu          sync.Mutex
}

func (b *BounceSystem) New(w *ecs.World) {
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.AudioSystem:
			b.audioSystem = sys
		}
	}
	b.mu = sync.Mutex{}

	go b.watch()
}

func (b *BounceSystem) Add(basic *ecs.BasicEntity, speed *SpeedComponent, space *common.SpaceComponent, audio *common.AudioComponent) {
	b.entities = append(b.entities, bounceEntity{basic, speed, space, audio})
}

func (b *BounceSystem) Remove(basic ecs.BasicEntity) {
}

func (b *BounceSystem) Update(dt float32) {
	next := []bounceEntity{}
	for _, e := range b.entities {
		if e.SpaceComponent.Position.Y > engo.GameHeight() {
			b.audioSystem.Add(e.BasicEntity, e.AudioComponent)
			e.AudioComponent.Player.Rewind()
			e.AudioComponent.Player.Play()

			b.mu.Lock()
			b.playing = append(b.playing, e)
			b.mu.Unlock()

			log.Println(e.AudioComponent.Player.URL())
		} else {
			next = append(next, e)
		}
	}
	b.entities = next
}

func (b *BounceSystem) watch() {
	t := time.NewTicker(time.Millisecond * 1000)
	for range t.C {
		next := []bounceEntity{}
		b.mu.Lock()
		for _, v := range b.playing {
			if v.Player.IsPlaying() {
				next = append(next, v)
			} else {
				b.audioSystem.Remove(*v.BasicEntity)
			}
		}
		b.playing = next
		b.mu.Unlock()

		log.Println("playing", len(b.playing))
	}
}
