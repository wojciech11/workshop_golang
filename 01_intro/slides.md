---
marp: true
theme: gaia
color: #000
colorSecondary: #333
backgroundColor: #fff
style: |
    section.lead h1 {
    text-align: center;
    }

---
<!-- _class: lead -->
# Golang Workshop

Introduction

---
<!-- _class: lead -->
## About me
<br />

- Wojciech Barczynski
- Tech lead
- Golang Developer with many years of experience

---
<!-- _class: lead -->
## What I like

- Simplicity
- Standard library
- Composition
- Pragmatism
- Readability

---
<!-- _class: lead -->
## Culture

Golang is simple.<br />
So we dive directly into talking about problems to solve.

---
<!-- _class: lead -->
## Birth

- Born out of frustration.
- Efficient compilation, efficient execution, or ease of programming?
- First sketches in September 2007 (Robert Griesemer, Rob Pike, Ken Thompson).
- Plenty of siblings (Rust, Elixir, Swift)

---
<!-- _class: lead -->
## Language goals

- Easy: compiled, statically typed, but with a dynamic feel.
- Modern: support for networked and multicore computing.
- Fast: at most a few seconds to build a large executable on a single computer.

---
<!-- _class: lead -->
## Language goals

"Go was intended as a language for writing server programs that would be easy to maintain over time." (golang.org)

"Clean procedural language designed for scalable cloud software." (Rob Pike)

---
<!-- _class: lead -->
## Core principles

- Reduce clutter and complexity.
- Everything is declared exactly once.
- Syntax is clean and light on keywords.
- No hierarchy: types just are, they don’t have to announce their relationships.
- Orthogonality of concepts.
- Orthogonality of features.

---
<!-- _class: lead -->
## Where is my feature?

- Go does not compete on features.
- As of Go ~~1~~ 2, the language is fixed.
- Features add complexity and hurt readability.
- Simple features that interact in simple ways.
- Simple means readable and predictable.
- Simple means easy to maintain.

---
<!-- _class: lead -->
## Simple things in Go

- garbage collection
- goroutines
- constants
- interfaces
- packages

Each hides complexity behind a simple facade.

---
<!-- _class: lead -->
## Provebs

- [https://go-proverbs.github.io/](https://go-proverbs.github.io/)

---
<!-- _class: lead -->
## Idiomatic Golang

- Simplicity
- Pragmatic
- Readability

---
<!-- _class: lead -->
## Idiomatic Golang

Examples:

- Test against real db
- Dislike mocks
- Dislike ORMs
- Avoid unneccessary abstractions

---
<!-- _class: lead -->
## Idiomatic Golang

- Talks of Rob Pike
- Peter Bourgon's talks and blog (e.g., [peter.bourgon.org/go-for-industrial-programming](https://peter.bourgon.org/go-for-industrial-programming))
- Dave Cheney

---
<!-- _class: lead -->
## Idiomatic Golang

- Join Golang meetups in your city
- Join the slack [invite.slack.golangbridge.org](https://invite.slack.golangbridge.org/)

---
<!-- _class: lead -->
## Where Golang shines

- macro/micro-services, especially on K8S
- Kubernetes/CloudNative ecosystem
- Serverless

---
<!-- _class: lead -->
## The most important

From https://go.dev/blog/gopher:

![](https://go.dev/blog/gopher/header.jpg)

---
<!-- _class: lead -->
## The most important

From https://github.com/ashleymcnamara/gophers:

![width:400px](https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GO_LEARN.png)

---
<!-- _class: lead -->
## Basics
<br />
Packages:

- collection of files in the same dir;
- all functions,...,variables are visible within the package.

---
<!-- _class: lead -->
## Basics
<br />
Repository:

- one or more modules;
- modules is a collection of packages;
- each module path indicates where to download it.

---
<!-- _class: lead -->
## References
<br />

- [Rob Pike - Simplicity is complicated](https://www.youtube.com/watch?v=rFejpH_tAHM)
- [Tomasz Grodzki - Why Golang looks like that](https://github.com/golangpoland/meetup_golang_warsaw/tree/master/2018/2018_21_Meetup/WhyGoLangLooksLikeThat)
