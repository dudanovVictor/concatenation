package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_existA(t *testing.T) {
	/*s1 := []string{"cat", "dogs", "catdogs"}
		res1 := []string{"catdogs"}
		assert.Equal(t, res1, findAllConcatenatedWordsInADict(s1))
	/**/
	s2 := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	res2 := []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}
	assert.Equal(t, res2, findAllConcatenatedWordsInADict(s2))
	/*
		s3 := []string{"a", "aaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaaaa"}
		res3 := []string{"aaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaaaa"}
		assert.Equal(t, res3, findAllConcatenatedWordsInADict(s3))

		s4 := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
		res4 := []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}
		assert.Equal(t, res4, findAllConcatenatedWordsInADict(s4))
		/**/
}
