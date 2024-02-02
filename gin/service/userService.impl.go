package service

import (
	"context"
	"errors"
	"example/curd/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func (u *UserServiceImpl) CreateUser(user *model.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*model.User, error) {
	var user *model.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, nil
}

func (u *UserServiceImpl) GetAll() ([]*model.User, error) {
	var users []*model.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)
	if len(users) == 0 {
		return nil, errors.New("document not found")
	}
	return users, nil
}

// 	var users []*model.User
// 	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
// 	if err != nil {
// 		return nil, err
// 	}
// 	for cursor.Next(u.ctx) {
// 		var user model.User
// 		err := cursor.Decode(&user)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, &user)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	cursor.Close(u.ctx)

// 	if len(users) == 0 {
// 		return nil, errors.New("documents not found")
// 	}
// 	return users, nil
// }

func (u *UserServiceImpl) UpdateUser(user *model.User) error {

	filter := bson.D{primitive.E{Key: "name", Value: user.Name}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: user.Name}, primitive.E{Key: "age", Value: user.Age}, primitive.E{Key: "address", Value: user.Adderss}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched documents for delete")
	}
	return nil
}
