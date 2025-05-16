package contact

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
)

type IContactController interface {
	CreateAppeal(ctx context.Context, appeal dto.Appeal) error
}
