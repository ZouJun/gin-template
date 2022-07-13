package pagination

import "math"

//切片分页
func SlicePage(pageIndex, pageSize, totalCount int) (sliceStart, sliceEnd, totalPage int) {
	if pageIndex <= 0 {
		pageIndex = 1
	}

	if pageSize < 0 {
		pageSize = 20
	}

	if pageSize > totalCount {
		return 0, totalCount, 1
	}

	// 总页数
	totalPage = int(math.Ceil(float64(totalCount) / float64(pageSize)))
	if pageIndex > totalPage {
		return 0, 0, 0
	}
	sliceStart = (pageIndex - 1) * pageSize
	sliceEnd = sliceStart + pageSize

	if sliceEnd > totalCount {
		sliceEnd = totalCount
	}

	return sliceStart, sliceEnd, totalPage
}
