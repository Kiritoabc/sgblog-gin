package vo

import "time"

type CommentVO struct {
	ID                int64        `json:"id"`
	ArticleID         int64        `json:"articleId"`
	RootID            int64        `json:"rootId"`
	Content           string       `json:"content"`
	ToCommentUserID   int64        `json:"toCommentUserId"`
	ToCommentUserName string       `json:"toCommentUserName"`
	ToCommentID       int64        `json:"toCommentId"`
	CreateBy          int64        `json:"createBy"`
	CreateTime        time.Time    `json:"createTime"`
	Username          string       `json:"username"`
	Children          []*CommentVO `json:"children"`
}
