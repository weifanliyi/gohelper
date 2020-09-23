package docm

import (
	"fmt"
	"os"
)

const (
	MOTHFORMAT       = "200601"
	DATEFORMAT       = "20060102"
	TIMEFORMAT       = "20060102150405"
	NORMALTIMEFORMAT = "2006-01-02 15:04:05"
)

func PathExists(p string) (bool, string) {
	_, err := os.Stat(p)
	if err == nil {
		return true, ""
	}

	//说明确认文件夹不存在
	if os.IsNotExist(err) {
		err = os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return false, fmt.Sprintf("DOCM,MKdirAll,log folder creation faild:%s", err.Error())
		}

		return true, fmt.Sprintf("DOCM,log folder created successfull")
	}

	return false, fmt.Sprintf("DOCM,log folder creation faild:%s", err.Error())
}
