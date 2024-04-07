package game

import "errors"

func findRoomExist(rooms []string, rid string) error {
	for _, v := range rooms {
		if v == rid {
			return nil
		}
	}

	return errors.New("roomId not found")
}
