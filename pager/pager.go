package pager

import (
	"math"
	"net/url"
	"strconv"
)

// Bring Data::Page and Data::Page::Navigation Impl and customize littie bit
// http://search.cpan.org/~lbrocard/Data-Page-2.02/,http://search.cpan.org/~kazeburo/Data-Page-Navigation-0.06/
type Pager struct {
	// TotalEntries is holding total of entries of the list
	TotalEntries uint64
	EntryPerPage uint64
	CurrentPage  uint64
	RequestURI   string
}

var PAGER_DEFAULT_ENTRY_PER_PAGE uint64 = 60

func NewPager(p string) *Pager {
	i, err := strconv.ParseUint(p, 10, 64)
	if err != nil {
		panic(err)
	}
	return &Pager{EntryPerPage: PAGER_DEFAULT_ENTRY_PER_PAGE, CurrentPage: i}
}

func (self *Pager) BuildPath(p uint64) string {
	u, _ := url.Parse(self.RequestURI)
	q := u.Query()
	q.Set("p", strconv.FormatUint(p, 10))
	u.RawQuery = q.Encode()
	return u.String()
}

// number of the first entry on the current page
func (self *Pager) First() uint64 {

	if self.TotalEntries == 0 {
		return 0
	} else {
		return ((self.CurrentPage - 1) * self.EntryPerPage) + 1
	}
}
func (self *Pager) Last() uint64 {
	if self.CurrentPage == self.LastPage() {
		return self.TotalEntries
	} else {
		return self.CurrentPage * self.EntryPerPage
	}
}

func (self *Pager) Offset() uint64 {
	skipped := self.First() - 1
	if skipped < 0 {
		return 0
	} else {
		return skipped
	}
}

func (self *Pager) FirstPage() uint64 {
	return 1
}

func (self *Pager) LastPage() uint64 {

	var pages float64 = float64(self.TotalEntries) / float64(self.EntryPerPage)
	var lastPage uint64 = 0

	if pages == math.Trunc(pages) {
		lastPage = uint64(pages)
	} else {
		lastPage = uint64(1 + math.Trunc(pages))
	}

	if lastPage < 1 {
		lastPage = 1
	}
	return lastPage

}

func (self *Pager) PreviousPage() uint64 {

	if self.CurrentPage > 1 {
		return self.CurrentPage - 1
	} else {
		return 0
	}
}

func (self *Pager) NextPage() uint64 {
	if self.CurrentPage < self.LastPage() {
		return self.CurrentPage + 1
	} else {
		// XXX
		return 0
	}
}

func (self *Pager) PagesInNavigation(pagesPerNavigation uint64) []uint64 {
	//var pagesPerNavigation uint64 = 3

	lastPage := self.LastPage()

	if pagesPerNavigation >= lastPage {
		nav := []uint64{}
		for i := self.FirstPage(); i <= lastPage; i++ {
			nav = append(nav, i)
		}
		return nav
	}

	var prev uint64 = self.CurrentPage - 1
	var next uint64 = self.CurrentPage + 1

	nav := []uint64{self.CurrentPage}

	i := 0
	for uint64(len(nav)) < pagesPerNavigation {
		if i%2 != 0 {
			if self.FirstPage() <= prev {
				nav = append([]uint64{prev}, nav...)
			}
			if prev > 0 {
				prev = prev - 1
			}

		} else {
			if self.LastPage() >= next {
				nav = append(nav, next)
				next = next + 1
			}

		}
		i = i + 1
	}
	return nav
}
