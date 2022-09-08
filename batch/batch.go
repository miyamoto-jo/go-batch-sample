package batch

type Batch interface {
	Execute([]string) error
}
