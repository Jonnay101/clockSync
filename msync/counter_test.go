package msync

import (
	"reflect"
	"testing"
)

func TestNewCounter(t *testing.T) {

	cntNum := counter(0)
	testCounter := &cntNum

	tests := []struct {
		name string
		want Counter
	}{
		{"create new counter", testCounter},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCounter()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCounter() = %v, want %v", got, tt.want)
			}

			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("NewCounter() type is %v, want %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
			}
		})
	}
}

func Test_counter_Get(t *testing.T) {

	cntNum := counter(3)
	testCounter := &cntNum

	tests := []struct {
		name string
		cnt  *counter
		want int64
	}{
		{"should get 3... the magic number", testCounter, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cnt.Get()

			if got != tt.want {
				t.Errorf("counter.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_Set(t *testing.T) {

	testCounter := new(counter)

	type args struct {
		num int64
	}
	tests := []struct {
		name string
		cnt  *counter
		args args
		want int64
	}{
		{"should set to 23", testCounter, args{num: 23}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cnt.Set(tt.args.num)

			if got := tt.cnt.Get(); got != tt.want {
				t.Errorf("counter value = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_Increment(t *testing.T) {

	testCounter := NewCounter()
	testCounter.Set(4)

	tests := []struct {
		name string
		cnt  Counter
		want int64
	}{
		{"increment from 4 to 5", testCounter, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cnt.Increment()

			if got := tt.cnt.Get(); got != tt.want {
				t.Errorf("counter value = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_IsSyncedWith(t *testing.T) {

	tests := []struct {
		name string
		num1 int64
		num2 int64
		want bool
	}{
		{"are in sync", 3, 3, true},
		{"are out of sync", 3, 7, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testCounter1 := NewCounter()
			testCounter1.Set(tt.num1)

			testCounter2 := NewCounter()
			testCounter2.Set(tt.num2)

			if got := testCounter1.IsSyncedWith(testCounter2); got != tt.want {
				t.Errorf("counter.IsSyncedWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_IsNewerThan(t *testing.T) {

	tests := []struct {
		name string
		num1 int64
		num2 int64
		want bool
	}{
		{"is newer than", 5, 1, true},
		{"is the same as", 567, 567, false},
		{"is not newer than", 4, 9, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testCounter1 := NewCounter()
			testCounter1.Set(tt.num1)

			testCounter2 := NewCounter()
			testCounter2.Set(tt.num2)

			if got := testCounter1.IsNewerThan(testCounter2); got != tt.want {
				t.Errorf("counter.IsNewerThan() = %v, want %v", got, tt.want)
			}
		})
	}
}
