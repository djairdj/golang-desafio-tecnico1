package product

import (
	"context"
	"github.com/djairdj/golang-desafio-tecnico1/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, name string) (*entity.Product, error)
	List(ctx context.Context) ([]entity.Product, error)
	GetOne(ctx context.Context, id string) (*entity.Product, error)
	Update(ctx context.Context, p *entity.Product) error
}
