package pagination

type Pagination struct {
	Offset int
	Limit  int
}

func New(offset, limit int) Pagination {
	if offset <= 0 {
		offset = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset = (offset - 1) * limit

	return Pagination{
		Offset: offset,
		Limit:  limit,
	}
}
