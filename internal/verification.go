package internal

import (
	"github.com/gocolly/colly"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func GenerateVerificationCode() string {
	n := 64
	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func CheckVerificationCode(u2ID int32, verificationCode string) bool {
	verified := false
	url := "https://u2.dmhy.org/userdetails.php?id=" + strconv.FormatInt(int64(u2ID), 10)
	header := http.Header{
		"Accept-Encoding": []string{""},
		"Cache-Control":   []string{"max-age=0"},
		"Cookie":          []string{os.Getenv("u2_cookie")},
		"Host":            []string{"u2.dmhy.org"},
		"Referer":         []string{"https://u2.dmhy.org/"},
	}

	c := colly.NewCollector(
		colly.AllowedDomains("u2.dmhy.org"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:83.0) Gecko/20100101 Firefox/83.0"),
	)

	c.OnHTML("#outer>.main", func(e *colly.HTMLElement) {
		idx := strings.Index(e.Text, verificationCode)
		if idx != -1 {
			verified = true
		}
	})

	log.Printf("Verifying U2 User ID = %d, code = %s", u2ID, verificationCode)
	c.Request("GET", url, nil, nil, header)

	return verified
}
