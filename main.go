package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"runtime"
	"sync"

	"github.com/mdh67899/w-spider/model"
	"github.com/mdh67899/w-spider/utils"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var (
	wg = new(sync.WaitGroup)
)

func main() {
	//var originUrl string = "https://h5.qzone.qq.com/weishi/feed/nxrCbaO5GkIqZzKe/wsfeed?_proxy=1&_wv=1&id=nxrCbaO5GkIqZzKe&spid=1524057614152341"
	originUrl := ""
	visitCount := 0
	threads := 1

	flag.StringVar(&originUrl, "url", "", "your shared video page url")
	flag.IntVar(&visitCount, "visit", 0, "your show num which wanted")
	flag.IntVar(&threads, "threads", 1, "your goroutine numbers")
	flag.Parse()

	if originUrl == "" {
		log.Println("you are not provides origin url, please use -url to make sure url")
		return
	}

	if visitCount == 0 {
		log.Println("your video show number is 0, please use -visit to make sure your show count")
		return
	}

	if threads < 1 {
		log.Println("your visit thread couldn't less than 0, please use -threads to make sure your visit thread")
		return
	}

	validUrl, err := url.Parse(originUrl)
	if err != nil {
		log.Println("url parse error:", originUrl)
		return
	}

	feedId := validUrl.Query().Get("id")
	if feedId == "" {
		log.Println("cann't get feed_id from origin_url", originUrl)
		return
	}

	fakeReport := model.Report{
		FeedId: feedId,
	}

	dataCh := make(chan model.Report)
	closeCh := make(chan struct{})
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go Visit(dataCh, closeCh, wg)
	}

	for i := 0; i < visitCount; i++ {
		dataCh <- fakeReport
	}

	close(closeCh)

	wg.Wait()
}

func Visit(r chan model.Report, closeCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data, exist := <-r:
			if !exist {
				return
			}

			resp, err := utils.PostJSON(model.ImpressionUrl, data)
			if err != nil {
				log.Println("Post json error:", err)
				continue
			}

			var Stats model.Reportstats
			err = json.Unmarshal(resp, &Stats)
			if err != nil {
				log.Println("Unmarshal response data error:", err)
				continue
			}

			log.Println(Stats.String())

		case <-closeCh:
			return
		}
	}
}
