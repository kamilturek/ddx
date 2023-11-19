package ddb

import (
	"encoding/json"
	"fmt"
	"strings"
)

type KeySchemaFormatter func(keySchemas []KeySchema) (*string, error)

func ToText(keySchemas []KeySchema) (*string, error) {
	var sb strings.Builder

	for _, ks := range keySchemas {
		sb.WriteString(fmt.Sprintf("Table Name:\t%s\n", *ks.TableName))
		sb.WriteString(fmt.Sprintf("Table Type:\t%s\n", ks.TableType))
		sb.WriteString(fmt.Sprintf("Parition Key:\t%s\n", *ks.PartitionKey))

		if ks.SortKey != nil {
			sb.WriteString(fmt.Sprintf("Sort Key:\t%s\n", *ks.SortKey))
		}

		sb.WriteString("\n")
	}

	out := sb.String()

	return &out, nil
}

func ToJSON(keySchemas []KeySchema) (*string, error) {
	out, err := json.MarshalIndent(keySchemas, "", "  ")
	if err != nil {
		return nil, err
	}

	outStr := string(out)

	return &outStr, nil
}
