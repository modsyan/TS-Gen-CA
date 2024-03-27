package cmd

import (
	"fmt"
	"ts-gen-ca/internal/generator/usecase"
	"ts-gen-ca/internal/utils"
)

func GenerateFiles(featureName, useCaseName string, properties []string, returnType string) error {

	filePath := fmt.Sprintf("%s/%s", featureName, useCaseName)

	if err := utils.FileExists(filePath); err == true {
		return fmt.Errorf("UseCase '%s' already exists. Exiting...\n", useCaseName)
	}

	if err := utils.CreateDirectory(filePath); err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	useCaseFileName := fmt.Sprintf("%s.usecase", useCaseName)
	requestFileName := fmt.Sprintf("%s.request", useCaseName)

	useCaseFilePath := fmt.Sprintf("%s/%s/%s.ts", featureName, useCaseName, useCaseFileName)
	requestFilePath := fmt.Sprintf("%s/%s/%s.ts", featureName, useCaseName, requestFileName)

	useCaseClassName := utils.ConvertToPascalCase(useCaseName) + "UseCase"
	requestClassName := utils.ConvertToPascalCase(useCaseName) + "Request"

	err := usecase.GenerateRequestFile(requestFilePath, requestClassName, properties)
	if err != nil {
		return fmt.Errorf("error generating request file: %v", err)
	}

	err = usecase.GenerateUseCaseFile(useCaseFilePath, useCaseClassName, requestClassName, requestFileName, returnType)
	if err != nil {
		return fmt.Errorf("error generating use case file: %v", err)
	}

	err = usecase.GenerateIndexFile(featureName, useCaseName, useCaseFileName, requestFileName)
	if err != nil {
		return fmt.Errorf("error generating index file: %v", err)
	}

	return nil
}
