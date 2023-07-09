package db

import (
	"database/sql"
	"github.com/punkestu/open_theunderground/shared/domain"
	"github.com/punkestu/open_theunderground/shared/exception"
	"github.com/savsgio/gotils/uuid"
)

func (p PostDB) GetAll() (*[]*domain.Post, error) {
	rows, err := p.conn.Query("SELECT p.id, p.topic, p.created_at, u.id, u.fullname, u.username, u.email FROM posts p JOIN users u on u.id = p.author_id")
	if err != nil {
		return nil, err
	}
	var mPosts []*domain.Post
	for rows.Next() {
		var mPost domain.Post
		err := rows.Scan(&mPost.ID, &mPost.Topic, &mPost.CreatedAt, &mPost.Author.ID, &mPost.Author.Fullname, &mPost.Author.Username, &mPost.Author.Email)
		if err != nil {
			return nil, err
		}
		mPosts = append(mPosts, &mPost)
	}
	return &mPosts, nil
}
func (p PostDB) GetByID(postId string) (*domain.Post, error) {
	var mPost domain.Post
	err := p.conn.QueryRow("SELECT p.id, p.topic, p.created_at, u.id, u.fullname, u.username, u.email FROM posts p JOIN users u on u.id = p.author_id WHERE p.id=?", postId).Scan(&mPost.ID, &mPost.Topic, &mPost.CreatedAt, &mPost.Author.ID, &mPost.Author.Fullname, &mPost.Author.Username, &mPost.Author.Email)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, exception.New("postId", "Id post is not found")
		}
		return nil, err
	}
	return &mPost, nil
}
func (p PostDB) GetByAuthor(authorId string) (*[]*domain.Post, error) {
	rows, err := p.conn.Query("SELECT p.id, p.topic, p.created_at, u.id, u.fullname, u.username, u.email FROM posts p JOIN users u on u.id = p.author_id WHERE u.id=?", authorId)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, exception.New("postId", "Id post is not found")
		}
		return nil, err
	}
	var mPosts []*domain.Post
	for rows.Next() {
		var mPost domain.Post
		err := rows.Scan(&mPost.ID, &mPost.Topic, &mPost.CreatedAt, &mPost.Author.ID, &mPost.Author.Fullname, &mPost.Author.Username, &mPost.Author.Email)
		if err != nil {
			return nil, err
		}
		mPosts = append(mPosts, &mPost)
	}
	return &mPosts, nil
}
func (p PostDB) Create(topic, authorId string) (*domain.Post, error) {
	id := uuid.V4()
	_, err := p.conn.Exec("INSERT INTO posts(id, topic, author_id) VALUE (?, ?, ?)", id, topic, authorId)
	if err != nil {
		return nil, err
	}

	return p.GetByID(id)
}
func (p PostDB) Update(_ string) (*domain.Post, error) {
	return nil, nil
}
