from PyQt5.QtGui import *
from PyQt5.QtCore import *
from PyQt5.QtWidgets import *
import sys
import socket

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
        self.toolbar = QToolBar("Test")
        self.toolbar.show()

        self.document = QPlainTextEdit(self)
        self.document.resize(QSize(320, 140))

        

# Main GUI application
class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        self.serverWindow = None

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
        # TODO
        ipAddr = f"{self.hostBox.text()}:{self.portBox.text()}"
        print(f"Server starting under {ipAddr}")
        self.serverWindow = ServerWindow()
        self.serverWindow.show()

if __name__ == '__main__':
    app = QApplication(sys.argv)
    gui = MainWindow()
    sys.exit(app.exec_())