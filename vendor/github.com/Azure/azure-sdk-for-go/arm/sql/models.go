package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/uuid"
)

// CreateMode enumerates the values for create mode.
type CreateMode string

const (
	// Copy specifies the copy state for create mode.
	Copy CreateMode = "Copy"
	// Default specifies the default state for create mode.
	Default CreateMode = "Default"
	// NonReadableSecondary specifies the non readable secondary state for
	// create mode.
	NonReadableSecondary CreateMode = "NonReadableSecondary"
	// OnlineSecondary specifies the online secondary state for create mode.
	OnlineSecondary CreateMode = "OnlineSecondary"
	// PointInTimeRestore specifies the point in time restore state for create
	// mode.
	PointInTimeRestore CreateMode = "PointInTimeRestore"
	// Recovery specifies the recovery state for create mode.
	Recovery CreateMode = "Recovery"
	// Restore specifies the restore state for create mode.
	Restore CreateMode = "Restore"
)

// DatabaseEditions enumerates the values for database editions.
type DatabaseEditions string

const (
	// Basic specifies the basic state for database editions.
	Basic DatabaseEditions = "Basic"
	// Business specifies the business state for database editions.
	Business DatabaseEditions = "Business"
	// DataWarehouse specifies the data warehouse state for database editions.
	DataWarehouse DatabaseEditions = "DataWarehouse"
	// Free specifies the free state for database editions.
	Free DatabaseEditions = "Free"
	// Premium specifies the premium state for database editions.
	Premium DatabaseEditions = "Premium"
	// Standard specifies the standard state for database editions.
	Standard DatabaseEditions = "Standard"
	// Stretch specifies the stretch state for database editions.
	Stretch DatabaseEditions = "Stretch"
	// Web specifies the web state for database editions.
	Web DatabaseEditions = "Web"
)

// ElasticPoolEditions enumerates the values for elastic pool editions.
type ElasticPoolEditions string

const (
	// ElasticPoolEditionsBasic specifies the elastic pool editions basic
	// state for elastic pool editions.
	ElasticPoolEditionsBasic ElasticPoolEditions = "Basic"
	// ElasticPoolEditionsPremium specifies the elastic pool editions premium
	// state for elastic pool editions.
	ElasticPoolEditionsPremium ElasticPoolEditions = "Premium"
	// ElasticPoolEditionsStandard specifies the elastic pool editions
	// standard state for elastic pool editions.
	ElasticPoolEditionsStandard ElasticPoolEditions = "Standard"
)

// ElasticPoolState enumerates the values for elastic pool state.
type ElasticPoolState string

const (
	// Creating specifies the creating state for elastic pool state.
	Creating ElasticPoolState = "Creating"
	// Disabled specifies the disabled state for elastic pool state.
	Disabled ElasticPoolState = "Disabled"
	// Ready specifies the ready state for elastic pool state.
	Ready ElasticPoolState = "Ready"
)

// RecommendedIndexActions enumerates the values for recommended index actions.
type RecommendedIndexActions string

const (
	// Create specifies the create state for recommended index actions.
	Create RecommendedIndexActions = "Create"
	// Drop specifies the drop state for recommended index actions.
	Drop RecommendedIndexActions = "Drop"
	// Rebuild specifies the rebuild state for recommended index actions.
	Rebuild RecommendedIndexActions = "Rebuild"
)

// RecommendedIndexStates enumerates the values for recommended index states.
type RecommendedIndexStates string

const (
	// Active specifies the active state for recommended index states.
	Active RecommendedIndexStates = "Active"
	// Blocked specifies the blocked state for recommended index states.
	Blocked RecommendedIndexStates = "Blocked"
	// Executing specifies the executing state for recommended index states.
	Executing RecommendedIndexStates = "Executing"
	// Expired specifies the expired state for recommended index states.
	Expired RecommendedIndexStates = "Expired"
	// Ignored specifies the ignored state for recommended index states.
	Ignored RecommendedIndexStates = "Ignored"
	// Pending specifies the pending state for recommended index states.
	Pending RecommendedIndexStates = "Pending"
	// PendingRevert specifies the pending revert state for recommended index
	// states.
	PendingRevert RecommendedIndexStates = "Pending Revert"
	// Reverted specifies the reverted state for recommended index states.
	Reverted RecommendedIndexStates = "Reverted"
	// Reverting specifies the reverting state for recommended index states.
	Reverting RecommendedIndexStates = "Reverting"
	// Success specifies the success state for recommended index states.
	Success RecommendedIndexStates = "Success"
	// Verifying specifies the verifying state for recommended index states.
	Verifying RecommendedIndexStates = "Verifying"
)

