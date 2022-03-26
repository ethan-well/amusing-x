package charonclient

import (
	"testing"
)

func ClientMock(t *testing.T) {
	InitClient(":20004")
}
