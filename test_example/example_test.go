package main

import "testing"

type TestCase struct {
	RoomType       string
	PersonCount    int
	Nights         int
	ExpectedResult int
}

var TestCases []TestCase = []TestCase{
	{RoomType: "standard", PersonCount: 2, Nights: 3, ExpectedResult: 1200000},
	{RoomType: "double", PersonCount: 2, Nights: 3, ExpectedResult: 1500000},
	{RoomType: "vip", PersonCount: 2, Nights: 3, ExpectedResult: 1800000},
	{RoomType: "standard", PersonCount: 2, Nights: 2, ExpectedResult: 800000},
}

//tdd: test driven development
func TestCalculateRoomPrice(t *testing.T) {
	actual := 0

	for _, v := range TestCases {
		actual = CalculateRoomPrice(v.RoomType, v.Nights, v.PersonCount)
		if actual != v.ExpectedResult {
			// t.Fail()
			t.Errorf("wrong result. actual:%v , expected: %v\n", actual, v.ExpectedResult)
		}
	}
}

func FuzzCalculateRoomPrice(f *testing.F) {
	f.Fuzz(func(t *testing.T, room_type string, nights int, person_count int) {
		CalculateRoomPrice(room_type, nights, person_count)
	})
}
