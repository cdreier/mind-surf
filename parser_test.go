package main

import (
	"io/ioutil"
	"testing"
)

func Test_parse(t *testing.T) {
	in, _ := ioutil.ReadFile("./debug.txt")

	rootItem := parse(string(in))
	if rootItem.Name != "mind surf!" {
		t.Error("wrong text in root item", rootItem.Name)
	}
	if len(rootItem.Children) != 3 {
		t.Error("wrong number of childrens", len(rootItem.Children))
	}

	if len(rootItem.Children[1].Children) != 2 {
		t.Error("wrong number of childrens for level 2", len(rootItem.Children[1].Children))
	}
}

func Test_getLevel(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "level root",
			want: 0,
			args: args{line: "root"},
		},
		{
			name: "first level",
			want: 1,
			args: args{line: "  first level with spaces"},
		},
		{
			name: "first level with spaces",
			want: 1,
			args: args{line: "  first   level with spaces"},
		},
		{
			name: "second level",
			want: 2,
			args: args{line: "    second level"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLevel(tt.args.line); got != tt.want {
				t.Errorf("getLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
