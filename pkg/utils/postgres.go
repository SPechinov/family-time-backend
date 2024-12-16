package utils

func GetPostgresConnectionString(user, password, host, port, dbName string, sslMode bool) string {
	connectionString := "postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + dbName
	if !sslMode {
		connectionString += "?sslmode=disable"
	}
	return connectionString
}
