package usecase

import (
	"gin-sample/domain"
	"gin-sample/usecase/port"
)

// DesignUseCase is
type DesignUseCase struct {
	designPort port.DesignPort
}

// NewDesignUseCase is
func NewDesignUseCase(designPort port.DesignPort) *DesignUseCase {
	return &DesignUseCase{designPort: designPort}
}

// FindDesigns is
func (d *DesignUseCase) FindDesigns(designs *[10]domain.Design, total *int, page int) error {
	return d.designPort.FindDesigns(designs, total, page)
}

// FindDesign isis
func (d *DesignUseCase) FindDesign(design *domain.Design, designID int) error {
	return d.designPort.FindDesign(design, designID)
}
