package draft07

import parseSchema "github.com/asyncapi/parser/pkg/schema"

var parser = parseSchema.NewParser(schema)

func Parse(v interface{}) error {
	return parser.Parse(v)
}
