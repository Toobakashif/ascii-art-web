# <b>ASCII ART WEB</b>
---
## <b>Description</b>
Browser based Ascii-art generator with a simple GUI
## <b>Authors</b>
* KeitiN <br>
* tooba <br>
* ailearner <br>
## <b>Usage: how to run</b>
* Run the program with the command: <br>
`go run main.go`
* Open your browser, and type in: <b>localhost:3000</b>
* To see the 200 OK Status code, press F12 and choose the Network tab. If you press "Submit" button and everything works fine, status code 200 should show up.
* To test the 404 error, simply type in a non-existing path, for example: localhost:3000/notapage
* To test the 400 error, go to the path localhost:3000/ascii-art without submitting any data
* To test the 500 error, try to use symbols, that are not covered in ascii banners, for example "ö", "ä", "ü" or "õ"
## <b>Implementation details: algorithm</b>
The program fires up a local web-server and allows the user to choose a font and input a text that the user wishes to see written in that font. When input is given and "Submit" button is pushed, the webpages gets directed to path /ascii-art, where the result is displayed.