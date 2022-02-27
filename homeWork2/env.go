package main

import "os"

type Env struct {
	Version string
}

/*
获取环境变量中的Version
*/
func getEnv() *Env {
	v := os.Getenv("VERSION")
	return &Env{Version: v}
}
