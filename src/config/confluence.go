package config

import (
	"confluence-poc/src/confluence"
	"confluence-poc/src/models"
)

func NewConfluenceClient(flags models.Flags) (*confluence.API, error) {
	return confluence.NewAPI("https://"+flags.Domain+".atlassian.net/wiki/rest/api", flags.Username, flags.ApiToken)
}
