package handler

import (
	"github.com/TvGelderen/budget-buddy/view/pages"
	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleHomePage(c echo.Context) error {
    userDto := apiCfg.GetUser(c.Request())

    return render(c, pages.Index(userDto)); 
}

func (apiCfg *ApiConfig) HandleDashboardPage(c echo.Context) error {
    userDto := apiCfg.GetUser(c.Request())

    return render(c, pages.Dashboard(userDto)); 
}

func (apiCfg *ApiConfig) HandleRegisterPage(c echo.Context) error {
    userDto := apiCfg.GetUser(c.Request())
    
    if userDto.Username != "" {
        return c.Redirect(302, "/")
    }

    return render(c, pages.Register(userDto));
}

func (apiCfg *ApiConfig) HandleLoginPage(c echo.Context) error {
    userDto := apiCfg.GetUser(c.Request())

    if userDto.Username != "" {
        return c.Redirect(302, "/")
    }

    return render(c, pages.Login(userDto));
}
