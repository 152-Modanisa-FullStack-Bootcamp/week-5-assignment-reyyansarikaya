package assignment

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {
	tests := []struct {
		name  string
		x     uint32
		y     uint32
		want  uint32
		want1 bool
	}{
		{"test1", math.MaxUint32, 1, 0, true},
		{"test2", 1, 1, 2, false},
		{"test3", 42, 2701, 2743, false},
		{"test4", 42, math.MaxUint32, 41, true},
		{"test5", 4294967290, 5, 4294967295, false},
		{"test6", 4294967290, 6, 0, true},
		{"test7", 4294967290, 10, 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AddUint32(tt.x, tt.y)
			assert.Equalf(t, tt.want, got, "AddUint32(%v, %v)", tt.x, tt.y)
			assert.Equalf(t, tt.want1, got1, "AddUint32(%v, %v)", tt.x, tt.y)
		})
	}
}

func TestCeilNumber(t *testing.T) {
	tests := []struct {
		name string
		f    float64
		want float64
	}{
		{"test1", 42.42, 42.50},
		{"test2", 42, 42},
		{"test3", 42.01, 42.25},
		{"test4", 42.24, 42.25},
		{"test5", 42.25, 42.25},
		{"test6", 42.26, 42.50},
		{"test7", 42.55, 42.75},
		{"test8", 42.75, 42.75},
		{"test9", 42.76, 43},
		{"test10", 42.99, 43},
		{"test11", 43.13, 43.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CeilNumber(tt.f), "CeilNumber(%v)", tt.f)
		})
	}
}

func TestAlphabetSoup(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test1", "hello", "ehllo"},
		{"test2", "", ""},
		{"test3", "h", "h"},
		{"test4", "ab", "ab"},
		{"test5", "ba", "ab"},
		{"test6", "bac", "abc"},
		{"test7", "cba", "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AlphabetSoup(tt.s), "AlphabetSoup(%v)", tt.s)
		})
	}
}

func TestStringMask(t *testing.T) {
	tests := []struct {
		name string
		s    string
		n    uint
		want string
	}{
		{"test1", "!mysecret*", 2, "!m********"},
		{"test2", "", 4, "*"},
		{"test3", "a", 1, "*"},
		{"test4", "string", 0, "******"},
		{"test5", "string", 3, "str***"},
		{"test6", "string", 5, "strin*"},
		{"test7", "string", 6, "******"},
		{"test8", "string", 7, "******"},
		{"test9", "s*r*n*", 3, "s*r***"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringMask(tt.s, tt.n), "StringMask(%v, %v)", tt.s, tt.n)
		})
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	tests := []struct {
		name string
		arr  [2]string
		want string
	}{
		{"test1", [2]string{"hellocat", words}, "hello,cat"},
		{"test2", [2]string{"catbat", words}, "cat,bat"},
		{"test3", [2]string{"yellowapple", words}, "yellow,apple"},
		{"test4", [2]string{"", words}, "not possible"},
		{"test5", [2]string{"notcat", words}, "not possible"},
		{"test6", [2]string{"bootcamprocks!", words}, "not possible"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, WordSplit(tt.arr), "WordSplit(%v)", tt.arr)
		})
	}
}

func TestVariadicSet(t *testing.T) {
	// FINAL BOSS ALERT :)
	//	Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )
	tests := []struct {
		name string
		i    []interface{}
		want []interface{}
	}{
		{"test1", VariadicSet(4, 2, 5, 4, 2, 4), []interface{}{4, 2, 5}},
		{"test2", VariadicSet("bootcamp", "rocks!", "really", "rocks!"), []interface{}{"bootcamp", "rocks!", "really"}},
		{"test3", VariadicSet(1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"), []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, VariadicSet(tt.i...), "VariadicSet(%v)", tt.i...)
		})
	}
}

var table = []struct {
	name string
	s    string
	n    uint
	want string
}{
	{"test1", "!mysecret*", 2, "!m********"},
	{"test2", "", 4, "*"},
	{"test3", "a", 1, "*"},
	{"test4", "string", 0, "******"},
	{"test5", "string", 3, "str***"},
	{"test6", "string", 5, "strin*"},
	{"test7", "string", 6, "******"},
	{"test8", "string", 7, "******"},
	{"test9", "s*r*n*", 3, "s*r***"},
}

func BenchmarkStringMask(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%T", v.s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringMask(v.s, v.n)
			}
		})
	}
}

func BenchmarkStringMask_SadumanSolve(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%T", v.s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringMask_SadumanSolve(v.s, v.n)
			}
		})
	}
}
