package usecase

import (
	. "gin-sample/domain"
	. "gin-sample/usecase/port"
)

type DesignUseCase struct {
	designPort DesignPort
}

func NewDesignUseCase(designPort DesignPort) *DesignUseCase {
	return &DesignUseCase{designPort: designPort}
}

func (this *DesignUseCase) FindDesigns(designs *[10]*Design, total *int, page int) error {
	return this.designPort.FindDesigns(designs, total, page)
}
