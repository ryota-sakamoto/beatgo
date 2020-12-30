package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

type KyeboardSystem struct {
}

func (k *KyeboardSystem) New(*ecs.World) {
	log.Println("KyeboardSystem.New")

	engo.Input.RegisterButton("push_1", engo.KeyZ)
	engo.Input.RegisterButton("push_2", engo.KeyX)
	engo.Input.RegisterButton("push_3", engo.KeyC)
	engo.Input.RegisterButton("push_4", engo.KeyJ)
	engo.Input.RegisterButton("push_5", engo.KeySpace)
	engo.Input.RegisterButton("push_6", engo.KeyK)
	engo.Input.RegisterButton("push_7", engo.KeyL)
}

func (k *KyeboardSystem) Update(dt float32) {
	for _, key := range []string{
		"push_1",
		"push_2",
		"push_3",
		"push_4",
		"push_5",
		"push_6",
		"push_7",
	} {
		if engo.Input.Button(key).JustPressed() {
			log.Println(key)
		}
	}
}

func (k *KyeboardSystem) Remove(ecs.BasicEntity) {
	log.Println("KyeboardSystem.Remove")
}
