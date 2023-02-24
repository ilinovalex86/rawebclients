package rawebclients

// Store - добавляет значение tcpId с ключом id, сохраняет json
func (cls *clients) Store(id string, tcpId string) {
	cls.mu.Lock()
	defer cls.mu.Unlock()
	cls.m[id] = tcpId
	go cls.dump()
}
