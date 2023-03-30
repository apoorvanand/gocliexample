/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// webserverCmd represents the webserver command
var webserverCmd = &cobra.Command{
	Use:   "webserver",
	Short: "Start a webserver on your system",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webserver Started at Port 8000")
		webservercmd :=exec.Command("python3","-m" ,"http.server")
		openfirefoxcmd := exec.Command("firefox","--new-tab","--url","http://localhost:8000")
		stderr,_ := webservercmd.StderrPipe()
		_= webservercmd.Start()
		_ = openfirefoxcmd.Run()
		_= webservercmd.Wait()
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
		}
		if stderr !=nil {
			log.Fatal(stderr)
		}

	},
}

func init() {
	rootCmd.AddCommand(webserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	webserverCmd.Flags().IntP("port", "p", 8000, "Help message for toggle")
}
