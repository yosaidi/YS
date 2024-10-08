package reload

import (
	"strconv"
)

func Base(s *[]string) {
	slice := *s

	for i := 1; i < len(slice); i++ {
		if slice[i] == "(bin)" {
			temp, err := strconv.ParseInt(slice[i-1], 2, 64)
			if err != nil {
				return
			}
			slice[i-1] = strconv.Itoa(int(temp))
			slice = append(slice[:i], slice[i+1:]...)

		} else if slice[i] == "(hex)" {
			temp, err := strconv.ParseInt(slice[i-1], 16, 64)
			if err != nil {
				return
			}
			slice[i-1] = strconv.Itoa(int(temp))
			slice = append(slice[:i], slice[i+1:]...)
		}
	}

	*s = slice
}
