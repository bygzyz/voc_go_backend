package models

type Article struct {
	Model

	EsDataId      string `gorm:"comment:'es存储id';size:128;not null;default:''" json:"es_data_id"`
	Title         string `gorm:"comment:'标题';size:256;not null;default:''" json:"title"`
	Content       string `gorm:"comment:'内容';size:9999;not null;default:''" json:"content"`
	ImageUrl      string `gorm:"comment:'图片url';size:512;not null;default:''" json:"image_url"`
	HotScore      int    `gorm:"comment:'热度值';not null;default:0" json:"hot_score"`
	ArticleOrigin string `gorm:"comment:'文章来源';size:16;not null;default:'case'" json:"article_origin"` // information：情报 case：案例
	ArticleType   string `gorm:"comment:'文章类型';size:32;not null;default:'其他'" json:"article_type"`     // 风险、业务等由管理员设置
	ViewCount     int    `gorm:"comment:'被浏览次数';not null;default:0" json:"view_count"`
	SpreadCount   int    `gorm:"comment:'转发次数';not null;default:0" json:"spread_count"`
}

//type Article struct {
//	Model
//
//	TagID int `json:"tag_id" gorm:"index"`
//	Tag   Tag `json:"tag"`
//
//	Title         string `json:"title"`
//	Desc          string `json:"desc"`
//	Content       string `json:"content"`
//	CoverImageUrl string `json:"cover_image_url"`
//	CreatedBy     string `json:"created_by"`
//	ModifiedBy    string `json:"modified_by"`
//	State         int    `json:"state"`
//}

//
//// GetRecommendArticle 获取推荐文章列表
//func GetRecommendArticle(userId int, articleOrigin string) ([]*Article, error) {
//	// 收藏的文章id
//	var collection Collection
//	err := initialize.db.Select("id")
//
//}
//
//// ExistArticleByID checks if an article exists based on ID
//func ExistArticleByID(id int) (bool, error) {
//	var article Article
//	err := initialize.db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return false, err
//	}
//
//	if article.ID > 0 {
//		return true, nil
//	}
//
//	return false, nil
//}
//
//// GetArticleTotal gets the total number of articles based on the constraints
//func GetArticleTotal(maps interface{}) (int, error) {
//	var count int
//	if err := initialize.db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
//		return 0, err
//	}
//
//	return count, nil
//}
//
//// GetArticles gets a list of articles based on paging constraints
//func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
//	var articles []*Article
//	err := initialize.db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, err
//	}
//
//	return articles, nil
//}
//
//// GetArticle Get a single article based on ID
//func GetArticle(id int) (*Article, error) {
//	var article Article
//	err := initialize.db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, err
//	}
//
//	err = initialize.db.Model(&article).Related(&article.Tag).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, err
//	}
//
//	return &article, nil
//}
//
//// EditArticle modify a single article
//func EditArticle(id int, data interface{}) error {
//	if err := initialize.db.Model(&Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// AddArticle add a single article
//func AddArticle(data map[string]interface{}) error {
//	article := Article{
//		TagID:         data["tag_id"].(int),
//		Title:         data["title"].(string),
//		Desc:          data["desc"].(string),
//		Content:       data["content"].(string),
//		CreatedBy:     data["created_by"].(string),
//		State:         data["state"].(int),
//		CoverImageUrl: data["cover_image_url"].(string),
//	}
//	if err := initialize.db.Create(&article).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// DeleteArticle delete a single article
//func DeleteArticle(id int) error {
//	if err := initialize.db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// CleanAllArticle clear all article
//func CleanAllArticle() error {
//	if err := initialize.db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
