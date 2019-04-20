package in_toto

import (
	"github.com/in-toto/in-toto-golang/in_toto"
    "os"
    "fmt"
    "strings"
)

type VerificationSetup struct {
    TargetType string
    Name string
    KeyPath string
    LayoutPath string
}

func saveImageID(ImageID string) error {
    file, err := os.Create("image_id")

    if err != nil {
        return err 
    }
    
    slices := strings.Split(ImageID, "//")

    if len(slices) != 2 {
        return fmt.Errorf("%v has an unexpected number of slices", slices)
    }

    file.WriteString(slices[1])
    file.Close()
    return nil
}

// ScanDefinition scans the provided resource definition.
func ScanContainer(setup *VerificationSetup, imageName string) (error) {

    linkDir := "./"

    var key in_toto.Key
    if err := key.LoadPublicKey(setup.KeyPath); err != nil {
        return err
    }

    var keyMap = map[string]in_toto.Key{
        key.KeyId: key,
    }

    var layout in_toto.Metablock
    err := layout.Load(setup.LayoutPath)

    var parameters = map[string]string {
        "IMAGE_ID": imageName,
    }
    saveImageID(imageName)

    if err != nil {
        return err
    }

    if _, err = in_toto.InTotoVerify(layout, keyMap, linkDir, 
            "toplevel", parameters); err != nil {
        return err
    }


	return nil
}
