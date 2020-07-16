package msync

// Clocker is the public interface for Clock
type Clocker interface {
	UpdateLocal()
	UpdateCloud()
	ResetClock()
	IncrementLocal()
	IncrementCloud()
	Local() Counter
	Cloud() Counter
	CloudIsAhead() bool
	CountersAreSynced() bool
	// comparison methods
	IsNewerThan(Clocker) bool
	IsSyncedWith(Clocker) bool
	IsNewerThanOrSyncedWith(Clocker) bool
	IsConflictingWith(Clocker) bool
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

func (c *clock) IncrementLocal() {
	c.local.Increment()
}

func (c *clock) IncrementCloud() {
	c.cloud.Increment()
}

func (c *clock) Local() Counter {
	return c.local
}

func (c *clock) Cloud() Counter {
	return c.cloud
}

func (c *clock) CloudIsAhead() bool {
	return c.cloud.Get() > c.local.Get()
}

func (c *clock) LocalIsAhead() bool {
	return c.local.Get() > c.cloud.Get()
}

func (c *clock) CountersAreSynced() bool {
	return c.local.Get() == c.cloud.Get()
}

// clock comparisons

func (c *clock) IsNewerThan(c2 Clocker) bool {
	var localNewer, cloudNewer bool

	if c.local.IsNewerThan(c2.Local()) {
		localNewer = true
	}

	if c.cloud.IsNewerThan(c2.Cloud()) {
		cloudNewer = true
	}

	return localNewer && cloudNewer
}

func (c *clock) IsSyncedWith(c2 Clocker) bool {
	var localSynced, cloudSynced bool

	if c.local.IsSyncedWith(c2.Local()) {
		localSynced = true
	}

	if c.cloud.IsSyncedWith(c2.Cloud()) {
		cloudSynced = true
	}

	return localSynced && cloudSynced
}

func (c *clock) IsNewerThanOrSyncedWith(c2 Clocker) bool {
	return c.IsNewerThan(c2) || c.IsSyncedWith(c2)
}

func (c *clock) IsConflictingWith(c2 Clocker) bool {
	if c.local.IsNewerThan(c2.Local()) && c2.Cloud().IsNewerThan(c.cloud) {
		return true
	}

	return c.cloud.IsNewerThan(c2.Cloud()) && c2.Local().IsNewerThan(c.local)
}
