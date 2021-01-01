package systems

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Ball struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	SpeedComponent
}

type LaneSystem struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	world *ecs.World

	turnNote  common.Texture
	whiteNote common.Texture
	blueNote  common.Texture

	before float32
}

func (l *LaneSystem) New(w *ecs.World) {
	log.Println("LaneSystem.New")
	engo.Mailbox.Listen(KeyboardMessage{}.Type(), l.PushHandler)

	l.world = w
	l.turnNote = common.NewTextureSingle(common.NewImageObject(common.ImageToNRGBA(getNote(0, 50, 30, 60), 100, 80)))
	l.whiteNote = common.NewTextureSingle(common.NewImageObject(common.ImageToNRGBA(getNote(31, 50, 48, 60), 100, 80)))
	l.blueNote = common.NewTextureSingle(common.NewImageObject(common.ImageToNRGBA(getNote(49, 50, 62, 60), 100, 200)))

	rand.Seed(rand.Int63())
}

func (l *LaneSystem) Update(dt float32) {
	l.before += dt
	if l.before > 0.3 {
		l.before = 0

		p := rand.Intn(512)
		for i := 0; i < 8; i++ {
			if (p>>i)&1 == 0 {
				continue
			}

			ball := Ball{BasicEntity: ecs.NewBasic()}
			if i == 0 {
				ball.RenderComponent = common.RenderComponent{
					Drawable: l.turnNote,
					Scale:    engo.Point{5, 5},
				}
				ball.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{
						X: 0,
						Y: -50,
					},
					Width:  l.whiteNote.Width() * ball.RenderComponent.Scale.X,
					Height: l.whiteNote.Height() * ball.RenderComponent.Scale.Y,
				}
			} else if i%2 == 1 {
				ball.RenderComponent = common.RenderComponent{
					Drawable: l.whiteNote,
					Scale:    engo.Point{5, 5},
				}
				ball.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{
						X: float32(150 + 150*(i/2)),
						Y: -50,
					},
					Width:  l.whiteNote.Width() * ball.RenderComponent.Scale.X,
					Height: l.whiteNote.Height() * ball.RenderComponent.Scale.Y,
				}
			} else {
				ball.RenderComponent = common.RenderComponent{
					Drawable: l.blueNote,
					Scale:    engo.Point{5, 5},
				}
				ball.SpaceComponent = common.SpaceComponent{
					Position: engo.Point{
						X: float32(85 + 150*(i/2)),
						Y: -50,
					},
					Width:  l.blueNote.Width() * ball.RenderComponent.Scale.X,
					Height: l.blueNote.Height() * ball.RenderComponent.Scale.Y,
				}
			}

			ball.SpeedComponent = SpeedComponent{Point: engo.Point{300, 300}}

			for _, system := range l.world.Systems() {
				switch sys := system.(type) {
				case *common.RenderSystem:
					sys.Add(&ball.BasicEntity, &ball.RenderComponent, &ball.SpaceComponent)
				case *SpeedSystem:
					sys.Add(&ball.BasicEntity, &ball.SpeedComponent, &ball.SpaceComponent)
				case *BounceSystem:
					sys.Add(&ball.BasicEntity, &ball.SpeedComponent, &ball.SpaceComponent)
				}
			}
		}
	}
}

func (l *LaneSystem) Remove(ecs.BasicEntity) {
}

func (l *LaneSystem) PushHandler(msg engo.Message) {
	log.Println("LaneSystem.PushHandler", msg)
}

func getNote(x0, y0, x1, y1 int) image.Image {
	f, err := os.Open("assets/note.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	if v, ok := src.(interface {
		SubImage(r image.Rectangle) image.Image
	}); ok {
		src = v.SubImage(image.Rect(x0, y0, x1, y1))
	}

	buff := bytes.Buffer{}
	if err := png.Encode(&buff, src); err != nil {
		panic(err)
	}

	src, _, err = image.Decode(&buff)
	if err != nil {
		panic(err)
	}

	return src
}
