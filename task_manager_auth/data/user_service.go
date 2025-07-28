...package data

import (
    "context"
    "errors"
    "time"

    "github.com/Lencho123/A2SV-Backend-Track-Tasks/models"
    "github.com/golang-jwt/jwt/v5"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection // set in init function
var jwtKey = []byte("secret") // move to env in production

func RegisterUser(ctx context.Context, user models.User) (*models.User, error) {
    count, _ := UserCollection.CountDocuments(ctx, bson.M{})
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)
    user.Role = "user"
    if count == 0 {
        user.Role = "admin"
    }
    _, err := UserCollection.InsertOne(ctx, user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func AuthenticateUser(ctx context.Context, login models.LoginInput) (string, error) {
    var user models.User
    err := UserCollection.FindOne(ctx, bson.M{"username": login.Username}).Decode(&user)
    if err != nil {
        return "", errors.New("user not found")
    }
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
    if err != nil {
        return "", errors.New("invalid credentials")
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "role":     user.Role,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })
    return token.SignedString(jwtKey)
}

func PromoteUser(ctx context.Context, username string) error {
    _, err := UserCollection.UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"role": "admin"}})
    return err
}