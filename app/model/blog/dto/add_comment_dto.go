package dto

import "time"

type AddCommentDto struct {
	ID              int64     `json:"id"`
	Type            string    `json:"type" description:"评论类型（0代表文章评论，1代表友链评论）"`
	ArticleID       int64     `json:"article_id" description:"文章id"`
	RootID          int64     `json:"root_id"`
	Content         string    `json:"content"`
	ToCommentUserID int64     `json:"to_comment_user_id"`
	ToCommentID     int64     `json:"to_comment_id"`
	CreateBy        int64     `json:"create_by"`
	CreateTime      time.Time `json:"create_time"`
	UpdateBy        int64     `json:"update_by"`
	UpdateTime      time.Time `json:"update_time"`
	DelFlag         int       `json:"del_flag" description:"删除标志（0代表未删除，1代表已删除）"`
}
