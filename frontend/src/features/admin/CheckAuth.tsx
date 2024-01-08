import { UserInfo } from './components';

type SetAuthUser = (user: UserInfo.UserType) => void;
type SetIsLoggedIn = (loggedIn: boolean) => void;

const CheckAuth = (setAuthUser: SetAuthUser, setIsLoggedin: SetIsLoggedIn) => {
  fetch('http://localhost:3000/api/session-info', { credentials: 'include' })
    .then((res) => {
      if (res.ok) {
        return res.json();
      } else {
        throw new Error('Not Authenticated');
      }
    })
    .then((data) => {
      if (data.authenticated) {
        setAuthUser({
          id: data.userID,
          email: data.email,
          name: data.name,
          avatar_url: data.avatar_url,
        });
        setIsLoggedin(true);
      } else {
        setIsLoggedin(false);
        window.location.href = 'http://localhost:3000/auth/google';
      }
    })
    .catch((error) => {
      console.log('Authentication error: ', error);
      setIsLoggedin(false);
      window.location.href = 'http://localhost:3000/auth/google';
    });
};

export default CheckAuth;
