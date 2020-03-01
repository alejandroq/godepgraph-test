# Go Dependency Graphs and Automated PR Processes

Upon a Git commit, the dependecy tree documentation is updated to reflect the current application's package topology. 
In an opinionated [`pull_request_template.md`](./.github/pull_request_template.md) this visual could be leveraged for review and preventing developers from accidently introducing any ungainly package relationships. 
This implementation is particular to Go.

For additional commands and capabilities, look to the [`Makefile`](./Makefile).

## Install `godepgraph`

`go get github.com/kisielk/godepgraph`

## Setup Commit Hook

`cp .github/pre-commit .git/hooks/`

## Application Package Topology

![dependencytree.png](./.github/dependencytree.png)

*Quick notes on the topology above:*
- Ideally a dependency tree should be short and flat - and this should be in ratio to program statements.
- Too few packages given a high statement count likely means the application is purely concrete and has not declared any bounded contexts (in the zone of pain).
- Too many interwoven packages, with a great height (provides the appearence of a hill), signifies incorrectly asserted bounded contexts (in the zone of pain).
- Short and flat provided an appropriate ratio of statements signifies an application codebase that is potentially a modular monolith, which is an ideal for shops expounding characteristics such as discovery and experimentation. This is despite the allergy the industry has taken to the word monolith. 
- An important emphasis with the latter is the complexity be maintained low or the code architecture verges being within the zone of uselessness.
