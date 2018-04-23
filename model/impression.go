package model

import (
	"fmt"
)

type Report struct {
	FeedId string `json:"feed_id"`
}

type Reportstats struct {
	Return int64           `json:"ret"`
	Msg    string          `json:"msg"`
	Data   ImpressionStats `json:"data"`
}

//{"ret":0,"msg":"","data":{"playnum":42,"_idc":"sh"}}
func (r Reportstats) String() string {
	return fmt.Sprintf("ret:%d, msg:%s, data:%s",
		r.Return,
		r.Msg,
		r.Data.String(),
	)
}

type ImpressionStats struct {
	PlayNum int64  `json:"playnum"`
	IDC     string `json:"_idc"`
}

func (i ImpressionStats) String() string {
	return fmt.Sprintf(
		"playnum:%d, _idc:%s",
		i.PlayNum,
		i.IDC,
	)
}

//'cookie: pgv_pvid=8268826880; pgv_info=ssid=s3380927418' -H 'origin: https://h5.qzone.qq.com' -H 'accept-encoding: gzip, deflate, br' -H 'accept-language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7' -H 'user-agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36' -H 'content-type: application/json' -H 'accept: application/json' -H 'referer: https://h5.qzone.qq.com/weishi/feed/nxrCbaO5GkIqZzKe/wsfeed?_proxy=1&_wv=1&id=nxrCbaO5GkIqZzKe&spid=1524057614152341&from=pc&orifrom='

const ImpressionUrl string = "https://h5.qzone.qq.com/webapp/json/weishi/ReportFeedPlay?&t=0.03073960283495647&g_tk="
