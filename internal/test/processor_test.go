package test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/Bofry/httparg"
	"github.com/Bofry/structproto"
)

func TestProcessor(t *testing.T) {
	query := "SHOW_DETAIL"
	postbody := []byte(`
	{
		"id": "F0003452",
		"type": "KNNS",
		"number": 280123412341234123,
		"tags": ["T","ER","XVV"],
		"detail": {
			"operator": "nami"
		}
	}`)
	contentType := "  application/json ;charset=utf8"
	arg := DummyRequestArg{}
	processor := httparg.NewProcessor(&arg, httparg.Option{
		ErrorHandler: func(err error) {
			t.Error(err)
		},
	})

	processor.
		ProcessQueryString(query).
		ProcessContent(postbody, contentType)

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
		if detail != nil {
			if detail.Operator != "nami" {
				t.Errorf("assert 'DummyRequestArg.Detail.Operator':: expected '%v', got '%v'", "nami", detail.Operator)
			}
		}
	}
}

func TestProcessNilValueOnRequiredField(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("the 'TestProcessNilValueOnRequiredField()' should throw '%s' error", "missing required symbol 'type'")
		} else {
			missingRequiredFieldError, ok := err.(*structproto.MissingRequiredFieldError)
			if !ok {
				t.Errorf("the error expected '%T', got '%T'", &structproto.MissingRequiredFieldError{}, err)
			}
			if missingRequiredFieldError.Field != "type" {
				t.Errorf("assert 'MissingRequiredFieldError.Field':: expected '%v', got '%v'", "type", missingRequiredFieldError.Field)
			}
		}
	}()

	query := "SHOW_DETAIL"
	postbody := []byte(`
	{
		"id": "F0003452",
		"type": null,
		"number": 280123412341234123,
		"tags": ["T","ER","XVV"],
		"detail": {
			"operator": "nami"
		}
	}`)
	contentType := "  application/json ;charset=utf8"
	arg := DummyRequestArg{}
	processor := httparg.ArgsWithOption(&arg, httparg.Option{
		ErrorHandler: func(err error) {
			panic(err)
		},
	})

	processor.
		ProcessQueryString(query).
		ProcessContent(postbody, contentType)
}

func TestWrongContentType(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("the 'TestWrongContentType()' should throw '%s' error", "cannot process specified content-type 'unknown_type'")
		}
	}()

	postbody := []byte("")
	contentType := "  unknown_type"

	arg := DummyRequestArg{}
	httparg.Args(&arg).
		ProcessContent(postbody, contentType)
}

func TestArgs(t *testing.T) {
	query := "SHOW_DETAIL"
	postbody := []byte(`
	{
		"id": "F0003452",
		"type": "KNNS",
		"number": 280123412341234123,
		"tags": ["T","ER","XVV"],
		"detail": {
			"operator": "nami"
		}
	}`)
	contentType := "  application/json ;charset=utf8"

	arg := DummyRequestArg{}
	httparg.Args(&arg).
		ProcessQueryString(query).
		ProcessContent(postbody, contentType)

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
		if detail != nil {
			if detail.Operator != "nami" {
				t.Errorf("assert 'DummyRequestArg.Detail.Operator':: expected '%v', got '%v'", "nami", detail.Operator)
			}
		}
	}
}

func TestValidatable(t *testing.T) {
	var buffer bytes.Buffer

	query := "SHOW_DETAIL"
	postbody := []byte(`
	{
		"id": "0",
		"type": "KNNS",
		"number": 280123412341234123,
		"tags": ["T","ER","XVV"],
		"detail": {
			"operator": "nami"
		}
	}`)
	contentType := "  application/json ;charset=utf8"

	if buffer.Len() != 0 {
		t.Errorf("buffer should be empty")
	}

	arg := DummyRequestArgValidatable{}
	httparg.ArgsWithOption(&arg, httparg.Option{
		ErrorHandler: func(err error) {
			fmt.Fprintf(&buffer, "%+v", err)
		},
	}).
		ProcessQueryString(query).
		ProcessContent(postbody, contentType).
		Validate()

	if buffer.Len() == 0 {
		t.Errorf("buffer should not be empty")
	}
	expectedErrMsg := `invalid argument "id"; cannot be 0`
	if buffer.String() != expectedErrMsg {
		t.Errorf("assert 'buffer.String()':: expected '%v', got '%v'", expectedErrMsg, buffer.String())
	}
}

