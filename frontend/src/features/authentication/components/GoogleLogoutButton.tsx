import { useAuth } from '../../../context/useAuth';
import { logoutGoogleUser } from '../../../services/authService';
import { styles } from '../../../styles';

function GoogleLogoutButton() {
    const { setAuthUser, setIsAuthenticated } = useAuth();

    const handleLogout = async () => {
        try {
            await logoutGoogleUser();

            // Update auth context
            setAuthUser(null);
            setIsAuthenticated(false);
        } catch (err) {
            console.error('Logout failed', err);
        }
    }

    return (
        <button className={styles.button} onClick={handleLogout}>
            Logout
        </button>
    );
}

export default GoogleLogoutButton;
