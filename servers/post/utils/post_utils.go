package utils

import (
	"post/db"
	"post/middleware"
	"post/schema"
)

type IPostUtils interface {
	GetPostByID(postID uint) (schema.Post, error)
	AddPost(post schema.Post) (schema.Post, error)
}

type PostUtils struct {
	DB          db.Database
	RedisClient middleware.RedisClientInterface
}

// NewPostUtils creates a new PostUtils
func NewPostUtils(DB db.Database, redisClient middleware.RedisClientInterface) *PostUtils {
	return &PostUtils{
		DB:          DB,
		RedisClient: redisClient,
	}
}

// func GetPostByID retrieves a post by its ID
func (pu *PostUtils) GetPostByID(postID uint) (schema.Post, error) {
	var post schema.Post
	err := pu.DB.First(&post, postID).Error
	if err != nil {
		return schema.Post{}, err
	}

	return post, nil
}

// func AddPost adds a post to the database
func (pu *PostUtils) AddPost(post schema.Post) (schema.Post, error) {
	err := pu.DB.Create(&post).Error
	if err != nil {
		return schema.Post{}, err
	}

	return post, nil
}
