package account

import (
	"context"
	"errors"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/eeQuillibrium/wobble/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (a *Account) Register(ctx context.Context, reg dto.Register, passHash []byte) (models.RegisteredUser, error) {
	var user models.RegisteredUser

	err := a.db.QueryRow(ctx, queryCreate, reg.Name, reg.Email, reg.Login, passHash).Scan(&user.Password, &user.Role)
	if err != nil {
		return models.RegisteredUser{}, errors.New("error w/ insert")
	}

	return user, nil
}

func (a *Account) GetLoginInfo(ctx context.Context, login string) (models.RegisteredUser, error) {
	var user models.RegisteredUser

	err := pgxscan.Get(ctx, a.db, &user, queryGetUser, login)
	if err != nil {
		return models.RegisteredUser{}, err
	}

	return user, nil
}
