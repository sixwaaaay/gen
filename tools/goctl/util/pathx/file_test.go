package pathx

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/sixwaaaay/gen/internal/version"
	"github.com/stretchr/testify/assert"
)

func TestGetTemplateDir(t *testing.T) {
	category := "foo"
	t.Run("before_have_templates", func(t *testing.T) {
		home := t.TempDir()
		RegisterHome("")
		RegisterHome(home)
		v := version.Version()
		dir := filepath.Join(home, v, category)
		err := MkdirIfNotExist(dir)
		if err != nil {
			return
		}
		tempFile := filepath.Join(dir, "bar.txt")
		err = ioutil.WriteFile(tempFile, []byte("foo"), os.ModePerm)
		if err != nil {
			return
		}
		templateDir, err := GetTemplateDir(category)
		if err != nil {
			return
		}
		assert.Equal(t, dir, templateDir)
		RegisterHome("")
	})

	t.Run("before_has_no_template", func(t *testing.T) {
		home := t.TempDir()
		RegisterHome("")
		RegisterHome(home)
		dir := filepath.Join(home, category)
		err := MkdirIfNotExist(dir)
		if err != nil {
			return
		}
		templateDir, err := GetTemplateDir(category)
		if err != nil {
			return
		}
		assert.Equal(t, dir, templateDir)
	})

	t.Run("default", func(t *testing.T) {
		RegisterHome("")
		dir, err := GetTemplateDir(category)
		if err != nil {
			return
		}
		assert.Contains(t, dir, version.BuildVersion)
	})
}

func TestGetGitHome(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	actual, err := GetGitHome()
	if err != nil {
		return
	}

	expected := filepath.Join(homeDir, genDir, gitDir)
	assert.Equal(t, expected, actual)
}

func TestGetHome(t *testing.T) {
	t.Run("gen_is_file", func(t *testing.T) {
		tmpFile := filepath.Join(t.TempDir(), "a.tmp")
		backupTempFile := tmpFile + ".old"
		err := ioutil.WriteFile(tmpFile, nil, 0o666)
		if err != nil {
			return
		}
		RegisterHome(tmpFile)
		home, err := GetHome()
		if err != nil {
			return
		}
		info, err := os.Stat(home)
		assert.Nil(t, err)
		assert.True(t, info.IsDir())

		_, err = os.Stat(backupTempFile)
		assert.Nil(t, err)
	})

	t.Run("gen_is_dir", func(t *testing.T) {
		RegisterHome("")
		dir := t.TempDir()
		RegisterHome(dir)
		home, err := GetHome()
		assert.Nil(t, err)
		assert.Equal(t, dir, home)
	})
}
