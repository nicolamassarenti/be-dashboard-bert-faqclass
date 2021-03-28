# BERT Faq Classification - dashboard backend

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/nicolamassarenti/be-dashboard-bert-faqclass)
[![HitCount](http://hits.dwyl.com/nicolamassarenti/be-dashboard-bert-faqclass.svg)](http://hits.dwyl.com/nicolamassarenti/be-dashboard-bert-faqclass)

Welcome in the repository of the back-end dashboard for the _bert-faqclass_ project! This repository contains the code to set-up a dashboard backend to control the knowledge base of your AI chatbot.

_Note_: check out the use case description in this [Medium article](https://nicola-massarenti.medium.com/end-to-end-machine-learning-bert-faqclass-81fd07d24058).

## About the design pattern

The code relies on the Clean Architecture, described by [Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), based on the _Dependency Rule_. Citing Uncle Bob'b blog:
>The concentric circles represent different areas of software. In general, the further in you go, the higher level the software becomes. The outer circles are mechanisms. The inner circles are policies.
>
>![Clean Architecture Code Design](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg "Clean Architecture")
>
>The overriding rule that makes this architecture work is The Dependency Rule. This rule says that source code dependencies can only point inwards. Nothing in an inner circle can know anything at all about something in an outer circle. In particular, the name of something declared in an outer circle must not be mentioned by the code in the an inner circle. That includes, functions, classes. variables, or any other named software entity.
> By the same token, data formats used in an outer circle should not be used by an inner circle, especially if those formats are generate by a framework in an outer circle. We don’t want anything in an outer circle to impact the inner circles.

The code is inspired by [Manuel Kiessling's blog example](https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/), 8 years old but still useful.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Before to install the software make sure you already have installed:

- [Golang](https://golang.org/): an open source programming language that makes it easy to build simple, reliable, and efficient software.

## Local development and running tests

It's easy! In order to run the code on your local machine you have to run the following command:

```bash
go run main.go
```

## Deployment and production build

To create a production build you need to run the following command:

```bash
go build main.go
```

## Authors

- **Nicola Massarenti**

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
