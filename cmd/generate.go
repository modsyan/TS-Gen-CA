package cmd

import (
	"fmt"
	"github.com/modsyan/TsCa/internal/generator/usecase"
	"github.com/modsyan/TsCa/internal/utils"
)

func GenerateFiles(featureName, useCaseName string, properties []string, returnType string) error {

	filePath := fmt.Sprintf("%s/%s", featureName, useCaseName)

	if err := utils.FileExists(filePath); err == true {
		return fmt.Errorf("UseCase '%s' already exists. Exiting...\n", useCaseName)
	}

	if err := utils.CreateDirectory(filePath); err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	useCaseFileName := fmt.Sprintf("%s-%s.usecase", useCaseName, featureName)
	requestFileName := fmt.Sprintf("%s-%s.request", useCaseName, featureName)

	useCaseFilePath := fmt.Sprintf("%s/%s/%s.ts", featureName, useCaseName, useCaseFileName)
	requestFilePath := fmt.Sprintf("%s/%s/%s.ts", featureName, useCaseName, requestFileName)

	useCaseClassName := utils.ConvertToPascalCase(useCaseName) + utils.ConvertToPascalCase(featureName) + "UseCase"
	requestClassName := utils.ConvertToPascalCase(useCaseName) + utils.ConvertToPascalCase(featureName) + "Request"

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
