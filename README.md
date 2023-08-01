# RED
RED is a **R**eal-time **E**diting **D**istribution software.
## Installation
1. Make sure you have `Go (most recent version)` installed.
2. You also need to have protobuf installed:
`go get google.golang.org/protobuf` and
`go get google.golang.org/protobuf/proto`

## How it works
Upon launching the GUI, users (processes) are given a blank text editor to work with. This serves as their "document". They can freely edit this document. To colloborate with other users, they must know the addresses of those users in order to communicate with them. This involves sending that user an invite upon which the receiveing user will be granted access to edit the now shared document. If the invited user has a document of their own, it will be wiped in order to connect with the document of the inviter.
