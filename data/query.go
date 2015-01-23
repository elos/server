package data

type AttrMap map[string]interface{}

type Query interface {
	Execute() (RecordIterator, error)

	Select(AttrMap) Query
	Limit(int) Query
	Skip(int) Query
	Batch(int) Query
}

type RecordIterator interface {
	Next(Record) bool
	Close() error
}
