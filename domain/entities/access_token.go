package entities

import base "altastore-api/domain/persistence"

//AccessToken base model
type AccessToken struct {
	base.Model
	AccountID uint64 `json:"account_id"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}
