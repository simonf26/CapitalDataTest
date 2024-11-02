package csv

import "fmt"

type Stat struct {
	CSVLines              int
	InvalidRecords        int
	UnsupportedDateLayout int
	InvalidFirstname      int
	InvalidLastname       int
	InvalidEmail          int
	JSONLines             int
}

// newStat creates a new Stat object.
func newStat() *Stat {
	return &Stat{
		CSVLines:       0,
		InvalidRecords: 0,
		JSONLines:      0,
	}
}

var stats *Stat

// Print prints out the stat's information.
func (stat *Stat) Print() {
	fmt.Println("Stats:")
	fmt.Printf("	CSV entries: %d\n", stat.CSVLines)
	fmt.Printf("	Invalid records: %d\n", stat.InvalidRecords)
	fmt.Printf("		Invalid firstname: %d\n", stat.InvalidFirstname)
	fmt.Printf("		Invalid lastname: %d\n", stat.InvalidLastname)
	fmt.Printf("		Invalid email: %d\n", stat.InvalidEmail)
	fmt.Printf("		Unsupported date layout: %d\n", stat.UnsupportedDateLayout)
	fmt.Printf("	JSON entries: %d\n", stat.JSONLines)
}
