package work

import (
	"fmt"
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
	
}

//保存短连接《-》长连接
func (shortPtr *Shortener) SaveShortUrl(shortUrl string) bool {
	longUrl := shortP.longUrl
	//保存
}




type ShortenerQuery struct {
}

func NewShortenerQuery() *ShortenerQuery {
	return &ShortenerQuery{}
}

//查询短连接
func (queryPtr *Shortener) GetLongUrlByShortUrl() string {
}



