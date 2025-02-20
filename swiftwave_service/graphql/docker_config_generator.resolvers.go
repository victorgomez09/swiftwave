package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	gitmanager "github.com/swiftwave-org/swiftwave/git_manager"
	"github.com/swiftwave-org/swiftwave/swiftwave_service/core"
	"github.com/swiftwave-org/swiftwave/swiftwave_service/graphql/model"
)

// DockerConfigGenerator is the resolver for the dockerConfigGenerator field.
func (r *queryResolver) DockerConfigGenerator(ctx context.Context, input model.DockerConfigGeneratorInput) (*model.DockerConfigGeneratorOutput, error) {
	if input.SourceType == model.DockerConfigSourceTypeSourceCode {
		if input.SourceCodeCompressedFileName == nil {
			return nil, errors.New("invalid source code provided")
		}
		filename := sanitizeFileName(*input.SourceCodeCompressedFileName)
		filename = filepath.Join(r.Config.LocalConfig.ServiceConfig.TarballDirectoryPath, filename)
		config, err := r.ServiceManager.DockerConfigGenerator.GenerateConfigFromSourceCodeTar(filename)
		if err != nil {
			return nil, errors.New("failed to generate docker config from source code")
		}
		return &model.DockerConfigGeneratorOutput{
			DockerFile:          &config.DockerFile,
			DetectedServiceName: &config.DetectedService,
			DockerBuildArgs:     convertMapToDockerConfigBuildArgs(config.Variables),
		}, nil
	} else if input.SourceType == model.DockerConfigSourceTypeGit {
		gitUsername := ""
		gitPassword := ""
		gitPrivateKey := ""
		if input.GitCredentialID != nil {
			var gitCredential core.GitCredential
			if err := gitCredential.FindById(ctx, r.ServiceManager.DbClient, *input.GitCredentialID); err != nil {
				return nil, errors.New("invalid git credential provided")
			}
			gitUsername = gitCredential.Username
			gitPassword = gitCredential.Password
			gitPrivateKey = gitCredential.SshPrivateKey
		}
		if input.RepositoryURL == nil || input.RepositoryBranch == nil {
			return nil, errors.New("invalid git url, repository owner, repository name or branch provided")
		}
		repoInfo, err := gitmanager.ParseGitRepoInfo(*input.RepositoryURL)
		if err != nil {
			return nil, err
		}
		if input.CodePath == nil {
			input.CodePath = new(string)
			*input.CodePath = ""
		}
		config, err := r.ServiceManager.DockerConfigGenerator.GenerateConfigFromGitRepository(repoInfo.URL(), *input.RepositoryBranch, *input.CodePath, gitUsername, gitPassword, gitPrivateKey)
		if err != nil {
			return nil, errors.New("failed to generate docker config from git repository")
		}
		return &model.DockerConfigGeneratorOutput{
			DockerFile:          &config.DockerFile,
			DetectedServiceName: &config.DetectedService,
			DockerBuildArgs:     convertMapToDockerConfigBuildArgs(config.Variables),
		}, nil
	} else if input.SourceType == model.DockerConfigSourceTypeCustom {
		if input.CustomDockerFile == nil {
			return nil, errors.New("invalid custom docker file provided")
		}
		dockerfile := strings.ReplaceAll(*input.CustomDockerFile, "\r\n", "\n")
		dockerfile = strings.ReplaceAll(dockerfile, "\r", "\n")
		dockerfile = strings.ReplaceAll(dockerfile, "\\n", "\r\n")
		config, err := r.ServiceManager.DockerConfigGenerator.GenerateConfigFromCustomDocker(dockerfile)
		if err != nil {
			return nil, errors.New("failed to generate docker config from custom docker file")
		}
		return &model.DockerConfigGeneratorOutput{
			DockerFile:          &config.DockerFile,
			DetectedServiceName: &config.DetectedService,
			DockerBuildArgs:     convertMapToDockerConfigBuildArgs(config.Variables),
		}, nil
	} else {
		return nil, fmt.Errorf("invalid source type")
	}
}

// AvailableDockerConfigs is the resolver for the availableDockerConfigs field.
func (r *queryResolver) AvailableDockerConfigs(ctx context.Context) ([]string, error) {
	return r.ServiceManager.DockerConfigGenerator.AvailableDockerConfigs(), nil
}

// DockerConfigFromServiceName is the resolver for the dockerConfigFromServiceName field.
func (r *queryResolver) DockerConfigFromServiceName(ctx context.Context, serviceName string) (*model.DockerConfigGeneratorOutput, error) {
	dockerConfig, err := r.ServiceManager.DockerConfigGenerator.DockerConfigFromServiceName(serviceName)
	if err != nil {
		return nil, err
	}
	return &model.DockerConfigGeneratorOutput{
		DockerFile:          &dockerConfig.DockerFile,
		DetectedServiceName: &dockerConfig.DetectedService,
		DockerBuildArgs:     convertMapToDockerConfigBuildArgs(dockerConfig.Variables),
	}, nil
}
