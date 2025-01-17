package constants

// All constants ending with 'Name' might have their value hard-coded in configmap/ directory
// Do not change their value.
const (
	BaseName = "lsst"
	AppLabel = "qserv"

	CmsdPort     = 2131
	CmsdPortName = string(CmsdName)

	CorePathVolumeName = "corepath"

	DashboardPort     = 8080
	DashboardPortName = "http"

	DataVolumeClaimTemplateName = QservName + "-data"

	MariadbPort     = 3306
	MariadbPortName = string(MariadbName)

	ProxyPort     = 4040
	ProxyPortName = string(ProxyName)

	QservName = "qserv"

	RedisName = "redis"

	ReplicationControllerPort     = 8080
	ReplicationControllerPortName = "http"

	WmgrPort     = 8080
	WmgrPortName = "http"

	XrootdAdminPathVolumeName = "xrootd-adminpath"
	XrootdPort                = 1094
	XrootdPortName            = string(XrootdName)

	GraceTime = 30

	ReplicationWorkerDefaultThreads = 16
)

// QservGID qserv user gid
var QservGID int64 = 1000

// QservUID qserv user uid
var QservUID int64 = 1000

// ContainerName name all containers
type ContainerName string

const (
	// CmsdName name for cmsd containers
	CmsdName ContainerName = "cmsd"
	// DashboardName name for dashboard container
	DashboardName ContainerName = "dashboard"
	// IngestDbName name for ingest database container
	IngestDbName ContainerName = "ingest-db"
	// InitDbName name for database initialization containers
	InitDbName ContainerName = "initdb"
	// MariadbName name for mariadb container
	MariadbName ContainerName = "mariadb"
	// ProxyName name for proxy container
	ProxyName ContainerName = "proxy"
	// ReplCtlName name for replication controller container
	ReplCtlName ContainerName = "repl-ctl"
	// ReplDbName name for replication database container
	ReplDbName ContainerName = "repl-db"
	// XrootdName name for xrootd container
	XrootdName ContainerName = "xrootd"
	// WmgrName name for worker manager container
	WmgrName ContainerName = "wmgr"
	// ReplWrkName name for replication worker container
	ReplWrkName ContainerName = "repl-wrk"
)

// PodClass name all classes of pod
// used to generate pod labels
type PodClass string

// GetDbContainerName return name of a database container for a given pod
func GetDbContainerName(pod PodClass) ContainerName {
	var dbName ContainerName
	if pod == ReplDb {
		dbName = ReplDbName
	} else if pod == IngestDb {
		dbName = IngestDbName
	} else {
		dbName = MariadbName
	}
	return dbName
}

const (
	// Czar name pods of class Czar
	Czar PodClass = "czar"
	// Dashboard name pods of class Dashboard
	Dashboard PodClass = PodClass(DashboardName)
	// IngestDb name pods of class Ingest database
	IngestDb PodClass = PodClass(IngestDbName)
	// ReplCtl name pods of class Replication controller
	ReplCtl PodClass = PodClass(ReplCtlName)
	// ReplDb name pods of class Replication database
	ReplDb PodClass = PodClass(ReplDbName)
	// Worker name pods of class Replication worker
	Worker PodClass = "worker"
	// XrootdRedirector name pods of class Xrootd redirector
	XrootdRedirector PodClass = "xrootd-redirector"
)

// ContainerConfigmaps contains names of all micro-services which require configmaps named:
// '<prefix>-<microservice-name>-etc' and '<prefix>-<microservice-name>-start'
var ContainerConfigmaps = []ContainerName{DashboardName, IngestDbName, MariadbName, XrootdName, ProxyName, WmgrName, ReplCtlName, ReplDbName, ReplWrkName}

// MicroserviceSecrets contains names of all micro-services which require secrets
var MicroserviceSecrets = []ContainerName{MariadbName, WmgrName, ReplDbName}

// Databases contains names of all Qserv pods which embed a database container
var Databases = []PodClass{Czar, ReplDb, Worker, IngestDb}

// Command contains the default command used to launch a container
var Command = []string{"/config-start/start.sh"}

// CommandDebug is a prerequisite for interactive debugging
var CommandDebug = []string{"sleep", "infinity"}
