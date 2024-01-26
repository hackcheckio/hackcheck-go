package hackcheck

import (
	"log"
	"os"
	"testing"
)

var (
	client *HackCheckClient
)

func TestMain(m *testing.M) {
	keyFile, err := os.ReadFile("TEST_API_KEY.txt")
	if err != nil {
		panic("failed to open TEST_API_KEY.txt")
	}

	if len(string(keyFile)) < 20 {
		panic("missin api key in TEST_API_KEY.txt")
	}

	client = New(string(keyFile))

	os.Exit(m.Run())
}

// just basic testing...
func TestHackcheckClient_Monitors(t *testing.T) {
	resp, err := client.GetMonitors()
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.AssetMonitors) != 1 && len(resp.DomainMonitors) != 1 {
		t.Fatal("no asset / domain monitors")
	}

	firsta, err := client.GetAssetMonitor(resp.AssetMonitors[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	firstd, err := client.GetDomainMonitor(resp.DomainMonitors[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	if firsta.Asset == "" || firstd.Domain == "" {
		t.Fatal("empty response")
	}
}

func TestHackCheckClient_Check(t *testing.T) {
	resp, err := client.Check(&CheckOptions{Field: SearchFieldEmail, Query: "hello@gmail.com"})
	if err != nil {
		t.Fatal(err)
	}

	if resp == false {
		t.Fatal("check returned false")
	}
}

func TestHackCheckClient_Search(t *testing.T) {
	searchQueries := [][]string{
		{SearchFieldDomain, "example.com"},
		{SearchFieldEmail, "hello@example.com"},
		{SearchFieldFullName, "John Smith"},
		{SearchFieldHash, "ae2b1fca515949e5d54fb22b8ed95575"},
		{SearchFieldIPAddress, "123.123.123.123"},
		{SearchFieldPassword, "somepassword"},
		{SearchFieldPhoneNumber, "1112223333"},
		{SearchFieldUsername, "walter"},
	}

	queryWithLotsOfResults := []string{SearchFieldUsername, "walter"}

	t.Log("testing search queries")
	for _, s := range searchQueries {
		field, query := s[0], s[1]

		t.Logf("searching %s %s", field, query)

		response, err := client.Search(&SearchOptions{Field: field, Query: query, Pagination: &SearchPaginationOptions{Limit: 100}})
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Results) < 1 {
			t.Error("results count less than 1")
		}
	}

	t.Log("testing pagination")

	{
		response1, err := client.Search(&SearchOptions{
			Field:      queryWithLotsOfResults[0],
			Query:      queryWithLotsOfResults[1],
			Pagination: &SearchPaginationOptions{Limit: 20, Offset: 0},
		})

		if err != nil {
			log.Fatal(err)
		}

		if response1.Pagination == nil {
			log.Fatal("response1 Pagination is nil")
		}

		if response1.Pagination.Next == nil {
			log.Fatal("response1 Pagination.Next is nil")
		}

		response2, err := client.Search(&SearchOptions{
			Field:      queryWithLotsOfResults[0],
			Query:      queryWithLotsOfResults[1],
			Pagination: &SearchPaginationOptions{Limit: 20, Offset: 1},
		})

		if err != nil {
			log.Fatal(err)
		}

		if response2.Pagination.Prev == nil {
			log.Fatal("response2 Pagination.Prev is nil")
		}
	}
}
