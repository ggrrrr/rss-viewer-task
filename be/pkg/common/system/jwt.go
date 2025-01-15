package system

import "github.com/ggrrrr/rss-viewer-task/be/pkg/common/auth"

func initJWT(system *System) error {
	if system.cfg.CrtKeyFile == "" {
		system.verifier = auth.NewMockVerifier()
		return nil
	}

	verifier, err := auth.NewVerifier(system.cfg.CrtKeyFile)
	if err != nil {
		return err
	}

	system.verifier = verifier

	return nil
}
