package test

import "github.com/Bofry/httparg"

type DummyRequestArg struct {
	ID          string                 `json:"*id"`
	Type        *string                `json:"*type"`
	Number      int64                  `json:"number"`
	ShowDetail  bool                   `json:"-"        query:"SHOW_DETAIL"`
	EnableDebug bool                   `json:"-"        query:"ENABLE_DEBUG"`
	Tags        []string               `json:"tags"`
	Detail      *DummyRequestArgDetail `json:"detail"`
}

type DummyRequestArgDetail struct {
	Operator string `json:"operator"`
}

type DummyRequestArgValidatable DummyRequestArg

func (v *DummyRequestArgValidatable) Validate() error {
	if v.ID == "0" {
		return &httparg.InvalidArgumentError{
			Name:   "id",
			Reason: "cannot be 0",
		}
	}
	return nil
}
