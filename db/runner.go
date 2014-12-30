package db

import (
	"fmt"
	"github.com/elos/server/util"
	"io/ioutil"
	"os/exec"
)

func StartMongo() error {
	out, err := exec.Command("mongod", "--config", "./mongo.conf").Output()

	if err != nil {
		fmt.Printf("%s", out)
		return err
	}

	util.Log("Mongo succesfully started")
	return nil
}

func StopMongo() error {
	bytes, err := ioutil.ReadFile("/tmp/mongodb.pid")
	if err != nil {
		return err
	}

	pid := string(bytes)

	cmd := exec.Command("kill", pid)

	err = cmd.Run()

	if err != nil {
		return err
	}

	util.Log("Mongo succesfully stopped")
	return nil
}
