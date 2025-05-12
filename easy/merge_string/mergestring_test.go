package mergestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	word1 := "abc"
	word2 := "pqr"
	expecting := "apbqcr"
	result := mergeAlternately(word1, word2)
	assert.Equal(t, expecting, result, "Expecting [%s] | Result [%s]", expecting, result)
}

func TestCase2(t *testing.T) {
	word1 := "ab"
	word2 := "pqrs"
	expecting := "apbqrs"
	result := mergeAlternately(word1, word2)
	assert.Equal(t, expecting, result, "Expecting [%s] | Result[%s]", expecting, result)
}

func TestCase3(t *testing.T) {
	word1 := "abcd"
	word2 := "pq"
	expecting := "apbqcd"
	result := mergeAlternately(word1, word2)
	assert.Equal(t, expecting, result, "Expecting [%s] | Result[%s]", expecting, result)
}

func TestCase4(t *testing.T) {
	word1 := ""
	word2 := "pq"
	expecting := "pq"
	result := mergeAlternately(word1, word2)
	assert.Equal(t, expecting, result, "Expecting [%s] | Result[%s]", expecting, result)
}
