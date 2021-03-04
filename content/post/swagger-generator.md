---
title: Go swagger generator
date: 2021-03-04T12:00:00-05:00
tags: [api]
---

I built a Go library that can generate Swagger and validation middleware automatically from Go code.

It's called CRUD and you can get it here: [https://github.com/jakecoffman/crud](https://github.com/jakecoffman/crud)

<!--more-->

## Getting started

Start by getting the package `go get github.com/jakecoffman/crud`

Then in your `main.go`:

1. Create a router with `NewRouter`
2. Add routes with `Add`
3. Then call `Serve`

Routes are specifications that look like this:

```go
crud.Spec{
	Method:      "PATCH",
	Path:        "/widgets/{id}",
	PreHandlers: []gin.HandlerFunc{Auth},
	Handler:     CreateHandler,
	Description: "Adds a widget",
	Tags:        []string{"Widgets"},
	Validate: crud.Validate{
		Path: crud.Object(map[string]crud.Field{
			"id": crud.Number().Required().Description("ID of the widget"),
        }),
		Body: crud.Object(map[string]crud.Field{
			"owner": crud.String().Required().Example("Bob").Description("Widget owner's name"),
		}),
	},
}
```

This will add a route `/widgets/:id` that responds to the PATCH method. It generates swagger and serves it at the root of the web application. It validates that the ID in the path is a number, so you don't have to. It also validates that the body is an object and has an "owner" property that is a string, again so you won't have to.

It mounts the swagger-ui at `/` and loads up the generated swagger.json.
