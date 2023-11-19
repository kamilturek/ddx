package cmd

import (
	"fmt"

	"github.com/kamilturek/ddx/ddb"
	"github.com/spf13/cobra"
)

var tableName string
var format string
var limit int

var formatters = map[string]ddb.KeySchemaFormatter{
	"json": ddb.ToJSON,
	"text": ddb.ToText,
}

var rootCmd = &cobra.Command{
	Use:   "ddx",
	Short: "Show DynamoDB table key schema",
	Long:  "A tool for quickly viewing your DynamoDB table and index key schemas",
	Args: func(cmd *cobra.Command, args []string) error {
		if _, ok := formatters[format]; !ok {
			return fmt.Errorf("invalid format specified: %s", format)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var tableNames []string

		if tableName != "" {
			tableNames = []string{tableName}
		} else {
			names, err := ddb.ListTables()
			tableNames = names
			if err != nil {
				panic(err)
			}
		}

		var keySchemas []ddb.KeySchema

		for i, tableName := range tableNames {
			if limit >= 0 && limit == i {
				break
			}

			result, err := ddb.GetKeySchemas(tableName)
			if err != nil {
				panic(err)
			}

			keySchemas = append(keySchemas, result...)
		}

		format := formatters[format]

		out, err := format(keySchemas)
		if err != nil {
			panic(err)
		}

		fmt.Println(*out)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&tableName, "table-name", "t", "", "table name whose key schema should be shown (all if not specified)")
	rootCmd.Flags().StringVarP(&format, "format", "f", "text", "output format, available: \"text\", \"json\"")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", -1, "maximum number of tables listed")
}

func Execute() {
	rootCmd.Execute()
}
