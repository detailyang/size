package size

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect Unit
	}{
		{
			"1B",
			B,
		},
		{
			"1KB",
			KB,
		},
		{
			"1.2KB",
			1200 * B,
		},
		{
			"1MB",
			1 * MB,
		},
		{
			"1MiB",
			1 * MiB,
		},
	} {
		a, err := Parse(c.input)
		require.Nil(t, err)
		require.Equal(t, c.expect, a)
	}
}

func TestString(t *testing.T) {
	a := 1 * B
	require.Equal(t, "1B", a.String())
	a = 1 * KB
	require.Equal(t, "1KB", a.String())
	a = 1 * MB
	require.Equal(t, "1MB", a.String())
	a = 1 * GB
	require.Equal(t, "1GB", a.String())
	a = 1 * TB
	require.Equal(t, "1TB", a.String())
	a = 1 * PB
	require.Equal(t, "1PB", a.String())
	a = 2 * PB
	require.Equal(t, "2PB", a.String())

	a = 1 * KiB
	require.Equal(t, "1KiB", a.String())
	a = 1 * MiB
	require.Equal(t, "1MiB", a.String())
	a = 1 * GiB
	require.Equal(t, "1GiB", a.String())
	a = 1 * TiB
	require.Equal(t, "1TiB", a.String())
	a = 1 * PiB
	require.Equal(t, "1PiB", a.String())
	a = 2 * PiB
	require.Equal(t, "2PiB", a.String())
}
