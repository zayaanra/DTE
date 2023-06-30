from PyQt5.QtGui import *
from PyQt5.QtCore import *
from PyQt5.QtWidgets import *

import sys
import socket
import subprocess

# SocketServer class. This will be used to send proto3 messages to the Golang backend.
class SocketServer():
    # TODO
    def __init__(self):
        self.socket = None

# ServerWindow class. It's opened when a user begins a new RED server.
class ServerWindow(QWidget):
    def __init__(self):
        super().__init__()
        self.server = None

        self.x, self.y, self.aw, self.ah, = 100, 300, 300, 150
        self.run()
    
    def run(self):
        self.setMinimumSize(QSize(320, 140))
        self.setWindowTitle("RED")

        # TODO - need to create a toolbar
        # self.toolbar = QToolBar("Test")
        # self.toolbar.show()

        self.document = QPlainTextEdit(self)
        self.document.textChanged.connect(self.fetch)
        self.document.resize(QSize(320, 140))
    
    # Fetches updates from the text editor
    # When this function is called, we will open a socket server from this Python process and send messages to the backend
    def fetch(self):
        text = self.document.toPlainText()
        print(text)

        

# Main GUI application
class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()

        # This is a dictionary because we should be able to spawn as many servers as we want
        # The key is the IP address and the value is the window associated it with that server
        self.serverWindows = {}

        self.x, self.y, self.aw, self.ah, = 100, 300, 300, 150
        self.run()

    def run(self):
        # self.setGeometry(self.x, self.y, self.aw, self.ah)
        self.setMinimumSize(QSize(320, 140))
        self.setWindowTitle("Configuration")
        
        self.hostLabel = QLabel('IP Address:', self)
        self.hostBox = QLineEdit(self)
        self.hostBox.move(100, 20)
        self.hostBox.resize(200, 32)
        self.hostLabel.move(20, 20)

        self.portLabel = QLabel('Port:', self)
        self.portBox = QLineEdit(self)
        self.portBox.move(100, 60)
        self.portBox.resize(200, 32)
        self.portLabel.move(20, 60)

        startBtn = QPushButton('Start', self)
        startBtn.clicked.connect(self.spawn)
        startBtn.resize(200, 32)
        startBtn.move(100, 100)

        self.show()
    
    # Spawns a Golang process to start the backend server
    def spawn(self):
        # TODO - spawn the golang process
        host, addr = f"{self.hostBox.text()}", f"{self.portBox.text()}"
        ipAddr = f"{host}:{addr}"
        # Don't spawn a new server if the address is already being used.
        # If a user wants to start a new server under an address that is already being used, then they must close the server associated with that address first
        if ipAddr in self.serverWindows:
            print("Address already in use")
        else:
            print(f"Server starting under {ipAddr}")
            self.serverWindows[ipAddr] = ServerWindow()
            self.serverWindows[ipAddr].show()
            exec = "../cmd/server/server"
            process = subprocess.Popen([exec, ipAddr], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            stderr, stdout = process.communicate()

            # Print stdout and stderr from spawned process
            print(stderr.decode())
            print(stdout.decode())

if __name__ == '__main__':
    app = QApplication(sys.argv)
    gui = MainWindow()
    sys.exit(app.exec_())