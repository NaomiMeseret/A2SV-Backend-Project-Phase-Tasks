package repositories

import (
	"context"
	"errors"
	domain "task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//mongoUserReposistory is the MongoDB implemntation of the UserRepository interface
type mongoUserRepository struct{
	collection *mongo.Collection
}
//NewMongoUserRepository creates a new UserRepository with the given MongoDB collection
func NewMongoUserReposistory(collection *mongo.Collection) domain.UserRepository{
	return &mongoUserRepository{collection:collection}
}

func (r *mongoUserRepository) CreateUser(user *domain.User)error{
	userData:=bson.M{
		"email":user.Email,
		"password":user.Password , 
		"role":user.Role,
	}
	result  , err:=r.collection.InsertOne(context.TODO() , userData)
	if err != nil{
		return err
	}
	objID , ok := result .InsertedID.(primitive.ObjectID)
	if ok{
		user.ID =objID.Hex()
	}
	return nil
}

func (r *mongoUserRepository)UserExists(email string) (bool , error){
	filter := bson.M{"email":email}
	err:=r.collection.FindOne(context.TODO(),filter).Err()
	if err == mongo.ErrNoDocuments{
		return false, nil
	}
	if err != nil{
		return false , err
	}
	return true, nil
}

func (r *mongoUserRepository)GetUserByID(id string) (*domain.User, error){
	objID , err:=primitive.ObjectIDFromHex(id)
	if err!=nil{
		return nil , errors.New("invaild user ID")
	}
	filter:=bson.M{"_id":objID}
	var userData bson.M
	err = r.collection.FindOne(context.TODO() , filter).Decode(&userData)
	if err != nil{
		return nil , err
	}
	user:=&domain.User{
		ID: id,
		Email: userData["Email"].(string),
		Password: userData["password"].(string),
		Role: userData["role"].(string),	
	}
	return user, nil
}

func (r *mongoUserRepository)CountUsers()(int , error){
	count , err :=r.collection.CountDocuments(context.TODO() , bson.M{})
	return int(count) , err
}

func (r *mongoUserRepository) GetUserByEmail(email string) (*domain.User, error) {
    filter := bson.M{"email": email}
    var userData bson.M
    err := r.collection.FindOne(context.TODO(), filter).Decode(&userData)
    if err != nil {
        return nil, err
    }
    id := ""
    if oid, ok := userData["_id"].(primitive.ObjectID); ok {
        id = oid.Hex()
    }
    user := &domain.User{
        ID:       id,
        Email:    userData["email"].(string),
        Password: userData["password"].(string),
        Role:     userData["role"].(string),
    }
    return user, nil
}

func (r *mongoUserRepository)GetAllUsers()([]*domain.User , error){
	cursor , err := r.collection.Find(context.TODO(), bson.M{})
	if err !=nil{
		return nil , err
	}
	defer cursor.Close(context.TODO())
	var users []*domain.User
	for cursor.Next(context.TODO()){
		var userData bson.M
		if err:= cursor.Decode(&userData);err!=nil{
			continue
		}
		id :=""
		if objID , ok:= userData["_id"].(primitive.ObjectID);ok{
			id = objID.Hex()
		}
		user:= &domain.User{
			ID: id,
			Email: userData["email"].(string),
			Password: userData["password"].(string),
			Role: userData["role"].(string),
		}
		users = append(users, user)

	}
	return users, nil
}

func (r *mongoUserRepository)PromoteUser (id string) error{
	objID , err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}
	filter :=bson.M{"_id":objID}
	update :=bson.M{"$set":bson.M{"role":"admin"}}
	_ , err =r.collection.UpdateOne(context.TODO() , filter, update)
	return err
}
