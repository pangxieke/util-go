package file

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// go channel 批量下载文件
func DownloadFiles(urls []string) ([][]byte, error) {
	dataset := make([][]byte, len(urls))
	done := make(chan map[int][]byte, len(urls))
	errch := make(chan error, len(urls))
	count := 0
	for i, url := range urls {
		if len(url) <= 0 {
			continue
		}
		count += 1
		go func(index int, url string) {
			b, err := download(url)
			if err != nil {
				done <- nil
				errch <- err
				return
			}
			done <- map[int][]byte{index: b}
			errch <- nil
		}(i, url)
	}

	for i := 0; i < count; i++ {
		if err := <-errch; err != nil {
			return nil, err
		}
		for k, v := range <-done {
			dataset[k] = v
		}
	}
	return dataset, nil
}

func download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failure to download: %s, error with status: %d", url, resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
