# Flux

We believe in GitOps:

- **You declaratively describe the entire desired state of your
  system in git.** This includes the apps, config, dashboards,
  monitoring and everything else.
- **What can be described can be automated.** Use YAMLs to enforce
  conformance of the system. You don't need to run `kubectl`, all changes go
  through git. Use diff tools to detect divergence between observed and
  desired state and get notifications.
- **You push code not containers.** Everything is controlled through
  pull requests. There is no learning curve for new devs, they just use
  your standard git PR process. The history in git allows you to recover
  from any snapshot as you have an sequence of transactions. It is much
  more transparent to make operational changes by pull request, e.g.
  fix a production issue via a pull request instead of making changes to
  the running system.

Flux is a tool that automatically ensures that the state of a cluster
matches the config in git. It uses an operator in the cluster to trigger
deployments inside Kubernetes, which means you don't need a separate CD tool.
It monitors all relevant image repositories, detects new images, triggers
deployments and updates the desired running configuration based on that
(and a configurable policy).

The benefits are: you don't need to grant your CI access to the cluster, every
change is atomic and transactional, git has your audit log. Each transaction
either fails or succeeds cleanly. You're entirely code centric and don't need
new infrastructure.

![Deployment Pipeline](site/images/deployment-pipeline.png)

[![CircleCI](https://circleci.com/gh/weaveworks/flux.svg?style=svg)](https://circleci.com/gh/weaveworks/flux)
[![GoDoc](https://godoc.org/github.com/weaveworks/flux?status.svg)](https://godoc.org/github.com/weaveworks/flux)

## What Flux does

Flux is most useful when used as a deployment tool at the end of a
Continuous Delivery pipeline. Flux will make sure that your new
container images and config changes are propagated to the cluster.

Among its features are:

- [Automated git → cluster synchronisation](/site/introduction.md#automated-git-cluster-synchronisation)
- [Automated deployment of new container images](/site/introduction.md#automated-deployment-of-new-container-images)
- [Integrations with other devops tools](/site/introduction.md#integrations-with-other-devops-tools) ([Helm](site/helm/helm-integration.md) and more)
- No additional service or infrastructure needed - Flux lives inside your
  cluster

## Get started with Flux

Get started by browsing through the documentation below:

- [Introduction to Flux](/site/introduction.md)
- [FAQ](/site/faq.md)
- [How it works](/site/how-it-works.md)
- [Installing Flux](/site/installing.md)
- [Using Flux](/site/using.md)
- [Upgrading to Flux v1](/site/upgrading-to-1.0.md)
- [Troubleshooting](/site/troubleshooting.md)

## Features

Flux is open-source and makes GitOps happen in your cluster in a secure and
obvious way: you don't need a separate CD tool, instead you make CD part of
the orchestration system in the cluster. Flux will ensure that the cluster
state always matches the config in git.

Weave Cloud builds on top of the great work in Flux and is our SaaS product
that along with a UI and alerts for deployments, gives you full
observability and insights into your cluster as well. It simplifies
deployment, monitoring and management for containers and microservices.
Read our full feature comparison to find out what's for you.

|                                                                                                | OSS Flux | Weave Cloud |
| ---------------------------------------------------------------------------------------------- |:--------:|:-----------:|
| &#x026AB; Speed up deployments of any workload into Kubernetes                                 | x        | x           |
| &#x026AB; Automated or manual deployments if you need more control                             | x        | x           |
| &#x026AB; Automatic or one-click manual rollbacks                                              | x        | x           |
| &#x026AB; Audit trail (who did what when) of deployments/commits/alerts/notifications          | x        | x           |
| &#x026AB; Batch operations - control workloads and their deployment strategies easily          | x        | x           |
| &#x026AB; Flux ensures that the cluster state matches the config in git                        | x        | x           |
| &#x026AB; Operators can lock a deployment and state why                                        | x        | x           |
| &#x026AB; Helm support for deployments (alpha)                                                 | x        | x           |
| &#x026AB; Overview of the system deployment state (version, automation status, locks, etc).    |          | x           |
| &#x026AB; All flux operations accessible at the click of a button                              |          | x           |
| &#x026AB; Straight-forward setup of your configuration                                         |          | x           |
| &#x026AB; Route deployment notifications to where your team reads them (slack, email, browser) |          | x           |
| &#x026AA; Start using monitoring and dashboards for workloads and nodes right after install    |          | x           |
| &#x026AA; Track latency and request duration of your workloads if instrumented for prometheus  |          | x           |
| &#x026AA; Use monitoring notebooks for groups of metrics you are interested in                 |          | x           |
| &#x026AA; Inspect and learn from 13 months of history of metrics and cluster state             |          | x           |
| &#x026AA; Use GCP Stackdriver and get logs and events easily aggregated                        |          | x           |
| &#x026AA; Debug and explore the state of your cluster in real-time                             |          | x           |
| &#x026AA; Use Time Travel to inspect and analyse the cluster and workloads history stats       |          | x           |
| &#x026AA; Invite team members to your cluster                                                  |          | x           |
| &#x026AA; Metrics of workloads are automatically* discovered and instantly available           |          | x           |

| Legend                                                                                             |
| -------------------------------------------------------------------------------------------------- |
| &#x026AB; GitOps feature                                                                           |
| &#x026AA; Observability feature                                                                    |
| <sup>* this includes many language- and framework-specific metrics, e.g. go, jvm, http, etc.</sup> |

Learn more about Weave Cloud, its features and see it in action on
[its homepage](https://www.weave.works/product/cloud/).

## Developer information

[Build documentation](/site/building.md)

[Release documentation](/internal_docs/releasing.md)

### Contribution

Flux follows a typical PR workflow.
All contributions should be made as PRs that satisfy the guidelines below.

### Guidelines

- All code must abide [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Names should abide [What's in a name](https://talks.golang.org/2014/names.slide#1)
- Code must build on both Linux and Darwin, via plain `go build`
- Code should have appropriate test coverage, invoked via plain `go test`

In addition, several mechanical checks are enforced.
See [the lint script](/lint) for details.

## <a name="help"></a>Getting Help

If you have any questions about Flux and continuous delivery:

- Read [the Weave Flux docs](https://github.com/weaveworks/flux/tree/master/site).
- Invite yourself to the <a href="https://weaveworks.github.io/community-slack/" target="_blank">Weave community</a> slack.
- Ask a question on the [#flux](https://weave-community.slack.com/messages/flux/) slack channel.
- Join the <a href="https://www.meetup.com/pro/Weave/"> Weave User Group </a> and get invited to online talks, hands-on training and meetups in your area.
- Send an email to <a href="mailto:weave-users@weave.works">weave-users@weave.works</a>
- <a href="https://github.com/weaveworks/flux/issues/new">File an issue.</a>

Your feedback is always welcome!
