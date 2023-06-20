import sys
from PyQt5.QtGui import *
from PyQt5.QtCore import *
from PyQt5.QtWidgets import *

class App(QMainWindow):
    def __init__(self):
        super().__init__()
        self.x, self.y, self.aw, self.ah, = 500, 1000, 1000, 750
        self.run()

    def run(self):
        self.setGeometry(500, 1000, 1000, 750)
        self.setWindowTitle("RED")
        
        self.hostBox = QLineEdit(self)
        self.hostBox.move(250, 100)
        self.portBox = QLineEdit(self)
        self.portBox.move(500, 100)

        self.hostLabel = QLabel('Host Name (or IP Address)', self)
        self.hostLabel.adjustSize()
        self.hostLabel.setBuddy(self.hostBox)
        self.portLabel = QLabel('Port', self)
        self.portLabel.setBuddy(self.portBox)

        self.show()
    


if __name__ == '__main__':
    app = QApplication(sys.argv)
    gui = App()
    sys.exit(app.exec_())

