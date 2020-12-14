package toolbox

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

//WriteImageWithTemplate puts an image in an html template
func WriteImageWithTemplate(w http.ResponseWriter, img *image.Image, htmlTemplate string) {

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Fatalln("unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	if tmpl, err := template.New("image").Parse(htmlTemplate); err != nil {
		log.Println("unable to parse image template.")
	} else {
		data := map[string]interface{}{"Image": str}
		if err = tmpl.Execute(w, data); err != nil {
			log.Println("unable to execute template.")
		}
	}
}

//GetImageFromFilePath opens an image at that file path
func GetImageFromFilePath(filePath string) image.Image {
	f, err := os.Open(filePath)
	if err != nil {
		ErrorHandler(err)
	}
	image, _, err := image.Decode(f)
	if err != nil {
		ErrorHandler(err)
	}
	return image
}

//GetTemplate opens an html string
func GetTemplate(filePath string) string {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		//deal with err here
	}
	return string(f)
}
