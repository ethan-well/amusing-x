package riskcontrol

import (
	"amusingx.fit/amusingx/mysqlstruct/amusingriskcontrol"
	"amusingx.fit/amusingx/services/callisto/conf"
	"context"
	"github.com/ItsWewin/superfactory/logger"
	"testing"
)

func TestQueryVerificationCodeRiskByPhone(t *testing.T) {
	conf.Mock()
	Mock()
	risk, err := QueryVerificationCodeRiskByPhone(context.Background(), "18710565589")
	if err != nil {
		t.Fatalf("Some error: %s", err)
	}

	t.Logf("Risk: %s", logger.ToJson(risk))
}

func TestVerificationCodeRiskUpdate(t *testing.T) {
	conf.Mock()
	Mock()

	risk := &amusingriskcontrol.VerificationCodeRisk{
		ID:        1,
		Phone:     "18710565589",
		UsedCount: 2,
		MaxCount:  10,
	}

	err := VerificationCodeRiskUpdate(context.Background(), risk)
	if err != nil {
		t.Fatalf("Some error: %s", err)
	}

	t.Logf("succeed")
}

func TestInsertOrUpdate(t *testing.T) {
	conf.Mock()
	Mock()

	risk := &amusingriskcontrol.VerificationCodeRisk{
		Phone:     "86-18710565589",
		UsedCount: 1,
		MaxCount:  11,
	}

	err := VerificationCodeRiskInsertOrUpdate(context.Background(), risk)
	if err != nil {
		t.Fatalf("Some error: %s", err)
	}

	t.Logf("succeed")
}
