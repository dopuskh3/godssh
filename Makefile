PROFILE_FILE=profile.out
COVERALLS_TOKEN=IVx3swS4YwoDCtOS5I1ETV6cY2R6ekGHV
TMP_COVER_DIR=tmp_cover

build:
	godep go build -v

test:
	godep go test -race -v `go list ./...` 

cover:
	set -x; \
	mkdir -p ${TMP_COVER_DIR}; \
	for pkg in `go list ./...`; do \
		echo $$pkg; \
		godep go test -v $$pkg -coverprofile=$$(mktemp -p ${TMP_COVER_DIR} -t coverXXX.out) || exit 1; \
	done; \
	echo "mode: set" > ${PROFILE_FILE}; \
	cat ${TMP_COVER_DIR}/cover*.out | grep -v "mode: set" >> ${PROFILE_FILE} || exit 1; \
	rm -rf ${TMP_COVER_DIR}

coveralls: cover
	goveralls -coverprofile=${PROFILE_FILE} -repotoken=${COVERALLS_TOKEN} || exit 1

