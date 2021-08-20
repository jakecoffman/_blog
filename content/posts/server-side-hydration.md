---
title: Server-Side Hydration Anti-Pattern
date: 2020-12-10T14:49:48-05:00
tags: [api]
---

When designing and implementing a REST API in a service-oriented or microservice architecture you may be tempted to fetch remote objects so the UI doesn't have to fetch them individually later. This is often called hydration or inflation.

I will now try to convince you that this is an anti-pattern to hydrate server-side.

<!--more-->

## An Example

Let's say you have a service for users, and your REST Resource contains a list of users. The naive approach is to store the user IDs in the list and then allow the UI to hydrate the users by calling to the user service.

Maybe you or the UI team thinks this is too much work for the UI to do! So you decide to hydrate your resource with the full user object by calling to the user service just before the response.

Suddenly you need to handle these new complexities:
- The server needs to run these calls concurrently or risk being very slow.
- The concurrent calls need to be rate limited, or you risk DDOSing your user service.

These are issues that browsers take care of for you automatically. The browser rate limits to (typically) six concurrent HTTP calls at a time.

Caching is another issue for the API. If you cache the data server-side you open a new can of worms: cache invalidation. Caching can also cause security issues if the user service denies external users from listing other external users, for example.

Having the UI fetch the resource makes much more sense from a design perspective as well. The initial call will complete quickly since it doesn't do any kind of hydration. This gives the UI some things to display while it fetches the user data in the list of user IDs. 

The UI is also in the position of knowing **WHICH** users to fetch. Consider a resource that has thousands of users. It doesn't make sense to fetch them all, the UI probably can only display a handful of user cards at a time anyway. So it can fetch the 5-10 that it needs, and the rest can wait. 

The UI also can now cache the user objects separately from the Resource. This way if the user is surfing the different resources, and it encounters the same users it doesn't need to re-hydrate. Cache invalidation is also simple: invalidate when the user refreshes. It can also use cache headers and a service-worker with the user service to further optimize caching.

## Keep the API Simple

Server-side hydration causes an explosion of complexity that is hard to get right. By pushing hydration onto the UI, the overall system is simpler, and the system that is in the best position to decide what data needs to be fetched is in control.
