package main

/*
 * IF sAMMAccountName matches Uid it is always the correct account
 */

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	. "./LDAPMethods"
	"time"
	"os"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func main() {
	log.Printf("LDAPSync is starting\r\n")
	var masterSettings string
	var slaveSettings string
	initializeViper()
	// logging init
	fmt.Println("logfile:", viper.GetString("logfile"))

	DebugMain("LDAPSync is starting", nil, 26)
	if len(os.Args) == 3 {
		masterSettings = os.Args[1]
		slaveSettings = os.Args[2]
		DebugMain(fmt.Sprintf("Set the master setting name to %s and the slave setting name to %s", masterSettings, slaveSettings), nil, 30)
	} else {
		fmt.Println("There was an error that caused the program to exit, check the logs for more information")
		DebugMain(fmt.Sprintf("program got: %d argument(s) want 2", len(os.Args)-1), nil, 33)
		return
	}
	Master.Initialize(masterSettings)
	Slave.Initialize(slaveSettings)
	CNGroups.Initialize()
	cycle := 0
	for {
		if cycle == 10 {
			cycle = 0
			fmt.Println("Cleaning slave")
			Slave.CleanSlave()
			DebugMain(fmt.Sprintf("reset the cycle counter to 0"), nil, 45)
		}
		cycle ++
		if cycle == 1 {
			Slave.ClearStruct()
			Slave.SearchLDAP()
			Slave.SortLdap()
			DebugMain(fmt.Sprintf("Refreshed the slave info"), nil, 52)
		}
		Master.ClearStruct()
		Master.SearchLDAP()
		Master.SortAd()
		Master.Compare()
		if Search == true {
			Slave.ClearStruct()
			Slave.SearchLDAP()
			Slave.SortLdap()
			Search = false
			DebugMain(fmt.Sprintf("The program made changes to the LDAP and refreshed the slave search"), nil, 63)
		}
		fmt.Println("cycle", cycle, "finished")
		DebugMain(fmt.Sprintf("The program finished a cycle %d/%d", cycle, 10), nil, 66)
		time.Sleep(1 * time.Hour)
	}
}

func initializeViper() {
	DebugMain(fmt.Sprintf("Initializing viper settings"), nil, 72)
	viper.SetConfigName("config")         // name of config.yaml file (without extension)
	viper.AddConfigPath("/etc/appname/")  // path to look for the config.yaml file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many searchLDAP paths
	viper.AddConfigPath(".")              // optionally look for config.yaml in the working directory
	err := viper.ReadInConfig() // Find and read the config.yaml file
	if err != nil {
		DebugMain(fmt.Sprintf("There was a problem with the config file"), err, 79)
		return
	}
	// setting debug parameters (true, false)
	DebugMainState = viper.GetBool("debugMain")
	DebugCompareState = viper.GetBool("debugCompare")
	DebugModifyState = viper.GetBool("debugModify")
	DebugDeleteState = viper.GetBool("debugDelete")
	DebugInitialState = viper.GetBool("debugInitial")
	DebugAddState = viper.GetBool("debugAdd")
	// set log file
	logfile := viper.GetString("logfile")
	log.SetOutput(&lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    5, // megabytes
		MaxBackups: 1,
		MaxAge:     28,    //days
		Compress:   false, // disabled by default
	})
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		DebugMain(fmt.Sprintf("config file changed: %s", e.Name), nil, 100)
		fmt.Println("Config file changed:", e.Name)
	})
	DebugMain(fmt.Sprintf("Done initializing viper"), nil, 103)
}
