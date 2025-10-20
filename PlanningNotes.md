# go-microsvc-template ʕ◔ϖ◔ʔ
Getting started with your very own Go microservice template.

## Base Reqs
**Project Structure** 🌳
- Where should the main app logic live?
- Which dirs are most relevant for beginning dev
- Should we send the pkg dir to the shadow realm?
- How do I handle standalone functionality? (ie. cron job, asynq workers, etc.)

**Observability** 👀 
- Logging (structured)
    - Should I use zap, logrus, or native log?
    - Global logger vs dependency injection vs struct embedding?
    - How should logging be implemented in the app? How do we integrate the same logger across the app (ie. db logs, redis logs, api logs, etc.)
- Tracing (find bottlenecks in microsvcs)
    - Opentelemtry (distributed tracing)
        - Collector of traces
        - Doesn't actually collect/visualize traces, need a sep backend for that (Jaeger / Prometheus?)
        - Can be deployed as a sep binary in Go
    - Maybe use DD tracing / send the traces to Otel
- Sentry
    - For unhandled errors
- All the above should live in a config struct?

**Peripherals / Maintenance** 👩‍💻🧑‍💻
- Naming conventions (variables, funcs etc.)
    - Why does Go prefer short and concise? (w *Writer) / (r *Reader) vs (reader *Reader)
- Database usage (ie. global vs other ways to handle it)
- Best practices for contexts?
    - Necessary to avoid unnecessary work within app 
        - (svc A -> svc B -> svc C)
        - If request fails at B, then don't call C
    - Don't make context objects large
    - Use it for ephemeral data (last as long as the request, etc.)
        - trace_id
        - auth_token
        - cookie_id 
- DB Migrations
    - Which ORMs to consider?
    - Which migration packages to consider?
- Testing
    - Unit testing / code coverage in Go
- Linting
- Precommit specific to Go

**Deployment** 🚀🚀🚀
- Build / Deployment (nothing Go specific here?)
<br><br><br>
## Case by Case Requirements
- Handling Async jobs/cron within Go

<br><br>
## Helpful Resources

### Go
- [Zen of Go](https://dave.cheney.net/2020/02/23/the-zen-of-go)
- [Go PR Review Tips](https://github.com/golang/go/wiki/CodeReviewComments#variable-names)

### Observablity
- [Intro to Tracing - Go specific](https://www.youtube.com/watch?v=idDu_jXqf4E)