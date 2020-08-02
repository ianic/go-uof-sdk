package uof

import (
	"encoding/xml"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	inner := fmt.Errorf("get failed")
	ae := APIError{URL: "url", Inner: inner}
	assert.Equal(t, inner, ae.Unwrap())

	e := E("api", ae)
	err := error(e)

	var s string
	var ae2 APIError
	if errors.As(err, &ae2) {
		s = ae2.Error()
	}
	assert.Equal(t, "uof api error url: url, inner: get failed", s)

	var e2 Error
	if errors.As(err, &e2) {
		s = e2.Error()
	}
	assert.Equal(t, "uof error op: api, inner: uof api error url: url, inner: get failed", s)

	ae = APIError{URL: "url", Inner: inner, StatusCode: 422, Response: "tee"}
	assert.Equal(t, "uof api error url: url, status code: 422, response: tee, inner: get failed", ae.Error())
}

func TestInnerError(t *testing.T) {
	inner := fmt.Errorf("some inner error")
	ue := Notice("operation", inner)
	err := ue.Unwrap()
	assert.Equal(t, inner, err)

	assert.Equal(t, "NOTICE uof error op: operation, inner: some inner error", ue.Error())
}

func TestParseResponse(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<response response_code="BAD_REQUEST">
	<action>Request for booking an event : sr:match:22868927 from bookmaker: 19638 received</action>
	<message>ERROR. Bad Request: The match does not belong to any available package</message>
</response>`

	var rsp UOFRsp
	err := xml.Unmarshal([]byte(data), &rsp)
	assert.Nil(t, err)
	//pp(rsp)
}
