package ram

type AssumedRoleUser struct {
	AssumedRoleId string
	Arn           string
}

type AssumedRoleUserCredentials struct {
	AccessKeySecret string
	AccessKeyId     string
	Expiration      string
	SecurityToken   string
}

type AssumeRoleResponse struct {
	RequestId       string
	AssumedRoleUser AssumedRoleUser
	Credentials     AssumedRoleUserCredentials
}
