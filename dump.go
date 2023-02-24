package rawebclients

import (
	"encoding/json"
	"io/ioutil"
)

// Сохранение cookies
func (cls *clients) dump() {
	cls.mu.RLock()
	jsonData, err := json.MarshalIndent(&cls.m, "", "  ")
	cls.mu.RUnlock()
	check(err)
	clientsDb.Lock()
	defer clientsDb.Unlock()
	err = ioutil.WriteFile(clientsDb.file, jsonData, 0644)
	check(err)
}
