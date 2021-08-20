---
title: Unit Testing
date: 2020-12-16T00:00:00-05:00
tags: [testing]
---

A unit test is a test that tests a single function with all external dependencies mocked or faked.

Unit testing can either be valuable or be a detriment to the code base, depending on what is being tested.

<!--more-->

## The Good Parts

Unit testing is valuable when the unit under test has few side effects and actually calculates a value. It can also be valuable to test difficult-to-recreate error scenarios like filesystem errors.

Cross is a great candidate for unit testing since it is pure:

```go
func (v Vector) Cross(other Vector) float64 {
	return v.X*other.Y - v.Y*other.X
}
```

This controller logic should be unit tested to ensure the error cases will always be handled. An integration test would work too but since there are no dependencies a unit test is preferable.

```javascript
async save(request, reply) {
  if (!request.payload.dataId && !request.payload.reservationId) {
    return Boom.badRequest(`Must be launched with one of: dataId, reservationId`)
  }
  if (request.payload.endDate && request.payload.endDate < new Date()) {
    return Boom.badRequest(`Already ended`)
  }
  if (!request.internal && request.payload.workId) {
    return Boom.badRequest(`External users may not provide workId`)
  }
  // ...
}
```

## Easier to test, harder to use

What happens when you write a test for this function?

```go
func (v Vector) Lerp(other Vector, t float64) Vector {
	return v.Mult(1.0 - t).Add(other.Mult(t))
}
```

Strictly speaking you've written an integration test, not a unit test!

So a test purist would say you need to introduce dependency injection to mock out the other units:

```go
func (v Vector) Lerp(other Vector, t float64, mult func(Vector, float64) Vector, add func(Vector, Vector) Vector) Vector {
    return add(mult(v, 1.0 - t), mult(other, t))
}
```

Or maybe some other similar scheme...

**PLEASE DON'T DO THIS.**

Just write integration tests then at that point, and put them in a different directory. Sometimes by making your code easier to *unit* test you're making it harder to use and reason about.

Dependency Injection is a good thing, but it needs to be applied in sensible ways. (TODO: Future blog post?)

## Testing the mocks

Consider this example:

```javascript
async update(request, reply) {
  const id = request.params.id
  const instance = await User.query().patchAndFetchById(id, request.payload)

  if (!instance) {
    return Boom.notFound(`Cannot find user with ID '${id}'`)
  }
  return instance
}
```

Mocking out the User object so that database interaction doesn't happen means you're essentially testing whether you got your mock right. Does this code correctly handle when the ID is not found? It either throws or returns null, but we don't know if we mock it.

Also consider if you get the mock right but later an update to the library changes the behavior. With a unit test you won't know something is wrong because *you're just testing your mocks*.

Here's another kind of function not worth unit testing that I often see:

```javascript
async createValue(a, b) {
  const c = await widgetService.fetchWidget(a, b)
  const trx = await transactionService.start(c)
  const result = await billingService.createBill(trx, c)
  return result
}
```

This is "glue code". Unit testing makes no sense and will slow you down when business needs change.
