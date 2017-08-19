/**
* @Author: anchen
* @Date:   2016-08-26 13:34:56
* @Last Modified by: anchen
* @Last Modified time: 2017-08-19 10:31:58
 */

package randomstrings

import (
	"math/rand"
	"sync"
	"time"
)

const (
	RANDOM_STRING_NUM     = "0123456789"
	RANDOM_STRING_LOWER   = "abcdefghijkmnpqrstuvwxyz"
	RANDOM_STRING_UPPER   = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	RANDOM_STRING_SPECIAL = "_!@#$"
)

var (
	randSeek = int64(1)
	l        sync.Mutex
)

func RandomString(n int, args interface{}) []byte {
	result := make([]byte, n)
	str := ""
	kind := 0

	if v, ok := args.(string); ok {
		str = v
	} else {
		if v, ok := args.(int); ok {
			kind = v
		}
		switch kind {
		case 0:
			str = RANDOM_STRING_NUM
		case 1:
			str = RANDOM_STRING_LOWER
		case 2:
			str = RANDOM_STRING_UPPER
		case 3:
			str = RANDOM_STRING_NUM + RANDOM_STRING_LOWER
		case 4:
			str = RANDOM_STRING_NUM + RANDOM_STRING_UPPER
		case 5:
			str = RANDOM_STRING_LOWER + RANDOM_STRING_UPPER
		case 6:
			str = RANDOM_STRING_NUM + RANDOM_STRING_LOWER + RANDOM_STRING_UPPER
		case 7:
			str = RANDOM_STRING_NUM + RANDOM_STRING_LOWER + RANDOM_STRING_UPPER + RANDOM_STRING_SPECIAL
		default:
			str = RANDOM_STRING_NUM + RANDOM_STRING_LOWER + RANDOM_STRING_UPPER
		}
	}
	size := len(str) - 1
	rand.Seed(getRandSeek())
	for i := 0; i < n; i++ {
		index := rand.Intn(size)
		result[i] = str[index]
	}
	return result
}

func RandomStringNum(n int) string {
	return string(RandomString(n, 0))
}

func RandomStringLower(n int) string {
	return string(RandomString(n, 1))
}

func RandomStringUpper(n int) string {
	return string(RandomString(n, 2))
}

func RandomStringNumLower(n int) string {
	return string(RandomString(n, 3))
}

func RandomStringNumUpper(n int) string {
	return string(RandomString(n, 4))
}

func RandomStringLowerUpper(n int) string {
	return string(RandomString(n, 5))
}

func RandomStringAll(n int) string {
	return string(RandomString(n, 6))
}

func RandomStringSpecial(n int) string {
	return string(RandomString(n, 7))
}

func getRandSeek() int64 {
	l.Lock()
	if randSeek >= 100000000 {
		randSeek = 1
	}
	randSeek++
	l.Unlock()
	return time.Now().UnixNano() + randSeek
}
