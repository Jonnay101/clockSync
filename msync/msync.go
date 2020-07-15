package msync

// Clocker is the public interface for Clock
type Clocker interface {
	updateLocal()
	updateCloud()
	resetClock()
}

type clock struct {
	local int64
	cloud int64
}

// NewClock creates a new clock
func NewClock() Clocker {
	return &clock{
		local: 0,
		cloud: 0,
	}
}

func (c *clock) updateLocal() {
	c.local++
}

func (c *clock) updateCloud() {
	c.cloud++
}

func (c *clock) resetClock() {
	c.local = 0
	c.cloud = 0
}
