// Create a aws rds aurora cli used to skip many replication errors pass in the endpint, username, password, and db name.

import (
	"fmt"
	"os"
	"strings"
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

// Set the flags.
var (
	endpoint = flag.String("endpoint", "", "endpoint")
	//region 	 =	flag.String("endpoint", "", "endpoint"
)

// Check the AWS RDS AURORA replication status and skip errors until it is in a good state.
func main() {
	flag.Parse()
	if *endpoint == "" {
		fmt.Println("endpoint is required")
		os.Exit(1)
	}

	// Create a new session with a config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a new RDS client
	svc := rds.New(sess)

	// Get the replication status.
	result, err := svc.DescribeDBClusterEndpoints(&rds.DescribeDBClusterEndpointsInput{
		DBClusterEndpointIdentifier: aws.String(*endpoint),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check the replication status and skip errors if on Duplicate entry.
	for _, status := range result.DBClusterEndpoints[0].StatusInfos {
		if strings.Contains(*status.Message, "Duplicate entry") {
			fmt.Println("Duplicate entry found, skipping")
			os.Exit(0)
		}
	}
}