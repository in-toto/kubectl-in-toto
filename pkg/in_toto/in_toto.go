package in_toto

import (
	"github.com/in-toto/in-toto-golang/in_toto"
)

type VerificationSetup struct {
    TargetType string
    Name string
    KeyPath string
    LayoutPath string
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

    if err != nil {
        return err
    }

    if _, err = in_toto.InTotoVerify(layout, keyMap, linkDir, "toplevel", nil); err != nil {
        return err
    }


	return nil
}
