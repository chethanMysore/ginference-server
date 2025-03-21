package prodconfig

import (
	"fmt"
	"net/url"
	"strconv"
)

const DBDomainURI = "mongodb://localhost:27017"
const WithTLS = true
const TLSCAFilePath = "E:\\softwares\\mongoDB\\certs\\server\\ca.pem"
const TLSCertificateKeyFilePath = "E:\\softwares\\mongoDB\\certs\\client\\client.pem"
const MongoDBURI = "mongodb://localhost:27017/?tls=true&tlsCAFile=E%3A%5Csoftwares%5CmongoDB%5Ccerts%5Cserver%5Cca.pem&tlsCertificateKeyFile=E%3A%5Csoftwares%5CmongoDB%5Ccerts%5Cclient%5Cclient.pem"
const DBName = "sentixDB"
const UserCollection = "users"
const ModelCollection = "models"
const APIDomain = "localhost"
const APIPort = "8080"

var DBConnectionStringWithTLS = fmt.Sprintf("%s/?tls=%s&tlsCAFile=%s&tlsCertificateKeyFile=%s", DBDomainURI, strconv.FormatBool(true), url.QueryEscape(TLSCAFilePath), url.QueryEscape(TLSCertificateKeyFilePath))
var DBConnectionString = fmt.Sprintf("%s/?tls=%s", DBDomainURI, strconv.FormatBool(false))
var APIDomainURI = fmt.Sprintf("%s:%s", APIDomain, APIPort)
