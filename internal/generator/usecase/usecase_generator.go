package usecase

import (
	"fmt"
	"os"
	"ts-gen-ca/internal/utils"
)

func GenerateUseCaseFile(useCaseFilePath string, useCaseClassName string, requestClassName string, requestFileName string, returnType string) error {

	useCaseClassContent := fmt.Sprintf(
		`import { BaseUseCase, UseCase } from '@application/shared';
import { %s } from './%s';
	
@UseCase()
class %s extends BaseUseCase<%s,%s> {

  constructor() {
	super();
  }

  public async performOperation({  }: %s): Promise<%s> {
	throw new Error('Method not implemented.');
  }
}

export { %s };`,
		requestClassName,
		requestFileName,
		useCaseClassName,
		requestClassName,
		returnType,
		requestClassName,
		returnType,
		useCaseClassName,
	)

	if err := utils.CreateFile(useCaseFilePath, useCaseClassContent); err != nil {
		fmt.Printf("Error creating usecase file: %v\n", err)
		os.Exit(1)
	}

	return nil
}
