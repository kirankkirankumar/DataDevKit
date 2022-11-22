package service

import (
	"bytes"
	"log"
	"os/exec"
)

func (s *service) GenerateModel() error {

	log.Println("Generating models")

	e := exec.Command("make", "generate")
	var errs bytes.Buffer
	var out bytes.Buffer
	e.Stderr = &errs
	e.Stdout = &out
	err := e.Run()
	if err != nil {
		log.Println(errs.String())
		log.Println("err->", err)
		return err
	}
	log.Println(out.String())

	log.Println("Generating complete")

	return nil
}