package usecase

import (
	"fmt"
	"github.com/modsyan/TS-Gen-CA/internal/utils"
)

func GenerateIndexFile(featureName string, useCaseName string, useCaseFileName string, requestFileName string) error {

	// Create use-case index file
	indexFilePath := fmt.Sprintf("%s/%s/index.ts", featureName, useCaseName)
	indexFileContent := fmt.Sprintf("export * from './%s';\nexport * from './%s';", useCaseFileName, requestFileName)

	if err := utils.CreateFile(indexFilePath, indexFileContent); err != nil {
		return fmt.Errorf("Error creating index file: %v\n", err)
	}

	// Append export statement to the feature index file
	featureIndexFilePath := featureName + "/index.ts"
	featureIndexFileContent := fmt.Sprintf("export * from './%s';\n", useCaseName)

	if err := utils.AppendToFile(featureIndexFilePath, featureIndexFileContent); err != nil {
		return fmt.Errorf("Error appending to feature index file: %v\n", err)
	}

	return nil
}
