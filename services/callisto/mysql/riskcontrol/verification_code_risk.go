package riskcontrol

import (
	"amusingx.fit/amusingx/mysqlstruct/amusingriskcontrol"
	"context"
	"database/sql"
	"github.com/ItsWewin/superfactory/xerror"
)

func QueryVerificationCodeRiskByPhone(ctx context.Context, phone string) (*amusingriskcontrol.VerificationCodeRisk, *xerror.Error) {
	query := `
   		SELECT id, phone, used_count, max_count
		FROM verification_code_risk
		WHERE phone = ?
    `

	var risk amusingriskcontrol.VerificationCodeRisk
	err := CallistoDB.GetContext(ctx, &risk, query, phone)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Sql execute error")
	}

	return &risk, nil
}

func VerificationCodeRiskUpdate(ctx context.Context, risk *amusingriskcontrol.VerificationCodeRisk) *xerror.Error {
	query := `
   		UPDATE verification_code_risk id
		set used_count = :used_count,
   			max_count = :max_count 
		WHERE id = :id
    `

	_, err := CallistoDB.NamedExecContext(ctx, query, risk)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Sql Execute failed. ")
	}

	return nil
}

func VerificationCodeRiskInsertOrUpdate(ctx context.Context, risk *amusingriskcontrol.VerificationCodeRisk) *xerror.Error {

	query := `
		INSERT INTO verification_code_risk (phone, used_count, max_count)
		VALUES (:phone, :used_count, :max_count)
		ON DUPLICATE KEY UPDATE
			used_count = VALUES(used_count),
			max_count = VALUES(max_count)
	`

	_, err := CallistoDB.NamedExecContext(ctx, query, risk)
	if err != nil {
		return xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Sql Execute failed. ")
	}

	return nil
}
