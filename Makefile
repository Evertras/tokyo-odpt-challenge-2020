################################################################################
# This Makefile contains useful commands and ways to grab data.  Requires the
# ODPT_TOKEN environment variable to be set.

# Stick this in a variable for easier conditionals later
TOKEN=${ODPT_TOKEN}

# Build everything
build: bin/importer valid-token
	# Available binaries in ./bin
	@ls bin

# Make sure we have a valid token
valid-token:
ifeq ($(TOKEN),)
	@echo "=====> ERROR: Please set ODPT_TOKEN environment variable"
	@exit 1
else
	@echo "You have a valid token ready to go!"
endif

# Delete all the stuff we've done
clean:
	rm -rf ./bin

fmt:
	@go fmt ./pkg/...
	@go fmt ./cmd/...

################################################################################
# Dependencies
GO_FILES = $(shell find . -type f -name '*.go')

################################################################################
# Compiling our CLI stuff
bin:
	mkdir -p bin

bin/importer: bin $(GO_FILES)
	go build -o bin/importer ./cmd/importer/main.go

################################################################################
# Elasticsearch helpers
clear-es:
	@echo "Deleting ES indices, some may not be found and that's fine"
	curl -XDELETE 'http://localhost:9200/passenger-survey'
	curl -XDELETE 'http://localhost:9200/bus-stop-pole'
	curl -XDELETE 'http://localhost:9200/bus-route-pattern'
	@echo

################################################################################
# Data dump API
#
# These are relatively large, static data dumps that can be downloaded once and
# then generally forgotten about.

DATA_DUMP_TYPES := Calendar Operator Station StationTimetable TrainTimetable TrainType RailDirection Railway RailwayFare PassengerSurvey BusTimetable BusroutePattern BusroutePatternFare BusstopPole BusstopPoleTimetable Airport AirportTerminal FlightSchedule FlightStatus
DATA_DUMP_FILES := $(foreach TYPE,$(DATA_DUMP_TYPES),data/$(TYPE).json)

data:
	mkdir -p data

$(DATA_DUMP_FILES): data
	@echo Retrieving $@
	curl -X GET -L https://api-tokyochallenge.odpt.org/api/v4/odpt:$(subst data/,,$@)?acl:consumerKey=$(TOKEN) > $@
	@touch $@

