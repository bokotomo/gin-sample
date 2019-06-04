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

func (this *DesignUseCase) FindDesigns(page int) ([]*Design, error) {
	return this.designPort.FindDesigns(page)
}
