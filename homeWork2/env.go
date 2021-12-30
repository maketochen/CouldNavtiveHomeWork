package main

import "os"

type Env struct {
	Version string
}

func getEnv() *Env {
	v := os.Getenv("Version")
	return &Env{Version: v}
}
