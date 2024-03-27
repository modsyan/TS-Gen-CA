package main

import (
	"flag"
	"fmt"
	"github.com/modsyan/TS-Gen-CA/cmd"
	"github.com/modsyan/TS-Gen-CA/internal/utils"
	"os"
)

func main() {
	//  flags
	featureFlag := flag.String("fn", "", "Name of the feature")
	useCaseFlag := flag.String("uc", "", "Name of the use case")
	propertyFlag := flag.String("p", "", "Properties of the use case")
	returnTypeFlag := flag.String("rt", "", "ReturnType of the use case")
	flag.Parse()

	// Check if feature name is provided
	if *featureFlag == "" {
		fmt.Println("Please provide a feature name using --feature or -fn flag")
		os.Exit(1)
	}

	// Check if use case name is provided
	if *useCaseFlag == "" {
		fmt.Println("Please provide a use case name using --usecase or -un flag")
		os.Exit(1)
	}
	// Check if property list is provided
	if *propertyFlag == "" {
		fmt.Println("Please provide properties using --property or -p flag")
		os.Exit(1)
	}

	if *returnTypeFlag == "" {
		fmt.Println("Please provide return type using --returnType or -rt flag")
		os.Exit(1)
	}

	featureName := *featureFlag
	useCaseName := *useCaseFlag
	properties := utils.ParseProperties(*propertyFlag)
	returnType := *returnTypeFlag

	err := cmd.GenerateFiles(featureName, useCaseName, properties, returnType)
	if err != nil {
		fmt.Printf("Error generating files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("TypeScript files '%s' and '%s' created successfully in the '%s' folder.\n", useCaseName+"Request.request.ts", useCaseName+"UseCase.usecase.ts", featureName)
}
