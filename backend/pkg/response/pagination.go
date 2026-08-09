package response

type PaginationMeta struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	TotalRows int64 `json:"total_rows"`
	TotalPage int   `json:"total_page"`
}

func CalculatePaginationMeta(page, limit int, totalRows int64) PaginationMeta {
	totalPage := int(totalRows) / limit
	if int(totalRows)%limit != 0 {
		totalPage++
	}

	return PaginationMeta{
		Page:      page,
		Limit:     limit,
		TotalRows: totalRows,
		TotalPage: totalPage,
	}
}
