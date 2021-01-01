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
	}

	return result, nil
}
