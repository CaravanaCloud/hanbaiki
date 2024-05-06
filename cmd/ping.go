package cmd

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/log"

	"golang.org/x/sync/errgroup"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/spf13/cobra"
)

func pingAWS() error {
	// Create a new AWS session
	sess, err := session.NewSession()
	if err != nil {
		return err
	}

	// Create an STS client
	svc := sts.New(sess)

	// Call the GetCallerIdentity operation to get details about the IAM identity
	_, err = svc.GetCallerIdentity(nil)
	if err != nil {
		log.Printf("Error calling GetCallerIdentity: %s", err)
		return err
	}

	log.Info("AWS GetCallerIdentity success.")
	return nil
}

func pingDB(db *sql.DB) error {
	log.Info("Database connection success.")
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging the database: %v", err)
	}

	var health string
	err = db.QueryRow("SELECT 'OK'").Scan(&health)
	if err != nil {
		return fmt.Errorf("error executing health check query: %v", err)
	}
	log.Info("Database health check:", health)

	return nil
}

func doPing(cmd *cobra.Command, args []string) {
	g, _ := errgroup.WithContext(cmd.Context())
	g.Go(func() error {
		return DBQuery(pingDB)
	})

	g.Go(func() error {
		return pingAWS()
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("Health check failed: %v", err)
	}

	log.Info("Health check success.")

}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A simple health check command",
	Long:  `Checks if database and dependencies are OK`,
	Run:   doPing,
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
