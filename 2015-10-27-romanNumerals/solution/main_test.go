package main

import "testing"

func Test_should_write_1(t *testing.T) {
	assertEquals(Roman(1), "I", t)
}

func Test_should_write_2(t *testing.T) {
	assertEquals(Roman(2), "II", t)
}

func Test_should_write_5(t *testing.T) {
	assertEquals(Roman(5), "V", t)
}

func Test_should_write_7(t *testing.T) {
	assertEquals(Roman(7), "VII", t)
}

func Test_should_write_4(t *testing.T) {
	assertEquals(Roman(4), "IV", t)
}

func Test_should_write_12(t *testing.T) {
	assertEquals(Roman(12), "XII", t)
}

func Test_should_write_33(t *testing.T) {
	assertEquals(Roman(33), "XXXIII", t)
}

func Test_should_write_19(t *testing.T) {
	assertEquals(Roman(19), "XIX", t)
}

func Test_should_write_24(t *testing.T) {
	assertEquals(Roman(24), "XXIV", t)
}

func Test_should_write_79(t *testing.T) {
	assertEquals(Roman(79), "LXXIX", t)
}

func Test_should_write_49(t *testing.T) {
	assertEquals(Roman(49), "XLIX", t)
}

func Test_should_write_1998(t *testing.T) {
	assertEquals(Roman(1998), "MCMXCVIII", t)
}


func assertEquals(s1 string,  s2 string, t *testing.T) {
	if (s1 != s2) {
		t.Fatalf("Assertion failed: '" + s1 + "' != '" + s2 + "'")
	}
}