// RecommendedIndexTypes enumerates the values for recommended index types.
type RecommendedIndexTypes string

const (
	// CLUSTERED specifies the clustered state for recommended index types.
	CLUSTERED RecommendedIndexTypes = "CLUSTERED"
	// CLUSTEREDCOLUMNSTORE specifies the clusteredcolumnstore state for
	// recommended index types.
	CLUSTEREDCOLUMNSTORE RecommendedIndexTypes = "CLUSTERED COLUMNSTORE"
	// COLUMNSTORE specifies the columnstore state for recommended index types.
	COLUMNSTORE RecommendedIndexTypes = "COLUMNSTORE"
	// NONCLUSTERED specifies the nonclustered state for recommended index
	// types.
	NONCLUSTERED RecommendedIndexTypes = "NONCLUSTERED"
)

// RestorePointTypes enumerates the values for restore point types.
type RestorePointTypes string

const (
	// CONTINUOUS specifies the continuous state for restore point types.
	CONTINUOUS RestorePointTypes = "CONTINUOUS"
	// DISCRETE specifies the discrete state for restore point types.
	DISCRETE RestorePointTypes = "DISCRETE"
)

// ServerVersion enumerates the values for server version.
type ServerVersion string

const (
	// OneTwoFullStopZero specifies the one two full stop zero state for
	// server version.
	OneTwoFullStopZero ServerVersion = "12.0"
	// TwoFullStopZero specifies the two full stop zero state for server
	// version.
	TwoFullStopZero ServerVersion = "2.0"
)

// ServiceObjectiveName enumerates the values for service objective name.
type ServiceObjectiveName string

const (
	// ServiceObjectiveNameBasic specifies the service objective name basic
	// state for service objective name.
	ServiceObjectiveNameBasic ServiceObjectiveName = "Basic"
	// ServiceObjectiveNameP1 specifies the service objective name p1 state
	// for service objective name.
	ServiceObjectiveNameP1 ServiceObjectiveName = "P1"
	// ServiceObjectiveNameP2 specifies the service objective name p2 state
	// for service objective name.
	ServiceObjectiveNameP2 ServiceObjectiveName = "P2"
	// ServiceObjectiveNameP3 specifies the service objective name p3 state
	// for service objective name.
	ServiceObjectiveNameP3 ServiceObjectiveName = "P3"
	// ServiceObjectiveNameS0 specifies the service objective name s0 state
	// for service objective name.
	ServiceObjectiveNameS0 ServiceObjectiveName = "S0"
	// ServiceObjectiveNameS1 specifies the service objective name s1 state
	// for service objective name.
	ServiceObjectiveNameS1 ServiceObjectiveName = "S1"
	// ServiceObjectiveNameS2 specifies the service objective name s2 state
	// for service objective name.
	ServiceObjectiveNameS2 ServiceObjectiveName = "S2"
	// ServiceObjectiveNameS3 specifies the service objective name s3 state
	// for service objective name.
	ServiceObjectiveNameS3 ServiceObjectiveName = "S3"
)

// TableType enumerates the values for table type.
type TableType string

const (
	// BaseTable specifies the base table state for table type.
	BaseTable TableType = "BaseTable"
	// View specifies the view state for table type.
	View TableType = "View"
)

// TargetDatabaseEditions enumerates the values for target database editions.
type TargetDatabaseEditions string

const (
	// TargetDatabaseEditionsBasic specifies the target database editions
	// basic state for target database editions.
	TargetDatabaseEditionsBasic TargetDatabaseEditions = "Basic"
	// TargetDatabaseEditionsDataWarehouse specifies the target database
	// editions data warehouse state for target database editions.
	TargetDatabaseEditionsDataWarehouse TargetDatabaseEditions = "DataWarehouse"
	// TargetDatabaseEditionsFree specifies the target database editions free
	// state for target database editions.
	TargetDatabaseEditionsFree TargetDatabaseEditions = "Free"
	// TargetDatabaseEditionsPremium specifies the target database editions
	// premium state for target database editions.
	TargetDatabaseEditionsPremium TargetDatabaseEditions = "Premium"
	// TargetDatabaseEditionsStandard specifies the target database editions
	// standard state for target database editions.
	TargetDatabaseEditionsStandard TargetDatabaseEditions = "Standard"
	// TargetDatabaseEditionsStretch specifies the target database editions
	// stretch state for target database editions.
	TargetDatabaseEditionsStretch TargetDatabaseEditions = "Stretch"
)

