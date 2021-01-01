package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type bounceEntity struct {
	*ecs.BasicEntity
	*SpeedComponent
	*common.SpaceComponent
}

type BounceSystem struct {
	entities []bounceEntity
}

func (b *BounceSystem) Add(basic *ecs.BasicEntity, speed *SpeedComponent, space *common.SpaceComponent) {
	b.entities = append(b.entities, bounceEntity{basic, speed, space})
}

func (b *BounceSystem) Remove(basic ecs.BasicEntity) {
}

func (b *BounceSystem) Update(dt float32) {
	// for _, e := range b.entities {
	// 	if e.SpaceComponent.Position.Y > engo.GameHeight() {
	// 		e.SpaceComponent.Position.Y = 0
	// 	}
	// }
}
