package Controllers

import (
	configs "cdn-project/Configs"
	middleware "cdn-project/Middleware"
	models "cdn-project/Models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Upload() http.HandlerFunc {
	return middleware.ProtectedHandler(func(w http.ResponseWriter, r *http.Request) {

		// Print JWT token
		tokenString := middleware.ExtractToken(r)
		token, err := middleware.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Erreur lors de la validation du token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Erreur lors de la récupération des claims", http.StatusUnauthorized)
			return
		}
		fmt.Println(claims["email"])

		// Requet userID from db with email
		var user models.Member
		err = configs.MemberCollection.FindOne(context.TODO(), primitive.M{"email": claims["email"]}).Decode(&user)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération de l'utilisateur", http.StatusUnauthorized)
			return
		}

		fmt.Println(user.Id)

		r.ParseMultipartForm(10 << 20)

		// ✅ Récupérer le fichier
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Erreur lors de la récupération du fichier", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// ✅ Créer le dossier `uploads/`
		uploadDir := "uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, os.ModePerm)
		}

		fileId := primitive.NewObjectID()
		filePath := filepath.Join(uploadDir, fileId.Hex()+"_"+handler.Filename)

		// ✅ Sauvegarder le fichier
		destFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Erreur lors de la création du fichier", http.StatusInternalServerError)
			return
		}
		defer destFile.Close()

		if _, err := io.Copy(destFile, file); err != nil {
			http.Error(w, "Erreur lors de la copie du fichier", http.StatusInternalServerError)
			return
		}

		// ✅ Enregistrer en base de données
		newFile := models.File{
			ID:          fileId,
			Filename:    handler.Filename,
			ContentType: handler.Header.Get("Content-Type"),
			Size:        handler.Size,
			UploadDate:  primitive.NewDateTimeFromTime(time.Now()),
			UserId:      user.Id,
		}

		_, err = configs.FileCollection.InsertOne(context.TODO(), newFile)
		if err != nil {
			http.Error(w, "Erreur lors de l'enregistrement en base", http.StatusInternalServerError)
			return
		}

		// ✅ Retourner la réponse JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Fichier uploadé avec succès",
			"file":    newFile,
		})
	})
}

func GetFiles() http.HandlerFunc {
	return middleware.ProtectedHandler(func(w http.ResponseWriter, r *http.Request) {
		// Print JWT token
		tokenString := middleware.ExtractToken(r)
		token, err := middleware.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Erreur lors de la validation du token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Erreur lors de la récupération des claims", http.StatusUnauthorized)
			return
		}
		fmt.Println(claims["email"])

		// Requet userID from db with email
		var user models.Member
		err = configs.MemberCollection.FindOne(context.TODO(), primitive.M{"email": claims["email"]}).Decode(&user)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération de l'utilisateur", http.StatusUnauthorized)
			return
		}

		fmt.Println(user.Id)

		// ✅ Rechercher les fichiers dans la base MongoDB
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var userFiles []models.File
		cursor, err := configs.FileCollection.Find(ctx, bson.M{"user_id": user.Id})
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des fichiers", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		// ✅ Lire les résultats et les ajouter à la liste
		for cursor.Next(ctx) {
			var file models.File
			if err := cursor.Decode(&file); err != nil {
				http.Error(w, "Erreur de décodage des fichiers", http.StatusInternalServerError)
				return
			}
			userFiles = append(userFiles, file)
		}

		// ✅ Vérifier si aucun fichier trouvé
		if len(userFiles) == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Aucun fichier trouvé pour cet utilisateur",
			})
			return
		}

		// ✅ Retourner la liste des fichiers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Fichiers récupérés avec succès",
			"files":   userFiles,
		})
	})
}

// func GetFiles() http.HandlerFunc {
// 	return middleware.ProtectedHandler(func(w http.ResponseWriter, r *http.Request) {

// 		// Print JWT token
// 		tokenString := middleware.ExtractToken(r)
// 		token, err := middleware.ValidateJWT(tokenString)
// 		if err != nil {
// 			http.Error(w, "Erreur lors de la validation du token", http.StatusUnauthorized)
// 			return
// 		}
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			http.Error(w, "Erreur lors de la récupération des claims", http.StatusUnauthorized)
// 			return
// 		}
// 		fmt.Println(claims["email"])

// 		// Requet userID from db with email
// 		var user models.Member
// 		err = configs.MemberCollection.FindOne(context.TODO(), primitive.M{"email": claims["email"]}).Decode(&user)
// 		if err != nil {
// 			http.Error(w, "Erreur lors de la récupération de l'utilisateur", http.StatusUnauthorized)
// 			return
// 		}

// 		fmt.Println(user.Id)

// 		// ✅ Récupérer les fichiers
// 		// Finf id from user
// 		cursor, err := configs.FileCollection.Find(context.TODO(), primitive.M{"user_id": user.Id})
// 		if err != nil {
// 			http.Error(w, "Erreur lors de la récupération des fichiers", http.StatusInternalServerError)
// 			return
// 		}
// 		defer cursor.Close(context.Background())

// 		var files []models.File
// 		for cursor.Next(context.Background()) {
// 			var file models.File
// 			cursor.Decode(&file)
// 			files = append(files, file)
// 		}

// 		// ✅ Retourner la réponse JSON
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"files": files,
// 		})
// 	})
// }
