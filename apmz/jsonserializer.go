package apmz

import (
	"bytes"
	"encoding/json"

	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
)

type telemetryBufferItems []*contracts.Envelope

func (tbis telemetryBufferItems) serialize() []byte {
	var result bytes.Buffer
	encoder := json.NewEncoder(&result)

	for _, item := range tbis {
		end := result.Len()
		if err := encoder.Encode(item); err != nil {
			diagnosticsWriter.Printf("Telemetry item failed to serialize: %s", err.Error())
			result.Truncate(end)
		}
	}

	return result.Bytes()
}