func TestMultipart(t *testing.T) {
	type (
		BookDetail struct {
			Price int      `json:"price"`
			Links []string `json:"links"`
		}

		Book struct {
			Title     string      `multipart:"*title"`
			Author    string      `multipart:"*author"`
			Reversion int         `multipart:"*reversion"`
			Category  string      `multipart:"category"`
			Detail    *BookDetail `multipart:"detail"`
		}
	)

	contentType := "multipart/mixed; boundary=foo"
	postbody := []byte(`
--foo
Content-Disposition: form-data; name="title"

Oliver Twist
--foo
Content-Disposition: form-data; name="author"

Dickens
--foo
Content-Disposition: form-data; name="reversion"

3
--foo
Content-Disposition: form-data; name="detail"
Content-Type: application/json

{
	"price": 25,
	"links": [
		"www.dickensbooks.com"
	]
}
--foo--
`)

	arg := Book{}
	processor := httparg.Args(&arg)
	processor.ProcessContent(postbody, contentType)

	{
		var expectedTitle string = "Oliver Twist"
		if arg.Title != expectedTitle {
			t.Errorf("assert 'Book.Title':: expected '%v', got '%v'", expectedTitle, arg.Title)
		}
		var expectedAuthor string = "Dickens"
		if arg.Author != expectedAuthor {
			t.Errorf("assert 'Book.Author':: expected '%v', got '%v'", expectedAuthor, arg.Author)
		}
		var expectedReversion int = 3
		if arg.Reversion != expectedReversion {
			t.Errorf("assert 'Book.Reversion':: expected '%v', got '%v'", expectedReversion, arg.Reversion)
		}
		var expectedCategory string = ""
		if arg.Category != expectedCategory {
			t.Errorf("assert 'Book.Category':: expected '%v', got '%v'", expectedCategory, arg.Category)
		}

		if arg.Detail == nil {
			t.Errorf("assert 'Book.Detail':: should not be nil")
		}
		{
			var detial = arg.Detail
			var expectedPrice int = 25
			if detial.Price != expectedPrice {
				t.Errorf("assert 'BookDetail.Price':: expected '%v', got '%v'", expectedPrice, detial.Price)
			}
			var expectedLinks []string = []string{"www.dickensbooks.com"}
			if !reflect.DeepEqual(detial.Links, expectedLinks) {
				t.Errorf("assert 'BookDetail.Links':: expected '%v', got '%v'", expectedLinks, detial.Links)
			}
		}
	}

	t.Logf("%+v\n", arg)        // {Title:Oliver Twist Author:Dickens Reversion:3 Category: Detail:0xc000086ae0
	t.Logf("%+v\n", arg.Detail) // &{Price:25 Links:[www.dickensbooks.com]}
}

