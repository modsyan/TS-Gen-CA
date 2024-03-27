package main

import (
	"flag"
	"fmt"
	"github.com/modsyan/TsCa/cmd"
	"github.com/modsyan/TsCa/internal/utils"
	"os"
)

func main() {
	//todo: adding subcommand generate and new

	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)

	//  flags
	featureFlag := generateCmd.String("feature-name", "", "Name of the feature")
	featureFlagShort := generateCmd.String("fn", "", "Name of the feature (short)")
	useCaseFlag := generateCmd.String("usecase-name", "", "Name of the use case")
	useCaseFlagShort := generateCmd.String("un", "", "Name of the use case (short)")
	propertiesFlag := generateCmd.String("property", "", "Properties of the use case")
	propertiesFlagShort := generateCmd.String("p", "", "Properties of the use case (short)")
	returnTypeFlag := generateCmd.String("return-type", "", "ReturnType of the use case")
	returnTypeFlagShort := generateCmd.String("rt", "", "ReturnType of the use case (short)")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a subcommand [ generate ]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		err := generateCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Parsing error, please try again.")
			os.Exit(1)
		}
	default:
		fmt.Printf("'%s' is not a valid subcommand\n", os.Args[1])
		os.Exit(1)
	}

	if generateCmd.Parsed() {
		if *featureFlag == "" && *featureFlagShort == "" {
			fmt.Println("Please provide a feature name using --feature-name or -fn flag")
			os.Exit(1)
		}

		if *useCaseFlag == "" && *useCaseFlagShort == "" {
			fmt.Println("Please provide a use case name using --usecase-name or -un flag")
			os.Exit(1)
		}

		if *propertiesFlag == "" && *propertiesFlagShort == "" {
			fmt.Println("Please provide properties using --property or -p flag")
			os.Exit(1)
		}

		if *returnTypeFlag == "" && *returnTypeFlagShort == "" {
			fmt.Println("Please provide return type using --return-type or -rt flag")
			os.Exit(1)
		}

		featureName := *featureFlag
		if *featureFlag == "" {
			featureName = *featureFlagShort
		}

		useCaseName := *useCaseFlag
		if *useCaseFlag == "" {
			useCaseName = *useCaseFlagShort
		}

		properties := utils.ParseProperties(*propertiesFlag)
		if *propertiesFlag == "" {
			properties = utils.ParseProperties(*propertiesFlagShort)
		}

		returnType := *returnTypeFlag
		if *returnTypeFlag == "" {
			returnType = *returnTypeFlagShort
		}

		err := cmd.GenerateFiles(featureName, useCaseName, properties, returnType)
		if err != nil {
			fmt.Printf("Error generating files: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("TypeScript files '%s' and '%s' created successfully in the '%s' folder.\n", useCaseName+"Request.request.ts", useCaseName+"UseCase.usecase.ts", featureName)
	}
}
