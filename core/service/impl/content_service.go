
package impl

/**
 * Copyright 2015 @ to2.net.
 * name : content_service
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */

import (
	"context"
	"go2o/core/domain/interface/content"
	"go2o/core/domain/interface/merchant"
	"go2o/core/query"
	"go2o/core/service/proto"
	"strconv"
)

var _ proto.ContentServiceServer = new(contentService)

type contentService struct {
	_contentRepo content.IContentRepo
	_query       *query.ContentQuery
	_sysContent  content.IContent
	serviceUtil
}

func NewContentService(rep content.IContentRepo, q *query.ContentQuery) *contentService {
	return &contentService{
		_contentRepo: rep,
		_query:       q,
		_sysContent:  rep.GetContent(0),
	}
}

func (cs *contentService) QueryPagingArticles(ctx context.Context, r *proto.PagingArticleRequest) (*proto.SPagingResult, error) {
	var total = 0
	var rows []*content.Article
	ic := cs._sysContent.ArticleManager().GetCategoryByAlias(r.Cat)
	if ic != nil && ic.GetDomainId() > 0 {
		total, rows = cs._query.PagedArticleList(ic.GetDomainId(), int(r.Begin), int(r.Size), "")
	}
	var arr = make([]*proto.SArticle, 0)
	for _, v := range rows {
		arr = append(arr, cs.parseArticle(v))
	}
	return cs.pagingResult(total, arr), nil
}

func (cs *contentService) QueryTopArticles(ctx context.Context, cat *proto.String) (r *proto.ArticlesResponse, err error) {
	var arr = make([]*proto.SArticle, 0)
	ic := cs._sysContent.ArticleManager().GetCategoryByAlias(cat.Value)
	if ic != nil && ic.GetDomainId() > 0 {
		_, rows := cs._query.PagedArticleList(ic.GetDomainId(), 0, 1, "")
		for _, v := range rows {
			arr = append(arr, cs.parseArticle(v))
		}
	}
	return &proto.ArticlesResponse{
		List: arr,
	}, nil
}

func (cs *contentService) parseArticle(v *content.Article) *proto.SArticle {
	return &proto.SArticle{
		Id:          v.ID,
		CatId:       v.CatId,
		Title:       v.Title,
		SmallTitle:  v.SmallTitle,
		Thumbnail:   v.Thumbnail,
		PublisherId: int32(v.PublisherId),
		Location:    v.Location,
		Priority:    int32(v.Priority),
		AccessKey:   v.AccessKey,
		Content:     v.Content,
		Tags:        v.Tags,
		ViewCount:   int32(v.ViewCount),
		SortNum:     int32(v.SortNum),
		CreateTime:  int32(v.CreateTime),
		UpdateTime:  int32(v.UpdateTime),
	}
}

// 获取页面
//todo: 取消mchId
func (cs *contentService) GetPage(mchId, id int32) *content.Page {
	c := cs._contentRepo.GetContent(mchId)
	page := c.GetPage(id)
	if page != nil {
		return page.GetValue()
	}
	return nil
}

// 根据标识获取页面
//todo: 取消mchId
func (cs *contentService) GetPageByIndent(userId int32, indent string) *content.Page {
	c := cs._contentRepo.GetContent(userId)
	page := c.GetPageByStringIndent(indent)
	if page != nil {
		return page.GetValue()
	}
	return nil
}

// 保存页面
func (cs *contentService) SavePage(mchId int32, v *content.Page) (int32, error) {
	c := cs._contentRepo.GetContent(mchId)
	var page content.IPage
	var err error
	if v.UserId != mchId {
		return -1, merchant.ErrMerchantNotMatch
	}

	if v.Id > 0 {
		page = c.GetPage(v.Id)
	} else {
		page = c.CreatePage(v)
	}
	err = page.SetValue(v)
	if err != nil {
		return 0, err
	}
	return page.Save()
}

// 删除页面
func (cs *contentService) DeletePage(mchId int32, pageId int32) error {
	c := cs._contentRepo.GetContent(mchId)
	return c.DeletePage(pageId)
}

// 获取所有栏目
func (cs *contentService) GetArticleCategories() []*content.ArticleCategory {
	list := cs._sysContent.ArticleManager().GetAllCategory()
	arr := make([]*content.ArticleCategory, len(list))
	for i, v := range list {
		val := v.GetValue()
		arr[i] = &val
	}
	return arr
}

// 获取文章栏目,可传入ID或者别名
func (cs *contentService) GetArticleCategory(cat string) content.ArticleCategory {
	mgr := cs._sysContent.ArticleManager()
	catId, err := strconv.Atoi(cat)
	var c content.ICategory
	if err == nil {
		c = mgr.GetCategory(int32(catId))
	} else {
		c = mgr.GetCategoryByAlias(cat)
	}
	if c != nil {
		return c.GetValue()
	}
	return content.ArticleCategory{}
}

// 保存文章栏目
func (cs *contentService) SaveArticleCategory(v *content.ArticleCategory) (int32, error) {
	m := cs._sysContent.ArticleManager()
	c := m.GetCategory(v.ID)
	if c == nil {
		c = m.CreateCategory(v)
	}
	err := c.SetValue(v)
	if err == nil {
		return c.Save()
	}
	return -1, err
}

// 删除文章分类
func (cs *contentService) DeleteArticleCategory(categoryId int32) error {
	return cs._sysContent.ArticleManager().DelCategory(categoryId)
}

// 获取文章
func (cs *contentService) GetArticle(id int32) *content.Article {
	a := cs._sysContent.ArticleManager().GetArticle(id)
	if a != nil {
		v := a.GetValue()
		return &v
	}
	return nil
}

// 删除文章
func (cs *contentService) DeleteArticle(id int32) error {
	return cs._sysContent.ArticleManager().DeleteArticle(id)
}

// 保存文章
func (cs *contentService) SaveArticle(e *content.Article) (int32, error) {
	m := cs._sysContent.ArticleManager()
	a := m.GetArticle(e.ID)
	if a == nil {
		a = m.CreateArticle(e)
	}
	err := a.SetValue(e)
	if err == nil {
		return a.Save()
	}
	return -1, err
}

func (cs *contentService) PagedArticleList(catAlias string, begin, size int,
	where string) (int, []*content.Article) {
	cat := cs._sysContent.ArticleManager().GetCategoryByAlias(catAlias)
	if cat == nil || cat.GetDomainId() <= 0 {
		return 0, []*content.Article{}
	}
	return cs._query.PagedArticleList(cat.GetDomainId(), begin, size, where)
}