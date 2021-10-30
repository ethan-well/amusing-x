package inventorycachedatasource

import (
	"context"
	"github.com/ItsWewin/superfactory/xerror"
	"golang.org/x/exp/rand"
)

type LocalVariablesSource struct {
}

func NewLocalVariablesSource() *LocalVariablesSource {
	return &LocalVariablesSource{}
}

func (s *LocalVariablesSource) GetInventories(ctx context.Context) (map[int64]int64, *xerror.Error) {
	maxId := 10000
	idInventoryMap := make(map[int64]int64, maxId)
	for i := 1; i <= 10000; i++ {
		idInventoryMap[int64(i)] = rand.Int63n(1000)
	}

	return idInventoryMap, nil
}
