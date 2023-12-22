import { styles } from '../../../styles';

function GoogleLoginButton() {
  const handleLogin = () => {
    // redirect to google login page
    window.location.href = 'http://localhost:3000/auth/google'
  }

  return (
    <button className={`${styles.button}`} onClick={handleLogin}>
      Login with Google
    </button>
  )
}

export default GoogleLoginButton
