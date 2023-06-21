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
        

# Main GUI application
class App(QMainWindow):
    def __init__(self):
        super().__init__()
        self.server = SocketServer()
        self.x, self.y, self.aw, self.ah, = 100, 300, 300, 150
        self.run()

    def run(self):
        # self.setGeometry(self.x, self.y, self.aw, self.ah)
        self.setMinimumSize(QSize(320, 140))
        self.setWindowTitle("RED")
        
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
        print("Button clicked")

if __name__ == '__main__':
    app = QApplication(sys.argv)
    gui = App()
    sys.exit(app.exec_())