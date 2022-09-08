package main

import (
	"autodeploy/common"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func main() {
	godotenv.Load(".env")
	f, err := os.OpenFile("autodeploy.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	commit, err := common.GetCommit()
	if err != nil {
		log.Errorln(err)
		return
	}

	// get id from .id file
	data, err := os.ReadFile(".id")
	if err != nil {
		log.Errorln(err)
		return
	}
	if string(data) != commit.ID {
		bashFilePath := os.Getenv("BashFilePath")
		err = os.Chmod(bashFilePath, 0777)
		if err != nil {
			log.Errorln(err)
			return
		}
		_, err := exec.Command("bash", bashFilePath).Output()
		if err != nil {
			log.Errorln(err)
		}
		common.SendMessage(common.GetTextMessage(commit), "8")
		err = os.WriteFile(".id", []byte(commit.ID), 0644)
	}
}
