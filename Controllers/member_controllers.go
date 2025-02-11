package Controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	configs "cdn-project/Configs"
	models "cdn-project/models"
	responses "cdn-project/responses"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var MemberCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

var validate = validator.New()

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Expected format: "Bearer <TOKEN>"
	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return ""
	}

	return splitToken[1]
}

func CreateMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var member models.Member
		defer cancel()

		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if validationErr := validate.Struct(&member); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		hashedPassword, err := HashPassword(member.Password)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Error hashing password"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		newUser := models.Member{
			Id:       primitive.NewObjectID(),
			Name:     member.Name,
			Email:    member.Email,
			Password: hashedPassword,
			City:     member.City,
		}

		result, err := MemberCollection.InsertOne(ctx, newUser)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.MemberResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result.InsertedID}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["id"]
		var user models.Member

		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := MemberCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.MemberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}}
		json.NewEncoder(rw).Encode(response)

	}
}

func GetAllMembers() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// ✅ Verify JWT token before allowing access
		tokenString := ExtractToken(r)
		if tokenString == "" {
			rw.WriteHeader(http.StatusUnauthorized)
			response := responses.MemberResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": "Missing token"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// ✅ Validate the token
		_, err := ValidateJWT(tokenString)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			response := responses.MemberResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": "Invalid token"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// ✅ Proceed if token is valid
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var members []models.Member
		defer cancel()

		results, err := MemberCollection.Find(ctx, bson.M{})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// ✅ Read from DB efficiently
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.Member
			if err = results.Decode(&singleUser); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
				return
			}
			members = append(members, singleUser)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.MemberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": members}}
		json.NewEncoder(rw).Encode(response)
	}
}

func UpdateMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["id"]
		var user models.Member

		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		update := bson.M{"name": user.Name, "email": user.Email, "city": user.City}

		result, err := MemberCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// Get Updated member details
		var updatedMember models.Member

		if result.MatchedCount == 1 {
			err := MemberCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedMember)

			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
				return
			}

		}

		rw.WriteHeader(http.StatusOK)
		response := responses.MemberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedMember}}
		json.NewEncoder(rw).Encode(response)

	}
}

func DeleteMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["id"]
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := MemberCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			rw.WriteHeader(http.StatusNotFound)
			response := responses.MemberResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Member Id Not found"}}
			json.NewEncoder(rw).Encode(response)
			return

		}

		rw.WriteHeader(http.StatusOK)
		response := responses.MemberResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": "Member deleted successfully"}}
		json.NewEncoder(rw).Encode(response)

	}
}

func HealthCheck() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		response := responses.MemberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Service is up and running"}}
		json.NewEncoder(rw).Encode(response)
	}
}

func LoginMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		defer cancel()

		if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid request body"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		var user models.Member
		err := MemberCollection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&user)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			response := responses.MemberResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": "Invalid email or password"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if !CheckPassword(user.Password, loginData.Password) {
			rw.WriteHeader(http.StatusUnauthorized)
			response := responses.MemberResponse{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": "Invalid email or password"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// Generate JWT token
		token, err := GenerateJWT(user.Email)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Error generating token"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.MemberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"token": token}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GenerateJWT(email string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
