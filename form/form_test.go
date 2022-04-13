package form

import (
	"reflect"
	"testing"

	"github.com/Bofry/structproto"
)

type DummyRequestArg struct {
	ID          string   `form:"*id"`
	Type        string   `form:"*type"`
	ShowDetail  bool     `form:"SHOW_DETAIL"`
	EnableDebug bool     `form:"ENABLE_DEBUG"`
	Tags        []string `form:"tags"`
}

func TestProcess(t *testing.T) {
	input := "id=F0003452&type=KNNS&SHOW_DETAIL&tags=T,ER,XVV"

	args := DummyRequestArg{}
	err := Process([]byte(input), &args)
	if err != nil {
		t.Error(err)
	}

	if args.ID != "F0003452" {
		t.Errorf("assert 'DummyRequestArg.ID':: expected '%v', got '%v'", "F0003452", args.ID)
	}
	if args.Type != "KNNS" {
		t.Errorf("assert 'DummyRequestArg.Type':: expected '%v', got '%v'", "KNNS", args.Type)
	}
	if args.ShowDetail != true {
		t.Errorf("assert 'DummyRequestArg.ShowDetail':: expected '%v', got '%v'", true, args.ShowDetail)
	}
	if args.EnableDebug != false {
		t.Errorf("assert 'DummyRequestArg.EnableDebug':: expected '%v', got '%v'", false, args.EnableDebug)
	}
	expectedTags := []string{"T", "ER", "XVV"}
	if !reflect.DeepEqual(args.Tags, expectedTags) {
		t.Errorf("assert 'character.Alias':: expected '%#v', got '%#v'", expectedTags, args.Tags)
	}
}

func TestProcess_WithMissingRequiredField(t *testing.T) {
	input := "id=F0003452&SHOW_DETAIL"

	args := DummyRequestArg{}
	err := Process([]byte(input), &args)

	if err == nil {
		t.Errorf("the 'ResolveQueryString()' should throw '%s' error", "with missing required symbol 'type'")
	}
	if e, ok := err.(*structproto.MissingRequiredFieldError); ok {
		if e.Field != "type" {
			t.Errorf("assert 'err.Field':: expected '%v', got '%v'", "type", e.Field)
		}
	} else {
		t.Errorf("the error except 'structprototype.MissingRequiredFieldError', got '%T'", err)
	}
}
