package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

var RunOptions struct {
	url      string
	duration int
}

var options = RunOptions

var rootCmd = &cobra.Command{
	Use:   "./hairpin-test",
	Short: "tests the accessibility of the endpoint over a period of time",
	Run: func(cmd *cobra.Command, args []string) {
		milliRunDuration := options.duration * 1000
		endTimeMillis := time.Now().UnixMilli() + int64(milliRunDuration)

		pass := 0
		fail := 0
		for time.Now().UnixMilli() < endTimeMillis {
			if err := makeEndpointCall(); err != nil {
				fail++
			} else {
				pass++
			}
		}

		returnCode := 0
		if pass == 0 || fail > 0 {
			returnCode = 1
		}
		fmt.Printf("pass: %d\tfail: %d", pass, fail)
		os.Exit(returnCode)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Execute()
}

func makeEndpointCall() error {
	resp, err := http.Get(options.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&options.url, "url", "", "url of the endpoint to query")
	rootCmd.PersistentFlags().IntVar(&options.duration, "duration", 10, "duration of the test in seconds")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("duration", rootCmd.PersistentFlags().Lookup("duration"))
	rootCmd.MarkPersistentFlagRequired("url")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
