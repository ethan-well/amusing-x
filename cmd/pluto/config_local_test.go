package main

import (
	"github.com/ItsWewin/superfactory/logger"
	"runtime"
	"testing"
)

func TestTingGetFileName(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	logger.Infof("file name: %s", filename)
}