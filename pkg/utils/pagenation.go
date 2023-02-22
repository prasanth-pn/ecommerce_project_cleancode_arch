package utils

import (
	"fmt"
	"math"
)

type Filter struct {
	Page     int
	PageSize int
}

type Metadata struct {
	CurrentPage  int
	PageSize     int
	FirstPage    int
	LastPage     int
	TotalRecords int
}

func (f *Filter) Limit() int {
	return f.PageSize
}

func (f *Filter) Offset() int {
	return int(math.Abs(float64((f.Page - 1) * f.PageSize)))
}
func ComputeMetadata(totalRecords, CurrentPage, pageSize *int) Metadata {
	fmt.Println(*pageSize, *CurrentPage, *totalRecords)
	if *totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  *CurrentPage,
		PageSize:     *pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(*totalRecords) / float64(*pageSize))),
		TotalRecords: *totalRecords,
	}
}
