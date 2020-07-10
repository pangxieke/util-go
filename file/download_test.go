package file

import (
	"testing"
)

func TestDownloadFiles(t *testing.T) {
	urls := []string{
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594386027941&di=2cc802eca03ea5ff71792d3c887eec7d&imgtype=0&src=http%3A%2F%2Fattachments.gfan.com%2Fforum%2Fattachments2%2Fday_110317%2F1103171814794ace68f8b856a8.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594386027941&di=4b196d6655314c8c851b4fe0ae1b7bdb&imgtype=0&src=http%3A%2F%2Fn.sinaimg.cn%2Fsinacn10%2F656%2Fw580h876%2F20180722%2F1553-hfqtahi6951449.jpg",
	}

	_, err := DownloadFiles(urls)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkDownloadFiles(b *testing.B) {
	urls := []string{
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594386027941&di=2cc802eca03ea5ff71792d3c887eec7d&imgtype=0&src=http%3A%2F%2Fattachments.gfan.com%2Fforum%2Fattachments2%2Fday_110317%2F1103171814794ace68f8b856a8.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594386027941&di=4b196d6655314c8c851b4fe0ae1b7bdb&imgtype=0&src=http%3A%2F%2Fn.sinaimg.cn%2Fsinacn10%2F656%2Fw580h876%2F20180722%2F1553-hfqtahi6951449.jpg",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := DownloadFiles(urls)
		if err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkDownload(b *testing.B) {
	url := "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594386027941&di=2cc802eca03ea5ff71792d3c887eec7d&imgtype=0&src=http%3A%2F%2Fattachments.gfan.com%2Fforum%2Fattachments2%2Fday_110317%2F1103171814794ace68f8b856a8.jpg"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := download(url)
		if err != nil {
			b.Fatal(err)
		}
	}

}
