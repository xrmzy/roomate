package common

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateID(prefix string) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randomNumber := random.Intn(99999) // Bisa disesuaikan dengan kebutuhan tim nanti

	return prefix + strconv.Itoa(randomNumber)
}
