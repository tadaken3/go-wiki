package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestServer(t *testing.T) {
    server := httptest.NewServer(routes())
    defer server.Close()

    resp, err := http.Get(server.URL)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        t.Fatal("status:", resp.StatusCode)
    }

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatal(err)
    }

    if string(respBody) != "Hello, world.\n" {
        t.Fatal("body:", string(respBody))
    }
}