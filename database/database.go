package database

var connnection string

func init() {
	connnection = "MySQL"
}

func GetDatabase() string {
	return connnection
}
