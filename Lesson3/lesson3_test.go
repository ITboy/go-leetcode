package lesson3

import "testing"

type testcase struct {
	s      string
	maxlen int
}

func TestFindNextLongestSubstrings(t *testing.T) {
	cases := []testcase{
		{s: "ababc", maxlen: 3},
		{s: "abcabcbb", maxlen: 3},
		{s: "bbbbb", maxlen: 1},
		{s: "pwwkew", maxlen: 3},
	}

	for _, c := range cases {
		if got := lengthOfLongestSubstring(c.s); c.maxlen != got {
			t.Errorf("lengthOfLongestSubstring(%v) = %v, expected: %v", c.s, got, c.maxlen)
		}
	}
}
