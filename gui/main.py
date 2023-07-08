from PyQt5.QtGui import *
from PyQt5.QtCore import *
from PyQt5.QtWidgets import *

import sys
import socket
import messages_pb2
import struct

# SocketServer class. This will be used to send proto3 messages to the Golang backend.
class SocketServer():
    def __init__(self, ipAddress):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.ipAddr = ipAddress
        self.host, self.port = self.ipAddr.split(":")
        self.port = int(self.port)

    # Sends the given message to the provided address
    def send(self, smsg):
        self.socket.connect((self.host, self.port))
        self.socket.send(smsg)
    
    # TODO - Receives a message
    def recv(self):
        pass

    # Closes the socket
    def close(self):
        self.socket.close()

        

# ServerWindow class. It's opened when a user begins a new RED server.
class ServerWindow(QWidget):
    def __init__(self, ipAddress):
        super().__init__()
        self.server = None
        self.process = None

        self.ipAddr = ipAddress

        self.x, self.y, self.aw, self.ah, = 100, 300, 300, 150
        self.run()
    
    def run(self):
        if self.process is None:
            self.setMinimumSize(QSize(320, 140))
            self.setWindowTitle("RED")

            # TODO - need to create a toolbar
            # self.toolbar = QToolBar("Test")
            # self.toolbar.show()

            self.document = QPlainTextEdit(self)
            self.document.textChanged.connect(self.fetch)
            self.document.resize(QSize(320, 140))

            self.killBtn = QPushButton('KILL', self)
            self.killBtn.clicked.connect(self.sendKill)

            # NOTE - we might want to get outputs from stderr and stdout
            print(f"Server starting under {self.ipAddr}")
            exec = "../cmd/server/server"
            self.process = QProcess()
            self.process.readyReadStandardOutput.connect(lambda: print(self.process.readAllStandardOutput().data().decode()))
            self.process.readyReadStandardError.connect(lambda: print(self.process.readAllStandardError().data().decode()))
            self.process.finished.connect(self.destroy)
            self.process.setProgram(exec)
            self.process.setArguments([self.ipAddr])
            self.process.start()


    # Fetches updates from the text editor
    # When this function is called, we will open a socket server from this Python process and send messages to the backend
    def fetch(self):
        text = self.document.toPlainText()
        print(text)
    
    def sendKill(self):
        # Create a protobuf message and pack the header into it
        # This is the equivalent of marshalling/unmarshalling we do on the backend
        smsg = messages_pb2.REDMessage()
        smsg.type = messages_pb2.KILL
        raw = smsg.SerializeToString()
        encoded = struct.pack('<H', len(raw)) + raw

        self.server = SocketServer(self.ipAddr)
        self.server.send(encoded)
        self.server.close()

    # Destroy the server once it's been terminated
    def destroy(self):
        print(f"Server under {self.ipAddr} shutting down...")
        self.process = None

# Main GUI application
class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()

        # This is a dictionary because we should be able to spawn as many servers as we want
        # The key is the IP address.
        # The value is a tuple holding the ServerWindow and Process object
        self.serverWindows = []

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
        startBtn.clicked.connect(self.spawnServerWindow)
        startBtn.resize(200, 32)
        startBtn.move(100, 100)

        self.show()
    
    # Spawns a new window for the created server
    def spawnServerWindow(self):
        host, addr = f"{self.hostBox.text()}", f"{self.portBox.text()}"
        ipAddr = f"{host}:{addr}"
        # Don't spawn a new window if the address is already being used.
        # If a user wants to start a new window under an address that is already being used, then they must close the window associated with that address first
        if ipAddr in self.serverWindows:
            print("Address already in use")
        else:            
            # TODO - if the server under a certain ip address is shutdown, how does this class find out about it?
            self.serverWindows.append(ServerWindow(ipAddr))
            self.serverWindows[len(self.serverWindows)-1].show()


if __name__ == '__main__':
    app = QApplication(sys.argv)
    gui = MainWindow()
    sys.exit(app.exec_())