package authservices

import userrepo "finder-job/repository/user"

type AuthService struct {
	UserRepo *userrepo.UserRepo
}


func NewAuthService(userrepo *userrepo.UserRepo) *AuthService {
   return &AuthService{UserRepo: userrepo}
}


func (s *AuthService) Login() {}
func (s *AuthService) Register() {}
func (s *AuthService) Logout() {}
