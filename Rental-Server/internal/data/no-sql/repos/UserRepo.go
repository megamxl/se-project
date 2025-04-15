package repos

import (
	"context"
	"errors"
	"github.com/google/uuid"
	dataInt "github.com/megamxl/se-project/Rental-Server/internal/data"
	"github.com/megamxl/se-project/Rental-Server/internal/data/no-sql/dao/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var _ dataInt.UserRepository = (*UserRepo)(nil)

type UserRepo struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func (u *UserRepo) GetAllUsers() ([]dataInt.RentalUser, error) {
	log.Println("👤 [Mongo] GetAllUsers")
	cursor, err := u.Collection.Find(u.Ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(u.Ctx)

	var results []dataInt.RentalUser
	for cursor.Next(u.Ctx) {
		var m model.RentalUser
		if err := cursor.Decode(&m); err != nil {
			return nil, err
		}
		user, err := convertModelToDataUser(m)
		if err != nil {
			return nil, err
		}
		results = append(results, user)
	}

	return results, nil
}

func (u *UserRepo) GetUserByEmail(email string) (dataInt.RentalUser, error) {
	log.Printf("👤 [Mongo] GetUserByEmail: %s", email)
	var result model.RentalUser

	err := u.Collection.FindOne(u.Ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return dataInt.RentalUser{}, errors.New("user not found")
		}
		return dataInt.RentalUser{}, err
	}

	return convertModelToDataUser(result)
}

func (u *UserRepo) GetUserById(id uuid.UUID) (dataInt.RentalUser, error) {
	log.Printf("👤 [Mongo] GetUserById: %s", id.String())
	var result model.RentalUser

	err := u.Collection.FindOne(u.Ctx, bson.M{"_id": id.String()}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return dataInt.RentalUser{}, errors.New("user not found")
		}
		return dataInt.RentalUser{}, err
	}

	return convertModelToDataUser(result)
}

func (u *UserRepo) UpdateUserById(id uuid.UUID, update dataInt.RentalUser) (dataInt.RentalUser, error) {
	log.Printf("👤 [Mongo] UpdateUserById: %s with %+v", id.String(), update)
	filter := bson.M{"_id": id.String()}
	updateDoc := bson.M{"$set": bson.M{
		"name":     update.Name,
		"email":    update.Email,
		"password": update.Password,
		"admin":    update.Admin,
	}}

	_, err := u.Collection.UpdateOne(u.Ctx, filter, updateDoc)
	if err != nil {
		return dataInt.RentalUser{}, err
	}

	return update, nil
}

func (u *UserRepo) UpdateUserByEmail(email string, update dataInt.RentalUser) (dataInt.RentalUser, error) {
	log.Printf("👤 [Mongo] UpdateUserByEmail: %s with %+v", email, update)
	filter := bson.M{"email": email}
	updateDoc := bson.M{"$set": bson.M{
		"name":     update.Name,
		"email":    update.Email,
		"password": update.Password,
		"admin":    update.Admin,
	}}

	_, err := u.Collection.UpdateOne(u.Ctx, filter, updateDoc)
	if err != nil {
		return dataInt.RentalUser{}, err
	}

	return update, nil
}

func (u *UserRepo) DeleteUserById(id uuid.UUID) error {
	log.Printf("👤 [Mongo] DeleteUserById: %s", id.String())
	filter := bson.M{"_id": id.String()}
	_, err := u.Collection.DeleteOne(u.Ctx, filter)
	return err
}

func (u *UserRepo) SaveUser(user dataInt.RentalUser) (dataInt.RentalUser, error) {
	log.Printf("👤 [Mongo] SaveUser: %+v", user)
	user.Id = uuid.New()

	doc := model.RentalUser{
		ID:       user.Id.String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Admin:    user.Admin,
	}

	_, err := u.Collection.InsertOne(u.Ctx, doc)
	if err != nil {
		return dataInt.RentalUser{}, err
	}

	return user, nil
}

// Helper

func NewUserRepo(ctx context.Context, db *mongo.Database) *UserRepo {
	log.Println("📦 [Mongo] NewUserRepo")
	return &UserRepo{
		Collection: db.Collection("users"),
		Ctx:        ctx,
	}
}

func convertModelToDataUser(m model.RentalUser) (dataInt.RentalUser, error) {
	parsedUUID, err := uuid.Parse(m.ID)
	if err != nil {
		return dataInt.RentalUser{}, err
	}

	return dataInt.RentalUser{
		Id:       parsedUUID,
		Name:     m.Name,
		Email:    m.Email,
		Password: m.Password,
		Admin:    m.Admin,
	}, nil
}
