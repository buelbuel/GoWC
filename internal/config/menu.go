package config

type MenuItem struct {
	Label string
	URL   string
}

type Menu struct {
	Items []MenuItem
}

var (
	FrontMenu = Menu{
		Items: []MenuItem{
			{Label: "Register", URL: "/register"},
			{Label: "Login", URL: "/login"},
		},
	}

	AppMenu = Menu{
		Items: []MenuItem{
			{Label: "Dashboard", URL: "/app/dashboard"},
			{Label: "Profile", URL: "/app/profile"},
			{Label: "Logout", URL: "/logout"},
		},
	}
)

func GetMenu(layout string) Menu {
	switch layout {
	case "AppLayout":
		return AppMenu
	default:
		return FrontMenu
	}
}
