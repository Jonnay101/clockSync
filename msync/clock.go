package msync

// Clocker is the public interface for Clock
type Clocker interface {
	SetLocal(int64)
	SetCloud(int64)
	ResetClock()
	IncrementLocal()
	IncrementCloud()
	Local() Counter
	Cloud() Counter
	LocalIsAhead() bool
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

// SetLocal sets the local counter to the given value
func (c *clock) SetLocal(num int64) {
	c.local.Set(num)
}

// SetCloud sets the cloud counter to the given value
func (c *clock) SetCloud(num int64) {
	c.cloud.Set(num)
}

// ResetClock resets both counters to zero value
func (c *clock) ResetClock() {
	c.local.Set(0)
	c.cloud.Set(0)
}

// IncrementLocal adds 1 to the value of the local counter
func (c *clock) IncrementLocal() {
	c.local.Increment()
}

// IncrementCloud adds 1 to the value of the cloud counter
func (c *clock) IncrementCloud() {
	c.cloud.Increment()
}

// Local returns the local counter
func (c *clock) Local() Counter {
	return c.local
}

// Cloud returns the cloud counter
func (c *clock) Cloud() Counter {
	return c.cloud
}

// Local is ahead returns true if local counter value is greater than cloud counter value
func (c *clock) LocalIsAhead() bool {
	return c.local.Get() > c.cloud.Get()
}

// Cloud is ahead returns true if cloud counter value is greater than local counter value
func (c *clock) CloudIsAhead() bool {
	return c.cloud.Get() > c.local.Get()
}

// CountersAreSynced returns true if both counter values are the equal
func (c *clock) CountersAreSynced() bool {
	return c.local.Get() == c.cloud.Get()
}

// CLOCK COMPARISONS

// IsNewerThan compares the clock to another clock to see which is 'newest'
func (c *clock) IsNewerThan(c2 Clocker) bool {
	var localNewer, cloudNewer bool

	if c.local.IsNewerThan(c2.Local()) {
		localNewer = true
	}

	if c.cloud.IsNewerThan(c2.Cloud()) {
		cloudNewer = true
	}

	if localNewer && !cloudNewer {
		if c.cloud.IsSyncedWith(c2.Cloud()) {
			cloudNewer = true
		}
	}

	if cloudNewer && !localNewer {
		if c.local.IsSyncedWith(c2.Local()) {
			localNewer = true
		}
	}

	return localNewer && cloudNewer
}

// IsSyncedWith checks if both the clocks counter values are equal
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

// IsNewerThanOrSyncedWith returns true if a clock is either newer or sync with another
func (c *clock) IsNewerThanOrSyncedWith(c2 Clocker) bool {
	return c.IsNewerThan(c2) || c.IsSyncedWith(c2)
}

// IsConflictingWith returns true if the clocks have conflicting states
// eg. Cloud ahead on one clock and Local ahead on the other
func (c *clock) IsConflictingWith(c2 Clocker) bool {
	conflict1 := c.local.IsNewerThan(c2.Local()) && c2.Cloud().IsNewerThan(c.cloud)

	conflict2 := c.cloud.IsNewerThan(c2.Cloud()) && c2.Local().IsNewerThan(c.local)

	return conflict1 || conflict2
}
