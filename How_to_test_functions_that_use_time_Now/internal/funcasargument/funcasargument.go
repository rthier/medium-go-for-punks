package funcasargument

import (
	"time"
)

func WhatTimeIsIt(funcTime func() time.Time) string {
	timeIs := funcTime()
	return timeIs.String()
}
