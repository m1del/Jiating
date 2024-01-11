import { loginGoogleUser } from '../../../services/authService';
import { styles } from '../../../styles';

function GoogleLoginButton() {
    const handleLogin = () => {
        loginGoogleUser();
    }

    return (
        <button className={styles.button} onClick={handleLogin}>
            Login with Google
        </button>
    );
}

export default GoogleLoginButton;
