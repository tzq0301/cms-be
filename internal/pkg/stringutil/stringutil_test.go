package stringutil

import (
	"reflect"
	"testing"
)

func TestFromBytes(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{[]byte("")}, ""},
		{"", args{[]byte("123abd")}, "123abd"},
		{"", args{[]byte("你好")}, "你好"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromBytes(tt.args.b); got != tt.want {
				t.Errorf("FromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{""}, true},
		{"", args{"\n"}, true},
		{"", args{"\n\t"}, true},
		{"", args{" nsd "}, false},
		{"", args{" nsd你好 "}, false},
		{"", args{" nsd你好\t"}, false},
		{"", args{"\nnsd你好 "}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.args.str); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantB []byte
	}{
		{"", args{""}, nil},
		{"", args{"123abd"}, []byte("123abd")},
		{"", args{"你好"}, []byte("你好")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := ToBytes(tt.args.s); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ToBytes() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
