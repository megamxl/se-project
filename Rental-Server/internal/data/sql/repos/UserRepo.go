package repos

import (
	"context"
	"errors"
	"github.com/google/uuid"
	dataInt "github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/model"
	"github.com/megamxl/se-project/Rental-Server/internal/data/sql/dao/query"
)

var _ dataInt.UserRepository = (*UserRepo)(nil)

type UserRepo struct {
	Q   *query.Query
	Ctx context.Context
}

func (u *UserRepo) GetUserByEmail(email string) (dataInt.RentalUser, error) {
	find, err := u.Q.WithContext(u.Ctx).RentalUser.Where(u.Q.RentalUser.Email.Eq(email)).Find()
	if err != nil {
		return dataInt.RentalUser{}, err
	}

	if len(find) != 1 {
		return dataInt.RentalUser{}, errors.New("RentalUser not found or duplicate Email")
	}

	return modelToIntRentalUser(find[0]), nil
}

func (u *UserRepo) GetUserById(id uuid.UUID) (dataInt.RentalUser, error) {
	find, err := u.Q.WithContext(u.Ctx).RentalUser.Where(u.Q.RentalUser.ID.Eq(id.String())).Find()
	if err != nil {
		return dataInt.RentalUser{}, err
	}

	if len(find) != 1 {
		return dataInt.RentalUser{}, errors.New("RentalUser not found or duplicate ID")
	}

	return modelToIntRentalUser(find[0]), nil
}

func (u *UserRepo) UpdateUserById(id uuid.UUID, update dataInt.RentalUser) (dataInt.RentalUser, error) {
	_, err := u.Q.WithContext(u.Ctx).RentalUser.Where(u.Q.RentalUser.ID.Eq(id.String())).Updates(update)
	if err != nil {
		return dataInt.RentalUser{}, err
	}

	find, err := u.Q.WithContext(u.Ctx).RentalUser.Where(u.Q.RentalUser.ID.Eq(id.String())).Find()
	return modelToIntRentalUser(find[0]), nil
}

func (u *UserRepo) UpdateUserByEmail(email string, update dataInt.RentalUser) (dataInt.RentalUser, error) {
	_, err := u.Q.WithContext(u.Ctx).RentalUser.Where(u.Q.RentalUser.Email.Eq(email)).Updates(update)
	if err != nil {
		return dataInt.RentalUser{}, err
	}

	find, err := u.Q.WithContext(u.Ctx).RentalUser.Where(u.Q.RentalUser.Email.Eq(email)).Find()
	return modelToIntRentalUser(find[0]), nil
}

func (u *UserRepo) DeleteUserById(id uuid.UUID) error {
	_, err := u.Q.WithContext(u.Ctx).RentalUser.Where(u.Q.RentalUser.ID.Eq(id.String())).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) SaveUser(user dataInt.RentalUser) (dataInt.RentalUser, error) {
	toSaveUser := intToModelRentalUser(user)
	err := u.Q.WithContext(u.Ctx).RentalUser.Save(toSaveUser)
	if err != nil {
		return dataInt.RentalUser{}, err
	}
	return modelToIntRentalUser(toSaveUser), nil
}

func intToModelRentalUser(rentalUser dataInt.RentalUser) *model.RentalUser {
	newRentalUser := &model.RentalUser{
		Name:     rentalUser.Name,
		Email:    rentalUser.Email,
		Password: rentalUser.Password,
	}
	return newRentalUser
}

func modelToIntRentalUser(newRentalUser *model.RentalUser) dataInt.RentalUser {
	savedRentalUser := dataInt.RentalUser{
		Id:       uuid.MustParse(newRentalUser.ID),
		Name:     newRentalUser.Name,
		Email:    newRentalUser.Email,
		Password: newRentalUser.Password,
	}
	return savedRentalUser
}
