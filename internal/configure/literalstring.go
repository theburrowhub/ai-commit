package configure

import (
	"gopkg.in/yaml.v3"
)

// LiteralString is a custom type to represent a string that should be serialized as a literal block.
type LiteralString string

// MarshalYAML implements the yaml.Marshaler interface for LiteralString.
func (s LiteralString) MarshalYAML() (interface{}, error) {
	node := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!str",
		Value: string(s),
		Style: yaml.LiteralStyle, // Uses the pipe style
	}
	return node, nil
}
