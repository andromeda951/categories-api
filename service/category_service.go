package service

import (
	"andromeda/belajar-golang-restful-api/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Upadate(ctx context.Context, request web.CategoryUpadateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindByid(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}