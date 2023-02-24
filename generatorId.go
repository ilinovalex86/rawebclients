package rawebclients

import (
	"math/rand"
	"strconv"
	"time"
)

// GeneratorId - Генерирует уникальный id для web клиента
func GeneratorId() string {
	rand.Seed(time.Now().UnixNano())
	for {
		id := strconv.Itoa(rand.Int())
		if len(id) < 19 {
			continue
		}
		if Clients.exist(id) {
			continue
		} else {
			return id
		}
	}
}
