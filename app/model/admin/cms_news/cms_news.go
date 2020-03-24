// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package cms_news

import (
	"gfast/app/model/admin/cms_category_news"
	"gfast/app/model/admin/user"
	"gfast/library/service"
	"gfast/library/utils"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

// Fill with you ideas below.

//添加文章参数
type ReqAddParams struct {
	NewsStatus    uint   `p:"status"    v:"in:0,1#状态只能为0或1"`       // 状态;1:已发布;0:未发布;
	IsTop         uint   `p:"IsTop"         v:"in:0,1#置顶只能为0或1"`   // 是否置顶;1:置顶;0:不置顶
	Recommended   uint   `p:"recommended"    v:"in:0,1#推荐只能为0或1"`  // 是否推荐;1:推荐;0:不推荐
	PublishedTime string `p:"published_time"`                      // 发布时间
	NewsTitle     string `p:"title"     v:"required#标题不能为空"`       // post标题
	NewsKeywords  string `p:"keywords"`                            // seo keywords
	NewsExcerpt   string `p:"excerpt"`                             // post摘要
	NewsSource    string `p:"source"  `                            // 转载文章的来源
	NewsContent   string `p:"content"   v:"required#文章内容不能为空"`     // 文章内容
	Thumbnail     string `p:"thumbnail"    `                       // 缩略图
	IsJump        uint   `p:"IsJump"        v:"in:0,1#跳转类型只能为0或1"` // 是否跳转地址
	JumpUrl       string `p:"JumpUrl"      `                       // 跳转地址
}

//文章搜索参数
type ReqListSearchParams struct {
	CateId             []int  `p:"cateId"`
	PublishedTimeStart string `p:"pubTimeStart"`
	PublishedTimeEnd   string `p:"pubTimeEnd"`
	KeyWords           string `p:"keyWords"`
	PageNum            int    `p:"page"`     //当前页码
	PageSize           int    `p:"pageSize"` //每页数
}

type ReqEditParams struct {
	Id int `p:"id" v:"integer|min:1#文章ID只能为整数|文章ID只能为正数"`
	ReqAddParams
}

//添加文章操作
func AddNews(req *ReqAddParams, cateIds []int, userId int) (insId int64, err error) {
	if len(cateIds) == 0 {
		err = gerror.New("栏目不能为空")
		return
	}
	tx, err := g.DB().Begin()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("添加失败")
		return
	}
	entity := &Entity{
		UserId:        gconv.Uint64(userId),
		NewsStatus:    req.NewsStatus,
		IsTop:         req.IsTop,
		Recommended:   req.Recommended,
		CreateTime:    gconv.Uint(gtime.Timestamp()),
		PublishedTime: gconv.Uint(utils.StrToTimestamp(req.PublishedTime)),
		NewsTitle:     req.NewsTitle,
		NewsKeywords:  req.NewsKeywords,
		NewsExcerpt:   req.NewsExcerpt,
		NewsSource:    req.NewsExcerpt,
		NewsContent:   req.NewsContent,
		Thumbnail:     req.Thumbnail,
		IsJump:        req.IsJump,
		JumpUrl:       req.JumpUrl,
	}
	res, e := entity.Save()
	if e != nil {
		g.Log().Error(e)
		err = gerror.New("添加文章失败")
		tx.Rollback()
		return
	}
	insId, err = res.LastInsertId()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("添加文章失败")
		tx.Rollback()
		return
	}
	//保存栏目与文章关联信息
	catNewsEntity := make([]cms_category_news.Entity, len(cateIds))
	for k, cateId := range cateIds {
		catNewsEntity[k].CategoryId = gconv.Uint64(cateId)
		catNewsEntity[k].NewsId = gconv.Uint64(insId)
	}
	_, err = cms_category_news.Model.Data(catNewsEntity).Insert()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("添加文章失败")
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

//修改文章操作
func EditNews(req *ReqEditParams, cateIds []int) (err error) {
	if len(cateIds) == 0 {
		err = gerror.New("栏目不能为空")
		return
	}
	tx, err := g.DB().Begin()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("添加失败")
		return
	}
	entity, err := Model.FindOne("id", req.Id)
	if err != nil {
		g.Log().Error(err)
	}
	if err != nil || entity == nil {
		err = gerror.New("文章信息获取失败")
		return
	}
	entity.NewsStatus = req.NewsStatus
	entity.IsTop = req.IsTop
	entity.Recommended = req.Recommended
	entity.UpdateTime = gconv.Uint(gtime.Timestamp())
	entity.PublishedTime = gconv.Uint(utils.StrToTimestamp(req.PublishedTime))
	entity.NewsTitle = req.NewsTitle
	entity.NewsKeywords = req.NewsKeywords
	entity.NewsExcerpt = req.NewsExcerpt
	entity.NewsSource = req.NewsExcerpt
	entity.NewsContent = req.NewsContent
	entity.Thumbnail = req.Thumbnail
	entity.IsJump = req.IsJump
	entity.JumpUrl = req.JumpUrl
	_, err = entity.Update()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("修改文章失败")
		tx.Rollback()
		return
	}
	//删除旧的栏目文章关联信息
	cnList, err := cms_category_news.GetCategoriesByNewsId(entity.Id)
	if err != nil {
		return
	}
	for _, cn := range cnList {
		_, err = cn.Delete()
		if err != nil {
			g.Log().Error(err)
			err = gerror.New("更新文章栏目所属信息失败")
			tx.Rollback()
			return
		}
	}
	//保存栏目与文章关联信息
	catNewsEntity := make([]cms_category_news.Entity, len(cateIds))
	for k, cateId := range cateIds {
		catNewsEntity[k].CategoryId = gconv.Uint64(cateId)
		catNewsEntity[k].NewsId = gconv.Uint64(req.Id)
	}
	_, err = cms_category_news.Model.Data(catNewsEntity).Insert()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("更新文章栏目所属信息失败")
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

//文章列表查询
func ListByPage(req *ReqListSearchParams) (total, page int, list gdb.Result, err error) {
	model := g.DB().Table(Table + " news")
	if req != nil {
		if len(req.CateId) > 0 {
			model = model.InnerJoin(cms_category_news.Table+" cn", "cn.news_id=news.id").Where("cn.category_id in(?)", req.CateId)
			model = model.Group("cn.news_id")
		}
		if req.KeyWords != "" {
			model = model.Where("news.news_title like ?", "%"+req.KeyWords+"%")
		}
		if req.PublishedTimeStart != "" {
			model = model.Where("news.published_time >=?", utils.StrToTimestamp(req.PublishedTimeStart))
		}
		if req.PublishedTimeEnd != "" {
			model = model.Where("news.published_time <=?", utils.StrToTimestamp(req.PublishedTimeEnd))
		}
	}
	model = model.LeftJoin(user.Table+" user", "news.user_id=user.id")
	total, err = model.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
		return
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	page = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = service.AdminPageNum
	}

	list, err = model.Page(page, req.PageSize).Fields("news.*,user.user_nickname").Order("published_time desc,news.id desc").All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	return
}

//通过文章id获取文章信息
func GetById(id int) (news *Entity, err error) {
	news, err = Model.FindOne("id", id)
	if err != nil {
		g.Log().Error(err)
	}
	if err != nil || news == nil {
		err = gerror.New("获取文章信息失败")
		return
	}
	return
}

func DeleteByIds(ids []int) error {
	_, err := Model.Delete("id in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("删除失败")
	}
	cms_category_news.Delete("news_id in (?)", ids)
	return nil
}
