package read_file_last_line

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLastLine(t *testing.T) {
	assert := assert.New(t)

	data, n, err := ReadLastLine("testdata/hello.txt")
	assert.NoError(err)

	line := string(data)
	assert.Equal("Goodbye World!", line)
	assert.Equal(13, n)
}

func TestReadLastLineNoEOL(t *testing.T) {
	assert := assert.New(t)

	data, n, err := ReadLastLine("testdata/noeol.txt")
	assert.NoError(err)

	line := string(data)
	assert.Equal("Goodbye World!", line)
	assert.Equal(13, n)
}
