package helper

import (
	"fmt"
	"html/template"
	"math"
	"net/url"
	"strconv"
	"strings"
)

type PaginatePage struct {
	Label      string
	Number     int
	IsActive   bool
	IsDisabled bool
}

type Paginator struct {
	Page     int
	Limit    int
	Count    int
	Interval int
	QueryURL url.Values
}

type Pager struct {
	Markup     template.HTML
	Count      int
	StartIndex int
	EndIndex   int
}

func (p *Paginator) Offset() int {
	if p.Page != 1 {
		return (p.Page - 1) * p.Limit
	}
	return 0
}

func (p *Paginator) PageCount() int {
	return (p.Count + p.Limit - 1) / p.Limit
}

func (p *Paginator) StartIndex() int {
	if p.Count != 0 {
		return (p.Limit * (p.Page - 1)) + 1
	}
	return 0
}

func (p *Paginator) EndIndex() int {
	if p.Count != 0 {
		a := p.StartIndex() + p.Limit - 1
		return minInt(a, p.Count)
	}
	return 0
}

func (p *Paginator) String() string {
	markup := []string{}
	markup = append(markup, `<ul class="pagination">`)

	for _, pp := range p.Paginate() {
		p.QueryURL.Set("page", strconv.Itoa(pp.Number))

		if pp.Label == "previous" {
			markup = append(markup, fmt.Sprintf(`<li><a href="?%s"><span>&laquo;</span></a></li>`, p.QueryURL.Encode()))
		} else if pp.Label == "next" {
			markup = append(markup, fmt.Sprintf(`<li><a href="?%s"><span>&raquo;</span></a></li>`, p.QueryURL.Encode()))
		} else if pp.Label == "..." {
			markup = append(markup, fmt.Sprintf(`<li class="disabled"><a href="javascript:void(%d)"><span>%s</span></a></li>`, pp.Number, pp.Label))
		} else {
			ppl := fmt.Sprintf(`<li><a href="?%s">%s</a></li>`, p.QueryURL.Encode(), pp.Label)
			if pp.IsActive {
				ppl = strings.Replace(ppl, `<li>`, `<li class="active">`, 1)
			}
			markup = append(markup, ppl)
		}
	}
	markup = append(markup, `</ul>`)
	return strings.Join(markup, "\n")
}

func (p *Paginator) Paginate() []PaginatePage {
	var pages []PaginatePage
	firstPage := 1
	lastPage := int(math.Ceil(float64(p.Count) / float64(p.Limit)))

	if p.Page < firstPage {
		p.Page = firstPage
	}

	if p.Page > lastPage {
		p.Page = lastPage
	}

	if (p.Page - 1) >= firstPage {
		pages = append(pages, PaginatePage{
			Label:  "previous",
			Number: p.Page - 1,
		})
	}

	pages = append(pages, PaginatePage{
		Label:    strconv.Itoa(firstPage),
		Number:   firstPage,
		IsActive: firstPage == p.Page,
	})

	if (p.Page - p.Interval) > (firstPage + 1) {
		if (p.Page - p.Interval - 1) == (firstPage + 1) {
			pages = append(pages, PaginatePage{
				Label:      strconv.Itoa(firstPage + 1),
				Number:     firstPage + 1,
				IsActive:   firstPage+1 == p.Page,
				IsDisabled: true,
			})
		} else {
			pages = append(pages, PaginatePage{
				Label:      "...",
				Number:     0,
				IsDisabled: true,
			})
		}
	}

	for pg := p.Page - p.Interval; pg <= (p.Page + p.Interval); pg++ {
		if pg > firstPage && pg < lastPage {
			pages = append(pages, PaginatePage{
				Label:    strconv.Itoa(pg),
				Number:   pg,
				IsActive: pg == p.Page,
			})
		}
	}

	if (p.Page + p.Interval) < (lastPage - 1) {
		if (p.Page + p.Interval + 1) == (lastPage - 1) {
			pages = append(pages, PaginatePage{
				Label:      strconv.Itoa(lastPage - 1),
				Number:     lastPage - 1,
				IsDisabled: true,
			})
		} else {
			pages = append(pages, PaginatePage{
				Label:      "...",
				Number:     0,
				IsDisabled: true,
			})
		}
	}

	if lastPage > firstPage {
		pages = append(pages, PaginatePage{
			Label:    strconv.Itoa(lastPage),
			Number:   lastPage,
			IsActive: lastPage == p.Page,
		})
	}

	if (p.Page + 1) <= lastPage {
		pages = append(pages, PaginatePage{
			Label:  "next",
			Number: p.Page + 1,
		})
	}
	return pages
}
