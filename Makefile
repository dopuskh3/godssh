PACKAGES="dssh"
PROFILE_FILE="profile.out"
COVERALLS_TOKEN="IVx3swS4YwoDCtOS5I1ETV6cY2R6ekGHV"

build:
	godep go build -v

test:
	godep go test -race -v ${PACKAGES}

cover:
	set -x; \
	mkdir -p cover_tmp; \
	for pkg in `go list ./...`; do \
		echo $$pkg; \
		godep go test -v $$pkg -coverprofile=$$(mktemp -p cover_tmp -t coverXXX.out); \
	done; \
	echo "mode: set" > ${PROFILE_FILE}; \
	cat tmp/cover*.out | grep -v "mode: set" >> ${PROFILE_FILE}; \
	rm -rf cover_tmp

coveralls: cover
	goveralls -coverprofile=${PROFILE_FILE} ${COVERALLS_TOKEN}

