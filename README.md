# gqlgen-todo
gqlgen-todo


```graphql
mutation {
  addProgramingLanguage(input: {appID: "1", programingLanguageIDs: [1]})
}
```

```graphql
mutation ($input: ProgramingLanguageInput!){
  addProgramingLanguage(input: $input)
}

```

```json
{
  "input": {
    "appID": "1",
    "programingLanguageIDs": [1]
  }
}
```