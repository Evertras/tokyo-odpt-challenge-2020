################################################################################
# This Makefile contains useful commands and ways to grab data.  Requires the
# ODPT_TOKEN environment variable to be set.

# Stick this in a variable for easier conditionals later
TOKEN=${ODPT_TOKEN}

# Make sure we have a valid token
valid-token:
ifeq ($(TOKEN),)
	@echo Please set ODPT_TOKEN environment variable
	@exit 1
else
	@echo hm
endif

# Delete all the stuff we've done
clean:


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
