package napodate

import (
	"context"
	//"fmt"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	srv, ctx := setup()
	s, err := srv.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	ok := s == "ok"
	if !ok {
		t.Errorf("expected service to be ok")
	}
}

func TestGet(t *testing.T) {
	srv, ctx := setup()
	d, err := srv.Get(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	time := time.Now()
	today := time.Format("02/01/2006")
	//fmt.Println("TextGet.1.today=", today)

	ok := today == d
	if !ok {
		t.Errorf("expected dates to be equal")
	}
}

func TestValidate(t *testing.T) {
	srv, ctx := setup()
	time := time.Now()
	today := time.Format("02/01/2006")
	//fmt.Println("TestValidate.1.today=", today)
	b, err := srv.Validate(ctx, today)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !b {
		t.Errorf("date should be valid")
	}

	today = "31/31/2019"
	//fmt.Println("TestValidate.2.today=", today)
	b, err = srv.Validate(ctx, today)
	if b {
		t.Errorf("date should be invalid: %v", b)
	}

	today = "12/31/2019"
	//fmt.Println("TestValidate.3.today=", today)
	b, err = srv.Validate(ctx, today)
	if b {
		t.Errorf("USA date should be invalid: %v", b)
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
