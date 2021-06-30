package herr

import (
	"errors"
	"net/http"
	"testing"
)

func TestError_As(t *testing.T) {
	type fields struct {
		CodeField        int
		StatusField      string
		ReasonField      string
		RequestIDField   string
		MessageField     string
		DescriptionField string
		DebugField       string
		DetailsField     map[string]interface{}
		ItemsField       []Item
		err              error
	}
	type args struct {
		target interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "valid target",
			fields: fields{
				CodeField:        http.StatusOK,
				StatusField:      http.StatusText(http.StatusOK),
				ReasonField:      "BAD_DAY",
				RequestIDField:   "66f6666c-6b66-6be6-66e6-66f666bb666b",
				MessageField:     "all is bad",
				DescriptionField: "It couldn't get any worse.",
				DebugField:       "ASS0117: hell is opened",
				DetailsField:     map[string]interface{}{"opps": "opps"},
				ItemsField:       []Item{{Name: "future", Code: "finished", Description: "Today."}},
				err:              errors.New("don't worry"),
			},
			args: args{&Error{}},
			want: true,
		},
		{
			name: "invalid target",
			fields: fields{
				CodeField:        http.StatusForbidden,
				StatusField:      http.StatusText(http.StatusForbidden),
				ReasonField:      "GOOD_DAY",
				RequestIDField:   "77f7777c-7b77-7be7-77e7-77f777bb777b",
				MessageField:     "all is good",
				DescriptionField: "It couldn't be better.",
				DebugField:       "CANDY0888: paradise is returned",
				DetailsField:     map[string]interface{}{"wow": "wow"},
				ItemsField:       []Item{{Name: "future", Code: "happened", Description: "Today."}},
				err:              errors.New("do you?"),
			},
			args: args{Error{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				CodeField:        tt.fields.CodeField,
				StatusField:      tt.fields.StatusField,
				ReasonField:      tt.fields.ReasonField,
				RequestIDField:   tt.fields.RequestIDField,
				MessageField:     tt.fields.MessageField,
				DescriptionField: tt.fields.DescriptionField,
				DebugField:       tt.fields.DebugField,
				DetailsField:     tt.fields.DetailsField,
				ItemsField:       tt.fields.ItemsField,
				err:              tt.fields.err,
			}
			if got := e.As(tt.args.target); got != tt.want {
				t.Errorf("Error.As() = %v, want %v", got, tt.want)
			}
		})
	}
}
