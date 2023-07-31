package leetcode0392

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, isSubsequence("abc", "ahbgdc"))
	a.Equal(false, isSubsequence("axc", "ahbgdc"))
	a.Equal(true, isSubsequence("", "ahbgdc"))
}
