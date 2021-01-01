package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

type KyeboardSystem struct {
	buttons map[rune]string
}

type KeyboardMessage struct {
	Key string
}

func (k KeyboardMessage) Type() string {
	return "KeyboardMessage"
}

func (k *KyeboardSystem) New(*ecs.World) {
	log.Println("KyeboardSystem.New")

	k.buttons = map[rune]string{
		122: "1",
		120: "2",
		99:  "3",
		106: "4",
		32:  "5",
		107: "6",
		108: "7",
	}

	engo.Mailbox.Listen(engo.TextMessage{}.Type(), k.PushHandler)
}

func (k *KyeboardSystem) Update(dt float32) {
}

func (k *KyeboardSystem) Remove(ecs.BasicEntity) {
	log.Println("KyeboardSystem.Remove")
}

func (k *KyeboardSystem) PushHandler(msg engo.Message) {
	textMsg := msg.(engo.TextMessage)
	if key, ok := k.buttons[textMsg.Char]; ok {
		engo.Mailbox.Dispatch(KeyboardMessage{
			Key: key,
		})
	}
}
