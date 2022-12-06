package domain

import "github.com/kuchensheng/capc/infrastructure/model"

//DomainIntf 最基本领域
type DomainIntf interface {
	GetRepository() model.Repository
}
