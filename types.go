package hackcheck

type Result struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	IP       string `json:"ip"`
	Phone    string `json:"phone"`
	Source   struct {
		Name        string `json:"name"`
		Date        string `json:"date"`
		Description string `json:"description"`
	} `json:"source"`
}

type Response struct {
	Found   int      `json:"found"`
	Elapsed string   `json:"elapsed"`
	Results []Result `json:"results"`
	Success bool     `json:"success"`
	Message string   `json:"message"`
}

type Field = string

const (
	FieldEmail    Field = "email"
	FieldUsername Field = "username"
	FieldName     Field = "name"
	FieldPassword Field = "password"
	FieldIP       Field = "ip"
	FieldPhone    Field = "phone"
	FieldDomain   Field = "domain"
)
