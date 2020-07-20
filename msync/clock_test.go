package msync

import (
	"reflect"
	"testing"
)

func TestNewClock(t *testing.T) {

	wantClock := &clock{
		local: NewCounter(),
		cloud: NewCounter(),
	}

	tests := []struct {
		name string
		want Clocker
	}{
		{"should create a clock", wantClock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clock_SetLocal(t *testing.T) {

	wantClock := &clock{NewCounter(), NewCounter()}
	wantClock.local.Set(3)

	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want Clocker
	}{
		{"local should be set to 3", args{num: 3}, wantClock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClock()
			c.SetLocal(tt.args.num)

			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("wanted %v, got %v", c, tt.want)
			}
		})
	}
}

func Test_clock_SetCloud(t *testing.T) {

	wantClock := &clock{NewCounter(), NewCounter()}
	wantClock.cloud.Set(3)

	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want Clocker
	}{
		{"local should be set to 3", args{num: 3}, wantClock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := NewClock()
			c.SetCloud(tt.args.num)

			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("wanted %v, got %v", c, tt.want)
			}
		})
	}
}

func Test_clock_ResetClock(t *testing.T) {

	wantClock := &clock{NewCounter(), NewCounter()}

	tests := []struct {
		name string
		want Clocker
	}{
		{"clock should be reset", wantClock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClock()
			c.ResetClock()

			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("wanted %v, got %v", c, tt.want)
			}
		})
	}
}

func Test_clock_IncrementLocal(t *testing.T) {

	wantClock := &clock{NewCounter(), NewCounter()}
	wantClock.local.Set(1)

	tests := []struct {
		name string
		want Clocker
	}{
		{"increment local counter", wantClock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClock()
			c.IncrementLocal()

			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("wanted %v, got %v", c, tt.want)
			}
		})
	}
}

func Test_clock_IncrementCloud(t *testing.T) {

	wantClock := &clock{NewCounter(), NewCounter()}
	wantClock.cloud.Set(1)

	tests := []struct {
		name string
		want Clocker
	}{
		{"should increment cloud", wantClock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClock()
			c.IncrementCloud()

			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("wanted %v, got %v", c, tt.want)
			}
		})
	}
}

func Test_clock_Local(t *testing.T) {

	wantCounter := NewCounter()
	wantCounter.Set(3)

	tests := []struct {
		name string
		num  int64
		want Counter
	}{
		{"should return counter val of 3", 3, wantCounter},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClock()
			c.SetLocal(tt.num)

			if got := c.Local(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clock.Local() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clock_Cloud(t *testing.T) {

	wantCounter := NewCounter()
	wantCounter.Set(5)

	tests := []struct {
		name string
		num  int64
		want Counter
	}{
		{"", 5, wantCounter},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClock()
			c.SetCloud(tt.num)

			if got := c.Cloud(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clock.Cloud() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clock_LocalIsAhead(t *testing.T) {

	tests := []struct {
		name     string
		localNum int64
		cloudNum int64
		want     bool
	}{
		{"local is ahead", 3, 2, true},
		{"cloud is ahead", 3, 4, false},
		{"counters are synced", 3, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClock()

			c.SetLocal(tt.localNum)
			c.SetCloud(tt.cloudNum)

			if got := c.LocalIsAhead(); got != tt.want {
				t.Errorf("clock.LocalIsAhead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clock_CloudIsAhead(t *testing.T) {

	tests := []struct {
		name     string
		localNum int64
		cloudNum int64
		want     bool
	}{
		{"cloud is ahead", 3, 4, true},
		{"local is ahead", 3, 2, false},
		{"counters are synced", 3, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClock()

			c.SetLocal(tt.localNum)
			c.SetCloud(tt.cloudNum)

			if got := c.CloudIsAhead(); got != tt.want {
				t.Errorf("clock.CloudIsAhead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clock_CountersAreSynced(t *testing.T) {

	tests := []struct {
		name     string
		localNum int64
		cloudNum int64
		want     bool
	}{
		{"counters are synced", 3, 3, true},
		{"local is ahead", 3, 2, false},
		{"cloud is ahead", 3, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClock()

			c.SetLocal(tt.localNum)
			c.SetCloud(tt.cloudNum)

			if got := c.CountersAreSynced(); got != tt.want {
				t.Errorf("clock.CountersAreSynced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clock_IsAheadOf(t *testing.T) {

	newerClock := NewClock()
	newerClock2 := NewClock()
	olderClock := NewClock()

	// cloud is newer
	newerClock.SetLocal(2)
	newerClock.SetCloud(3)

	// local is newer
	newerClock2.SetLocal(3)
	newerClock2.SetCloud(2)

	olderClock.SetLocal(2)
	olderClock.SetCloud(2)

	type args struct {
		c2 Clocker
	}
	tests := []struct {
		name string
		c    Clocker
		args args
		want bool
	}{
		{"cloud is newer", newerClock, args{olderClock}, true},
		{"local is newer", newerClock2, args{olderClock}, true},
		{"is not newer", olderClock, args{newerClock}, false},
		{"are the same", olderClock, args{olderClock}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsAheadOf(tt.args.c2); got != tt.want {
				t.Errorf("clock.IsAheadOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clock_IsSyncedWith(t *testing.T) {

	newerClock := NewClock()
	newerClock2 := NewClock()
	olderClock := NewClock()

	// cloud is newer
	newerClock.SetLocal(2)
	newerClock.SetCloud(3)

	// local is newer
	newerClock2.SetLocal(3)
	newerClock2.SetCloud(2)

	olderClock.SetLocal(2)
	olderClock.SetCloud(2)

	type args struct {
		c2 Clocker
	}
	tests := []struct {
		name string
		c    Clocker
		args args
		want bool
	}{

		{"are the same", olderClock, args{olderClock}, true},
		{"cloud is newer", newerClock, args{olderClock}, false},
		{"local is newer", newerClock2, args{olderClock}, false},
		{"is behind", olderClock, args{newerClock}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsSyncedWith(tt.args.c2); got != tt.want {
				t.Errorf("clock.IsSyncedWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
