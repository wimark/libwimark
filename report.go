package libwimark

import (
	"time"
)

const (
	COLL_STAT_REPORT        = "reports"
	COLL_STAT_REPORT_RESULT = "report_results"
)

type timeBounds struct {
	Start int64 `json:"start" bson:"start"`
	Stop  int64 `json:"stop" bson:"stop"`
}

// StatReport struct for represent auto reports
type StatReport struct {
	ID string `json:"id" bson:"_id"`

	// basic fields
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// location dat
	Location string `json:"-" bson:"location"`
	// subject of report (CPEs, Clients, Events, Custom)
	Subject ReportSubject `json:"subject" bson:"subject"`
	// CustomSubject

	// report type (current, summary)
	Type ReportType `json:"type" bson:"type"`

	// collect period (now, day, week, month)
	Period ReportPeriod `json:"collect_period" bson:"collect_period"`

	// period timebounds if Period == now
	TimeBounds timeBounds `json:"timebounds" bson:"timebounds"`

	// report format
	Format ReportFormat `json:"format" bson:"format"`

	// time to do post action (CRON string)
	ActionTime string `json:"action_time" bson:"action_time"`

	// type of post action (email, script)
	ActionType ReportActionType `json:"action_type" bson:"action_type"`

	// post action address dests
	ActionDest []string `json:"action_dest" bson:"action_dest"`

	// add charts
	Charts bool `json:"charts" bson:"charts"`
}

// StatReportResult struct for store collected reports
type StatReportResult struct {
	ID string `json:"id" bson:"_id"`

	Report StatReport  `json:"report_id" bson:"report_id"`
	Data   interface{} `json:"data" bson:"data"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`
}
