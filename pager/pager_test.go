package pager

import (
	"reflect"
	"testing"
)

func TestNewPager(t *testing.T) {

	if NewPager(int(2)).CurrentPage != 2 {
		t.Fail()
	}
	if NewPager(int64(2)).CurrentPage != 2 {
		t.Fail()
	}
	if NewPager("2").CurrentPage != 2 {
		t.Fail()
	}
	if NewPager(uint64(2)).CurrentPage != 2 {
		t.Fail()
	}

	pager := NewPager(1)

	if reflect.TypeOf(pager).String() != "*pager.Pager" {
		t.Fatalf("type must be *pager.Pager", reflect.TypeOf(pager).Name())
	}

	if pager.EntryPerPage != 60 {
		t.Fatal("default should be 60")
	}

	// change default
	PAGER_DEFAULT_ENTRY_PER_PAGE = 35

	if v := pager.EntryPerPage; v != 60 {
		t.Fatal("default should 60 on this instance", v)
	}

	{
		pager := NewPager("1")
		if v := pager.EntryPerPage; v != 35 {
			t.Fatal("default should change to 35", v)
		}

	}

}

func TestPagerLogic(t *testing.T) {
	PAGER_DEFAULT_ENTRY_PER_PAGE = 10

	{
		pager := NewPager("1")
		pager.TotalEntries = 100
		pager.RequestURI = "/foo/?hoge=hoge"

		if v := pager.Offset(); v != 0 {
			t.Fatal("Error Offset()", v)
		}

		if v := pager.First(); v != 1 {
			t.Fatal("Error First()", v)
		}

		if v := pager.Last(); v != 10 {
			t.Fatal("Error Last()", v)
		}

		if v := pager.PreviousPage(); v != 0 {
			t.Fatal("Error PreviousPage()", v)
		}

		if v := pager.NextPage(); v != 2 {
			t.Fatal("Error NextPage()", v)
		}

		//--

		if v := pager.LastPage(); v != 10 {
			t.Fatal("Error Last()", v)
		}

		if v := pager.BuildPath(1); v != "/foo/?hoge=hoge&p=1" {
			t.Fatal("Error BuildPath()", v)
		}

		if v := pager.PagesInNavigation(3); !reflect.DeepEqual(v, []uint64{1, 2, 3}) {
			t.Fatal("PagesInNavigation()", v)
		}
	}

	{
		pager := NewPager("3")
		pager.TotalEntries = 100
		pager.RequestURI = "/foo/?hoge=hoge"

		if v := pager.Offset(); v != 20 {
			t.Fatal("Error 2 Offset()", v)
		}

		if v := pager.First(); v != 21 {
			t.Fatal("Error 2 First()", v)
		}

		if v := pager.Last(); v != 30 {
			t.Fatal("Error 2 Last()", v)
		}

		if v := pager.PreviousPage(); v != 2 {
			t.Fatal("Error 2 PreviousPage()", v)
		}

		if v := pager.NextPage(); v != 4 {
			t.Fatal("Error 2 NextPage()", v)
		}
		if v := pager.PagesInNavigation(3); !reflect.DeepEqual(v, []uint64{2, 3, 4}) {
			t.Fatal("PagesInNavigation()", v)
		}

		if v := pager.PagesInNavigation(5); !reflect.DeepEqual(v, []uint64{1, 2, 3, 4, 5}) {
			t.Fatal("PagesInNavigation()", v)
		}

		if v := pager.PagesInNavigation(20); !reflect.DeepEqual(v, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
			t.Fatal("PagesInNavigation()", v)
		}

	}

}
