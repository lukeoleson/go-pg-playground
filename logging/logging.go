package logging

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-pg/pg/v10"
)

type QueryLogger struct{}

func (logger QueryLogger) BeforeQuery(ctx context.Context, queryEvent *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (logger QueryLogger) AfterQuery(ctx context.Context, queryEvent *pg.QueryEvent) error {
	formattedQuery, err := queryEvent.FormattedQuery()
	if err != nil {
		fmt.Printf("error printing query: %v", err)
	}

	prettyQuery := prettifyQuery(formattedQuery)

	fmt.Printf("\nGenerated Query\n---------------")
	fmt.Printf("%s\n", prettyQuery)
	return nil
}

func prettifyQuery(formattedQuery []byte) string {
	queryString := fmt.Sprintf("%s", formattedQuery)

	queryString = strings.Replace(queryString, "\"", "", -1)

	requireLineBreaks := []string{"FROM", "SET", "WHERE", "RETURNING", "LIMIT", "GROUP BY", "ORDER BY", "SELECT", "UPDATE"}
	for _, directive := range requireLineBreaks {
		queryString = strings.Replace(queryString, directive, "\n"+directive, -1)
	}

	return queryString
}
