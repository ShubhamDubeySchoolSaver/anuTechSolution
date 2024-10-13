package utility

import (
	"fmt"
	"strings"
	"time"
)

func GetBlobName(fileName string) string {
	t := time.Now()
	splitedFileName := strings.Split(fileName, ".")
	ext := splitedFileName[len(splitedFileName)-1]
	return fmt.Sprintf("%s_%s.%s", splitedFileName[0], t.Format("20060102"), ext)
}
