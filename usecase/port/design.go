package port

import (
	. "gin-sample/domain"
)

type DesignPort interface {
	FindDesigns(designs *[10]Design, total *int, page int) error
	FindDesign(design *Design, designId int) error
}
