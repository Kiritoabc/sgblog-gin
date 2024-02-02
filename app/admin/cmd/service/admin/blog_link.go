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
