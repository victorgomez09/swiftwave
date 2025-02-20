package core

import (
	"database/sql/driver"
	"encoding/json"
)

// ************************************************************************************* //
//                                Swiftwave System Configuration 		   			     //
// ************************************************************************************* //

// UserRole : role of the registered user
type UserRole string

const (
	// AdministratorRole : admin user can perform any operation on the system
	AdministratorRole UserRole = "admin"
	// ManagerRole : manager user can perform any operation on the system
	// except user management, system configuration and server management
	ManagerRole UserRole = "manager"
)

// ServerStatus : status of the server
type ServerStatus string

const (
	ServerNeedsSetup ServerStatus = "needs_setup"
	ServerPreparing  ServerStatus = "preparing"
	ServerOnline     ServerStatus = "online"
	ServerOffline    ServerStatus = "offline"
)

// SwarmMode : mode of the swarm
type SwarmMode string

const (
	SwarmManager SwarmMode = "manager"
	SwarmWorker  SwarmMode = "worker"
)

// ProxyType : type of the proxy
type ProxyType string

const (
	BackupProxy ProxyType = "backup"
	ActiveProxy ProxyType = "active"
)

// ProxyConfig : hold information about proxy configuration
type ProxyConfig struct {
	Enabled      bool      `json:"enabled" gorm:"default:false"`
	SetupRunning bool      `json:"setup_running" gorm:"default:false"` // just to show warning to user, that's it
	Type         ProxyType `json:"type" gorm:"default:'active'"`
}

// ************************************************************************************* //
//                                Application Level Table       		   			     //
// ************************************************************************************* //

// UpstreamType : type of source for the codebase of the application
type UpstreamType string

const (
	UpstreamTypeGit        UpstreamType = "git"
	UpstreamTypeSourceCode UpstreamType = "sourceCode"
	UpstreamTypeImage      UpstreamType = "image"
)

// GitProvider : type of git provider
type GitProvider string

const (
	GitProviderNone   GitProvider = "none"
	GitProviderGithub GitProvider = "github"
	GitProviderGitlab GitProvider = "gitlab"
)

// DomainSSLStatus : status of the ssl certificate for a domain
type DomainSSLStatus string

const (
	DomainSSLStatusNone    DomainSSLStatus = "none"
	DomainSSLStatusPending DomainSSLStatus = "pending"
	DomainSSLStatusFailed  DomainSSLStatus = "failed"
	DomainSSLStatusIssued  DomainSSLStatus = "issued"
)

// DeploymentStatus : status of the deployment
type DeploymentStatus string

const (
	DeploymentStatusPending       DeploymentStatus = "pending"
	DeploymentStatusDeployPending DeploymentStatus = "deployPending"
	DeploymentStatusLive          DeploymentStatus = "live"
	DeploymentStatusStopped       DeploymentStatus = "stopped"
	DeploymentStatusFailed        DeploymentStatus = "failed"
	DeploymentStalled             DeploymentStatus = "stalled"
)

// GitType type of git credential
type GitType string

const (
	GitHttp GitType = "http"
	GitSsh  GitType = "ssh"
)

// ProtocolType : type of protocol for ingress rule
type ProtocolType string

const (
	HTTPProtocol  ProtocolType = "http"
	HTTPSProtocol ProtocolType = "https"
	TCPProtocol   ProtocolType = "tcp"
	UDPProtocol   ProtocolType = "udp"
)

// IngressRuleTargetType : type of target for ingress rule
type IngressRuleTargetType string

const (
	ApplicationIngressRule     IngressRuleTargetType = "application"
	ExternalServiceIngressRule IngressRuleTargetType = "externalService"
)

// IngressRuleStatus : status of the ingress rule
type IngressRuleStatus string

const (
	IngressRuleStatusPending  IngressRuleStatus = "pending"
	IngressRuleStatusApplied  IngressRuleStatus = "applied"
	IngressRuleStatusFailed   IngressRuleStatus = "failed"
	IngressRuleStatusDeleting IngressRuleStatus = "deleting"
)

// RedirectRuleStatus : status of the redirect rule
type RedirectRuleStatus string

const (
	RedirectRuleStatusPending  RedirectRuleStatus = "pending"
	RedirectRuleStatusApplied  RedirectRuleStatus = "applied"
	RedirectRuleStatusFailed   RedirectRuleStatus = "failed"
	RedirectRuleStatusDeleting RedirectRuleStatus = "deleting"
)

// DeploymentMode : mode of deployment of application (replicated or global)
type DeploymentMode string

const (
	DeploymentModeReplicated DeploymentMode = "replicated"
	DeploymentModeGlobal     DeploymentMode = "global"
)

// ApplicationUpdateResult : result of application update
type ApplicationUpdateResult struct {
	RebuildRequired bool
	ReloadRequired  bool
	DeploymentId    string
}

// DeploymentUpdateResult : result of deployment update
type DeploymentUpdateResult struct {
	RebuildRequired bool
	DeploymentId    string
}

// BackupType : type of backup
type BackupType string

const (
	LocalBackup BackupType = "local"
	S3Backup    BackupType = "s3"
)

// BackupStatus : status of backup
type BackupStatus string

const (
	BackupPending BackupStatus = "pending"
	BackupFailed  BackupStatus = "failed"
	BackupSuccess BackupStatus = "success"
)