// TargetElasticPoolEditions enumerates the values for target elastic pool
// editions.
type TargetElasticPoolEditions string

const (
	// TargetElasticPoolEditionsBasic specifies the target elastic pool
	// editions basic state for target elastic pool editions.
	TargetElasticPoolEditionsBasic TargetElasticPoolEditions = "Basic"
	// TargetElasticPoolEditionsPremium specifies the target elastic pool
	// editions premium state for target elastic pool editions.
	TargetElasticPoolEditionsPremium TargetElasticPoolEditions = "Premium"
	// TargetElasticPoolEditionsStandard specifies the target elastic pool
	// editions standard state for target elastic pool editions.
	TargetElasticPoolEditionsStandard TargetElasticPoolEditions = "Standard"
)

// TransparentDataEncryptionActivityStates enumerates the values for
// transparent data encryption activity states.
type TransparentDataEncryptionActivityStates string

const (
	// Decrypting specifies the decrypting state for transparent data
	// encryption activity states.
	Decrypting TransparentDataEncryptionActivityStates = "Decrypting"
	// Encrypting specifies the encrypting state for transparent data
	// encryption activity states.
	Encrypting TransparentDataEncryptionActivityStates = "Encrypting"
)

// TransparentDataEncryptionStates enumerates the values for transparent data
// encryption states.
type TransparentDataEncryptionStates string

const (
	// TransparentDataEncryptionStatesDisabled specifies the transparent data
	// encryption states disabled state for transparent data encryption
	// states.
	TransparentDataEncryptionStatesDisabled TransparentDataEncryptionStates = "Disabled"
	// TransparentDataEncryptionStatesEnabled specifies the transparent data
	// encryption states enabled state for transparent data encryption states.
	TransparentDataEncryptionStatesEnabled TransparentDataEncryptionStates = "Enabled"
)

// Column is represents an Azure SQL Database table column.
type Column struct {
	Name              *string             `json:"name,omitempty"`
	ID                *string             `json:"id,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	*ColumnProperties `json:"properties,omitempty"`
}

// ColumnProperties is represents the properties of an Azure SQL Database
// table column.
type ColumnProperties struct {
	ColumnType *string `json:"columnType,omitempty"`
}

// Database is represents an Azure SQL Database.
type Database struct {
	autorest.Response   `json:"-"`
	Name                *string             `json:"name,omitempty"`
	ID                  *string             `json:"id,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Location            *string             `json:"location,omitempty"`
	Tags                *map[string]*string `json:"tags,omitempty"`
	*DatabaseProperties `json:"properties,omitempty"`
}

// DatabaseListResult is represents the response to a List Azure SQL Database
// request.
type DatabaseListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Database `json:"value,omitempty"`
}

// DatabaseMetric is represents Azure SQL Database metrics.
type DatabaseMetric struct {
	ResourceName  *string    `json:"resourceName,omitempty"`
	DisplayName   *string    `json:"displayName,omitempty"`
	CurrentValue  *float64   `json:"currentValue,omitempty"`
	Limit         *float64   `json:"limit,omitempty"`
	Unit          *string    `json:"unit,omitempty"`
	NextResetTime *date.Time `json:"nextResetTime,omitempty"`
}

// DatabaseMetricListResult is represents the response to a List Azure SQL
// Database metrics request.
type DatabaseMetricListResult struct {
	autorest.Response `json:"-"`
	Value             *[]DatabaseMetric `json:"value,omitempty"`
}

// DatabaseProperties is represents the properties of an Azure SQL Database.
type DatabaseProperties struct {
	Collation                     *string                      `json:"collation,omitempty"`
	CreationDate                  *date.Time                   `json:"creationDate,omitempty"`
	ContainmentState              *int64                       `json:"containmentState,omitempty"`
	CurrentServiceObjectiveID     *uuid.UUID                   `json:"currentServiceObjectiveId,omitempty"`
	DatabaseID                    *string                      `json:"databaseId,omitempty"`
	EarliestRestoreDate           *date.Time                   `json:"earliestRestoreDate,omitempty"`
	CreateMode                    CreateMode                   `json:"createMode,omitempty"`
	SourceDatabaseID              *string                      `json:"sourceDatabaseId,omitempty"`
	Edition                       DatabaseEditions             `json:"edition,omitempty"`
	MaxSizeBytes                  *string                      `json:"maxSizeBytes,omitempty"`
	RequestedServiceObjectiveID   *uuid.UUID                   `json:"requestedServiceObjectiveId,omitempty"`
	RequestedServiceObjectiveName ServiceObjectiveName         `json:"requestedServiceObjectiveName,omitempty"`
	ServiceLevelObjective         ServiceObjectiveName         `json:"serviceLevelObjective,omitempty"`
	Status                        *string                      `json:"status,omitempty"`
	ElasticPoolName               *string                      `json:"elasticPoolName,omitempty"`
	DefaultSecondaryLocation      *string                      `json:"defaultSecondaryLocation,omitempty"`
	ServiceTierAdvisors           *[]ServiceTierAdvisor        `json:"serviceTierAdvisors,omitempty"`
	UpgradeHint                   *UpgradeHint                 `json:"upgradeHint,omitempty"`
	Schemas                       *[]Schema                    `json:"schemas,omitempty"`
	TransparentDataEncryption     *[]TransparentDataEncryption `json:"transparentDataEncryption,omitempty"`
	RecommendedIndex              *[]RecommendedIndex          `json:"recommendedIndex,omitempty"`
}

