package links

type LinkDA interface {
	Get(id string) (*Link, error)
	GetByUrl(url string) (*Link, error)
	Save(link Link) (*Link, error)
}

type ErrNotFound struct {
}

func (err ErrNotFound) Error() string {
	return "link not found"
}

type ErrUrlDuplicate struct {
}

func (err ErrUrlDuplicate) Error() string {
	return "duplicate link"
}

type ErrUniqueIdGenerationFailed struct {
}

func (err ErrUniqueIdGenerationFailed) Error() string {
	return "failed to generate unique ID for link"
}

type ErrRepositoryError struct {
}

func (err ErrRepositoryError) Error() string {
	return "internal Server error"
}
