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
				"channel" : 11,
				"dbm" : 10
			  },
			  {
				"cpe_id" : "643ba589-88af-4786-1111-f09fc2f3069e",
				"channel" : 7,
				"dbm" : 20
			  }
			]
	    }
	`)

	var data RRMStatusData
	json.Unmarshal(bt, &data)

	if l := len(data.Cpes); l != 2 {
		t.Error()
	}

	if data.Cpes[0].CpeID != "643ba589-88af-4786-0000-f09fc2f3069e" {
		t.Error()
	}

	if data.Cpes[1].Channel != 7 {
		t.Error()
	}

	if data.Cpes[1].Dbm != 20 {
		t.Error()
	}
}
