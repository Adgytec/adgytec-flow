package pagination

import (
	"testing"
	"time"
)

type testPaginationItem struct{}

func (t testPaginationItem) GetCreatedAt() time.Time {
	return time.Now()
}

func TestNewPaginationResponse(t *testing.T) {
	// sample test items
	items := []testPaginationItem{{}}

	// no next and prev page
	var nilNext *testPaginationItem
	var nilPrev *testPaginationItem
	nilRes := newPaginationResponse(items, nilNext, nilPrev)

	if nilRes.PageInfo.HasNextPage || nilRes.PageInfo.HasPrevPage {
		t.Errorf("Expected next page and prev page to be false but got true instead.")
	}

	// with next and prev page
	next := &testPaginationItem{}
	prev := &testPaginationItem{}
	res := newPaginationResponse(items, next, prev)

	if !res.PageInfo.HasNextPage || !res.PageInfo.HasPrevPage {
		t.Errorf("Expected next page and prev page to be true but got false instead.")
	}

}
