package main

import (
	"context"
	"crypto/sha256"
	"database/sql"

	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi"
	"github.com/gobuffalo/packr"
	"github.com/spf13/viper"
	"github.com/stephenafamo/expense-tracker/auth"
	"github.com/stephenafamo/expense-tracker/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mailgun/mailgun-go.v1"

	"github.com/volatiletech/authboss"
	_ "github.com/volatiletech/authboss/auth"
	_ "github.com/volatiletech/authboss/lock"
	_ "github.com/volatiletech/authboss/logout"
	_ "github.com/volatiletech/authboss/remember"

	"github.com/volatiletech/authboss/defaults"
)

type server struct {
	router       *chi.Mux
	DB           *sql.DB
	auth         *authboss.Authboss
	mailer       mailgun.Mailgun
	Files        *packr.Box
	BundledFiles *packr.Box
	Handlers     Handlers
}

func NewServer(db *sql.DB) *server {
	s := &server{
		DB: db,
	}

	s.configure()

	return s
}

func (s *server) configure() {
	s.configureMailer()
	s.configureRouter()
	s.configureAuth()
	s.loadFiles()
	s.registerHandlers()
}

func (s *server) addAdminUser() {
	ctx := context.Background()
	email := viper.GetString("ADMIN_EMAIL")

	user, err := models.Users(qm.Where("email=?", email)).One(ctx, s.DB)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		// Add the user
		user = &models.User{}
		user.Email = email

		password, err := bcrypt.GenerateFromPassword([]byte(viper.GetString("ADMIN_PASSWORD")), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		user.Password = null.StringFrom(string(password))
		err = user.Insert(ctx, s.DB, boil.Infer())
		if err != nil {
			panic(err)
		}
	}

}

func (s *server) configureMailer() {
	mg := mailgun.NewMailgun(viper.GetString("mailgun_domain"), viper.GetString("mailgun_priv_key"), viper.GetString("mailgun_pub_key"))
	s.mailer = mg
}

func (s *server) configureRouter() {
	r := chi.NewRouter()
	s.router = r
	s.loadRoutes()
}

func (s *server) loadFiles() {
	Box := packr.NewBox("./views/")
	s.Files = &Box
}

func (s *server) configureAuth() {
	s.auth = authboss.New()
	hashKey := sha256.Sum256([]byte(viper.GetString("hash_key")))
	blockKey := sha256.Sum256([]byte(viper.GetString("block_key")))

	// authboss.LoadClientStateMiddleware
	s.auth.Config.Storage.Server = auth.NewUserStore(s.DB)
	s.auth.Config.Storage.SessionState = auth.NewSessionStore(s.DB, hashKey[:], blockKey[:])
	s.auth.Config.Storage.CookieState = auth.NewCookieStore(hashKey[:], blockKey[:])

	s.auth.Config.Paths.Mount = "/auth"
	s.auth.Config.Paths.RootURL = getBaseUrl().String()

	s.auth.Config.Mail.From = viper.GetString("mail_from")
	s.auth.Config.Mail.FromName = viper.GetString("mail_from_name")
	s.auth.Config.Modules.LogoutMethod = "GET"

	s.auth.Config.Core.ViewRenderer = AuthRenderer{Base: "auth", server: s}
	s.auth.Config.Core.MailRenderer = AuthRenderer{Base: "auth/mail", server: s}

	// Set up defaults for basically everything besides the ViewRenderer/MailRenderer in the HTTP stack
	defaults.SetCore(&s.auth.Config, false, false)

	s.auth.Config.Core.Mailer = AuthMailer{}

	if err := s.auth.Init(); err != nil {
		panic(err)
	}
}

func (s *server) getTemplate(name string) (string, error) {
	body, err := s.Files.FindString(name + ".html")
	return body, err
}

func (s *server) render(templateName string, data ...interface{}) (string, error) {
	provider := PartialProvider{}

	template, err := s.getTemplate(templateName)
	checkError(err)

	return mustache.RenderPartials(template, provider, data...)
}
