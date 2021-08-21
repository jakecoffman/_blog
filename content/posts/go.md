---
title: Convincing my team to switch to Go 
date: 2021-08-21T00:00:00-05:00 
tags: [story]
---

A story on how I convinced my team to switch to Go from Node.js.

<!--more-->

## Scenario

When I first joined the company, we were using Groovy on Grails. I actually was hired because I had experienced in this
niche language and framework.

After a few years though, the writing was on the wall that this framework was not going to be supported for much longer.
So the greater organization decided to put it to a vote on what language to use next, and they chose Node.js.

At the time I had been using Go in my spare time. I was apprehensive to advocate for it because it had no generics, and
I was worried it would be too difficult for people to learn. I had created a few projects at work where Go fit
perfectly, mostly proxying requests.

## Task

I was asked to join a smaller team to work on new features around the [Apache Guacamole](../guacamole-client-go)
project. I knew that Go would be a good fit for the project since it's a streaming service.

## Action

I proposed to the manager of the new team that we begin using Go for this project and for other new projects rather than
Node.js. I had a myriad of reasons why Go is superior to Node.js for the work we do.

Go is a simpler language than JavaScript. After having used Node.js for a while, I learned my previous thoughts that Go
is harder to learn were wrong, even without generics. Mostly this is because of the asynchronous nature of Node.js. I've
seen bug after bug related to a missing `await`, or confusion around Promises or callbacks.

Go is higher performance than Node.js out-of-the-box without doing any extra work. We sometimes need a high performance
proxy that makes decisions about small packets of data. Node.js is a nonstarter in those situations since we'd have to
run so many more servers to meet demand.

Go often doesn't require any external dependencies. The standard library is so robust, often there are no dependencies
required. Go also has a culture of resisting dependencies. With Node.js, programmers have been taught to grab
dependencies for even the simplest things. I've experienced a number of bugs around broken dependencies due to a myriad
of reasons:

- The author of the dependency introduced a breaking change and didn't bump the major version.
- The dependency was abandoned, and we have to move up to a newer unsupported version of Node.js, so we have to fork it
  to fix it.
- Developers didn't use a package-lock, causing confusion around why a new build with no changes is breaking.
- The dependency is not robust or well tested and used by very few people, so we find bugs when it gets unexpected
  inputs.

Go is statically typed, which leads to a better developer experience (DX). Our editors know more about what we're doing
and can make better suggestions. The compiler uses types to catch errors earlier. It removes the possibility that our
function that we expect a string argument receives an integer. Of course with Node.js we can use TypeScript which
introduces types, but this leads to even more problems:

- TypeScript requires a build process to transpile into JavaScript, where with JavaScript directly we could just run it.
  This build step can be complex and confusing to set up for developers that aren't experienced with it. I myself have
  tried this and often get stuck.
- Since TypeScript is transpiled, it's machine generated JavaScript that's running in production. So when something goes
  wrong you may have to debug code that is messier and harder to read.
- TypeScript the language is getting very complicated. The authors are constantly adding new features, and the language
  also supports all JavaScript features. This is a recipe for madness! We can't fix a language by continuously adding to
  it. In a lot of ways, TypeScript is committing the same errors that C++ did.
- I often see developers just use `any` which completely defeats the purpose.

### Drawbacks

Being a pragmatic programmer I also have openly acknowledged the drawbacks of Go.

The largest drawback to me is Go lacks generics (although some are coming soon!) and this has had a big impact on the usability of the language. In
JavaScript, we can do:

```javascript
const items = [{id: 1, name: 'one'}, {id: 2, name: 'two'}]
// find by ID (number)
const result1 = items.find(item => item.id === 1)
// find by name (string)
const result2 = items.find(item => item.name === 'one')
console.assert(result1 === result2) // true
```

In Go, there is no built-in `Find`, so we'd have to write a function like this:

```go
package main_test

import (
	"reflect"
	"testing"
)

type Item struct {
	ID   int
	Name string
}

func TestFindItem(t *testing.T) {
	items := []Item{
		{1, "one"},
		{2, "two"},
	}
	result1 := FindItem(items, func(item *Item) bool {
		return item.ID == 1
	})
	result2 := FindItem(items, func(item *Item) bool {
		return item.Name == "one"
	})
	if !reflect.DeepEqual(result1, result2) {
		t.Error("Unexpected results")
	}
}

// FindItem returns the item in the array that results from finder returning true.
func FindItem(items []Item, finder func(item *Item) bool) *Item {
	for i := range items {
		if finder(&items[i]) {
			return &items[i]
		}
	}
	return nil
}
```

While this is not complex to do, we have to write a Find function for each type. This adds to the verbosity of the language, and is just frustrating when most other languages have these features built-in. They even have a page called [Slice Tricks](https://github.com/golang/go/wiki/SliceTricks) which is a huge smell that they missed something in the design of the language.

To me this has been the largest drawback to the language. With the upcoming Go 1.18 release, we should get generics and a generic package to handle most of these problems with slices and other things.

Some other smaller drawbacks to the language:
- The standard library's http Handlers are missing a couple of critical features: url parameters and chained middleware. So when trying to use the builtin http package I often end up writing a small framework around it. So instead of doing that, I often reach for a router dependency which has nicer features, which ends up pulling in more dependencies, which is not ideal.
- Some developers don't want to learn anything new, and Go is not as widely known as JavaScript.
- Error handling is very simple, but verbose. It's something you have to come to love, but it leads to a lack of stack traces which makes debugging harder.
- Concurrency primitives are great, but there are no higher level concurrency libraries due to lack of generics. After generics lands expect to see some kind of Streams library.

## Results

The other senior developer on the team who was pushing TypeScript immediately understood the value of Go and was behind my efforts 100%.

A self-taught programmer who had been doing Node.js for a while try Go, and he picked it up quickly and misses it when he is back on a Node.js project.

A more seasoned team lead tried it and is starting all his new projects in Go.

I rewrote the Apache Guacamole client in Go and have no regrets. I have a number of other projects launched, and I am finding good patterns for newer developers to use.

We currently have 40 Go repos in our GitHub.

Once generics lands I want to approach the larger team and offer this information as a case study!
