package connection

var (
	//ConnectionString is using to connect MongoDB instance
	ConnectionString string
	//DBName is using to determine which DB is used to store images
	DBName string
	//MaxPoolSize is using to connect MongoDB instance
	MaxPoolSize uint64
)

var (
	//AzureURI is using to connect Azure Blob Storage instance
	AzureURI string
	//AccountName is using to connect Azure Blob Storage instance
	AccountName string
	//AccountKey is using to connect Azure Blob Storage instance
	AccountKey string
)
