package types

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/logger"
	"strconv"
)

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.Error(err)
	}

	return i
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}
