---
title: The Role of Algorithms in Computing
slug: the-role-of-algorithms-in-computing
description: The Role of Algorithms in Computing
date: 2024-07-25
tags: alien, wine
---

This is the first of a long series of posts. Once I have started to read the
[Introduction of Algorithms](http://mitpress.mit.edu/books/introduction-algorithms)
by Thomas H. Cormen, I am going to publish a new post for every read chapter.

This one, in particular, are the answer for the first chapter:
_The Role of Algorithms in Computing_.

I have been used some useful tools to help me get and format the answer. They
are the following:

- [WolframAlpha](http://www.wolframalpha.com)
- [MathJax](http://www.mathjax.org/)

Also, I will be using [Go](http://golang.org) as the implementing language, once I have started to
learn it.

> Give an example of an application that requires algorithmic content at the
application level, and discuss the function of the algorithms involved.

**Answer:** When we need to find a route in Google Maps. The algorithm can vary among the
shortest route, route without tolls, without ferries, etc.

> Suppose we are comparing implementations of insertion sort and merge sort on
the same machine. For inputs of size n, insertion sort runs in 8n2 steps, while
merge sort runs in 64n lg n steps. For which values of n does insertion sort
beat merge sort?

**Answer:** Merge sort beats insertion sort when `n > 43`. See the plot below.

> What is the smallest value of n such that an algorithm whose running time is
100n2 runs faster than an algorithm whose running time is 2n on the same machine?

**Answer:** The first algorith runs faster at `n > 14`.
