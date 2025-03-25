package repos

import (
	"github.com/google/uuid"
	dataInt "github.com/megamxl/se-project/Rental-Server/internal/data"
)

var _ dataInt.UserRepository = (*UserRepo)(nil)

type UserRepo struct {
}

func (u *UserRepo) UpdateUserByEmail(email string, update dataInt.RentalUser) (dataInt.RentalUser, error) {
	_, err := u.Q.WithContext(u.Ctx).RentalUser.Where(u.Q.RentalUser.Email.Eq(email)).Updates(update)
	if err != nil {
		return dataInt.RentalUser{}, err
	}

func (u UserRepo) GetUserById(id uuid.UUID) (dataInt.RentalUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) UpdateUserById(id uuid.UUID, update dataInt.RentalUser) (dataInt.RentalUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) UpdateUserByEmail(email string, update dataInt.RentalUser) (dataInt.RentalUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteUserById(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) SaveUser(user dataInt.RentalUser) (dataInt.RentalUser, error) {
	//TODO implement me
	panic("implement me")
}
