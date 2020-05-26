package maestro

import (
	"errors"
	"time"

	"github.com/open-service-mesh/osm/pkg/logger"
)

// TestResult is the type for the test result enum
type TestResult int

const (
	// TestsPassed is used for tests that passed.
	TestsPassed TestResult = iota + 1

	// TestsFailed is used for tests that failed.
	TestsFailed

	// TestsTimedOut is used for tests that timed out.
	TestsTimedOut

	// KubeConfigEnvVar is the environment variable for KUBECONFIG.
	KubeConfigEnvVar = "KUBECONFIG"

	// OSMNamespaceEnvVar is the environment variable for the OSM namespace.
	OSMNamespaceEnvVar = "K8S_NAMESPACE"

	// BookbuyerNamespaceEnvVar is the environment variable for the Bookbuyer namespace.
	BookbuyerNamespaceEnvVar = "BOOKBUYER_NAMESPACE"

	// BookthiefNamespaceEnvVar is the environment variable for the Bookbuyer namespace.
	BookthiefNamespaceEnvVar = "BOOKTHIEF_NAMESPACE"

	// BookstoreNamespaceEnvVar is the environment variable for the Bookbuyer namespace.
	BookstoreNamespaceEnvVar = "BOOKSTORE_NAMESPACE"

	// WaitForPodTimeSecondsEnvVar is the environment variable for the time we will wait on the pod to be ready.
	WaitForPodTimeSecondsEnvVar = "WAIT_FOR_POD_TIME_SECONDS"

	// OsmIDEnvVar is the environment variable for the ID of an OSM instance
	OsmIDEnvVar = "OSM_ID"
)

var (
	// WaitForPod is the time we wait for a pod to become ready
	WaitForPod = 5 * time.Second

	// PollLogsFromTimeSince is the interval we go back in time to get pod logs
	PollLogsFromTimeSince = 2 * time.Second

	// FailureLogsFromTimeSince is the interval we go back in time to get pod logs
	FailureLogsFromTimeSince = 10 * time.Minute

	log            = logger.New("ci/maestro")
	errNoPodsFound = errors.New("no pods found")
)
