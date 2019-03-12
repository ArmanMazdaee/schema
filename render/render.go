package render

import (
	"io/ioutil"
	"os"

	"github.com/resumic/schema/schema"
)

func RenderHTML(resumeJSON []byte, themeDir string) ([]byte, error) {
	if err := schema.ValidateResume(resumeJSON); err != nil {
		return nil, err
	}

	siteDir, err := ioutil.TempDir(os.TempDir(), "resumic")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(siteDir)

	if err := hugoInitSite(siteDir); err != nil {
		return nil, err
	}

	resumeName := "resume"
	if err := hugoWriteResumeJSON(resumeJSON, resumeName, siteDir); err != nil {
		return nil, err
	}

	if err := hugoBuild(themeDir, siteDir); err != nil {
		return nil, err
	}

	return hugoReadResumeHTML(resumeName, siteDir)
}
