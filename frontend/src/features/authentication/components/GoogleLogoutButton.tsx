import { styles } from '../../../styles';


import { useAuth } from '../../../context/AuthContext';

function GoogleLogoutButton() {
    const { setAuthUser, setIsLoggedin } = useAuth();

    const handleLogout = () => {
        try {
            // redirect to backend logout route
            window.location.href = 'http://localhost:3000/auth/logout/google';
            // update auth context
            setAuthUser(null);
            setIsLoggedin(false);
            
            // clear cookies
            document.cookie.split(";").forEach((c) => {
                document.cookie = c.trim().split("=")[0] + "=;expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
            });

        } catch (err) {
            console.error('Logout failed', err);
        }
    }

  return (
    <button className={`${styles.button}`} onClick={handleLogout}>
        Logout
    </button>
  )
}

export default GoogleLogoutButton
