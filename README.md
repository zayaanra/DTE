# DTE
# Idea
This project will be a distributed text editor. Users (processes) can create text files and edit them as desired. Users may also invite other users to edit the text file. Users that are not invited should not have permission to edit the text file.

## Concerns
1. What happens when a process fails?
2. Are there any distributed systems protocols that could be useful here? (for example, a consensus algorithm may be useful so that participating users agree on the state of the document)

## Notes
1. If a user saves a document, then the document should be saved for all participating users on that document.
2. Only the user who created the document may delete it. If a document is deleted, then all participating users on that document should not be able to access that document.
3. We may use a consensus algorithm so that the set of participating users for a document agree on the state of the document at any given time.
