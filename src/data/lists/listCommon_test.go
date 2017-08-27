package lists

import "testing"

type fakeDummy struct {
	a int
}

var fakeA, fakeB, fakeC = 1, 1, 1

var fakeTable = [][]interface{}{
	{1.0, 2.2, 3.14},
	{-1000, 0, 1000},
	{"str1", "str2"},
	{true, false, true},
	{fakeDummy{1}, fakeDummy{2}},
	{1, true, "str"},
	{nil},
	{nil, nil},
	{nil, nil, nil, nil, nil},
	{&fakeA, &fakeB, &fakeC},
}

func helperInitIsEmpty(listType string, l ListCommon, t *testing.T) {
	if l.IsEmpty() == false {
		t.Errorf("new %v isEmpty() expected true, got false", listType)
	}

	if l.HasElement() {
		t.Errorf("new %v hasElement(), it should be empty", listType)
	}

	len := l.Len()

	if len != 0 {
		t.Errorf("new %s len expected 0, got %v", listType, len)
	}
}
