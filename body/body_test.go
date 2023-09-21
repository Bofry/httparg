package body_test

import (
	"reflect"
	"testing"

	"github.com/Bofry/httparg/body"
)

type DummyRequestArg struct {
	TextBody  string `body:"*text_body"`
	BytesBody []byte `body:"*bytes_body"`
}

func TestProcess(t *testing.T) {
	input := "id=F0003452&type=KNNS&SHOW_DETAIL&tags=T,ER,XVV"

	args := DummyRequestArg{}
	err := body.Process([]byte(input), &args)
	if err != nil {
		t.Error(err)
	}

	expected := DummyRequestArg{
		TextBody:  "id=F0003452&type=KNNS&SHOW_DETAIL&tags=T,ER,XVV",
		BytesBody: []byte("id=F0003452&type=KNNS&SHOW_DETAIL&tags=T,ER,XVV"),
	}

	if !reflect.DeepEqual(args, expected) {
		t.Errorf("assert DummyRequestArg:: expected '%#v', got '%#v'", expected, args)
	}
}
