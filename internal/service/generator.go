package service

type IdGenerator interface {
	Next() (int64, error)
	BatchNext(size int) ([]int64, error)
}
