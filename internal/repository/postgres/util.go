package postgres

func getPageNumber(pageSize, pageNo uint) (offset, limit int) {
	offset = int((pageNo - 1) * pageSize)
	limit = int(pageSize)

	return offset, limit
}
