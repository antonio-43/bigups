package models

type Results struct {
	Tested []Mapper
}

type Mapper struct {
	Target   string
	DnsIp    string
	TargetIp []string
	Resolved bool
}
