package usecase

import (
	"fmt"
	"github.com/modsyan/TS-Gen-CA/internal/utils"
	"strings"
)

func GenerateRequestFile(requestFilePath, requestClassName string, properties []string) error {

	requestClassContent := generateRequestClassContent(requestClassName, properties)

	if err := utils.CreateFile(requestFilePath, requestClassContent); err != nil {
		return fmt.Errorf("Error creating request file: %v\n", err)
	}

	return nil
}

func generateRequestClassContent(className string, properties []string) string {
	propertiesContent := ""
	params := ""
	paramsWithoutType := ""
	propertiesInitialization := ""

	for _, prop := range properties {
		propParts := strings.Split(prop, ":")
		propName := propParts[0]
		propType := propParts[1]

		propertiesContent += fmt.Sprintf("  readonly %s: %s;\n", propName, propType)
		params += fmt.Sprintf("    %s: %s,\n", propName, propType)
		paramsWithoutType += fmt.Sprintf("      %s,\n", propName)
		propertiesInitialization += fmt.Sprintf("    this.%s = %s;\n", propName, propName)
	}

	propertiesContent = utils.RemoveLastNewline(propertiesContent)
	params = utils.RemoveLastNewline(params)
	paramsWithoutType = utils.RemoveLastNewline(paramsWithoutType)
	propertiesInitialization = utils.RemoveLastNewline(propertiesInitialization)

	return fmt.Sprintf(`
import { UseCaseRequest } from '@application/shared';
import { TriggeredBy } from '@domain/shared/entities/triggered-by';
// import { InvalidParameterException } from '@domain/shared/exceptions';

class %s extends UseCaseRequest {

%s

  // Constructor Section
  constructor(
    triggeredBy: TriggeredBy,
%s
  ) {
    super(triggeredBy);
%s
  }

  public static create(
    triggeredBy: TriggeredBy,
%s
  ): %s {
    return new %s(
      triggeredBy,
%s
    );
  } 

  // Validate here using EnsureClass
  protected validatePayload(): void {
  }
}

export { %s };
`, className, propertiesContent, params, propertiesInitialization, params, className, className, paramsWithoutType, className)
}
