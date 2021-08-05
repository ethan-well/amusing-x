package amusingriskcontrol

type VerificationCodeRisk struct {
	ID         int64  `db:"id" json:"id" `
	Phone      string `db:"phone" json:"phone"`
	UsedCount  int64  `db:"used_count" json:"used_count"`
	MaxCount   int64  `db:"max_count" json:"max_count"`
	CreateTime string `db:"create_time" json:"create_time"`
	UpdateTime string `db:"update_time" json:"update_time"`
}
