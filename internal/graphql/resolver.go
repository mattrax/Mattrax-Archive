package graphql

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) CurrentUser(ctx context.Context) (User, error) {
	panic("not implemented")
}
func (r *queryResolver) GetDevice(ctx context.Context, uuid string) (Device, error) {
	panic("not implemented")
}