// ElasticPool is represents an Azure SQL Database elastic pool.
type ElasticPool struct {
	autorest.Response      `json:"-"`
	Name                   *string             `json:"name,omitempty"`
	ID                     *string             `json:"id,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Location               *string             `json:"location,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
	*ElasticPoolProperties `json:"properties,omitempty"`
}

// ElasticPoolActivity is represents the activity on an Azure SQL Elastic Pool.
type ElasticPoolActivity struct {
	Name                           *string             `json:"name,omitempty"`
	ID                             *string             `json:"id,omitempty"`
	Type                           *string             `json:"type,omitempty"`
	Location                       *string             `json:"location,omitempty"`
	Tags                           *map[string]*string `json:"tags,omitempty"`
	*ElasticPoolActivityProperties `json:"properties,omitempty"`
}

// ElasticPoolActivityListResult is represents the response to a List Azure
// SQL Elastic Pool Activity request.
type ElasticPoolActivityListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ElasticPoolActivity `json:"value,omitempty"`
}

// ElasticPoolActivityProperties is represents the properties of an Azure SQL
// Elastic Pool.
type ElasticPoolActivityProperties struct {
	EndTime                   *date.Time `json:"endTime,omitempty"`
	ErrorCode                 *int32     `json:"errorCode,omitempty"`
	ErrorMessage              *string    `json:"errorMessage,omitempty"`
	ErrorSeverity             *int32     `json:"errorSeverity,omitempty"`
	Operation                 *string    `json:"operation,omitempty"`
	OperationID               *string    `json:"operationId,omitempty"`
	PercentComplete           *int32     `json:"percentComplete,omitempty"`
	RequestedDatabaseDtuMax   *int32     `json:"requestedDatabaseDtuMax,omitempty"`
	RequestedDatabaseDtuMin   *int32     `json:"requestedDatabaseDtuMin,omitempty"`
	RequestedDtu              *int32     `json:"requestedDtu,omitempty"`
	RequestedElasticPoolName  *string    `json:"requestedElasticPoolName,omitempty"`
	RequestedStorageLimitInGB *int64     `json:"requestedStorageLimitInGB,omitempty"`
	ElasticPoolName           *string    `json:"elasticPoolName,omitempty"`
	ServerName                *string    `json:"serverName,omitempty"`
	StartTime                 *date.Time `json:"startTime,omitempty"`
	State                     *string    `json:"state,omitempty"`
}

// ElasticPoolDatabaseActivity is represents the activity on an Azure SQL
// Elastic Pool.
type ElasticPoolDatabaseActivity struct {
	Name                                   *string             `json:"name,omitempty"`
	ID                                     *string             `json:"id,omitempty"`
	Type                                   *string             `json:"type,omitempty"`
	Location                               *string             `json:"location,omitempty"`
	Tags                                   *map[string]*string `json:"tags,omitempty"`
	*ElasticPoolDatabaseActivityProperties `json:"properties,omitempty"`
}

// ElasticPoolDatabaseActivityListResult is represents the response to a List
// Azure SQL Elastic Pool Database Activity request.
type ElasticPoolDatabaseActivityListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ElasticPoolDatabaseActivity `json:"value,omitempty"`
}

