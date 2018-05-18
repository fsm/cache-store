package cachestore

type cachedTraverser struct {
	uuid         string
	currentState string
	platform     string
	Data         map[string]interface{}
}

// ===== UUID =====

func (c *cachedTraverser) UUID() string {
	return c.uuid
}

func (c *cachedTraverser) SetUUID(newUUID string) {
	c.uuid = newUUID
}

// ===== Platform =====

func (c *cachedTraverser) Platform() string {
	return c.platform
}

func (c *cachedTraverser) SetPlatform(platform string) {
	c.platform = platform
}

// ===== State =====

func (c *cachedTraverser) CurrentState() string {
	return c.currentState
}

func (c *cachedTraverser) SetCurrentState(newState string) {
	c.currentState = newState
}

// ===== Other Data =====

func (c *cachedTraverser) Upsert(key string, value interface{}) {
	c.Data[key] = value
}

func (c *cachedTraverser) Fetch(key string) interface{} {
	if val, ok := c.Data[key]; ok {
		return val
	}
	return nil
}

func (c *cachedTraverser) Delete(key string) {
	delete(c.Data, key)
}
