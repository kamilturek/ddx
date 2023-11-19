package ddb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TableType string

const (
	TableTypeBase TableType = "BASE"
	TableTypeGSI  TableType = "GSI"
	TableTypeLSI  TableType = "LSI"
)

type KeySchema struct {
	TableName    *string
	TableType    TableType
	PartitionKey *string
	SortKey      *string
}

type TableReader interface {
	GetKeySchema(tableName string) KeySchema
	ListTables()
}

type DDBTableReader struct {
	client *dynamodb.Client
}

func NewDDBTableReader(ctx context.Context) (*DDBTableReader, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &DDBTableReader{client: dynamodb.NewFromConfig(cfg)}, nil
}

func (ddb *DDBTableReader) GetKeySchema(ctx context.Context, tableName string) ([]KeySchema, error) {
	in := &dynamodb.DescribeTableInput{TableName: &tableName}
	out, err := ddb.client.DescribeTable(ctx, in)
	if err != nil {
		return nil, err
	}

	partitionKey, sortKey := getKeys(out.Table.KeySchema)
	keySchemas := []KeySchema{
		{
			TableName:    out.Table.TableName,
			TableType:    TableTypeBase,
			PartitionKey: partitionKey,
			SortKey:      sortKey,
		},
	}

	for _, lsi := range out.Table.LocalSecondaryIndexes {
		partitionKey, sortKey = getKeys(lsi.KeySchema)
		keySchemas = append(keySchemas, KeySchema{
			TableName:    lsi.IndexName,
			TableType:    TableTypeGSI,
			PartitionKey: partitionKey,
			SortKey:      sortKey,
		})
	}

	for _, gsi := range out.Table.GlobalSecondaryIndexes {
		partitionKey, sortKey = getKeys(gsi.KeySchema)
		keySchemas = append(keySchemas, KeySchema{
			TableName:    gsi.IndexName,
			TableType:    TableTypeGSI,
			PartitionKey: partitionKey,
			SortKey:      sortKey,
		})
	}

	return keySchemas, nil
}

func getKeys(ks []types.KeySchemaElement) (*string, *string) {
	var partitionKey *string
	var sortKey *string

	for _, v := range ks {
		if v.KeyType == types.KeyTypeHash {
			partitionKey = v.AttributeName
		} else if v.KeyType == types.KeyTypeRange {
			sortKey = v.AttributeName
		}
	}

	return partitionKey, sortKey
}

func (ddb *DDBTableReader) ListTables(ctx context.Context) ([]string, error) {
	in := &dynamodb.ListTablesInput{}
	out, err := ddb.client.ListTables(ctx, in)
	if err != nil {
		return nil, err
	}

	return out.TableNames, nil
}

func GetKeySchemas(tableName string) ([]KeySchema, error) {
	ctx := context.Background()

	reader, err := NewDDBTableReader(ctx)
	if err != nil {
		return nil, err
	}

	return reader.GetKeySchema(ctx, tableName)
}

func ListTables() ([]string, error) {
	ctx := context.Background()

	reader, err := NewDDBTableReader(ctx)
	if err != nil {
		return nil, err
	}

	return reader.ListTables(ctx)
}
