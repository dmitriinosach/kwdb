package displacement

type Policy interface {
	Cut()
	Use(string string)
}
