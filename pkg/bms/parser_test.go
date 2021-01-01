package bms_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ryota-sakamoto/beatgo/pkg/bms"
)

func TestParse(t *testing.T) {
	tests := []struct {
		data string
		want *bms.BMS
	}{
		{
			data: `#PLAYER 1
#GENRE MYSTIC SEQUENCER
#TITLE 冥界帰航(ANOTHER)
#ARTIST ZUN(Arr.sun3)
#BPM 200
#PLAYLEVEL 12
#RANK 2
#TOTAL 400

#WAV02 bd.wav
#WAV03 bdl.wav`,
			want: &bms.BMS{
				Header: bms.Header{
					Player:    1,
					Genre:     "MYSTIC SEQUENCER",
					Artist:    "ZUN(Arr.sun3)",
					BPM:       200,
					Playlevel: 12,
					Rank:      2,
					Total:     400,
					Wav: []bms.Wav{
						{
							Index: "02",
							File:  "bd.wav",
						},
						{
							Index: "03",
							File:  "bdl.wav",
						},
					},
				},
				Data: bms.Data{},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result, err := bms.Parse(tt.data)
			if !assert.NoError(t, err) {
				t.Log(err)
				t.FailNow()
			}

			assert.Equal(t, tt.want, result)
		})
	}
}
