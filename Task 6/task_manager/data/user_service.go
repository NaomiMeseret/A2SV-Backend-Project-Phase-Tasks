package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)
var UserCollection *mongo.Collection

func CreateUser(user models.User)(*models.User , error){
	ctx, cancel :=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Input validation
	if user.UserName == "" {
		return nil, errors.New("username cannot be empty")
	}
	if user.Password == "" {
		return nil, errors.New("password cannot be empty")
	}
	if len(user.Password) < 4 {
		return nil, errors.New("password must be at least 4 characters long")
	}
	
	userNameCount , _ := UserCollection.CountDocuments(ctx,bson.M{"username":user.UserName})
	if userNameCount>0{
		return nil, errors.New("username already taken")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil{
		return nil , err
	}
	user.Password = string(hash)
	totalUser  ,_:=UserCollection.CountDocuments(ctx, bson.M{})
	if totalUser == 0 {
		user.Role = "admin"
	}else{
		user.Role = "user"
	}
	res, err := UserCollection.InsertOne(ctx, user)
	if err !=nil{
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return &user, nil	
}
func AuthenticateUser(username, password string)(*models.User, error){
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	
	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.User
	err:=UserCollection.FindOne(ctx, bson.M{"username":username}).Decode(&user)
	if err != nil {
		return nil , errors.New("invalid username or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err!=nil{
		return nil , errors.New("invalid username or password")
	}
	return &user , nil

}
func PromoteUser(id string)error{
	objID , err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}
	ctx, cancel:=context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancel()
	_, err = UserCollection.UpdateOne(ctx, bson.M{"_id":objID},
	bson.M{"$set":bson.M{"role":"admin"}})
	return err
}

func GetAllUsers() ([]models.User, error) {
    var users []models.User
    cursor, err := UserCollection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())
    for cursor.Next(context.TODO()) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}