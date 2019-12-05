package apmz

import "testing"

func TestTelemetryConfiguration(t *testing.T) {
	testKey := "test"
	defaultEndpoint := "https://dc.services.visualstudio.com/v2/track"

	config := NewTelemetryConfiguration(testKey)

	if config.InstrumentationKey != testKey {
		t.Errorf("InstrumentationKey is %s, want %s", config.InstrumentationKey, testKey)
	}

	if config.EndpointURL != defaultEndpoint {
		t.Errorf("EndpointURL is %s, want %s", config.EndpointURL, defaultEndpoint)
	}
}