// ElasticPoolDatabaseActivityProperties is represents the properties of an
// Azure SQL Elastic Pool Database Activity.
type ElasticPoolDatabaseActivityProperties struct {
	DatabaseName              *string    `json:"databaseName,omitempty"`
	EndTime                   *date.Time `json:"endTime,omitempty"`
	ErrorCode                 *int32     `json:"errorCode,omitempty"`
	ErrorMessage              *string    `json:"errorMessage,omitempty"`
	ErrorSeverity             *int32     `json:"errorSeverity,omitempty"`
	Operation                 *string    `json:"operation,omitempty"`
	OperationID               *string    `json:"operationId,omitempty"`
	PercentComplete           *int32     `json:"percentComplete,omitempty"`
	RequestedElasticPoolName  *string    `json:"requestedElasticPoolName,omitempty"`
	CurrentElasticPoolName    *string    `json:"currentElasticPoolName,omitempty"`
	CurrentServiceObjective   *string    `json:"currentServiceObjective,omitempty"`
	RequestedServiceObjective *string    `json:"requestedServiceObjective,omitempty"`
	ServerName                *string    `json:"serverName,omitempty"`
	StartTime                 *date.Time `json:"startTime,omitempty"`
	State                     *string    `json:"state,omitempty"`
}

// ElasticPoolListResult is represents the response to a List Azure SQL
// Elastic Pool request.
type ElasticPoolListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ElasticPool `json:"value,omitempty"`
}

// ElasticPoolProperties is represents the properties of an Azure SQL Elastic
// Pool.
type ElasticPoolProperties struct {
	CreationDate   *date.Time          `json:"creationDate,omitempty"`
	State          ElasticPoolState    `json:"state,omitempty"`
	Edition        ElasticPoolEditions `json:"edition,omitempty"`
	Dtu            *int32              `json:"dtu,omitempty"`
	DatabaseDtuMax *int32              `json:"databaseDtuMax,omitempty"`
	DatabaseDtuMin *int32              `json:"databaseDtuMin,omitempty"`
	StorageMB      *int32              `json:"storageMB,omitempty"`
}

// OperationImpact is represents impact of an operation, both in absolute and
// relative terms.
type OperationImpact struct {
	Name                *string  `json:"name,omitempty"`
	Unit                *string  `json:"unit,omitempty"`
	ChangeValueAbsolute *float64 `json:"changeValueAbsolute,omitempty"`
	ChangeValueRelative *float64 `json:"changeValueRelative,omitempty"`
}

// RecommendedDatabaseProperties is represents the properties of a recommended
// Azure SQL Database being upgraded.
type RecommendedDatabaseProperties struct {
	Name                        *string                `json:"Name,omitempty"`
	TargetEdition               TargetDatabaseEditions `json:"TargetEdition,omitempty"`
	TargetServiceLevelObjective *string                `json:"TargetServiceLevelObjective,omitempty"`
}

// RecommendedElasticPool is represents an Azure SQL Recommended Elastic Pool.
type RecommendedElasticPool struct {
	autorest.Response                 `json:"-"`
	Name                              *string             `json:"name,omitempty"`
	ID                                *string             `json:"id,omitempty"`
	Type                              *string             `json:"type,omitempty"`
	Location                          *string             `json:"location,omitempty"`
	Tags                              *map[string]*string `json:"tags,omitempty"`
	*RecommendedElasticPoolProperties `json:"properties,omitempty"`
}

// RecommendedElasticPoolListMetricsResult is represents the response to a
// List Azure SQL Recommended Elastic Pool metrics request.
type RecommendedElasticPoolListMetricsResult struct {
	autorest.Response `json:"-"`
	Value             *[]RecommendedElasticPoolMetric `json:"value,omitempty"`
}

// RecommendedElasticPoolListResult is represents the response to a List Azure
// SQL Recommended Elastic Pool request.
type RecommendedElasticPoolListResult struct {
	autorest.Response `json:"-"`
	Value             *[]RecommendedElasticPool `json:"value,omitempty"`
}

// RecommendedElasticPoolMetric is represents Azure SQL recommended elastic
// pool metric.
type RecommendedElasticPoolMetric struct {
	DateTime *date.Time `json:"dateTime,omitempty"`
	Dtu      *float64   `json:"dtu,omitempty"`
	SizeGB   *float64   `json:"sizeGB,omitempty"`
}

