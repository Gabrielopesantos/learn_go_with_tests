package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name            string
		Input           interface{}
		ExpectableCalls []string
	}{
		{
			"Struct with one string field",
			Person{Name: "Chris"},
			[]string{"Chris"},
		},
		// {
		// 	"Struct with two string fields",
		// 	Person{
		// 		Name: "Chris",
		// 		Profile: Profile{
		// 			City: "London",
		// 		},
		// 	},
		// 	[]string{"Chris", "London"},
		// },
		{
			"Nested fields",
			Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Lisbon"},
			},
			[]string{"London", "Lisbon"},
		},
		{
			"Slices",
			[2]Profile{
				{33, "London"},
				{34, "Lisbon"},
			},
			[]string{"London", "Lisbon"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectableCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectableCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Lisbon"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Lisbon"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Lisbon"}
		}

		var got []string
		want := []string{"Berlin", "Lisbon"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expetected %+v to contain %q but it didn't", haystack, needle)
	}
}
