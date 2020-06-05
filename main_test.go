package main

import (
	"testing"
)

func TestBar(t *testing.T) {
	result := Bar()
	if result != "bar" {
		t.Errorf("expecting bar, got %s", result)
	}
}

func TestFoo(t *testing.T) {
	result := Foo()
	if result != "foo" {
		t.Errorf("expecting foo, got %s", result)
	}
}

func TestQuz(t *testing.T) {
	result := Qux("foo")
	if result != "foo" {
		t.Errorf("expecting foo, got %s", result)
	}

	result = Qux("bar")
	if result != "bar" {
		t.Errorf("expecting bar, got %s", result)
	}

	result = Qux("qux")
	if result != "INVALID" {
		t.Errorf("expecting INVALID, got %s", result)
	}
}
