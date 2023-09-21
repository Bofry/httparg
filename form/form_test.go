package form_test

import (
	"reflect"
	"testing"

	"github.com/Bofry/httparg/form"
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
	err := form.Process([]byte(input), &args)
	if err != nil {
		t.Error(err)
	}

	expected := DummyRequestArg{
		ID:          "F0003452",
		Type:        "KNNS",
		ShowDetail:  true,
		EnableDebug: false,
		Tags:        []string{"T", "ER", "XVV"},
	}

	if !reflect.DeepEqual(args, expected) {
		t.Errorf("assert DummyRequestArg:: expected '%#v', got '%#v'", expected, args)
	}
}

func TestProcess_WithMissingRequiredField(t *testing.T) {
	input := "id=F0003452&SHOW_DETAIL"

	args := DummyRequestArg{}
	err := form.Process([]byte(input), &args)

	if err == nil {
		t.Errorf("the 'ResolveQueryString()' should throw '%s' error", "with missing required symbol 'type'")
	}
	if e, ok := err.(*structproto.MissingRequiredFieldError); ok {
		if e.Field != "type" {
			t.Errorf("assert 'err.Field':: expected '%v', got '%v'", "type", e.Field)
		}
	} else {
		t.Errorf("the error expect 'structprototype.MissingRequiredFieldError', got '%T'", err)
	}
}
