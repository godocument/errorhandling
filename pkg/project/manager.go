package project

import (
	"context"

	"github.com/godocument/errorhandling/pkg/project/dao"
)

type Manager interface {
	Get(ctxt context.Context, id int64) (*dao.Project, error)
}

func New() Manager {
	return &manager{
		dao: dao.New(),
	}
}

var _ Manager = &manager{}

type manager struct {
	dao dao.DAO
}

func (m *manager) Get(ctxt context.Context, id int64) (*dao.Project, error) {
	return m.dao.Get(ctxt, id)
}
