// Create a aws rds aurora cli used to skip many replication errors pass in the endpint.
package main

import (
	"flag"
	"fmt"
	"os/exec"
)

// Set the default values for the flags.
var dbEndPoint string

func execute() {
	//fmt.Printf("dbEndPoint = %v\n", dbEndPoint)
	bashCommand := fmt.Sprintf(`mysql -h %v -t -e "show slave status\G"`, dbEndPoint)
	//fmt.Printf("Exicuting command:\n%v\n\n", bashCommand)
	out, err := exec.Command("bash", "-c", bashCommand).Output()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("Command Successfully Executed")
	fmt.Println(string(out))
}

func fixit() {
	//fmt.Printf("dbEndPoint = %v\n", dbEndPoint)
	bashCommand2 := fmt.Sprintf(`mysql -h %v -t -e "CALL mysql.rds_skip_repl_error()"`, dbEndPoint)
	out, err := exec.Command("bash", "-c", bashCommand2).Output()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("Command Successfully Executed")
	fmt.Println(string(out))
}

// main function
func main() {
	flag.StringVar(&dbEndPoint, "h", "", "Specify the AWS EndPoint.")
	flag.Parse()

	if dbEndPoint != "" {
		execute()
		return
	}
	flag.PrintDefaults()
	return
	// if show slave status has Duplicate entry error then loop calling until CALL mysql.rds_skip_repl_error() until no error.
	input := "Duplicate entry"
	
	if  strings.Contains(execute(), input) {
		fmt.Printf("Duplicate entry error found.\n")
		for strings.Contains(execute(), input) {
			fmt.Printf("Duplicate entry error found.\n")
			fixit()
		}else {
			fmt.Printf("Duplicate entry error not found.\n")
			// Do I need a break to get out of the loop?
			
		}	
	}
}
