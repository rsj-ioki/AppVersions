package usecase

import (
	"fmt"
	"testing"
)

func TestAndroidAppInfoSuccess(t *testing.T) {
	want := App{
		Id:       "com.ioki.hamburg",
		Name:     "ioki Hamburg",
		Url:      "https://play.google.com/store/apps/details?id=com.ioki.hamburg",
		ImageSrc: "https://play-lh.googleusercontent.com/08VmA2BAm0Jiy3tatZU3cQ2y07HH98q8qvdwUjGaPLJTW2uwVLHfpikqxzI4Av71btc",
		Error:    false,
	}
	got := androidAppInfo("com.ioki.hamburg")

	fmt.Printf("\nGot the following Android App info:\n%+v\n\n", got)
	if want.Id != got.Id {
		t.Errorf("%s is not %s", want.Id, got.Id)
	}
	if want.Name != got.Name {
		t.Errorf("%s is not %s", want.Name, got.Name)
	}
	if "" == got.Version {
		t.Errorf("Version is empty")
	}
	if "" == got.Rating {
		t.Errorf("Rating is empty")
	}
	if want.Url != got.Url {
		t.Errorf("%s is not %s", want.Url, got.Url)
	}
	if want.ImageSrc != got.ImageSrc {
		t.Errorf("%s is not %s", want.ImageSrc, got.ImageSrc)
	}
	if want.Error != got.Error {
		t.Errorf("%t is not %t", want.Error, got.Error)
	}
}

func TestAndroidAppInfoFailure(t *testing.T) {
	want := App{
		Id:       "com.ioki.does.not.exist",
		Name:     "",
		Version:  "",
		Rating:   "",
		Url:      "https://play.google.com/store/apps/details?id=com.ioki.does.not.exist",
		ImageSrc: "",
		Error:    true,
	}
	got := androidAppInfo("com.ioki.does.not.exist")

	fmt.Printf("\nGot the following Android App info:\n%+v\n\n", got)
	if want.Id != got.Id {
		t.Errorf("%s is not %s", want.Id, got.Id)
	}
	if want.Name != got.Name {
		t.Errorf("%s is not %s", want.Name, got.Name)
	}
	if "" != got.Version {
		t.Errorf("Version is not empty")
	}
	if "" != got.Rating {
		t.Errorf("Rating is not empty")
	}
	if want.Url != got.Url {
		t.Errorf("%s is not %s", want.Url, got.Url)
	}
	if want.ImageSrc != got.ImageSrc {
		t.Errorf("%s is not %s", want.ImageSrc, got.ImageSrc)
	}
	if want.Error != got.Error {
		t.Errorf("%t is not %t", want.Error, got.Error)
	}
}
