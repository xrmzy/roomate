package common

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomId(prefix string) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randomNumber := random.Intn(99999)
	return prefix + strconv.Itoa(randomNumber)
}
