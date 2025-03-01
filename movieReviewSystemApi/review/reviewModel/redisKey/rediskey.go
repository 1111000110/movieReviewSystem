package rediskey

import "strconv"

func ReviewKey(baseId int64) string {
	return "review:baseId:" + strconv.FormatInt(baseId, 10)
}
