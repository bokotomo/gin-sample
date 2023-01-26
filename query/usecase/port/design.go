package port

import (
	"gin-sample/domain"
)

// DesignPort is
type DesignPort interface {
	FindDesigns(designs *[10]domain.Design, total *int, page int) error
	FindDesign(design *domain.Design, designID int) error
}
