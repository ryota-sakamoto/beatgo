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
	"github.com/ryota-sakamoto/beatgo/pkg/bms"
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

	before    float32
	baseSpeed float32
}

func (l *LaneSystem) New(w *ecs.World) {
	log.Println("LaneSystem.New")
	engo.Mailbox.Listen(KeyboardMessage{}.Type(), l.PushHandler)

	l.world = w
	l.turnNote = common.NewTextureSingle(common.NewImageObject(common.ImageToNRGBA(getNoteImage(0, 50, 30, 60), 100, 80)))
	l.whiteNote = common.NewTextureSingle(common.NewImageObject(common.ImageToNRGBA(getNoteImage(31, 50, 48, 60), 100, 80)))
	l.blueNote = common.NewTextureSingle(common.NewImageObject(common.ImageToNRGBA(getNoteImage(49, 50, 62, 60), 100, 200)))
	l.baseSpeed = 1000

	rand.Seed(rand.Int63())

	f, err := os.Open("assets/meikai(ANOTHER).bme")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := bms.Parse(f)
	if err != nil {
		panic(err)
	}

	l.PlaceNote(data)
}

func (l *LaneSystem) PlaceNote(data *bms.BMS) {
	for _, v := range data.Data {
		list := l.GetNote(&v)
		if len(list) == 0 {
			continue
		}

		for _, v := range list {
			for _, system := range l.world.Systems() {
				switch sys := system.(type) {
				case *common.RenderSystem:
					sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
				case *SpeedSystem:
					sys.Add(&v.BasicEntity, &v.SpeedComponent, &v.SpaceComponent)
				case *BounceSystem:
					sys.Add(&v.BasicEntity, &v.SpeedComponent, &v.SpaceComponent)
				}
			}
		}
	}
}

func (l *LaneSystem) GetNote(data *bms.Data) []*Ball {
	result := []*Ball{}

	basis := l.baseSpeed / float32(len(data.Note))
	for i, v := range data.Note {
		if v == "00" {
			continue
		}

		ball, x := l.getNote(data.Channel)
		if ball == nil {
			continue
		}

		ball.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{
				X: x,
				Y: (l.baseSpeed*float32(data.Bar) + basis*float32(i)) * -1,
			},
			Width:  l.whiteNote.Width() * ball.RenderComponent.Scale.X,
			Height: l.whiteNote.Height() * ball.RenderComponent.Scale.Y,
		}

		result = append(result, ball)
	}

	return result
}

func (l *LaneSystem) getNote(channel int) (*Ball, float32) {
	ball := &Ball{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Scale: engo.Point{5, 5},
		},
		SpeedComponent: SpeedComponent{Point: engo.Point{l.baseSpeed, l.baseSpeed}},
	}
	var x float32

	switch channel {
	case 16:
		x = 0
		ball.Drawable = l.turnNote
	case 11:
		x = 150
		ball.Drawable = l.whiteNote
	case 12:
		x = 235
		ball.Drawable = l.blueNote
	case 13:
		x = 300
		ball.Drawable = l.whiteNote
	case 14:
		x = 385
		ball.Drawable = l.blueNote
	case 15:
		x = 450
		ball.Drawable = l.whiteNote
	case 17:
		x = 535
		ball.Drawable = l.blueNote
	case 18:
		x = 600
		ball.Drawable = l.whiteNote
	default:
		return nil, 0
	}

	return ball, x
}

func (l *LaneSystem) Update(dt float32) {
	// l.before += dt
	// if l.before > 0.3 {
	// 	l.before = 0
	// }
}

func (l *LaneSystem) Remove(ecs.BasicEntity) {
}

func (l *LaneSystem) PushHandler(msg engo.Message) {
	log.Println("LaneSystem.PushHandler", msg)
}

func getNoteImage(x0, y0, x1, y1 int) image.Image {
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