// RecommendedElasticPoolProperties is represents the properties of an Azure
// SQL Recommended Elastic Pool.
type RecommendedElasticPoolProperties struct {
	DatabaseEdition        ElasticPoolEditions             `json:"databaseEdition,omitempty"`
	Dtu                    *float64                        `json:"dtu,omitempty"`
	DatabaseDtuMin         *float64                        `json:"databaseDtuMin,omitempty"`
	DatabaseDtuMax         *float64                        `json:"databaseDtuMax,omitempty"`
	StorageMB              *float64                        `json:"storageMB,omitempty"`
	ObservationPeriodStart *date.Time                      `json:"observationPeriodStart,omitempty"`
	ObservationPeriodEnd   *date.Time                      `json:"observationPeriodEnd,omitempty"`
	MaxObservedDtu         *float64                        `json:"maxObservedDtu,omitempty"`
	MaxObservedStorageMB   *float64                        `json:"maxObservedStorageMB,omitempty"`
	Databases              *[]Database                     `json:"databases,omitempty"`
	Metrics                *[]RecommendedElasticPoolMetric `json:"metrics,omitempty"`
}

// RecommendedIndex is represents an Azure SQL Database recommended index.
type RecommendedIndex struct {
	Name                        *string             `json:"name,omitempty"`
	ID                          *string             `json:"id,omitempty"`
	Type                        *string             `json:"type,omitempty"`
	Location                    *string             `json:"location,omitempty"`
	Tags                        *map[string]*string `json:"tags,omitempty"`
	*RecommendedIndexProperties `json:"properties,omitempty"`
}

// RecommendedIndexProperties is represents the properties of an Azure SQL
// Database recommended index.
type RecommendedIndexProperties struct {
	Action          RecommendedIndexActions `json:"action,omitempty"`
	State           RecommendedIndexStates  `json:"state,omitempty"`
	Created         *date.Time              `json:"created,omitempty"`
	LastModified    *date.Time              `json:"lastModified,omitempty"`
	IndexType       RecommendedIndexTypes   `json:"indexType,omitempty"`
	Schema          *string                 `json:"schema,omitempty"`
	Table           *string                 `json:"table,omitempty"`
	Columns         *[]string               `json:"columns,omitempty"`
	IncludedColumns *[]string               `json:"includedColumns,omitempty"`
	IndexScript     *string                 `json:"indexScript,omitempty"`
	EstimatedImpact *[]OperationImpact      `json:"estimatedImpact,omitempty"`
	ReportedImpact  *[]OperationImpact      `json:"reportedImpact,omitempty"`
}

// Resource is resource properties
type Resource struct {
	Name     *string             `json:"name,omitempty"`
	ID       *string             `json:"id,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// RestorePoint is represents an Azure SQL Database restore point.
type RestorePoint struct {
	Name                    *string             `json:"name,omitempty"`
	ID                      *string             `json:"id,omitempty"`
	Type                    *string             `json:"type,omitempty"`
	Location                *string             `json:"location,omitempty"`
	Tags                    *map[string]*string `json:"tags,omitempty"`
	*RestorePointProperties `json:"properties,omitempty"`
}

// RestorePointListResult is represents the response to a List Azure SQL
// Database restore points request.
type RestorePointListResult struct {
	autorest.Response `json:"-"`
	Value             *[]RestorePoint `json:"value,omitempty"`
}

// RestorePointProperties is represents the properties of an Azure SQL
// Database restore point.
type RestorePointProperties struct {
	RestorePointType         RestorePointTypes `json:"restorePointType,omitempty"`
	RestorePointCreationDate *date.Time        `json:"restorePointCreationDate,omitempty"`
	EarliestRestoreDate      *date.Time        `json:"earliestRestoreDate,omitempty"`
}

// Schema is represents an Azure SQL Database schema.
type Schema struct {
	Name              *string             `json:"name,omitempty"`
	ID                *string             `json:"id,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	*SchemaProperties `json:"properties,omitempty"`
}

// SchemaProperties is represents the properties of an Azure SQL Database
// schema.
type SchemaProperties struct {
	Tables *[]Table `json:"tables,omitempty"`
}

// Server is represents an Azure SQL server.
type Server struct {
	autorest.Response `json:"-"`
	Name              *string             `json:"name,omitempty"`
	ID                *string             `json:"id,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	*ServerProperties `json:"properties,omitempty"`
}

// ServerListResult is represents the response to a Get Azure SQL server
// request.
type ServerListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Server `json:"value,omitempty"`
}

// ServerMetric is represents Azure SQL server metrics.
type ServerMetric struct {
	ResourceName  *string    `json:"resourceName,omitempty"`
	DisplayName   *string    `json:"displayName,omitempty"`
	CurrentValue  *float64   `json:"currentValue,omitempty"`
	Limit         *float64   `json:"limit,omitempty"`
	Unit          *string    `json:"unit,omitempty"`
	NextResetTime *date.Time `json:"nextResetTime,omitempty"`
}

// ServerMetricListResult is represents the response to a List Azure SQL
// server metrics request.
type ServerMetricListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ServerMetric `json:"value,omitempty"`
}

// ServerProperties is represents the properties of an Azure SQL server.
type ServerProperties struct {
	FullyQualifiedDomainName   *string       `json:"fullyQualifiedDomainName,omitempty"`
	Version                    ServerVersion `json:"version,omitempty"`
	AdministratorLogin         *string       `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *string       `json:"administratorLoginPassword,omitempty"`
}

