# DTE
## Idea
This project will be a distributed text editor. Users (processes) can create text files and edit them as desired. Users may also invite other users to edit the text file. Users that are not invited should not have permission to edit the text file.

## Concerns
1. What happens when a process fails?
2. Are there any distributed systems protocols that could be useful here? (for example, a consensus algorithm may be useful so that participating users agree on the state of the document)

## Notes
1. If a user saves a document, then the document should be saved for all participating users on that document.
2. Only the user who created the document may delete it. If a document is deleted, then all participating users on that document should not be able to access that document.
3. We may use a consensus algorithm so that the set of participating users for a document agree on the state of the document at any given time.

## Implementation
- Front-End
  - Python will be used a front-end. We'll use HTML/CSS and PyScript. The front-end will begin the GUI process. Users can (visually) edit documents from here.
- Back-End
  - Golang will be used a back-end. We'll use Gorilla for this. The back-end will receive edit messages and process them. It will send back "OK" messages if a user has permission to edit a document.
- Project Layout
```bash
my-editor/
├── frontend/
│   ├── app.py
│   ├── templates/
│   │   ├── index.html
│   │   └── ...
│   ├── static/
│   │   ├── css/
│   │   │   └── style.css
│   │   ├── js/
│   │   │   └── script.js
│   │   └── ...
│   └── ...
├── backend/
│   ├── main.go
│   ├── api/
│   │   ├── handlers.go
│   │   └── ...
│   ├── models/
│   │   ├── document.go
│   │   └── ...
│   └── ...
└── ...
```
