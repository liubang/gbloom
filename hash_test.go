package gbloom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	test_cases = map[string]uint32{
		"OatoXR1oWQ27aghr":    3936657155,
		"FsuHfYSASnISaLXK":    3587760215,
		"0u56dghh6A2VpoqW":    2287161574,
		"$Q$0pjppVzm367Vr":    3093058208,
		"5wA#jO9nuLEUBKWf":    2815475469,
		"lfb8M5E4ZqqgFLUJ":    2274938967,
		"gF#580VAxlU8SYjl":    1398756171,
		"helloworld":          2651853088,
		"$poskBwgBjiR2%dFL":   103479080,
		"N#zvGO2EDbOKast9gY":  270766449,
		"lJe1H3X6BM5kB!#BFR8": 3731866214,
	}
)

func TestHash(t *testing.T) {
	for str, val := range test_cases {
		byte_arr := []byte(str)
		assert.Equal(t, Hash(byte_arr, uint32(len(byte_arr)), 0xbc9f1d34), val)
	}
}
