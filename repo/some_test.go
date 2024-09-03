package tests

import (
	"testing"

//	"github.com/na50r/repo/action1"
	"github.com/na50r/repo/action2"
)

// var wordTable = []struct {
// 	word  string
// 	count int
// }{
// 	{"hello", 5},
// 	{"world", 5},
// 	{"test", 4},
// }

// func TestWordCountEndpoint(t *testing.T) {
// 	for _, tt := range wordTable {
// 		t.Run(tt.word, func(t *testing.T) {
// 			action1.CalcWordCount(t, tt.word, tt.count)
// 		})
// 	}
// }

var userTable = []struct {
	userId string
	name   string
	age    int
}{
	{"u1", "John", 30},
	{"u2", "Cathy", 18},
	{"u3", "Foo", 69},
}


func TestUserData(t *testing.T) {
	for _, tt := range userTable {
		t.Run(tt.userId, func(t *testing.T) {
			action2.CheckUserData(t, tt.userId, tt.name, tt.age)
		})
	}
}
