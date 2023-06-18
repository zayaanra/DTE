# RED
RED is a **R**eal-time **E**diting **D**istribution software.
## Installation
1. Make sure you have `Go (most recent version)` installed.
2. You also need to have protobuf installed:
`go get google.golang.org/protobuf` and
`go get google.golang.org/protobuf/proto`

## Idea
This project will be a distributed text editor. Users (processes) can create text files and edit them as desired. Users may also invite other users to edit the text file. Users that are not invited should not have permission to edit the text file. For now, servers can only be hosted locally (from the same machine).

## Concerns
1. What happens when a process fails?
2. Are there any distributed systems protocols that could be useful here? (for example, a consensus algorithm may be useful so that participating users agree on the state of the document)

## Notes
1. If a user saves a document, then the document should be saved for all participating users on that document.
2. Only the user who created the document may delete it. If a document is deleted, then all participating users on that document should not be able to access that document.
3. We may use a consensus algorithm so that the set of participating users for a document agree on the state of the document at any given time.

## Implementation
- Front-End
  - We'll use a simple GUI built in Golang. Users can choose a port (hosted locally) to start their server on. Users can also invite other peers once the server is started.
- Back-End
  - Golang will be used a back-end. We'll use Gorilla for this. The back-end will receive edit messages and process them. It will send back "OK" messages if a user has permission to edit a document.
- Project Layout
