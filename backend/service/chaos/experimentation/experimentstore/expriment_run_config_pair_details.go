package experimentstore

import (
	"database/sql"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/any"
	experimentation "github.com/lyft/clutch/backend/api/chaos/experimentation/v1"
	"strconv"
	"strings"
	"time"
)

func NewRunConfigPairDetails(fetchedID uint64, startTime sql.NullTime, endTime sql.NullTime, scheduledEndTime sql.NullTime, creationTime time.Time, details string) (*experimentation.ExperimentRunConfigPairDetails, error) {
	runConfigPair := &experimentation.ExperimentRunConfigPairDetails{}
	runConfigPair.RunId = fetchedID
	runConfigPair.Form = &experimentation.Form{}
	runConfigPair.Status = timesToStatus(startTime, endTime, scheduledEndTime)
	runConfigPair.GetForm().Fields = []*experimentation.TextField{
		{Label: "Run Identifier", Value: strconv.FormatUint(fetchedID, 10)},
		{Label: "Status", Value: statusToString(runConfigPair.Status)},
		{Label: "Start Time", Value: timeToString(startTime)},
	}

	if runConfigPair.Status == experimentation.Status_COMPLETED {
		endTimeField := &experimentation.TextField{Label: "Scheduled End Time", Value: timeToString(endTime)}
		runConfigPair.GetForm().Fields = append(runConfigPair.GetForm().Fields, endTimeField)
	}

	scheduledEndTimeField := &experimentation.TextField{Label: "Scheduled End Time", Value: timeToString(endTime)}
	runConfigPair.GetForm().Fields = append(runConfigPair.GetForm().Fields, scheduledEndTimeField)

	anyConfig := &any.Any{}
	err := jsonpb.Unmarshal(strings.NewReader(details), anyConfig)
	if err != nil {
		return nil, err
	}

	runConfigPair.Config = anyConfig
	return runConfigPair, err
}

func timeToString(t sql.NullTime) string {
	layout := "01/02/06 15:04:05"
	if t.Valid {
		return t.Time.Format(layout)
	} else {
		return "Undefined"
	}
}

func timesToStatus(startTime sql.NullTime, endTime sql.NullTime, scheduledEndTime sql.NullTime) experimentation.Status {
	now := time.Now()
	if startTime.Valid && !endTime.Valid {
		if now.After(startTime.Time) {
			return experimentation.Status_RUNNING
		} else {
			return experimentation.Status_SCHEDULED
		}
	} else if !startTime.Valid && endTime.Valid {
		if now.Before(endTime.Time) {
			return experimentation.Status_RUNNING
		} else {
			return experimentation.Status_COMPLETED
		}
	} else {
		if now.Before(startTime.Time) {
			return experimentation.Status_SCHEDULED
		} else if now.After(startTime.Time) && now.Before(endTime.Time) {
			return experimentation.Status_RUNNING
		} else {
			return experimentation.Status_COMPLETED
		}
	}

	return experimentation.Status_UNKNOWN
}

func statusToString(status experimentation.Status) string {
	switch status {
	case experimentation.Status_UNKNOWN:
		return "Unknown"
	case experimentation.Status_SCHEDULED:
		return "Scheduled"
	case experimentation.Status_RUNNING:
		return "Running"
	case experimentation.Status_COMPLETED:
		return "Completed"
	}

	return status.String()
}
