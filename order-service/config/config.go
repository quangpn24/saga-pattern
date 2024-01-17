package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppEnv                        string `envconfig:"APP_ENV"`
	Port                          string `envconfig:"PORT"`
	TokenSecret                   string `envconfig:"TOKEN_SECRET"`
	SentryDSN                     string `envconfig:"SENTRY_DSN"`
	UserPoolID                    string `envconfig:"USER_POOL_ID"`
	CognitoIssuer                 string `envconfig:"COGNITO_ISSUER"`
	CognitoURLGetJWKS             string `envconfig:"COGNITO_URL_GET_JWKS"`
	AllowOrigins                  string `envconfig:"ALLOW_ORIGINS"`
	SMTPUsername                  string `envconfig:"SMTP_USERNAME"`
	SMTPPass                      string `envconfig:"SMTP_PASS"`
	SMTPEndpoint                  string `envconfig:"SMTP_ENDPOINT"`
	URLExpireIn                   int    `envconfig:"URL_EXPIRE_IN"`
	URLExpectationFormExpireInDay int    `envconfig:"URL_EXPECTATION_FORM_EXPIRE_IN_DAY"`
	LPHostDefault                 string `envconfig:"LP_HOST_DEFAULT"`
	AppHostDefault                string `envconfig:"APP_HOST_DEFAULT"`
	HIFMailAddress                string `envconfig:"HIF_MAIL_ADDRESS"`
	UrlFormLoan                   string `envconfig:"URL_FORM_LOAN"`
	UrlSettingResona              string `envconfig:"URL_SETTING_RESONA"`
	UrlResonaSuccess              string `envconfig:"URL_RESONA_SUCCESS"`
	UrlResonaError                string `envconfig:"URL_RESONA_ERROR"`
	UrlVerifyEmail                string `envconfig:"URL_VERIFY_EMAIL"`

	DB struct {
		Name      string `envconfig:"DB_NAME"`
		Host      string `envconfig:"DB_HOST"`
		Port      int    `envconfig:"DB_PORT"`
		User      string `envconfig:"DB_USER"`
		Pass      string `envconfig:"DB_PASS"`
		EnableSSL bool   `envconfig:"ENABLE_SSL"`
	}

	RabbitMQ struct {
		DSN string `envconfig:"RABBITMQ_DSN"`
	}

	Worker struct {
		MaxConcurrency int    `envconfig:"WORKER_MAX_CONCURRENCY"`
		MaxCapacity    int    `envconfig:"WORKER_MAX_CAPACITY"`
		QueueName      string `envconfig:"QUEUE_NAME"`
	}

	AWSConfig struct {
		Region    string `envconfig:"AWS_REGION"`
		AccessKey string `envconfig:"AWS_ACCESS_KEY"`
		SecretKey string `envconfig:"AWS_SECRET_KEY"`
	}

	EKYCConfig struct {
		SenderAddress        string `envconfig:"SENDER_ADDRESS"`
		PathUrlMailEkyc      string `envconfig:"PATH_URL_MAIL_EKYC"`
		APIKey               string `envconfig:"API_KEY"`
		RedirectPathURL      string `envconfig:"REDIRECT_PATH_URL"`
		ErrorRedirectPathURL string `envconfig:"ERROR_REDIRECT_PATH_URL"`
		UrlLiquid            string `envconfig:"URL_LIQUID"`
		UrlSuccess           string `envconfig:"URL_SUCCESS_EKYC"`
	}
	FinanceConfig struct {
		PathUrlMailIncomeProof string `envconfig:"PATH_URL_MAIL_INCOME_PROOF"`
		PathUrlMailApprove     string `envconfig:"PATH_URL_MAIL_APPROVE"`
	}

	Crypto struct {
		CryptoSecretKey string `envconfig:"CRYPTO_SECRET_KEY"`
	}
	S3Config struct {
		BucketName string `envconfig:"BUCKET_NAME"`
	}

	Resona struct {
		PasswordResona    string `envconfig:"PASSWORD_RESONA"`
		CompanyCode       string `envconfig:"COMPANY_CODE"`
		ProductIdentifyID string `envconfig:"PRODUCT_IDENTIFY_ID"`
		SettingType       string `envconfig:"SETTING_TYPE"`
		AccountType       string `envconfig:"ACCOUNT_TYPE"`
		UrlResona         string `envconfig:"URL_RESONA"`
	}
}

func LoadConfig() (*Config, error) {
	// load default .env file, ignore the error
	_ = godotenv.Load()

	cfg := new(Config)
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("load config error: %v", err)
	}

	return cfg, nil
}
