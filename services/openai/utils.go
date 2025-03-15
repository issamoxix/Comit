package ai

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateComitId() string {
	now := time.Now()
	timestamp := fmt.Sprintf("%d", now.Unix())
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)
	suffix := make([]byte, 3)
	for i := range suffix {
		suffix[i] = letters[r.Intn(len(letters))]
	}
	return timestamp + string(suffix)
}
