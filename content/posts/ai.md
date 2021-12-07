---
title: My AI Saga
date: 2021-08-20T00:00:00-05:00
tags: [programming]
---

I started getting into AI in 2020 when a coworker asked me if I wanted to compete on the [CodinGame](https://www.codingame.com) Spring Challenge. After 11 days, [I finished 112th](https://www.codingame.com/contests/spring-challenge-2020/leaderboard/global?column=keyword&value=jke) out of about 5,000 contestants. I was hooked!

<!--more-->

## 2020 Spring Challenge

{{< youtube TA2hfUBZs84 >}}

My basic strategy in this challenge was to use A* to find the most valuable route to the larger pellets, or just to find a valuable route that my opponent probably won't beat me to.   

Unfortunately I can't reveal my code since anyone could copy and paste it into the game. I did however create a repo called [graph](https://github.com/jakecoffman/graph) that has generic versions of algorithms in Go that you can copy and paste into your project and use.

Over the next year I joined in 2 more competitions.

## 2020 Fall Challenge

{{< youtube UiFTRy2mdCo  >}}

In my time since the Spring Challenge I basically forgot the competition existed, so when I started this one I only had A* in my arsenal. At first glance it looked like I wouldn't be able to use it, but once you realize any problem can be reduced to a graph or a tree, and then traverse it using A*. So that's what I did for this competition.

To my surprise, I ranked lower this time: 248th out of 7,011. I started reading about Beam Search which many of the other competitors said they used. I had trouble finding good code examples. The write-ups for Beam were in more academic papers which contained mathematical notation rather than code snippets.

## 2021 Spring Challenge

{{< youtube KcfwfAJxgyc >}}

This time I was prepared to try Beam Search, but I was still having trouble finding nice examples for that algorithm. Someone in the chat said they were using a variant called Chokudai (named after another competitor!) which is a variant that uses depth-first search rather than breadth-first. I found a very clear and simple C++ program explaining the algorithm, so I was off!

Unfortunately I had a worse placement than before: 413th out of 6,867. This time between competitions I decided to learn more, and capture what I already knew in my [graph](https://github.com/jakecoffman/graph) repo.

## Genetic algorithm

After I made some clear examples of Beam and Chokudai, I decided to learn about Genetic Algorithms (GA) which had always interested me.

{{< youtube 81mh4uJ7WHA >}}

In Code vs Zombies you can see my GA is corralling zombies in for a huge combo. In this optimization problem [I am ranked 56th out of 10,000+](https://www.codingame.com/multiplayer/optimization/code-vs-zombies/leaderboard?column=keyword&value=jke)!

{{< youtube UVmh9adAtc4 >}}

In Mars Lander my GA navigates difficult terrain to land at the right speed on flat ground.

For GAs it's important to visualize what each generation is doing. Here's what generation 0 might look like:

![Generation 0](/blog/mars/gen0.svg)

There are some wispy lines that I pre-generate to get better results, and the rest are random.

Generation 0 is my initial population. After selection, breeding, and mutating, we get Generation 1:

![Generation 1](/blog/mars/gen1.svg)

Now there appears to be fewer lines, I suspect there is a lot of overlap. However now the lines are beginning to move in the direction of the landing site.

Jumping up to Generation 10:

![Generation 10](/blog/mars/gen10.svg)

Generation 10 is laser focused the landing site, but appears to be crashing straight into it. There should be more of a curve, so it lands safely.

Jumping forward again to the final and best Generation:

![Generation 154](/blog/mars/gen154.svg)

Now we see that the GA is slowing the lander down, so it can land safely.

## What's next?

2021 Fall Challenge is coming up. I'm excited and hopeful to place high in the competition, assuming I can carve out enough time.

I also have some more algorithms I need to try out, like Minimax.
