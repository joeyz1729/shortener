package sequence

const sqlReplaceStub = `REPLACE INTO sequence (stub) VALUES('a')`

type Sequence interface {
	Next() (uint64, error)
}
