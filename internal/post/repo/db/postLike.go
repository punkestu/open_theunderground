package db

import (
	"database/sql"
	"github.com/punkestu/open_theunderground/shared/domain"
)

func (p PostDB) GetLikeByPostID(postId string) ([]*domain.PostLike, error) {
	ds, err := p.conn.Query("SELECT u.id, u.username, u.email, u.fullname, pl.created_at FROM post_likes pl JOIN users u ON pl.user_id = u.id WHERE post_id = ?", postId)
	if err != nil {
		return nil, err
	}
	var rows []*domain.PostLike
	for ds.Next() {
		row := domain.PostLike{}
		err := ds.Scan(&row.User.ID, &row.User.Username, &row.User.Email, &row.User.Fullname, &row.CreatedAt)
		if err != nil {
			return nil, err
		}
		rows = append(rows, &row)
	}
	return rows, nil
}

func (p PostDB) GetLikeByPostAndUserID(postId, userId string) (*domain.PostLike, error) {
	row := domain.PostLike{
		Post: &domain.Post{},
	}
	err := p.conn.QueryRow("SELECT u.id, p.id FROM post_likes pl JOIN users u ON pl.user_id = u.id JOIN posts p ON pl.post_id = p.id WHERE post_id = ? AND user_id = ?", postId, userId).Scan(&row.User.ID, &row.Post.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &row, nil
}

func (p PostDB) CreateLike(postId, userId string) (*domain.PostLike, error) {
	if _, err := p.conn.Exec("INSERT INTO post_likes(post_id, user_id) VALUE (?, ?)", postId, userId); err != nil {
		return nil, err
	}

	return p.GetLikeByPostAndUserID(postId, userId)
}

func (p PostDB) DeleteLike(postId, userId string) error {
	if _, err := p.conn.Exec("DELETE FROM post_likes WHERE post_id = ? AND user_id = ?", postId, userId); err != nil {
		return err
	}

	return nil
}
