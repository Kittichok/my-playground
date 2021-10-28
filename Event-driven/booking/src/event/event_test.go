package event

import (
	"context"
	"testing"
)

func TestSubmitEvent(t *testing.T) {

	e := NewEventConnection()
	ctx := context.Background()
	e.SubmitMessage(ctx, "BookingSubmit", "{detail: [{pid:1,q:2}]}")

}
