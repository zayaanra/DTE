# TODO - Use socket library to send proto3 messages to the golang side.
# TODO - Use selenium. We can use selenium to continously fetch updates from the text document the user is currently editing.
# We will send those updates to our Golang backend.

from bs4 import BeautifulSoup
import requests


URL = "http://localhost:8080"
page = requests.get(URL)

soup = BeautifulSoup(page.content, "html.parser")

while True:
    results = soup.find(id="doc")
    print(results)