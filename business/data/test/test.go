package test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	db "github.com/thetnaingtn/go-dermacare-service/business/sys/database/mongo"
	"github.com/thetnaingtn/go-dermacare-service/foundation/docker"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

type DBContainer struct {
	Image string
	Port  string
	Args  []string
}

func NewUnit(t *testing.T, dbc DBContainer) (*mongo.Database, func()) {
	// dbc.Port is the port listen on the container
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	c := docker.StartContainer(t, dbc.Image, dbc.Port, dbc.Args...)
	db, cleanup := db.CreateDatabase(db.DBConfig{
		Host: c.Host,
		Name: "dermacare",
	})

	teardown := func() {
		docker.StopContainer(t, c.ID)
		cleanup()
		w.Close()

		var buf bytes.Buffer
		io.Copy(&buf, r)
		os.Stdout = old

		fmt.Println("******************** LOGS ********************")
		fmt.Print(buf.String())
		fmt.Println("******************** LOGS ********************")
	}

	return db, teardown
}
