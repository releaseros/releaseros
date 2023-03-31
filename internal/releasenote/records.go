package releasenote

import (
	"fmt"
	"regexp"
	"sort"

	"releaseros/internal/context"

	"github.com/rs/zerolog"
)

type Record struct {
	CommitHash string
	Message    string
}

func (record Record) String() string {
	return fmt.Sprintf("%s %s", record.CommitHash, record.Message)
}

type Records []Record

func (records Records) Filter(ctx *context.Context) (Records, error) {
	for _, filter := range ctx.Config.Filters.Exclude {
		regexp, err := regexp.Compile(filter)
		if err != nil {
			return records, err
		}

		records = records.deleteMatchedMessage(regexp)
	}
	return records, nil
}

func (records Records) deleteMatchedMessage(filter *regexp.Regexp) Records {
	result := Records{}
	for _, record := range records {
		if !filter.MatchString(record.Message) {
			result = append(result, record)
		}
	}
	return result
}

func (records Records) Sort(ctx *context.Context) Records {
	direction := ctx.Config.Sort
	if direction == "" {
		return records
	}

	sort.Slice(records, func(i, j int) bool {
		if direction == "asc" {
			return records[i].Message < records[j].Message
		}
		return records[i].Message > records[j].Message
	})

	return records
}

func (records Records) MarshalZerologArray(a *zerolog.Array) {
	for _, record := range records {
		a.Str(record.String())
	}
}
