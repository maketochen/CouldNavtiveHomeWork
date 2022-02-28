package main

import "os"

type Env struct {
	Version string
}

/*
获取环境变量中的Version
*/
func getEnv() *Env {
	os.Setenv("Version","0.0.1")
	v := os.Getenv("Version")
	return &Env{Version: v}
}
