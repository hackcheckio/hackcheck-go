package hackcheck

type SearchResult struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	IPAddress   string `json:"ip_address"`
	PhoneNumber string `json:"phone_number"`
	Hash        string `json:"hash"`
	Source      struct {
		Name        string `json:"name"`
		Date        string `json:"date"`
		Description string `json:"description"`
	} `json:"source"`
}

// External use
type SearchResponse struct {
	Databases  int                       `json:"databases"`
	Results    []SearchResult            `json:"results"`
	Pagination *SearchResponsePagination `json:"pagination"`
}

// Internal use
type searchResponse struct {
	Databases  int                       `json:"databases"`
	Found      int                       `json:"found"`
	Elapsed    string                    `json:"elapsed"`
	Error      string                    `json:"error"`
	Results    []SearchResult            `json:"results"`
	Pagination *SearchResponsePagination `json:"pagination"`
}

type SearchResponsePagination struct {
	DocumentCount int `json:"document_count"`
	Next          *struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"next"`
	Prev *struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"prev"`
}

type SearchFilterType = string

const (
	SearchFilterTypeUse    = "use"
	SearchFilterTypeIgnore = "ignore"
)

type SearchFilterOptions struct {
	Type      SearchFilterType
	Databases []string
}

type SearchPaginationOptions struct {
	Offset int
	Limit  int
}

type SearchOptions struct {
	Field      SearchField
	Query      string
	Filter     *SearchFilterOptions
	Pagination *SearchPaginationOptions
}

type SearchField = string

const (
	SearchFieldEmail       SearchField = "email"
	SearchFieldUsername    SearchField = "username"
	SearchFieldFullName    SearchField = "full_name"
	SearchFieldPassword    SearchField = "password"
	SearchFieldIPAddress   SearchField = "ip_address"
	SearchFieldPhoneNumber SearchField = "phone_number"
	SearchFieldDomain      SearchField = "domain"
	SearchFieldHash        SearchField = "hash"
)
