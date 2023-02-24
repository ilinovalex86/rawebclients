package rawebclients

// Load - возвращает значение и статус по id
func (cls *clients) Load(id string) (string, bool) {
	cls.mu.RLock()
	defer cls.mu.RUnlock()
	tcpId, ok := cls.m[id]
	return tcpId, ok
}
