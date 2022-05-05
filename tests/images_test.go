package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ku2482/alfred-aws-icons/awsutil"
)

func TestImages(t *testing.T) {
	awsServices := awsutil.ParseAwsServices("../services.yaml")
	for _, awsService := range awsServices {
		// PNG
		image_path, err := filepath.Abs("../images/png/" + awsService.Id + ".png")
		if err != nil {
			t.Fatal(err.Error())
		}
		_, err = os.Stat(image_path)
		if err != nil {
			t.Fatal(err.Error())
		}

		// SVG
		image_path, err = filepath.Abs("../images/svg/" + awsService.Id + ".svg")
		if err != nil {
			t.Fatal(err.Error())
		}
		_, err = os.Stat(image_path)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}
