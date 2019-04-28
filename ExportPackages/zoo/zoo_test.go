package zoo

import "testing"

func TestLion(t *testing.T) {
	expect := "GAH"
	actual := Lion()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestDog(t *testing.T) {
	expect := "HAHA"
	actual := Dog()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
