package util

import (
	"github.com/rs/zerolog/log"
	"strconv"
)

func StrToInt(s string) int {
	if i, e := strconv.Atoi(s); e != nil {
		log.Warn().Msgf("无法转换:[%s]不是整形", s)
		return 0
	} else {
		return i
	}
}
