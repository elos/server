package data

type AttrMap map[string]interface{}

type Query interface {
	Execute() (ModelIterator, error)

	Select(AttrMap) Query
	Limit(int) Query
	Skip(int) Query
	Batch(int) Query
}

type ModelIterator interface {
	Next(Model) bool
	Close() error
}
