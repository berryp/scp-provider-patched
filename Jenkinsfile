// Build the patched SCP terraform provider.
//
// Output: /data1/deployment/scp-provider-patched/terraform-provider-samsungcloudplatform
//
// Three patches against upstream v3.16.1 are baked in (see README.md):
//   1. kubernetes_node_pool composite-id Importer (engineId/nodePoolId)
//   2. readNodePool populates schema-declared attributes
//   3. kubernetes_namespace composite-id Importer (engineId/namespaceName)
//
// Consumers (e.g. scp-ske-statefix in the terragrunt-profiles repo) read the
// binary from the fixed workspace path above via a terraformrc dev_overrides
// entry. They can trigger this job on a cache miss via `build job:
// 'build-scp-provider-patched', wait: true`. See the consumer Jenkinsfile for
// the wiring.

pipeline {
    agent {
        node {
            label 'scp'
            // Fixed workspace — this directory IS the published binary location.
            // Jenkins' default workspace locking prevents concurrent builds from
            // clobbering one another. SCM checkout lands here, the build writes
            // the binary here, and consumers read from here.
            customWorkspace '/data1/deployment/scp-provider-patched'
        }
    }
    options {
        // Discard old build records (the workspace persists; we just don't need
        // an unbounded list of run history for a build job that produces a
        // single shared binary).
        buildDiscarder(logRotator(numToKeepStr: '20'))
        ansiColor('xterm')
        timestamps()
        timeout(time: 10, unit: 'MINUTES')
    }
    triggers {
        // Poll the repo every 15 minutes for new commits. Webhooks would be
        // better but polling is the path of least resistance for a private
        // Azure DevOps remote.
        pollSCM('H/15 * * * *')
    }
    environment {
        BINARY_NAME = 'terraform-provider-samsungcloudplatform'
        GOTOOLCHAIN = 'go1.26.4'
    }
    stages {
        stage('Build') {
            steps {
                sh '''#!/bin/bash
set -euo pipefail

if ! command -v go >/dev/null 2>&1; then
    echo "ERROR: go toolchain not found on PATH; install Go on the Jenkins agent." >&2
    exit 1
fi

# Build into a temp file in the same directory as the final binary so the
# atomic rename below works on the same filesystem. We never overwrite the
# live binary until the new one has passed every sanity check.
TMP_BIN="$(mktemp -p "${WORKSPACE}" "${BINARY_NAME}.new.XXXXXX")"
trap 'rm -f "${TMP_BIN}"' EXIT

# Install required go toolchain
go install golang.org/dl/${GOTOOLCHAIN}@latest
echo "Using $(go version)"

echo "Running go build"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \\
    go build -mod=vendor -o "${TMP_BIN}" .

# Every patched Importer + a sanity-check upstream string must be in the
# binary. If any is missing, something replaced one of the source files before
# the build — fail fast instead of promoting a broken binary on top of a
# working one.
MARKERS=(
    'import id must be in format <engineId>/<nodePoolId>'
    'import id must be in format <engineId>/<namespaceName>'
    'desired_node_count is not supported when auto_scale is enabled'
)
for m in "${MARKERS[@]}"; do
    if ! grep -aq "$m" "${TMP_BIN}"; then
        echo "ERROR: built binary is missing required marker: $m" >&2
        exit 1
    fi
done

# Atomic promote: rename(2) is atomic on the same filesystem, so a consumer
# reading the live path either sees the old binary or the new — never half.
LIVE_BIN="${WORKSPACE}/${BINARY_NAME}"
chmod 0755 "${TMP_BIN}"
mv -f "${TMP_BIN}" "${LIVE_BIN}"
trap - EXIT

echo
echo "── Promoted ─────────────────────────────────────────────────────────────"
ls -lh "${LIVE_BIN}"
echo "  sha256: $(sha256sum "${LIVE_BIN}" | awk '{print $1}')"
echo "  git:    $(git rev-parse HEAD)  ($(git describe --always --dirty 2>/dev/null || echo no-tags))"
echo "  date:   $(date -u +%Y-%m-%dT%H:%M:%SZ)"
echo "─────────────────────────────────────────────────────────────────────────"
'''
            }
        }
    }
}
