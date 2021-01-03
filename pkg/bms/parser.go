package bms

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Parse(r io.Reader) (*BMS, error) {
	result := &BMS{
		Header: Header{
			Wav: []Wav{},
		},
	}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		v := scanner.Text()

		if len(v) == 0 || v[0] != '#' {
			continue
		}

		v = v[1:]

		if strings.HasPrefix(v, "PLAYER ") {
			v = strings.Replace(v, "PLAYER ", "", 1)
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.Player = n
		}

		if strings.HasPrefix(v, "GENRE ") {
			v = strings.Replace(v, "GENRE ", "", 1)
			result.Header.Genre = v
		}

		if strings.HasPrefix(v, "ARTIST ") {
			v = strings.Replace(v, "ARTIST ", "", 1)
			result.Header.Artist = v
		}

		if strings.HasPrefix(v, "BPM ") {
			v = strings.Replace(v, "BPM ", "", 1)
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.BPM = n
		}

		if strings.HasPrefix(v, "PLAYLEVEL ") {
			v = strings.Replace(v, "PLAYLEVEL ", "", 1)
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.Playlevel = n
		}

		if strings.HasPrefix(v, "RANK ") {
			v = strings.Replace(v, "RANK ", "", 1)
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.Rank = n
		}

		if strings.HasPrefix(v, "TOTAL ") {
			v = strings.Replace(v, "TOTAL ", "", 1)
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.Total = n
		}

		if strings.HasPrefix(v, "WAV") {
			v = strings.Replace(v, "WAV", "", 1)
			wav := strings.Split(v, " ")
			if len(wav) != 2 {
				return nil, fmt.Errorf("invalid wav: %+v", wav)
			}
			result.Header.Wav = append(result.Header.Wav, Wav{
				Index: wav[0],
				File:  wav[1],
			})
		}
	}

	return result, nil
}
