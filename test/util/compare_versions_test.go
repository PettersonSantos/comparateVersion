package util

import (
	"compareVersions/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCompareVersionAndReturnNotError(t *testing.T) {
	assert.Equal(t, -1, util.CompareVersion("0.1", "1.1"))
	assert.Equal(t, 1, util.CompareVersion("1.0.1", "1"))
	assert.Equal(t, -1, util.CompareVersion("7.5.2.4", "7.5.3"))
	assert.Equal(t, 0, util.CompareVersion("1.01", "1.001"))
	assert.Equal(t, 0, util.CompareVersion("1.0", "1.0.0"))
}
