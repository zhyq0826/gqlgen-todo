package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/zhyq0826/gqlgen-todo/graph"
	"github.com/zhyq0826/gqlgen-todo/graph/generated"
)

func test1(){
	query := `mutation ($input: ProgramingLanguageInput!) {
		addProgramingLanguage(input: $input)
	}`
	variables := map[string]interface{}{
		"input": map[string]interface{}{
			"appID":                 "1",
			"programingLanguageIDs": []int{1},
		},
	}

	exec(query, variables)
}

func test2(){
	query := `mutation{
	  addProgramingLanguage(input: {
		appID: "1",
		programingLanguageIDs: [1]
	  })
	}`
	exec(query, nil)
}

func exec(query string, variables map[string]interface{}){
	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	myExecutor := executor.New(schema)
	ctx := context.Background()
	ctx = graphql.StartOperationTrace(ctx)

	optCtx, err := myExecutor.CreateOperationContext(ctx, &graphql.RawParams{
		Query:     query,
		Variables: variables,
	})

	if err != nil {
		fmt.Println(" create request error ===>", err)
	}
	respHandle, ctx := myExecutor.DispatchOperation(ctx, optCtx)
	resp := respHandle(ctx)
	result, err1 := resp.Data.MarshalJSON()
	fmt.Println("result = >", string(result))
	fmt.Println("err =>", err1)
}


func main() {
	test1()
	test2()
}
