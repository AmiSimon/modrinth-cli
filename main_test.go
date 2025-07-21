package main

import (
	"reflect"
	"testing"
)

func TestGetFlags(t *testing.T) {
	args := []string{"exec", "search", "phosphor", "-c", "cursed", "--category", "adventure", "--match-any", "-v", "1.21.8"}
	parsedArgs, keyword := GetFlags(args)
	wantsParsedArgs := []Flag{
		{"-c", "cursed"},
		{"--category", "adventure"},
		{"--match-any", ""},
		{"-v", "1.21.8"},
	}
	wantsKeyword := "phosphor"
	if !reflect.DeepEqual(parsedArgs, wantsParsedArgs) || wantsKeyword != keyword {
		t.Fatalf("GetFlags(args) = %v, %v, wants %v, %v", parsedArgs, keyword, wantsParsedArgs, wantsKeyword)
	}
}