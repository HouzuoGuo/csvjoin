package main

import (
	"reflect"
	"testing"
)

func TestJoin(t *testing.T) {
	left := [][]string{
		[]string{"a", "1", "foo"},
		[]string{"b", "2", "bar"},
	}
	right := [][]string{
		[]string{"alpha", "1", "zulu"},
		[]string{"beta", "2", "yankee"},
	}

	expected := [][]string{
		[]string{"a", "1", "foo", "alpha", "zulu"},
		[]string{"b", "2", "bar", "beta", "yankee"},
	}
	actual := Join(left, right, 1, 1, false)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected\n%+v\nActual\n%+v\n", expected, actual)
	}
}

func TestShortJoin(t *testing.T) {
	left := [][]string{
		[]string{"a", "1", "foo"},
		[]string{"b", "2", "bar"},
	}
	right := [][]string{
		[]string{"alpha", "1", "zulu"},
	}

	expected := [][]string{
		[]string{"a", "1", "foo", "alpha", "zulu"},
		[]string{"b", "2", "bar"},
	}
	actual := Join(left, right, 1, 1, false)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected\n%+v\nActual\n%+v\n", expected, actual)
	}
}

func TestInsensitiveJoin(t *testing.T) {
	left := [][]string{
		[]string{"A", "1", "foo"},
		[]string{" B ", "2", "bar"},
	}
	right := [][]string{
		[]string{"  a  ", "1", "zulu"},
		[]string{"   b   ", "2", "yankee"},
	}

	expected := [][]string{
		[]string{"A", "1", "foo", "1", "zulu"},
		[]string{" B ", "2", "bar", "2", "yankee"},
	}
	actual := Join(left, right, 0, 0, true)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected\n%+v\nActual\n%+v\n", expected, actual)
	}
}
