package admin

import (
	"github.com/jinzhu/copier"
	"sgblog-go/app/admin/cmd/global"
	"sgblog-go/app/model/blog"
	"sgblog-go/app/model/blog/vo"
	"sgblog-go/common/constants"
)

type LinkService struct{}

func (s *LinkService) GetAllLink() ([]*vo.LinkVo, error) {
	// 查询所有审核通过的
	var linkVoList []*vo.LinkVo

	var links []blog.SgLink
	err := global.SG_BLOG_DB.Model(&blog.SgLink{}).
		Where("status = ?", constants.LinkStatusNormal).
		Find(&links).Error

	if err != nil {
		return nil, err
	}

	err = copier.Copy(&linkVoList, &links)
	if err != nil {
		return nil, err
	}

	return linkVoList, nil
}

func (s *LinkService) SelectLinkPage(link blog.SgLink,
	pageNum int, pageSize int) ([]*blog.SgLink, int64, error) {
	var links []*blog.SgLink

	query := global.SG_BLOG_DB.Model(&blog.SgLink{})

	// build query
	if link.Name != "" {
		query = query.Where("name LIKE ?", "%"+link.Name+"%")
	}
	if link.Status == "" {
		query = query.Where("status = ?", link.Status)
	}

	// page query
	offset := (pageNum - 1) * pageSize
	limit := pageSize
	var total int64
	if err := query.Count(&total).Offset(offset).
		Limit(limit).Find(&links).Error; err != nil {
		return nil, 0, err
	}

	// return results
	return links, total, nil
}
