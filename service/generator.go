package service

import (
	"bytes"
	"log"
	"os/exec"
)

func (s *service) GenerateModel() error {
	e := exec.Command("make", "generate")
	var out bytes.Buffer
	e.Stderr = &out
	err := e.Run()
	if err != nil {
		log.Println(out.String())
		log.Println("err->", err)
		return err
	}
	// fmt.Printf("Output: %q", out.String())
	return nil
}