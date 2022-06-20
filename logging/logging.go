package logging

import (
	"context"
	"fmt"
	"strings"

	pgV9 "github.com/go-pg/pg/v9"

	pgV10 "github.com/go-pg/pg/v10"
)

type QueryLoggerV10 struct{}

func (logger QueryLoggerV10) BeforeQuery(ctx context.Context, queryEvent *pgV10.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (logger QueryLoggerV10) AfterQuery(ctx context.Context, queryEvent *pgV10.QueryEvent) error {
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

type QueryLoggerV9 struct{}

func (logger QueryLoggerV9) BeforeQuery(ctx context.Context, queryEvent *pgV9.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (logger QueryLoggerV9) AfterQuery(ctx context.Context, queryEvent *pgV9.QueryEvent) error {
	formattedQuery, err := queryEvent.FormattedQuery()
	if err != nil {
		fmt.Printf("error printing query: %v", err)
	}

	prettyQuery := prettifyQuery([]byte(formattedQuery))

	fmt.Printf("\nGenerated Query\n---------------")
	fmt.Printf("%s\n", prettyQuery)
	return nil
}
