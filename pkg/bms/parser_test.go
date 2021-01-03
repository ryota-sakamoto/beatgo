package bms_test

import (
	"strings"
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
#WAV03 bdl.wav

#00101:00
#00101:2R
#00113:0000002100000000`,
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
				Data: []bms.Data{
					{
						Bar:     1,
						Channel: 1,
						Note: []string{
							"00",
						},
					},
					{
						Bar:     1,
						Channel: 1,
						Note: []string{
							"2R",
						},
					},
					{
						Bar:     1,
						Channel: 13,
						Note: []string{
							"00",
							"00",
							"00",
							"21",
							"00",
							"00",
							"00",
							"00",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result, err := bms.Parse(strings.NewReader(tt.data))
			if !assert.NoError(t, err) {
				t.Log(err)
				t.FailNow()
			}

			assert.Equal(t, tt.want, result)
		})
	}
}
