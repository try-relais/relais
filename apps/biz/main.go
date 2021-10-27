// Package cmd contains an entrypoint for running an ion-sfu instance.
package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/pion/ion-log"
	biz "github.com/pion/ion/apps/biz/server"
	"github.com/spf13/viper"
)

var (
	conf = biz.Config{}
	file string
)

func showHelp() {
	fmt.Printf("Usage:%s {params}\n", os.Args[0])
	fmt.Println("      -c {config file}")
	fmt.Println("      -h (show help info)")
}

func unmarshal(rawVal interface{}) bool {
	if err := viper.Unmarshal(rawVal); err != nil {
		fmt.Printf("config file %s loaded failed. %v\n", file, err)
		return false
	}
	return true
}

func load() bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}

	viper.SetConfigFile(file)
	viper.SetConfigType("toml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file %s read failed. %v\n", file, err)
		return false
	}

	if !unmarshal(&conf) {
		return false
	}
	if err != nil {
		fmt.Printf("config file %s loaded failed. %v\n", file, err)
		return false
	}

	fmt.Printf("config %s load ok!\n", file)

	return true
}

func parse() bool {
	flag.StringVar(&file, "c", "configs/biz.toml", "config file")
	help := flag.Bool("h", false, "help info")
	flag.Parse()
	if !load() {
		return false
	}

	if *help {
		showHelp()
		return false
	}
	return true
}

func main() {
	if !parse() {
		showHelp()
		os.Exit(-1)
	}
	log.Init(conf.Log.Level)
	log.Infof("--- Starting Biz Node ---")
	node := biz.NewBIZ(conf.Node.NID)
	if err := node.Start(conf); err != nil {
		log.Errorf("biz init start: %v", err)
		os.Exit(-1)
	}
	defer node.Close()
	select {}
}
