package file

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// go channel 批量上传文件
func UploadFiles(urls []string, datas [][]byte, types []string) error {
	errch := make(chan error, len(urls))
	for i, url := range urls {
		go func(url string, data []byte, _type string) {
			if len(url) <= 0 || len(data) <= 0 {
				errch <- nil
				return
			}
			err := upload(url, data, _type)
			if err != nil {
				errch <- err
				return
			}
			errch <- nil
		}(url, datas[i], types[i])
	}
	for _, _ = range urls {
		if err := <-errch; err != nil {
			return err
		}
	}
	return nil
}

func upload(url string, data []byte, dataType string) error {
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	if dataType != "" {
		req.Header.Set("Content-Type", dataType)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		fmt.Errorf("upload response error: %s", string(bodyBytes))
	}
	return nil
}
