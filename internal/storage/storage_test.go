package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewDataStorage(t *testing.T) {
	require := require.New(t)
	log := zap.NewNop()
	fileName := os.TempDir() + string(os.PathSeparator) + "test_123456789.json"

	t.Run("storage without file", func(t *testing.T) {
		ds := NewDataStorage("", 0, log)
		require.NotNil(ds)
	})

	t.Run("storage with file", func(t *testing.T) {
		ds := NewDataStorage(fileName, 0, log)
		require.NotNil(ds)
	})

}
