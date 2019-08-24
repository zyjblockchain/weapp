package weapp

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestBankCardByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiBankcard, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiBankcard {
			t.Fatalf("Except to path '%s',get '%s'", apiBankcard, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"type", "access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"id": "622213XXXXXXXXX"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := bankCardByURL(ts.URL+apiBankcard, "mock-access-token", ts.URL+"/mediaurl", RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}

	_, err = bankCardByURL(ts.URL+apiBankcard, "mock-access-token", ts.URL+"/mediaurl", RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBankCardByFile(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiBankcard {
			t.Fatalf("Except to path '%s',get '%s'", apiBankcard, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"type", "access_token"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"id": "622213XXXXXXXXX"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := bankCardByFile(ts.URL+apiBankcard, "mock-access-token", testIMGName, RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}
	_, err = bankCardByFile(ts.URL+apiBankcard, "mock-access-token", testIMGName, RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriverLicenseByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiDriving, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDriving {
			t.Fatalf("Except to path '%s',get '%s'", apiDriving, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"id_num": "660601xxxxxxxx1234",
			"name": "张三",
			"sex": "男",
			"nationality": "中国",
			"address": "广东省东莞市xxxxx号",
			"birth_date": "1990-12-21",
			"issue_date": "2012-12-21",
			"car_class": "C1",
			"valid_from": "2018-07-06",
			"valid_to": "2020-07-01",
			"official_seal": "xx市公安局公安交通管理局"
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := driverLicenseByURL(ts.URL+apiDriving, "mock-access-token", ts.URL+"/mediaurl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDriverLicenseByFile(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDriving {
			t.Fatalf("Except to path '%s',get '%s'", apiDriving, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"id_num": "660601xxxxxxxx1234",
			"name": "张三",
			"sex": "男",
			"nationality": "中国",
			"address": "广东省东莞市xxxxx号",
			"birth_date": "1990-12-21",
			"issue_date": "2012-12-21",
			"car_class": "C1",
			"valid_from": "2018-07-06",
			"valid_to": "2020-07-01",
			"official_seal": "xx市公安局公安交通管理局"
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := driverLicenseByFile(ts.URL+apiDriving, "mock-access-token", testIMGName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIDCardByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiIDCard, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiIDCard {
			t.Fatalf("Except to path '%s',get '%s'", apiIDCard, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"type", "access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"type": "Front",
			"id": "44XXXXXXXXXXXXXXX1"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := idCardByURL(ts.URL+apiIDCard, "mock-access-token", ts.URL+"/mediaurl", RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}

	_, err = idCardByURL(ts.URL+apiIDCard, "mock-access-token", ts.URL+"/mediaurl", RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIDCardByFile(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiIDCard {
			t.Fatalf("Except to path '%s',get '%s'", apiIDCard, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		queries := []string{"type", "access_token"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"type": "Front",
			"id": "44XXXXXXXXXXXXXXX1"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := idCardByFile(ts.URL+apiIDCard, "mock-access-token", testIMGName, RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}
	_, err = idCardByFile(ts.URL+apiIDCard, "mock-access-token", testIMGName, RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVehicleLicenseByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiDriving, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDriving {
			t.Fatalf("Except to path '%s',get '%s'", apiDriving, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"type", "access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"vhicle_type": "小型普通客⻋",
			"owner": "东莞市xxxxx机械厂",
			"addr": "广东省东莞市xxxxx号",
			"use_character": "非营运",
			"model": "江淮牌HFCxxxxxxx",
			"vin": "LJ166xxxxxxxx51",
			"engine_num": "J3xxxxx3",
			"register_date": "2018-07-06",
			"issue_date": "2018-07-01",
			"plate_num_b": "粤xxxxx",
			"record": "441xxxxxx3",
			"passengers_num": "7人",
			"total_quality": "2700kg",
			"prepare_quality": "1995kg"
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := vehicleLicenseByURL(ts.URL+apiDriving, "mock-access-token", ts.URL+"/mediaurl", RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}

	_, err = vehicleLicenseByURL(ts.URL+apiDriving, "mock-access-token", ts.URL+"/mediaurl", RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVehicleLicenseByFile(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDriving {
			t.Fatalf("Except to path '%s',get '%s'", apiDriving, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		queries := []string{"type", "access_token"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"vhicle_type": "小型普通客⻋",
			"owner": "东莞市xxxxx机械厂",
			"addr": "广东省东莞市xxxxx号",
			"use_character": "非营运",
			"model": "江淮牌HFCxxxxxxx",
			"vin": "LJ166xxxxxxxx51",
			"engine_num": "J3xxxxx3",
			"register_date": "2018-07-06",
			"issue_date": "2018-07-01",
			"plate_num_b": "粤xxxxx",
			"record": "441xxxxxx3",
			"passengers_num": "7人",
			"total_quality": "2700kg",
			"prepare_quality": "1995kg"
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := vehicleLicenseByFile(ts.URL+apiDriving, "mock-access-token", testIMGName, RecognizeModePhoto)
	if err != nil {
		t.Fatal(err)
	}
	_, err = vehicleLicenseByFile(ts.URL+apiDriving, "mock-access-token", testIMGName, RecognizeModeScan)
	if err != nil {
		t.Fatal(err)
	}
}
