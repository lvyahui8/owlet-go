package utils

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCopyFile(t *testing.T) {

}


func TestGetCommonPrefix(t *testing.T) {
    commonPrefix := GetCommonPrefix([]string{
        "a.b.c",
        "a.b",
        "a.b.x",
    })
    assert.True(t, len(commonPrefix) > 0)
    t.Log(commonPrefix)
}