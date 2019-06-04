package port

import (
	. "gin-sample/domain"
)

type DesignPort interface {
	FindDesigns(page int) ([]*Design, error)
}
