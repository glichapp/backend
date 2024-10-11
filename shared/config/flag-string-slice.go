package config

import (
	"flag"
	"fmt"
)

// FlagStringSlice implement flag.Value and implies the use for getting array of flag values.
type FlagStringSlice []string

var _ flag.Value = (*FlagStringSlice)(nil)

func (ss *FlagStringSlice) String() string {
	return fmt.Sprintf("%v", *ss)
}

func (ss *FlagStringSlice) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}
