package admin

import (
	"github.com/jinzhu/copier"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/dto"
	"sgblog-go/app/model/blog/vo"
)

type TagService struct{}

func (s *TagService) PageTagList(tagListDto dto.TagListDto,
	pageNum int, pageSize int) ([]*blog.SgTag, int64, error) {
	var tagList []*blog.SgTag

	query := global.SG_BLOG_DB.Model(&blog.SgTag{})

	// build query
	if tagListDto.Name != "" {
		query = query.
			Where("name like ?", "%"+tagListDto.Name+"%")
	}
	if tagListDto.Remark != "" {
		query = query.
			Where("remark like ?", "%"+tagListDto.Remark+"%")
	}

	// page query
	offset := (pageNum - 1) * pageSize
	limit := pageSize

	var total int64
	if err := query.Count(&total).Offset(offset).Limit(limit).Find(&tagList).Error; err != nil {
		return nil, 0, err
	}

	return tagList, total, nil
}

func (s *TagService) ListAllTag() ([]*vo.TagVo, error) {

	var tagList []*blog.SgTag
	var tagVoList []*vo.TagVo

	if err := global.SG_BLOG_DB.Model(&blog.SgTag{}).Find(&tagList).Error; err != nil {
		return nil, err
	}

	err := copier.Copy(&tagVoList, &tagList)
	if err != nil {
		return nil, err
	}
	return tagVoList, nil
}
