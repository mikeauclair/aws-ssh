// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package directoryserviceiface provides an interface for the AWS Directory Service.
package directoryserviceiface

import (
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/service/directoryservice"
)

// DirectoryServiceAPI is the interface type for directoryservice.DirectoryService.
type DirectoryServiceAPI interface {
	ConnectDirectoryRequest(*directoryservice.ConnectDirectoryInput) (*service.Request, *directoryservice.ConnectDirectoryOutput)

	ConnectDirectory(*directoryservice.ConnectDirectoryInput) (*directoryservice.ConnectDirectoryOutput, error)

	CreateAliasRequest(*directoryservice.CreateAliasInput) (*service.Request, *directoryservice.CreateAliasOutput)

	CreateAlias(*directoryservice.CreateAliasInput) (*directoryservice.CreateAliasOutput, error)

	CreateComputerRequest(*directoryservice.CreateComputerInput) (*service.Request, *directoryservice.CreateComputerOutput)

	CreateComputer(*directoryservice.CreateComputerInput) (*directoryservice.CreateComputerOutput, error)

	CreateDirectoryRequest(*directoryservice.CreateDirectoryInput) (*service.Request, *directoryservice.CreateDirectoryOutput)

	CreateDirectory(*directoryservice.CreateDirectoryInput) (*directoryservice.CreateDirectoryOutput, error)

	CreateSnapshotRequest(*directoryservice.CreateSnapshotInput) (*service.Request, *directoryservice.CreateSnapshotOutput)

	CreateSnapshot(*directoryservice.CreateSnapshotInput) (*directoryservice.CreateSnapshotOutput, error)

	DeleteDirectoryRequest(*directoryservice.DeleteDirectoryInput) (*service.Request, *directoryservice.DeleteDirectoryOutput)

	DeleteDirectory(*directoryservice.DeleteDirectoryInput) (*directoryservice.DeleteDirectoryOutput, error)

	DeleteSnapshotRequest(*directoryservice.DeleteSnapshotInput) (*service.Request, *directoryservice.DeleteSnapshotOutput)

	DeleteSnapshot(*directoryservice.DeleteSnapshotInput) (*directoryservice.DeleteSnapshotOutput, error)

	DescribeDirectoriesRequest(*directoryservice.DescribeDirectoriesInput) (*service.Request, *directoryservice.DescribeDirectoriesOutput)

	DescribeDirectories(*directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error)

	DescribeSnapshotsRequest(*directoryservice.DescribeSnapshotsInput) (*service.Request, *directoryservice.DescribeSnapshotsOutput)

	DescribeSnapshots(*directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error)

	DisableRadiusRequest(*directoryservice.DisableRadiusInput) (*service.Request, *directoryservice.DisableRadiusOutput)

	DisableRadius(*directoryservice.DisableRadiusInput) (*directoryservice.DisableRadiusOutput, error)

	DisableSsoRequest(*directoryservice.DisableSsoInput) (*service.Request, *directoryservice.DisableSsoOutput)

	DisableSso(*directoryservice.DisableSsoInput) (*directoryservice.DisableSsoOutput, error)

	EnableRadiusRequest(*directoryservice.EnableRadiusInput) (*service.Request, *directoryservice.EnableRadiusOutput)

	EnableRadius(*directoryservice.EnableRadiusInput) (*directoryservice.EnableRadiusOutput, error)

	EnableSsoRequest(*directoryservice.EnableSsoInput) (*service.Request, *directoryservice.EnableSsoOutput)

	EnableSso(*directoryservice.EnableSsoInput) (*directoryservice.EnableSsoOutput, error)

	GetDirectoryLimitsRequest(*directoryservice.GetDirectoryLimitsInput) (*service.Request, *directoryservice.GetDirectoryLimitsOutput)

	GetDirectoryLimits(*directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error)

	GetSnapshotLimitsRequest(*directoryservice.GetSnapshotLimitsInput) (*service.Request, *directoryservice.GetSnapshotLimitsOutput)

	GetSnapshotLimits(*directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error)

	RestoreFromSnapshotRequest(*directoryservice.RestoreFromSnapshotInput) (*service.Request, *directoryservice.RestoreFromSnapshotOutput)

	RestoreFromSnapshot(*directoryservice.RestoreFromSnapshotInput) (*directoryservice.RestoreFromSnapshotOutput, error)

	UpdateRadiusRequest(*directoryservice.UpdateRadiusInput) (*service.Request, *directoryservice.UpdateRadiusOutput)

	UpdateRadius(*directoryservice.UpdateRadiusInput) (*directoryservice.UpdateRadiusOutput, error)
}
