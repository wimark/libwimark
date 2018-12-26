package libwimark

import (
	"encoding/json"
	"testing"
)

func TestParsingRRMEvent(t *testing.T) {
	bt := []byte(`
		{
			"cpes": [
			  {
				"cpe_id" : "643ba589-88af-4786-0000-f09fc2f3069e",
				"channel" : {
				  "channel" : 1
				},
				"power" : 10
			  },
			  {
				"cpe_id" : "643ba589-88af-4786-1111-f09fc2f3069e",
				"channel" : {
				  "channel" : 7
				},
				"power" : 20
			  }
			]
	    }
	`)

	var data RRMStatusData
	e := json.Unmarshal(bt, &data)
	if e != nil {
		t.Error(e)
	}

	if l := len(data.Cpes); l != 2 {
		t.Error()
		return
	}

	if data.Cpes[0].CpeID != "643ba589-88af-4786-0000-f09fc2f3069e" {
		t.Error()
	}

	if data.Cpes[1].Channel.Channel != 7 {
		t.Error()
	}

	if data.Cpes[1].Power != 20 {
		t.Error()
	}
}