func TestMultipart_Nested(t *testing.T) {
	type (
		BookDetail struct {
			Price int      `json:"price"`
			Links []string `json:"links"`
		}

		BookPreview struct {
			Prolog         string `multipart:"prolog"`
			TableOfContent string `multipart:"tof"`
		}

		Book struct {
			Title     string       `multipart:"*title"`
			Author    string       `multipart:"*author"`
			Reversion int          `multipart:"*reversion"`
			Category  string       `multipart:"category"`
			Detail    *BookDetail  `multipart:"detail"`
			Preview   *BookPreview `multipart:"preview"`
		}
	)

	contentType := "multipart/mixed; boundary=foo"
	postbody := []byte(`
--foo
Content-Disposition: form-data; name="title"

Oliver Twist
--foo
Content-Disposition: form-data; name="author"

Dickens
--foo
Content-Disposition: form-data; name="reversion"

3
--foo
Content-Disposition: form-data; name="detail"
Content-Type: application/json

{
	"price": 25,
	"links": [
		"www.dickensbooks.com"
	]
}
--foo
Content-Disposition: form-data; name="preview"
Content-Type: multipart/mixed; boundary=bar

--bar
Content-Disposition: form-data; name="prolog"

Prolog .....
--bar
Content-Disposition: form-data; name="tof"

Table of Content .....
--bar--
--foo--
`)

	arg := Book{}
	processor := httparg.Args(&arg)
	processor.ProcessContent(postbody, contentType)

	{
		var expectedTitle string = "Oliver Twist"
		if arg.Title != expectedTitle {
			t.Errorf("assert 'Book.Title':: expected '%v', got '%v'", expectedTitle, arg.Title)
		}
		var expectedAuthor string = "Dickens"
		if arg.Author != expectedAuthor {
			t.Errorf("assert 'Book.Author':: expected '%v', got '%v'", expectedAuthor, arg.Author)
		}
		var expectedReversion int = 3
		if arg.Reversion != expectedReversion {
			t.Errorf("assert 'Book.Reversion':: expected '%v', got '%v'", expectedReversion, arg.Reversion)
		}
		var expectedCategory string = ""
		if arg.Category != expectedCategory {
			t.Errorf("assert 'Book.Category':: expected '%v', got '%v'", expectedCategory, arg.Category)
		}

		if arg.Detail == nil {
			t.Errorf("assert 'Book.Detail':: should not be nil")
		}
		{
			var detial = arg.Detail
			var expectedPrice int = 25
			if detial.Price != expectedPrice {
				t.Errorf("assert 'BookDetail.Price':: expected '%v', got '%v'", expectedPrice, detial.Price)
			}
			var expectedLinks []string = []string{"www.dickensbooks.com"}
			if !reflect.DeepEqual(detial.Links, expectedLinks) {
				t.Errorf("assert 'BookDetail.Links':: expected '%v', got '%v'", expectedLinks, detial.Links)
			}
		}

		if arg.Preview == nil {
			t.Errorf("assert 'Book.Preview':: should not be nil")
		}
		{
			var preview = arg.Preview
			var expectedProlog string = "Prolog ....."
			if preview.Prolog != expectedProlog {
				t.Errorf("assert 'BookPreview.Prolog':: expected '%v', got '%v'", expectedProlog, preview.Prolog)
			}
			var expectedTableOfContent string = "Table of Content ....."
			if preview.TableOfContent != expectedTableOfContent {
				t.Errorf("assert 'BookPreview.TableOfContent':: expected '%v', got '%v'", expectedTableOfContent, preview.TableOfContent)
			}
		}
	}

	t.Logf("%+v\n", arg)         // {Title:Oliver Twist Author:Dickens Reversion:3 Category: Detail:0xc000005ac0 Preview:0xc000005c60}
	t.Logf("%+v\n", arg.Detail)  // &{Price:25 Links:[www.dickensbooks.com]}
	t.Logf("%+v\n", arg.Preview) // &{Prolog:Prolog ..... TableOfContent:Table of Content .....}
}

// NOTE: the test must put the last test
func TestSetUpErrorHandleFunc(t *testing.T) {
	var buffer bytes.Buffer

	httparg.RegistryService.SetupErrorHandler(func(err error) {
		fmt.Fprintf(&buffer, "%+v", err)
	})
	query := "SHOW_DETAIL"
	postbody := []byte(`
	{
		"id": "F0003452",
		"type": null,
		"number": 280123412341234123,
		"tags": ["T","ER","XVV"],
		"detail": {
			"operator": "nami"
		}
	}`)
	contentType := "  application/json ;charset=utf8"

	if buffer.Len() != 0 {
		t.Errorf("buffer should be empty")
	}

	arg := DummyRequestArg{}
	httparg.Args(&arg).
		ProcessQueryString(query).
		ProcessContent(postbody, contentType)

	if buffer.Len() == 0 {
		t.Errorf("buffer should not be empty")
	}
	if buffer.String() != "missing required symbol 'type'" {
		t.Errorf("assert 'buffer.String()':: expected '%v', got '%v'", "missing required symbol 'type'", buffer.String())
	}
}
