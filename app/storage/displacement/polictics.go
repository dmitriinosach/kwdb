package displacement

type Policy interface {
	Cut()
	//Push(string string, cell *storage.Cell)
}
