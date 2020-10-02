package upload

import (
	"fmt"
	"os"
)

func UploadFile(key string, val []byte) (ret string, err error) {
	file, e := os.Create("/www/redources/" + key)
	if e != nil {
		err = e
		return
	}
	_, err = file.Write(val)
	if err != nil {
		return
	}
	ret = fmt.Sprintf("http://momoman.cn/redources/%s", key)
	return
}
