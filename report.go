package libwimark

import (
	"time"
)

const (
	COLL_STAT_REPORT        = "reports"
	COLL_STAT_REPORT_RESULT = "report_results"
)

// StatReport struct for represent auto reports
type StatReport struct {
	ID string `json:"id" bson:"_id"`

	// basic fields
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`

	// subject of report (CPEs, Clients, Events)
	Subject string `json:"subject" bson:"subject"`

	// report type (current, summary)
	Type string `json:"type" bson:"type"`

	// collect period (now, day, week, month)
	CollectPeriod string `json:"collect_period" bson:"collect_period"`

	// report format
	Format string `json:"format" bson:"format"`

	// time to do post action (CRON string)
	ActionTime string `json:"action_time" bson:"action_time"`

	// type of post action (email, script)
	ActionType string `json:"action_type" bson:"action_type"`

	// post action address dests
	ActionDest []string `json:"action_dest" bson:"action_dest"`
}

// StatReportResult struct for store collected reports
type StatReportResult struct {
	ID string `json:"id" bson:"_id"`

	Report StatReport    `json:"report_id" bson:"report_id"`
	Data   []interface{} `json:"data" bson:"data"`

	CreateAt time.Time `json:"create_at" bson:"create_at"`
}
