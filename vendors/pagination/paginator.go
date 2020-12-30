package pagination

import (
	"blog/vendors/config"
	"blog/vendors/types"
	"strings"
)

type Paginator struct {
	PagerData  PagerData
	OnEachSide int
}

type Slider struct {
	first  []string
	slider []string
	last   []string
}

type Link struct {
	Url    string
	Label  string
	Active bool
}

func CreatePaginator(data PagerData, onEachSide int) *Paginator {
	return &Paginator{
		PagerData:  data,
		OnEachSide: onEachSide,
	}
}

func (p *Paginator) Links() []Link {
	links := []Link{{Url: p.PrevUrl(), Label: "‹", Active: false}}
	slider := p.Slider()

	links = append(links, p.RangeLinks(slider.first)...)

	if slider.slider != nil && len(slider.slider) > 0 {
		links = append(links, Link{Url: "", Label: "...", Active: false})
		links = append(links, p.RangeLinks(slider.slider)...)
		links = append(links, Link{Url: "", Label: "...", Active: false})
	}
	if len(slider.last) > 0 {
		if len(slider.slider) == 0 {
			links = append(links, Link{Url: "", Label: "...", Active: false})
		}
		links = append(links, p.RangeLinks(slider.last)...)
	}
	links = append(links, Link{Url: p.NextUrl(), Label: "›", Active: false})
	return links
}

func (p *Paginator) RangeLinks(slider []string) []Link {
	var links []Link
	for pa, link := range slider {
		page := pa + 1
		links = append(links, Link{
			Url:    link,
			Label:  types.IntToString(page),
			Active: page == p.PagerData.CurrentPage,
		})
	}
	return links
}

func (p *Paginator) Slider() Slider {
	if p.PagerData.TotalPage < (p.OnEachSide*2)+8 {
		return p.GetSmallSlider()
	}
	return p.GetUrlSlider()
}

func (p Paginator) GetUrlSlider() Slider {
	window := p.OnEachSide + 4
	if !p.hasPages() {
		return Slider{}
	}
	if p.PagerData.CurrentPage <= window {
		return p.GetSliderTooCloseToBeginning(window)
	} else if p.PagerData.CurrentPage > p.PagerData.TotalPage-window {
		return p.GetSliderTooCloseToEnding(window)
	}
	return p.GetFullSlider()
}

func (p Paginator) GetSliderTooCloseToEnding(window int) Slider {
	return Slider{
		first: p.GetUrlRange(1, 2),
		last:  p.GetUrlRange(p.PagerData.TotalPage-(window+(p.OnEachSide-1)), p.PagerData.TotalPage),
	}
}

func (p *Paginator) GetSliderTooCloseToBeginning(window int) Slider {
	return Slider{
		first: p.GetUrlRange(1, window+p.OnEachSide),
	}
}

func (p *Paginator) GetFullSlider() Slider {
	return Slider{
		first:  p.GetUrlRange(1, 2),
		slider: p.GetUrlRange(p.PagerData.CurrentPage-p.OnEachSide, p.PagerData.CurrentPage+p.OnEachSide),
		last:   p.GetUrlRange(p.PagerData.TotalPage-1, p.PagerData.TotalPage),
	}
}

func (p *Paginator) GetSmallSlider() Slider {
	return Slider{
		first: p.GetUrlRange(1, p.PagerData.TotalPage),
	}
}

func (p *Paginator) GetUrlRange(start, end int) (pages []string) {
	for i := start; i <= end; i++ {
		pages = append(pages, p.Url(i))
	}
	return pages
}

// Determine if the underlying paginator being presented has pages to show.
func (p *Paginator) hasPages() bool {
	return p.PagerData.TotalPage > 1
}

func (p *Paginator) PrevUrl() string {
	if p.PagerData.CurrentPage > 1 {
		return p.Url(p.PagerData.CurrentPage - 1)
	}
	return ""
}

func (p *Paginator) NextUrl() string {
	if p.PagerData.CurrentPage < p.PagerData.TotalPage {
		return p.Url(p.PagerData.CurrentPage + 1)
	}
	return ""
}

//Url
func (p *Paginator) Url(currentPage int) string {
	if strings.Contains(p.PagerData.BaseUrl, "?") {
		return p.PagerData.BaseUrl + "&" + p.getPageName() + "=" + types.IntToString(currentPage)
	}
	return p.PagerData.BaseUrl + "?" + p.getPageName() + "=" + types.IntToString(currentPage)
}

func (p *Paginator) getPageName() string {
	if p.PagerData.PageName == "" {
		return config.GetString("pagination.url_query")
	}
	return p.PagerData.PageName
}
