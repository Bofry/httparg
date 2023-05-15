package json

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Bofry/structproto"
)

type DummyRequestArg struct {
	ID          string           `json:"*id"`
	Type        *string          `json:"*type"`
	Number      int64            `json:"number"`
	ShowDetail  bool             `json:"showDetail"`
	EnableDebug bool             `json:"enableDebug"`
	Tags        []string         `json:"tags"`
	Raw         *json.RawMessage `json:"raw"`

	Detail    *DummyRequestArgDetail `json:"detail"`
	ExtraInfo map[string]interface{} `json:"extraInfo"`
}

type DummyRequestArgDetail struct {
	Operator string `json:"operator"`
}

func TestProcess_WithStruct(t *testing.T) {
	input := `{
		"id": "F0003452",
		"type": "KNNS",
		"number": 280123412341234123,
		"showDetail": true,
		"tags": ["T","ER","XVV"],
		"detail": {
			"operator": "nami"
		},
		"extraInfo": {
			"alias": "Cat Burglar",
			"age"  : 18
		}
	}`

	arg := DummyRequestArg{}
	err := Process([]byte(input), &arg)
	if err != nil {
		t.Error(err)
	}

	var expectedID string = "F0003452"
	if arg.ID != expectedID {
		t.Errorf("assert 'DummyRequestArg.ID':: expected '%v', got '%v'", expectedID, arg.ID)
	}
	var expectedType string = "KNNS"
	if *arg.Type != expectedType {
		t.Errorf("assert 'DummyRequestArg.Type':: expected '%v', got '%v'", expectedType, arg.Type)
	}
	var expectedNumber int64 = 280123412341234123
	if arg.Number != expectedNumber {
		t.Errorf("assert 'DummyRequestArg.Number':: expected '%v', got '%v'", expectedNumber, arg.Number)
	}
	var expectedShowDetail bool = true
	if arg.ShowDetail != expectedShowDetail {
		t.Errorf("assert 'DummyRequestArg.ShowDetail':: expected '%v', got '%v'", expectedShowDetail, arg.ShowDetail)
	}
	var expectedEnableDebug bool = false
	if arg.EnableDebug != expectedEnableDebug {
		t.Errorf("assert 'DummyRequestArg.EnableDebug':: expected '%v', got '%v'", expectedEnableDebug, arg.EnableDebug)
	}
	var expectedTags = []string{"T", "ER", "XVV"}
	if !reflect.DeepEqual(arg.Tags, expectedTags) {
		t.Errorf("assert 'character.Alias':: expected '%#v', got '%#v'", expectedTags, arg.Tags)
	}
	{
		if arg.Detail == nil {
			t.Error("assert 'DummyRequestArg.Detail':: should not be nil")
		}
		var detail = arg.Detail
		if detail.Operator != "nami" {
			t.Errorf("assert 'DummyRequestArg.Detail.Operator':: expected '%v', got '%v'", "nami", detail.Operator)
		}
	}
	var expectedExtraInfo map[string]interface{} = map[string]interface{}{
		"alias": "Cat Burglar",
		"age":   json.Number("18"),
	}
	if arg.ExtraInfo == nil || len(arg.ExtraInfo) != len(expectedExtraInfo) {
		t.Errorf("assert 'DummyRequestArg.ExtraInfo':: expected '%+v', got '%+v'", expectedExtraInfo, arg.ExtraInfo)
	}
	{
		for k, v := range expectedExtraInfo {
			expectedExtraInfoValue := v
			actualExtraInfoValue, ok := arg.ExtraInfo[k]
			if !ok {
				t.Errorf("assert 'DummyRequestArg.ExtraInfo':: missing key '%s'", k)
			}
			if !reflect.DeepEqual(actualExtraInfoValue, expectedExtraInfoValue) {
				t.Errorf("assert 'DummyRequestArg.ExtraInfo[%s]':: expected '%+v', got '%+v'", k, expectedExtraInfoValue, actualExtraInfoValue)
			}
		}
	}
}

