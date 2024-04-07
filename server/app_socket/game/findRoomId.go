package game

import (
	"errors"
	"strings"
)

func findRoomId(querys []string) (string, error) {
	for _, q := range querys {
		tmp := strings.Split(q, "=")
		k, v := tmp[0], tmp[1]

		if k == roomIdKey {
			return v, nil
		}

	}
	return "", errors.New("roomId key not found")
}
