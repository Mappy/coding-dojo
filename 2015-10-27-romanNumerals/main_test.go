package main

import "testing"

func Test_should_write_1(t *testing.T) {
	assertEquals(Roman(1), "I", t)
}

func Test_should_write_3(t *testing.T) {
	assertEquals(Roman(3), "III", t)
}

func Test_should_write_4(t *testing.T) {
	assertEquals(Roman(4), "IV", t)
}

func Test_should_write_5(t *testing.T) {
	assertEquals(Roman(5), "V", t)
}

func Test_should_write_6(t *testing.T) {
	assertEquals(Roman(6), "VI", t)
}

func Test_should_write_9(t *testing.T) {
	assertEquals(Roman(9), "IX", t)
}

func Test_should_write_10(t *testing.T) {
	assertEquals(Roman(10), "X", t)
}

func Test_should_write_13(t *testing.T) {
	assertEquals(Roman(13), "XIII", t)
}

func Test_should_write_16(t *testing.T) {
	assertEquals(Roman(16), "XVI", t)
}

func Test_should_write_20(t *testing.T) {
	assertEquals(Roman(20), "XX", t)
}

func Test_should_write_50(t *testing.T) {
	assertEquals(Roman(50), "L", t)
}

func Test_should_write_100(t *testing.T) {
	assertEquals(Roman(100), "C", t)
}

func Test_should_write_490(t *testing.T) {
	assertEquals(Roman(490), "CDXC", t)
}


func Test_should_write_501(t *testing.T) {
	assertEquals(Roman(501), "DI", t)
}

func Test_should_write_96(t *testing.T) {
	assertEquals(Roman(96), "XCVI", t)
}

func Test_should_write_1000(t *testing.T) {
	assertEquals(Roman(1000), "DI", t)
}


func assertEquals(s1 string, s2 string, t *testing.T) {
	if (s1 != s2) {
		t.Fatalf("Assertion failed: '" + s1 + "' != '" + s2 + "'")
	}
}