// ServiceObjective is represents an Azure SQL Database Service Objective.
type ServiceObjective struct {
	autorest.Response           `json:"-"`
	Name                        *string `json:"name,omitempty"`
	ID                          *string `json:"id,omitempty"`
	*ServiceObjectiveProperties `json:"properties,omitempty"`
}

// ServiceObjectiveListResult is represents the response to a Get Azure SQL
// Database Service Objectives request.
type ServiceObjectiveListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ServiceObjective `json:"value,omitempty"`
}

// ServiceObjectiveProperties is represents the properties of an Azure SQL
// Database Service Objective.
type ServiceObjectiveProperties struct {
	ServiceObjectiveName *string `json:"serviceObjectiveName,omitempty"`
	IsDefault            *bool   `json:"isDefault,omitempty"`
	IsSystem             *bool   `json:"isSystem,omitempty"`
	Description          *string `json:"description,omitempty"`
	Enabled              *bool   `json:"enabled,omitempty"`
}

// ServiceTierAdvisor is represents a Service Tier Advisor.
type ServiceTierAdvisor struct {
	autorest.Response             `json:"-"`
	Name                          *string `json:"name,omitempty"`
	ID                            *string `json:"id,omitempty"`
	*ServiceTierAdvisorProperties `json:"properties,omitempty"`
}

// ServiceTierAdvisorListResult is represents the response to a list service
// tier advisor request.
type ServiceTierAdvisorListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ServiceTierAdvisor `json:"value,omitempty"`
}

// ServiceTierAdvisorProperties is represents the properties of a Service Tier
// Advisor.
type ServiceTierAdvisorProperties struct {
	ObservationPeriodStart                                 *date.Time        `json:"observationPeriodStart,omitempty"`
	ObservationPeriodEnd                                   *date.Time        `json:"observationPeriodEnd,omitempty"`
	ActiveTimeRatio                                        *float64          `json:"activeTimeRatio,omitempty"`
	MinDtu                                                 *float64          `json:"minDtu,omitempty"`
	AvgDtu                                                 *float64          `json:"avgDtu,omitempty"`
	MaxDtu                                                 *float64          `json:"maxDtu,omitempty"`
	MaxSizeInGB                                            *float64          `json:"maxSizeInGB,omitempty"`
	ServiceLevelObjectiveUsageMetrics                      *[]SloUsageMetric `json:"serviceLevelObjectiveUsageMetrics,omitempty"`
	CurrentServiceLevelObjective                           *uuid.UUID        `json:"currentServiceLevelObjective,omitempty"`
	CurrentServiceLevelObjectiveID                         *uuid.UUID        `json:"currentServiceLevelObjectiveId,omitempty"`
	UsageBasedRecommendationServiceLevelObjective          *string           `json:"usageBasedRecommendationServiceLevelObjective,omitempty"`
	UsageBasedRecommendationServiceLevelObjectiveID        *uuid.UUID        `json:"usageBasedRecommendationServiceLevelObjectiveId,omitempty"`
	DatabaseSizeBasedRecommendationServiceLevelObjective   *string           `json:"databaseSizeBasedRecommendationServiceLevelObjective,omitempty"`
	DatabaseSizeBasedRecommendationServiceLevelObjectiveID *uuid.UUID        `json:"databaseSizeBasedRecommendationServiceLevelObjectiveId,omitempty"`
	DisasterPlanBasedRecommendationServiceLevelObjective   *string           `json:"disasterPlanBasedRecommendationServiceLevelObjective,omitempty"`
	DisasterPlanBasedRecommendationServiceLevelObjectiveID *uuid.UUID        `json:"disasterPlanBasedRecommendationServiceLevelObjectiveId,omitempty"`
	OverallRecommendationServiceLevelObjective             *string           `json:"overallRecommendationServiceLevelObjective,omitempty"`
	OverallRecommendationServiceLevelObjectiveID           *uuid.UUID        `json:"overallRecommendationServiceLevelObjectiveId,omitempty"`
	Confidence                                             *float64          `json:"confidence,omitempty"`
}

