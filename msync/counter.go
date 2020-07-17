package msync

// Counter -
type Counter interface {
	Get() int64
	Set(int64)
	Increment()
	IsSyncedWith(cnt2 Counter) bool
	IsNewerThan(cnt2 Counter) bool
}

type counter int64

// NewCounter returns a new zero value counter
func NewCounter() Counter {
	cntConv := counter(0)
	return &cntConv
}

// Get returns the int64 value of the counter
func (cnt *counter) Get() int64 {
	return int64(*cnt)
}

// Set the value of the counter
func (cnt *counter) Set(num int64) {
	*cnt = counter(num)
}

func (cnt *counter) Increment() {
	plusOne := cnt.Get() + 1
	*cnt = counter(plusOne)
}

func (cnt *counter) IsSyncedWith(cnt2 Counter) bool {
	return cnt.Get() == cnt2.Get()
}

func (cnt *counter) IsNewerThan(cnt2 Counter) bool {
	return cnt.Get() > cnt2.Get()
}
