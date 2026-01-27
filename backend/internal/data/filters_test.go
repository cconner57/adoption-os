package data

import (
	"testing"

	"github.com/cconner57/adoption-os/backend/internal/validator"
)

func TestCalculateMetadata(t *testing.T) {
	tests := []struct {
		name         string
		totalRecords int
		page         int
		pageSize     int
		wantLastPage int
	}{
		{"zero records", 0, 1, 10, 0},
		{"single page", 5, 1, 10, 1},
		{"exact page", 10, 1, 10, 1},
		{"multiple pages", 11, 1, 10, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			meta := calculateMetadata(tt.totalRecords, tt.page, tt.pageSize)
			if meta.LastPage != tt.wantLastPage {
				t.Errorf("want last page %d; got %d", tt.wantLastPage, meta.LastPage)
			}
			if meta.TotalRecords != tt.totalRecords {
				t.Errorf("want total records %d; got %d", tt.totalRecords, meta.TotalRecords)
			}
		})
	}
}

func TestValidateFilters(t *testing.T) {
	tests := []struct {
		name      string
		filters   Filters
		wantValid bool
	}{
		{
			name: "valid filters",
			filters: Filters{
				Page:         1,
				PageSize:     10,
				Sort:         "id",
				SortSafelist: []string{"id", "name"},
			},
			wantValid: true,
		},
		{
			name: "invalid page",
			filters: Filters{
				Page:         0,
				PageSize:     10,
				Sort:         "id",
				SortSafelist: []string{"id", "name"},
			},
			wantValid: false,
		},
		{
			name: "invalid page size",
			filters: Filters{
				Page:         1,
				PageSize:     101,
				Sort:         "id",
				SortSafelist: []string{"id", "name"},
			},
			wantValid: false,
		},
		{
			name: "invalid sort",
			filters: Filters{
				Page:         1,
				PageSize:     10,
				Sort:         "invalid",
				SortSafelist: []string{"id", "name"},
			},
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.New()
			ValidateFilters(v, tt.filters)
			if v.Valid() != tt.wantValid {
				t.Errorf("want valid %v; got %v", tt.wantValid, v.Valid())
			}
		})
	}
}
