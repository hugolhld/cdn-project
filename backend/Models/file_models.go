package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

// FileMetadata - Stores metadata about uploaded files
type File struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Filename    string             `bson:"filename" json:"filename"`
	ContentType string             `bson:"content_type" json:"content_type"`
	Size        int64              `bson:"size" json:"size"`
	UploadDate  primitive.DateTime `bson:"upload_date" json:"upload_date"`
	UserId      primitive.ObjectID `bson:"user_id" json:"user_id"`
}
