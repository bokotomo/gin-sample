package port

import (
	. "gin-sample/domain"
)

type DesignPort interface {
	FindAllDesigns() ([]*Design, error)
}