// RestoreType : type of restore
type RestoreType string

const (
	LocalRestore RestoreType = "local"
)

// RestoreStatus : status of restore
type RestoreStatus string

const (
	RestorePending RestoreStatus = "pending"
	RestoreFailed  RestoreStatus = "failed"
	RestoreSuccess RestoreStatus = "success"
)

// PersistentVolumeType : type of persistent volume
type PersistentVolumeType string

const (
	PersistentVolumeTypeLocal PersistentVolumeType = "local"
	PersistentVolumeTypeNFS   PersistentVolumeType = "nfs"
	PersistentVolumeTypeCIFS  PersistentVolumeType = "cifs"
)

// NFSConfig : configuration for NFS Storage
type NFSConfig struct {
	Host    string `json:"host,omitempty"`
	Path    string `json:"path,omitempty"`
	Version int    `json:"version,omitempty"`
}

// CIFSConfig : configuration for CIFS Storage
type CIFSConfig struct {
	Share    string `json:"share"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	FileMode string `json:"file_mode"`
	DirMode  string `json:"dir_mode"`
	Uid      int    `json:"uid" gorm:"default:0"`
	Gid      int    `json:"gid" gorm:"default:0"`
}

var RequiredServerDependencies = []string{
	"init",
	"awk",
	"curl",
	"unzip",
	"git",
	"tar",
	"iproute",
	"nfs",
	"cifs",
	"rsync",
	"docker",
}

var DependencyCheckCommands = map[string]string{
	"init":    "echo hi", // dummy command
	"awk":     "which awk",
	"curl":    "which curl",
	"unzip":   "which unzip",
	"git":     "which git",
	"tar":     "which tar",
	"iproute": "which ip",
	"nfs":     "which nfsstat",
	"cifs":    "which mount.cifs",
	"rsync":   "which rsync",
	"docker":  "which docker",
}

var DebianDependenciesInstallCommands = map[string]string{
	"init":    "apt -y update",
	"awk":     "apt install -y gawk",
	"curl":    "apt install -y curl",
	"unzip":   "apt install -y unzip",
	"git":     "apt install -y git",
	"tar":     "apt install -y tar",
	"iproute": "apt install -y iproute2",
	"nfs":     "apt install -y nfs-common && systemctl stop rpcbind.socket && systemctl stop rpcbind && systemctl disable rpcbind.socket && systemctl disable rpcbind",
	"cifs":    "apt install -y cifs-utils",
	"rsync":   "apt install -y rsync",
	"docker":  "curl -fsSL get.docker.com | sh -",
}
var FedoraDependenciesInstallCommands = map[string]string{
	"init":    "dnf -y update",
	"awk":     "dnf install -y gawk",
	"curl":    "dnf install -y curl",
	"unzip":   "dnf install -y unzip",
	"git":     "dnf install -y git",
	"tar":     "dnf install -y tar",
	"iproute": "dnf install -y iproute",
	"nfs":     "dnf install -y nfs-utils && systemctl stop rpcbind.socket && systemctl stop rpcbind && systemctl disable rpcbind.socket && systemctl disable rpcbind",
	"cifs":    "dnf install -y cifs-utils",
	"rsync":   "dnf install -y rsync",
	"docker":  "curl -fsSL get.docker.com | sh -",
}

// ConsoleTarget : type of console target
type ConsoleTarget string

const (
	ConsoleTargetTypeServer      ConsoleTarget = "server"
	ConsoleTargetTypeApplication ConsoleTarget = "application"
)

// ************************************************************************************* //
//                              	Server Related Stats       		   			         //
// ************************************************************************************* //

type ServerDiskStat struct {
	Path       string  `json:"path"`
	MountPoint string  `json:"mount_point"`
	TotalGB    float32 `json:"total_gb"`
	UsedGB     float32 `json:"used_gb"`
}

type ServerDiskStats []ServerDiskStat

// Scan implement value scanner interface for gorm
func (d *ServerDiskStats) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), d)
}

// Value implement driver.Valuer interface for gorm
func (d ServerDiskStats) Value() (driver.Value, error) {
	return json.Marshal(d)
}

type ServerMemoryStat struct {
	TotalGB  float32 `json:"total_gb"`
	UsedGB   float32 `json:"used_gb"`
	CachedGB float32 `json:"cached_gb"`
}

type ServerNetStat struct {
	SentKB   uint64 `json:"sent_kb"`
	RecvKB   uint64 `json:"recv_kb"`
	SentKBPS uint64 `json:"sent_kbps"`
	RecvKBPS uint64 `json:"recv_kbps"`
}

// ************************************************************************************* //
//                            	Application Related Stats       		   			     //
// ************************************************************************************* //

type ApplicationServiceNetStat struct {
	SentKB   uint64 `json:"sent_kb"`
	RecvKB   uint64 `json:"recv_kb"`
	SentKBPS uint64 `json:"sent_kbps"`
	RecvKBPS uint64 `json:"recv_kbps"`
}

// ************************************************************************************* //
//                               	Application Related       		    		 	     //
// ************************************************************************************* //

type ApplicationResourceLimit struct {
	MemoryMB int `json:"memory_mb" gorm:"default:0"`
}

type ApplicationReservedResource struct {
	MemoryMB int `json:"memory_mb" gorm:"default:0"`
}
