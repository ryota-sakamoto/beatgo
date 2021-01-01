package bms

import (
	"strconv"
	"strings"
)

func Parse(data string) (*BMS, error) {
	result := &BMS{
		Header: Header{},
	}

	for _, v := range strings.Split(data, "\n") {
		if strings.HasPrefix(v, "#PLAYER ") {
			v = strings.ReplaceAll(v, "#PLAYER ", "")
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.Player = n
		}

		if strings.HasPrefix(v, "#GENRE ") {
			v = strings.ReplaceAll(v, "#GENRE ", "")
			result.Header.Genre = v
		}

		if strings.HasPrefix(v, "#ARTIST ") {
			v = strings.ReplaceAll(v, "#ARTIST ", "")
			result.Header.Artist = v
		}

		if strings.HasPrefix(v, "#BPM ") {
			v = strings.ReplaceAll(v, "#BPM ", "")
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.BPM = n
		}

		if strings.HasPrefix(v, "#PLAYLEVEL ") {
			v = strings.ReplaceAll(v, "#PLAYLEVEL ", "")
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.Playlevel = n
		}

		if strings.HasPrefix(v, "#RANK ") {
			v = strings.ReplaceAll(v, "#RANK ", "")
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.Rank = n
		}

		if strings.HasPrefix(v, "#TOTAL ") {
			v = strings.ReplaceAll(v, "#TOTAL ", "")
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result.Header.Total = n
		}
	}

	return result, nil
}
