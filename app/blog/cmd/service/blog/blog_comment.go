package blog

import (
	"errors"
	"github.com/jinzhu/copier"
	"sgblog-go/app/blog/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
	"sgblog-go/common/constants"
)

type CommentService struct{}

func (s *CommentService) CommentList(commentType string,
	articleId int64, pageNum int, pageSize int) ([]*vo.CommentVO, int64, error) {
	var comments []*blog.SgComment
	var total int64
	// 1.查询对应文章的评论
	tx := global.SG_BLOG_DB.Model(&blog.SgComment{})
	// 2.对articleId进行判断
	if constants.ArticleComment == commentType {
		tx = tx.Where("article_id = ?", articleId)
	}
	// 3.根评论 rootId 为-1
	// 4.评论类型
	// 5. 分页查询
	err := tx.Where("root_id = ? and type = ?", -1, commentType).
		Order("create_time asc").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Count(&total).
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	// 6.转换成 commentVoList
	commentVoList, err := s.toCommentVoList(comments)
	if err != nil {
		return nil, 0, err
	}
	// 7.查询所有根评论对应的子评论集合，并且赋值给对应的属性
	for _, commentVo := range commentVoList {
		children, err := s.getChildren(commentVo.ID)
		if err != nil {
			return nil, 0, err
		}
		// 赋值
		commentVo.Children = children
	}
	return commentVoList, total, nil
}

func (s *CommentService) toCommentVoList(comments []*blog.SgComment) ([]*vo.CommentVO, error) {
	var commentVoList = []*vo.CommentVO{}
	copier.Copy(&commentVoList, &comments)
	for _, commentVo := range commentVoList {
		var user = blog.SysUser{}
		err := global.SG_BLOG_DB.Model(&blog.SysUser{}).
			Where("id = ?", commentVo.CreateBy).
			Find(&user).Error
		if err != nil {
			return nil, err
		}

		commentVo.Username = user.NickName

		// 如果toCommentUserId 不为-1才进行查询
		if commentVo.ToCommentUserID != -1 {
			err := global.SG_BLOG_DB.Model(&blog.SysUser{}).
				Where("id = ?", commentVo.CreateBy).
				Find(&user).Error
			if err != nil {
				return nil, err
			}
			commentVo.ToCommentUserName = user.NickName
		}
	}
	return commentVoList, nil
}

func (s *CommentService) getChildren(id int64) (vos []*vo.CommentVO, err error) {
	var comments []*blog.SgComment
	if err = global.SG_BLOG_DB.Model(&blog.SgComment{}).
		Where("root_id", id).
		Order("create_time desc").
		Find(&comments).Error; err != nil {
		return
	}
	vos, err = s.toCommentVoList(comments)
	if err != nil {
		return
	}
	return
}

func (s *CommentService) AddComment(comment *blog.SgComment) error {
	if comment.Content == "" {
		return errors.New("评论内容不能为空")
	}
	err := global.SG_BLOG_DB.Save(comment).Error
	if err != nil {
		return err
	}
	return nil
}
