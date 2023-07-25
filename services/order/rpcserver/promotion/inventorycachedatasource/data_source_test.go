package inventorycachedatasource

import (
	"context"
	"testing"
)

func TestNewDataSource(t *testing.T) {
	if testing.Short() {
		t.Skip("skip ...")
	}

	inventories := map[string]int64{
		LocalVariableSource: 10000,
	}

	for st, iv := range inventories {
		source, err := NewDataSource(st)
		if err != nil {
			t.Fatal(err)
		}

		data, err := source.GetInventories(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if int64(len(data)) != iv {
			t.Fatal(err)
		}
	}
}
