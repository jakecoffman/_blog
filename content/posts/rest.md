---
title: Go REST API
date: 2029-12-02T12:00:00-05:00
tags: [programming, api]
---

I [developed a library called crud](https://github.com/jakecoffman/crud) for building out a Swagger spec which also adds a middleware that validates the inputs against the spec. This is nice because whatever you write in your Swagger is automatically enforced and it cuts down on all the boilerplate.

Here are examples for a [SQL approach](https://github.com/jakecoffman/go-rest) and a [NoSQL approach](https://github.com/jakecoffman/go-mongo-rest)

<!--more-->

## example

Here's an example spec:

{{< gist jakecoffman df3cf0607284a2e21974cca8ec12b793 >}}

In addition to building a nice Swagger UI, the validation enforces:
- The `id` in the path is an integer
- The `born` field is a Date format
- The `books` field is an Array containing a `title` and `genre`
- The `genre` field value must exist in an enum

That means all of that above doesn't need to be validated in the handler!

Validations are available for the url query, path parameters, the body, and headers. The Swagger that is generated is exported as a struct, so you can also make modifications that aren't supported by the library. 

The library also has options to error when unexpected/unknown fields are encountered or just strip the unknown fields. So if the user typos a query parameter it can error, or if they submit some secret field in the body you can strip it.

The library re-serializes the body and puts it back into the Request.Body so your handlers look the same even without the validation middleware. This means if you decide to adopt this library or move away from it, you don't have to change your handler.

## library vs framework

I took a "library" approach to developing this rather than a "framework," so in order to use it you will need to grab an [adaptor](https://github.com/jakecoffman/crud/tree/master/adapters) for a router or write your own. 

Unfortunately the built-in `http.ServeMux` can't be supported since it requires method routing and path parameters. Doing `/authors/{authorId}/books/{bookId}` is just not possible without creating a router. I wanted this library to have a focus on doing the Swagger and validation. Implementing a router would bring this into the framework territory.

There's [another project](https://github.com/emicklei/go-restful) that is similar (I didn't know it existed when I started mine!) but it takes more of a framework approach since they implement their own router. If you prefer this then go for it!
