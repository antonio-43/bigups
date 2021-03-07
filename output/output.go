package output

import (
	"bgs/models"
	"fmt"
)

func LogIt(r models.Results) {
	for _, pen := range r.Tested {
		fmt.Printf("%s ::: %s ::: %t ::: %s\n", pen.Target, pen.DnsIp, pen.Resolved, pen.TargetIp)
	}
}

func LogClean(r models.Results) {
	for _, ipList := range r.Tested {
		for _, ip := range ipList.TargetIp {
			fmt.Println(ip)
		}
	}
}
