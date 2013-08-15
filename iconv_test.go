package iconv

import (
	"testing"
)

func TestIconvOpenClose(t *testing.T) {
	i, err := Open("GB18030", "UTF-8")
	if err != nil {
		t.Fatalf("Open Error: %s\n", err.Error())
	}
	err = i.Close()
	if err != nil {
		t.Fatalf("Close Error: %s\n", err.Error())
	}
}

func TestConvEmptyString(t *testing.T) {
	i, err := Open("GB18030", "UTF-8")
	if err != nil {
		t.Fatalf("Open Error: %s\n", err.Error())
	}
	defer i.Close()
	s, err := i.ConvString("")
	if err != nil {
		t.Fatalf("Conv Empty String Error: %s\n", err.Error())
	}
	if s != "" {
		t.Fatalf("Conv Empty String Fail: %s\n", s)
	}
}

func TestConvGB18030(t *testing.T) {
	i, err := Open("GB18030", "UTF-8")
	if err != nil {
		t.Fatalf("Open Error: %s\n", err.Error())
	}
	defer i.Close()
	gb := "\xd6\xd0\xce\xc4"
	utf8, err := i.ConvString(gb)
	if err != nil {
		t.Fatalf("Conv GB18030 to UTF-8 Fail: %s\n", err.Error())
	}
	if utf8 != "\xe4\xb8\xad\xe6\x96\x87" {
		t.Fatalf("Conv GB18030 to UTF-8 Fail: %s\n", utf8)
	}
}
