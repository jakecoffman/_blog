---
title: Highly Available Cron-as-a-Service
date: 2021-08-21T00:01:33-05:00
tags: [story]
---

Cloud Foundry has a problem. There's no builtin way to run a single job on a schedule. So I developed one.

<!--more-->

## Problem

For HA reasons we run multiple datacenters. So if one of them loses power, catches fire, or blows away, the other is still being load balanced to and no one notices.

So the problem with running a single job every hour is all the instances across datacenters will try to fire the job at the same time, causing havoc.

## Solution

To solve this problem I used a Message Queue (MQ). I created a durable queue and inserted a single message into it. 

When the application starts up, it tries to grab the single message from the queue. Whichever instance gets the message then becomes the "executor" and installs the cron jobs and begins executing them. 

The executor also processes cron CRUD operations since they immediately need to be scheduled, rescheduled, or unscheduled. The non-executors send CRUD operations to another queue that the executor processes. The executor changes the schedule and then writes the change to the database.

The executor never acknowledges the single message in the MQ. If it shuts down during a deployment or for maintenance, it can nack the message and the broker will send it to another app. If it crashes, the broker will see it lost connection and then resend to another app that is still running. In this way, we've achieved High Availability!

## Result

We have over 350 cron jobs, many running every minute!

We rarely see a problem where during maintenance of the MQ for some reason the message is consumed. We simply check for that and insert the message again if it's missing.

Otherwise, this has been very reliable.

## Other Solutions

Another solution I thought of was a single database row that the application locks when it starts up.

I would need to look more into the situation where the application crashes whether the row lock is released. This would only work for databases that have row or object locking functionality.

The reason I didn't go this route is we only had older versions of Mongo that didn't have this kind of functionality.
