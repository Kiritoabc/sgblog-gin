package vo

import "time"

type CommentVO struct {
	ID                int64        `json:"id"`
	ArticleID         int64        `json:"article_id"`
	RootID            int64        `json:"root_id"`
	Content           string       `json:"content"`
	ToCommentUserID   int64        `json:"to_comment_user_id"`
	ToCommentUserName string       `json:"to_comment_user_name"`
	ToCommentID       int64        `json:"to_comment_id"`
	CreateBy          int64        `json:"create_by"`
	CreateTime        time.Time    `json:"create_time"`
	Username          string       `json:"username"`
	Children          []*CommentVO `json:"children"`
}
