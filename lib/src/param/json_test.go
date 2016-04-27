package param

import (
	"testing"
)

func TestJsonArray(t *testing.T) {
	req := make(map[string]interface{}, 1)
	req["a"] = "[\"abc\",\"def\"]"
	re, okRe := GetJsonArrayWithKey(req, "a")
	t.Logf("re %v, okRe %v", re, okRe)
}