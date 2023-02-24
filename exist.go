package rawebclients

// Проверяет есть ли клиент с таким id
func (cls *clients) exist(id string) bool {
	cls.mu.RLock()
	defer cls.mu.RUnlock()
	for k, _ := range cls.m {
		if k == id {
			return true
		}
	}
	return false
}
