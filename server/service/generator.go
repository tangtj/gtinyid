package service

type IdGenService interface {
	Next() (int64, error)
	BatchNext(size int) ([]int64, error)
}