func TestProcess_WithStructError(t *testing.T) {
	input := `{
		"id": "F0003452",
		"type": null,
		"number": 280123412341234123,
		"showDetail": true,
		"tags": ["T","ER","XVV"],
		"detail": {
			"operator": "nami"
		}
	}`

	arg := DummyRequestArg{}
	err := Process([]byte(input), &arg)
	if err == nil {
		t.Errorf("the 'Process()' should throw '%s' error", "missing required symbol 'type'")
	} else {
		missingRequiredFieldError, ok := err.(*structproto.MissingRequiredFieldError)
		if !ok {
			t.Errorf("the error expected '%T', got '%T'", &structproto.MissingRequiredFieldError{}, err)
		}
		if missingRequiredFieldError.Field != "type" {
			t.Errorf("assert 'MissingRequiredFieldError.Field':: expected '%v', got '%v'", "type", missingRequiredFieldError.Field)
		}
	}
}

func TestProcess_WithArray(t *testing.T) {
	input := `[{
		"id": "F0003452",
		"type": "KNNS",
		"number": 280123412341234123,
		"showDetail": true,
		"tags": ["T","ER","XVV"],
		"detail": {
			"operator": "nami"
		}
	}]`

	args := []DummyRequestArg{}
	err := Process([]byte(input), &args)
	if err != nil {
		t.Error(err)
	}

	if len(args) != 1 {
		t.Errorf("assert 'len([]DummyRequestArg{})':: expected '%v', got '%v'", 1, len(args))
	}

	arg := args[0]

	if arg.ID != "F0003452" {
		t.Errorf("assert 'DummyRequestArg.ID':: expected '%v', got '%v'", "F0003452", arg.ID)
	}
	if *arg.Type != "KNNS" {
		t.Errorf("assert 'DummyRequestArg.Type':: expected '%v', got '%v'", "KNNS", arg.Type)
	}
	if arg.Number != 280123412341234123 {
		t.Errorf("assert 'DummyRequestArg.Number':: expected '%v', got '%v'", 280123412341234123, arg.Number)
	}
	if arg.ShowDetail != true {
		t.Errorf("assert 'DummyRequestArg.ShowDetail':: expected '%v', got '%v'", true, arg.ShowDetail)
	}
	if arg.EnableDebug != false {
		t.Errorf("assert 'DummyRequestArg.EnableDebug':: expected '%v', got '%v'", false, arg.EnableDebug)
	}
	expectedTags := []string{"T", "ER", "XVV"}
	if !reflect.DeepEqual(arg.Tags, expectedTags) {
		t.Errorf("assert 'character.Alias':: expected '%#v', got '%#v'", expectedTags, arg.Tags)
	}
	{
		if arg.Detail == nil {
			t.Error("assert 'DummyRequestArg.Detail':: should not be nil")
		}
		var detail = arg.Detail
		if detail.Operator != "nami" {
			t.Errorf("assert 'DummyRequestArg.Detail.Operator':: expected '%v', got '%v'", "nami", detail.Operator)
		}
	}
}

func TestProcess_WithRaw(t *testing.T) {
	input := `{
		"id": "F0003452",
		"type": "KNNS",
		"number": 280123412341234123,
		"showDetail": true,
		"tags": ["T","ER","XVV"],
		"detail": {
			"operator": "nami"
		},
		"raw": { "key": "value" }
	}`

	arg := DummyRequestArg{}
	err := Process([]byte(input), &arg)
	if err != nil {
		t.Error(err)
	}

	expectedRow := `{"key":"value"}`
	if string(*arg.Raw) != expectedRow {
		t.Errorf("assert 'DummyRequestArg.Raw':: expected '%v', got '%v'", `{"key":"value"}`, string(*arg.Raw))
	}
}
