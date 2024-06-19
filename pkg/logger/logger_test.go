package logger

import "testing"

func TestEntry_string(t *testing.T) {
	type fields struct {
		Message  string
		Severity Severity
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"given 'test message' and INFO severity produces test message and INFO severity in json format",
			fields{"test message", Info},
			"{\"message\":\"test message\",\"severity\":\"INFO\"}",
			false,
		},
		{
			"given empty message and severity still produces logger message and INFO severity",
			fields{"", Severity("")},
			"{\"message\":\"logger: No message provided\",\"severity\":\"INFO\"}",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Entry{
				Message:  tt.fields.Message,
				Severity: tt.fields.Severity,
			}
			got, err := e.string()
			if (err != nil) != tt.wantErr {
				t.Errorf("Entry.string() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Entry.string() = %v, want %v", got, tt.want)
			}
		})
	}
}