// SloUsageMetric is represents a Slo Usage Metric.
type SloUsageMetric struct {
	Name                    *string              `json:"name,omitempty"`
	ID                      *string              `json:"id,omitempty"`
	Type                    *string              `json:"type,omitempty"`
	Location                *string              `json:"location,omitempty"`
	Tags                    *map[string]*string  `json:"tags,omitempty"`
	ServiceLevelObjective   ServiceObjectiveName `json:"serviceLevelObjective,omitempty"`
	ServiceLevelObjectiveID *uuid.UUID           `json:"serviceLevelObjectiveId,omitempty"`
	InRangeTimeRatio        *float64             `json:"inRangeTimeRatio,omitempty"`
}

// SubResource is subresource properties
type SubResource struct {
	Name *string `json:"name,omitempty"`
	ID   *string `json:"id,omitempty"`
}

// Table is represents an Azure SQL Database table.
type Table struct {
	Name             *string             `json:"name,omitempty"`
	ID               *string             `json:"id,omitempty"`
	Type             *string             `json:"type,omitempty"`
	Location         *string             `json:"location,omitempty"`
	Tags             *map[string]*string `json:"tags,omitempty"`
	*TableProperties `json:"properties,omitempty"`
}

// TableProperties is represents the properties of an Azure SQL Database table.
type TableProperties struct {
	TableType          TableType           `json:"tableType,omitempty"`
	Columns            *[]Column           `json:"columns,omitempty"`
	RecommendedIndexes *[]RecommendedIndex `json:"recommendedIndexes,omitempty"`
}

// TransparentDataEncryption is represents an Azure SQL Database Transparent
// Data Encryption .
type TransparentDataEncryption struct {
	autorest.Response                    `json:"-"`
	Name                                 *string `json:"name,omitempty"`
	ID                                   *string `json:"id,omitempty"`
	*TransparentDataEncryptionProperties `json:"properties,omitempty"`
}

// TransparentDataEncryptionActivity is represents an Azure SQL Database
// Transparent Data Encryption Scan.
type TransparentDataEncryptionActivity struct {
	Name                                         *string `json:"name,omitempty"`
	ID                                           *string `json:"id,omitempty"`
	*TransparentDataEncryptionActivityProperties `json:"properties,omitempty"`
}

// TransparentDataEncryptionActivityListResult is represents the response to a
// List Azure SQL Database Transparent Data Encryption Activity request.
type TransparentDataEncryptionActivityListResult struct {
	autorest.Response `json:"-"`
	Value             *[]TransparentDataEncryptionActivity `json:"value,omitempty"`
}

// TransparentDataEncryptionActivityProperties is represents the properties of
// an Azure SQL Database Transparent Data Encryption Scan.
type TransparentDataEncryptionActivityProperties struct {
	Status          TransparentDataEncryptionActivityStates `json:"status,omitempty"`
	PercentComplete *float64                                `json:"percentComplete,omitempty"`
}

// TransparentDataEncryptionProperties is represents the properties of an
// Azure SQL Database Transparent Data Encryption.
type TransparentDataEncryptionProperties struct {
	Status TransparentDataEncryptionStates `json:"status,omitempty"`
}

// UpgradeHint is represents a Upgrade Hint.
type UpgradeHint struct {
	Name                          *string             `json:"name,omitempty"`
	ID                            *string             `json:"id,omitempty"`
	Type                          *string             `json:"type,omitempty"`
	Location                      *string             `json:"location,omitempty"`
	Tags                          *map[string]*string `json:"tags,omitempty"`
	TargetServiceLevelObjective   *string             `json:"targetServiceLevelObjective,omitempty"`
	TargetServiceLevelObjectiveID *uuid.UUID          `json:"targetServiceLevelObjectiveId,omitempty"`
}

// UpgradeRecommendedElasticPoolProperties is represents the properties of a
// Azure SQL Recommended Elastic Pool being upgraded.
type UpgradeRecommendedElasticPoolProperties struct {
	Name                *string                   `json:"Name,omitempty"`
	Edition             TargetElasticPoolEditions `json:"Edition,omitempty"`
	Dtu                 *int32                    `json:"Dtu,omitempty"`
	StorageMb           *int32                    `json:"StorageMb,omitempty"`
	DatabaseDtuMin      *int32                    `json:"DatabaseDtuMin,omitempty"`
	DatabaseDtuMax      *int32                    `json:"DatabaseDtuMax,omitempty"`
	DatabaseCollection  *[]string                 `json:"DatabaseCollection,omitempty"`
	IncludeAllDatabases *bool                     `json:"IncludeAllDatabases,omitempty"`
}