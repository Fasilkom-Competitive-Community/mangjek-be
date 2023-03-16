package http

import (
	gcs "cloud.google.com/go/storage"
	"fmt"
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"net/http"
	"os"
)

type gscUploadProfile struct {
	bkt *gcs.BucketHandle
}

func (d HTTPFileDelivery) uploadProfile(c *gin.Context) {
	var r gscUploadProfile
	img, err := imageupload.Process(c.Request, "file")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("aman load image")

	obj := r.bkt.Object(img.Filename)
	w := obj.NewWriter(c)

	if _, err := w.Write(img.Data); err != nil {
		fmt.Println(err)
	}
	fmt.Println("bucket aman")

	if err := w.Close(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("close writer aman")

	c.JSON(http.StatusCreated, httpCommon.Response{
		Data: img,
	})
	fmt.Println("response aman")
}

func ensureDir(dirName string) error {

	err := os.MkdirAll(dirName, os.ModeDir)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func (d HTTPFileDelivery) addFile(c *gin.Context) {
	img, err := imageupload.Process(c.Request, "file")

	if err != nil {
		fmt.Println(err)
		//panic(err)
	}

	var osFile *os.File

	if img.ContentType == "image/jpeg" {
		err = ensureDir("./images")
		if err != nil {
			print("content type salah")
			fmt.Println(err)
		}

		osFile, err = os.Create("./images/" + img.Filename + ".jpeg")
		if err != nil {
			fmt.Println(err)
		}
	} else if img.ContentType == "pdf" {
		err = ensureDir("./pdf")
		if err != nil {
			fmt.Println(err)
		}
		osFile, err = os.Create("./pdf/" + img.Filename + ".pdf")
		if err != nil {
			fmt.Println(err)
		}
	}

	osFile.Write(img.Data)

	//img.Write(c.Writer)

	c.JSON(http.StatusCreated, httpCommon.Response{
		Data: img,
	})
}
