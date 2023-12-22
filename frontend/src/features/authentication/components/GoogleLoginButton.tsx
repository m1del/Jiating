function GoogleLoginButton() {
  const handleLogin = () => {
    // redirect to google login page
    window.location.href = 'http://localhost:3000/auth/google'
  }

  return (
    <button onClick={handleLogin} className="text-white">
      Login with Google
    </button>
  )
}

export default GoogleLoginButton
