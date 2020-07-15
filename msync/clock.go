package msync

// Clocker is the public interface for Clock
type Clocker interface {
	UpdateLocal()
	UpdateCloud()
	ResetClock()
}

type clock struct {
	local Counter
	cloud Counter
}

// NewClock creates a new clock
func NewClock() Clocker {
	return &clock{
		local: NewCounter(),
		cloud: NewCounter(),
	}
}

func (c *clock) UpdateLocal() {
	c.local.Increment()
}

func (c *clock) UpdateCloud() {
	c.cloud.Increment()
}

func (c *clock) ResetClock() {
	c.local.Set(0)
	c.cloud.Set(0)
}

func (c *clock) IsNewerThan(c2 *clock) bool {
	var localNewer, cloudNewer bool

	if c.local.IsNewerThan(c2.local) {
		localNewer = true
	}

	if c.cloud.IsNewerThan(c2.cloud) {
		cloudNewer = true
	}

	return localNewer && cloudNewer
}

func (c *clock) IsNewerThanOrSyncedWith(c2 *clock) bool {
	var localNewerOrSynced, cloudNewerOrSynced bool

	if c.local.IsNewerThan(c2.local) {
		localNewerOrSynced = true
	}

	if c.local.IsSyncedWith(c2.local) {
		localNewerOrSynced = true
	}

	if c.cloud.IsNewerThan(c2.cloud) {
		cloudNewerOrSynced = true
	}

	if c.cloud.IsSyncedWith(c2.cloud) {
		cloudNewerOrSynced = true
	}

	return localNewerOrSynced && cloudNewerOrSynced
}
