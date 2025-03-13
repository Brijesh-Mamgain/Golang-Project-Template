package service

import (
	"fmt"
	"project-auth-service/domain"
	"project-auth-service/model"
	appexception "project-common/exception"
	"project-common/logger"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	Login(model.LoginRequest) (*model.LoginResponse, *appexception.AppError)
	Verify(urlParams map[string]string) *appexception.AppError
	Refresh(request model.RefreshTokenRequest) (*model.LoginResponse, *appexception.AppError)
}

type DefaultAuthService struct {
	repo            domain.AuthRepository
	rolePermissions domain.RolePermissions
}

func (s DefaultAuthService) Refresh(request model.RefreshTokenRequest) (*model.LoginResponse, *appexception.AppError) {
	if vErr := request.IsAccessTokenValid(); vErr != nil {
		if vErr.Errors == jwt.ValidationErrorExpired {
			// continue with the refresh token functionality
			var appErr *appexception.AppError
			if appErr = s.repo.RefreshTokenExists(request.RefreshToken); appErr != nil {
				return nil, appErr
			}
			// generate a access token from refresh token.
			var accessToken string
			if accessToken, appErr = domain.NewAccessTokenFromRefreshToken(request.RefreshToken); appErr != nil {
				return nil, appErr
			}
			return &model.LoginResponse{AccessToken: accessToken}, nil
		}
		return nil, appexception.AuthenticationError("invalid token")
	}
	return nil, appexception.AuthenticationError("cannot generate a new access token until the current one expires")
}

func (s DefaultAuthService) Login(req model.LoginRequest) (*model.LoginResponse, *appexception.AppError) {
	var appErr *appexception.AppError
	var login *domain.Login

	if login, appErr = s.repo.FindBy(req.Username, req.Password); appErr != nil {
		return nil, appErr
	}

	claims := login.ClaimsForAccessToken()
	authToken := domain.NewAuthToken(claims)

	var accessToken, refreshToken string
	if accessToken, appErr = authToken.NewAccessToken(); appErr != nil {
		return nil, appErr
	}

	if refreshToken, appErr = s.repo.GenerateAndSaveRefreshTokenToStore(authToken); appErr != nil {
		return nil, appErr
	}

	return &model.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s DefaultAuthService) Verify(urlParams map[string]string) *appexception.AppError {
	// convert the string token to JWT struct
	if jwtToken, err := jwtTokenFromString(urlParams["token"]); err != nil {
		return appexception.AuthenticationError(err.Error())
	} else {
		/*
		   Checking the validity of the token, this verifies the expiry
		   time and the signature of the token
		*/
		if jwtToken.Valid {
			// type cast the token claims to jwt.MapClaims
			claims := jwtToken.Claims.(*domain.AccessTokenClaims)
			/* if Role if user then check if the account_id and customer_id
			   coming in the URL belongs to the same token
			*/
			if claims.IsUserRole() {
				if !claims.IsRequestVerifiedWithTokenClaims(urlParams) {
					return appexception.AuthenticationError("request not verified with the token claims")
				}
			}
			// verify of the role is authorized to use the route
			isAuthorized := s.rolePermissions.IsAuthorizedFor(claims.Role, urlParams["routeName"])
			if !isAuthorized {
				return appexception.AuthenticationError(fmt.Sprintf("%s role is not authorized", claims.Role))
			}
			return nil
		} else {
			return appexception.AuthenticationError("Invalid token")
		}
	}
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}

func NewLoginService(repo domain.AuthRepository, permissions domain.RolePermissions) DefaultAuthService {
	return DefaultAuthService{repo, permissions}
}
