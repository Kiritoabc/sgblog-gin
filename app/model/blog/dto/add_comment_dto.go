package dto

import "time"

type AddCommentDto struct {
	ID              int64     `json:"id"`
	Type            int       `json:"type" description:"评论类型（0代表文章评论，1代表友链评论）"`
	ArticleID       int64     `json:"articleId" description:"文章id"`
	RootID          int64     `json:"rootId"`
	Content         string    `json:"content"`
	ToCommentUserID int64     `json:"toCommentUserId"`
	ToCommentID     int64     `json:"toCommentId"`
	CreateBy        int64     `json:"createBy"`
	CreateTime      time.Time `json:"createTime"`
	UpdateBy        int64     `json:"updateBy"`
	UpdateTime      time.Time `json:"updateTime"`
	DelFlag         int       `json:"delFlag" description:"删除标志（0代表未删除，1代表已删除）"`
}
