package db

import (
	"database/sql"
	"github.com/punkestu/open_theunderground/shared/domain"
	"github.com/punkestu/open_theunderground/shared/error/invalid"
	"github.com/savsgio/gotils/uuid"
)

func (p PostDB) CreateComment(postId, userId, comment string) (*domain.PostComment, error) {
	id := uuid.V4()
	_, err := p.conn.Exec("INSERT INTO post_comments(id, comment, user_id, post_id) VALUE (?, ?, ?, ?)", id, comment, userId, postId)
	if err != nil {
		return nil, err
	}

	return p.GetCommentByID(id)
}

func (p PostDB) GetCommentByID(commentId string) (*domain.PostComment, error) {
	var mPostComment domain.PostComment
	err := p.conn.QueryRow("SELECT pc.id, pc.comment, pc.created_at, u.id, u.fullname, u.username, u.email FROM post_comments pc JOIN users u on u.id = pc.user_id WHERE pc.id=?", commentId).Scan(&mPostComment.ID, &mPostComment.Comment, &mPostComment.CreatedAt, &mPostComment.User.ID, &mPostComment.User.Fullname, &mPostComment.User.Username, &mPostComment.User.Email)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, invalid.New("postId", "Id post is not found")
		}
		return nil, err
	}
	return &mPostComment, nil
}

func (p PostDB) GetCommentByPostID(postId string) ([]*domain.PostComment, error) {
	rows, err := p.conn.Query("SELECT pc.id, pc.comment, pc.created_at, u.id, u.fullname, u.username, u.email FROM post_comments pc JOIN users u on u.id = pc.user_id WHERE pc.post_id=?", postId)
	if err != nil {
		return nil, err
	}
	var mPostComments []*domain.PostComment
	for rows.Next() {
		var mPostComment domain.PostComment
		err := rows.Scan(&mPostComment.ID, &mPostComment.Comment, &mPostComment.CreatedAt, &mPostComment.User.ID, &mPostComment.User.Fullname, &mPostComment.User.Username, &mPostComment.User.Email)
		if err != nil {
			return nil, err
		}
		mPostComments = append(mPostComments, &mPostComment)
	}
	return mPostComments, nil
}
