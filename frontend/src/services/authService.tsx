// This is a new file where you handle the logout logic
export const logoutGoogleUser = async () => {
    // redirect to backend logout route
    window.location.href = 'http://localhost:3000/auth/logout/google';

    // ideally need to handle cookies and other side-effects here from what I read
    // although i did do some cookie handling server side
}

export const loginGoogleUser = async () => {
    // redirect to google login page
    window.location.href = 'http://localhost:3000/auth/google'
}