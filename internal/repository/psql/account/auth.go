package account

import (
	"context"
	"errors"
	"github.com/eeQuillibrium/wobble/internal/dto"
)

func (a *Account) Register(ctx context.Context, reg dto.Register, passHash []byte) (uint64, error) {
	var id uint64

	err := a.db.QueryRow(ctx, queryCreate, reg.Name, reg.Email, reg.Login, passHash).Scan(&id)
	if err != nil {
		return 0, errors.New("error w/ insert")
	}

	return id, nil
}

func (a *Account) Login(ctx context.Context, login string) ([]byte, uint64, error) {
	var user struct {
		ID       uint64
		Password []byte
	}

	err := a.db.QueryRow(ctx, queryGetUser, login).Scan(&user.ID, &user.Password)
	if err != nil {
		return nil, 0, err
	}

	return user.Password, user.ID, nil
}
