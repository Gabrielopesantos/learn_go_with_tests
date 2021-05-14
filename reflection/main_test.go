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
}
