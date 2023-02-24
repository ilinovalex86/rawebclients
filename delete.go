package rawebclients

// Удаляет web клиента, сохраняет json
func (cls *clients) delete(id string) {
	cls.mu.Lock()
	defer cls.mu.Unlock()
	delete(cls.m, id)
	go cls.dump()
}
