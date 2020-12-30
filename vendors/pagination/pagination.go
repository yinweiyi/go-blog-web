package pagination

import (
	"blog/vendors/config"
	"blog/vendors/types"
	"math"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

// PageData 同视图渲染的数据
type PagerData struct {
	//当前页
	CurrentPage int
	//每页数量
	PerPage int
	// 数据库的内容总数量
	TotalCount int64
	// 总页数
	TotalPage int
	// 基础路径
	BaseUrl string
	//分页名称
	PageName string
}

// Pagination 分页对象
type Pagination struct {
	PerPage     int
	Page        int
	Count       int64
	db          *gorm.DB
	association string
	baseUrl     string
	pageName    string
}

// New 分页对象构建器
// r —— 用来获取分页的 URL 参数，默认是 page，可通过 config/pagination.go 修改
// db —— GORM 查询句柄，用以查询数据集和获取数据总数
// PerPage —— 每页条数，传参为小于或者等于 0 时为默认值  10，可通过 config/pagination.go 修改
// association —— 关联的对象
func New(r *http.Request, db *gorm.DB, PerPage int, association, pageName string) *Pagination {
	// 默认每页数量
	if PerPage <= 0 {
		PerPage = config.GetInt("pagination.perPage")
	}

	// 实例对象
	p := &Pagination{
		db:          db,
		PerPage:     PerPage,
		Page:        1,
		Count:       -1,
		association: association,
		pageName:    pageName,
	}
	p.baseUrl = p.getBaseUrl(r)

	// 设置当前页码
	p.SetPage(p.GetPageFromRequest(r))

	return p
}

// Paging 返回渲染分页所需的数据
func (p *Pagination) Paging() PagerData {
	return PagerData{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalCount:  p.TotalCount(),
		TotalPage:   p.TotalPage(),
		BaseUrl:     p.baseUrl,
		PageName:    p.pageName,
	}
}

// SetPage 设置当前页
func (p *Pagination) SetPage(page int) {
	if page <= 0 {
		page = 1
	}

	p.Page = page
}

// CurrentPage 返回当前页码
func (p Pagination) CurrentPage() int {
	totalPage := p.TotalPage()
	if totalPage == 0 {
		return 0
	}

	if p.Page > totalPage {
		return totalPage
	}

	return p.Page
}

// Results 返回请求数据，请注意 data 参数必须为 GROM 模型的 Slice 对象
func (p Pagination) Results(data interface{}) error {
	var err error
	var offset int
	page := p.CurrentPage()
	if page == 0 {
		return err
	}

	if page > 1 {
		offset = (page - 1) * p.PerPage
	}

	if p.isNotAssociated() {
		return p.db.Limit(p.PerPage).Offset(offset).Association(p.association).Find(data)
	}
	return p.db.Limit(p.PerPage).Offset(offset).Find(data).Error
}

// TotalCount 返回的是数据库里的条数
func (p *Pagination) TotalCount() int64 {
	if p.Count == -1 {
		var count int64

		if p.isNotAssociated() {
			count = p.db.Association(p.association).Count()
		} else {
			if err := p.db.Count(&count).Error; err != nil {
				return 0
			}
		}
		p.Count = count
	}

	return p.Count
}

// HasPages 总页数大于 1 时会返回 true
func (p *Pagination) HasPages() bool {
	n := p.TotalCount()
	return n > int64(p.PerPage)
}

// HasNext returns true if current page is not the last page
func (p Pagination) HasNext() bool {
	totalPage := p.TotalPage()
	if totalPage == 0 {
		return false
	}

	page := p.CurrentPage()
	if page == 0 {
		return false
	}

	return page < totalPage
}

// PrevPage 前一页码，0 意味着这就是第一页
func (p Pagination) PrevPage() int {
	hasPrev := p.HasPrev()

	if !hasPrev {
		return 0
	}

	page := p.CurrentPage()
	if page == 0 {
		return 0
	}

	return page - 1
}

// NextPage 下一页码，0 的话就是最后一页
func (p Pagination) NextPage() int {
	hasNext := p.HasNext()
	if !hasNext {
		return 0
	}

	page := p.CurrentPage()
	if page == 0 {
		return 0
	}

	return page + 1
}

// HasPrev 如果当前页不为第一页，就返回 true
func (p Pagination) HasPrev() bool {
	page := p.CurrentPage()
	if page == 0 {
		return false
	}

	return page > 1
}

// TotalPage 返回总页数
func (p Pagination) TotalPage() int {
	count := p.TotalCount()
	if count == 0 {
		return 0
	}

	nums := int64(math.Ceil(float64(count) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}

	return int(nums)
}

// GetPageFromRequest 从 URL 中获取 page 参数
func (p Pagination) GetPageFromRequest(r *http.Request) int {
	page := r.URL.Query().Get(p.getPageName())
	if page == "" {
		return 1
	}
	pageInt := types.StringToInt(page)
	if pageInt <= 0 {
		return 1
	}
	return pageInt
}

//isNotAssociated 是否未关联
func (p *Pagination) isNotAssociated() bool {
	return p.association != "" && strings.ToLower(p.association) != p.db.Statement.Table
}

//isNotAssociated 是否未关联
func (p Pagination) getBaseUrl(r *http.Request) string {
	var query []string
	for key, value := range r.URL.Query() {
		if key != p.getPageName() {
			query = append(query, key+"="+value[0])
		}
	}
	if len(query) > 0 {
		return r.URL.Path + "?" + strings.Join(query, "&")
	}
	return r.URL.Path
}

func (p Pagination) getPageName() string {
	if p.pageName != "" {
		return p.pageName
	}
	return config.GetString("pagination.url_query")
}
