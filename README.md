# TsCa (Typscript Clean Architechture)

TS-CA is a command-line tool written in Go that generates TypeScript code for [express-typescript-skeleton](https://github.com/borjapazr/express-typescript-skeleton).

## Features

- Generates use-case files with customizable features
- Provides a modular and extensible structure for generating code
- Easy-to-use command-line interface

## Installation

To install TS-CA, you need to have Go installed on your system. Then, you can install it using the following command:

```bash
go install github.com/yourusername/TsCa@v0.0.1
```

## Usage

To generate Usecase TypeScript code for your Clean Architechture Typscript Backend application follows Use-Cases pattern, you can use the following command:

```bash
tsca generate -fn [--feature-name] <feature-name> -un [--usecase-name] <use-case-namee> -p <--property> -rt [--return-type] <returnType>
```

- `<featureName>`: Name of the feature (e.g., "user", "authentication")
- `<useCaseName>`: Name of the use case (e.g., "createUser", "authenticate-user")
- `<properties>`: Properties of the use case separated by spaces (e.g., "name:string email:string password:string")
- `<returnType>`: Return type of the use case (e.g., "User", "AuthenticationResponse")

## Directory Structure

```
.
├── bin
│   └── tsca
├── cmd
│   └── generate.go
├── go.mod
├── internal
│   ├── generator
│   │   └── usecase
│   │       ├── index_generator.go => append to feature-name/index.ts and createa featrue-name/usecase-name/index.ts
│   │       ├── request_generator.go => generae a usecase-name.request.ts where it's the payload of the requesed usedcase
│   │       └── usecase_generator.go => generaea a usecase file in usecase-name.usecase.ts
│   └── utils
│       ├── file_utils.go
│       └── string_utils.go
├── LICENSE
└── main.go
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Express TypeScript Skeleton for @borjapazr](https://github.com/borjapazr/express-typescript-skeleton) for providing the boilerplate template.
- Go community for creating and maintaining amazing tools and libraries.
