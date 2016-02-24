package work

import (
	_ "fmt"
)

type Shortener struct {
	longUrl string
}

func NewShortener(longUrl string) *Shortener {
	return &Shortener{longUrl}
}

func (shortPtr *Shortener) IsUrl() bool {
	return true
}

//生成短连接
func (shortPtr *Shortener) GenShortUrl() string {
	return ""
}

//保存短连接《-》长连接
func (shortPtr *Shortener) SaveShortUrl(shortUrl string) bool {
	//_ := shortPtr.longUrl
	//保存
	return true
}

type ShortenerQuery struct {
}

func NewShortenerQuery() *ShortenerQuery {
	return &ShortenerQuery{}
}

//查询短连接
func (queryPtr *Shortener) GetLongUrlByShortUrl() string {
	return ""
}
