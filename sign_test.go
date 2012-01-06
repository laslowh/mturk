package mturk_test

import (
	. "launchpad.net/gocheck"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/mturk"
)

// Mechanical Turk REST authentication docs: http://goo.gl/wrzfn

var testAuth = aws.Auth{"user", "secret"}

func (s *S) TestBasicSignature(c *C) {
	params := map[string]string{}
	mturk.Sign(testAuth, "AWSMechanicalTurkRequester", "SearchHITS", "2011-12-30T15:49:33Z", params)
	expected := "vqSn4jCjvGA0tFu+VHukWolEhEU="
	c.Assert(params["Signature"], Equals, expected)
}

