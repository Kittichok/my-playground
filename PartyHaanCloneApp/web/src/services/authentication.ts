

function login(username: string, password: string) {
    // const requestOptions = {
    //     method: 'POST',
    //     headers: { 'Content-Type': 'application/json' },
    //     body: JSON.stringify({ username, password })
    // };

    // return fetch(`${config.apiUrl}/users/authenticate`, requestOptions)
    //     .then(handleResponse)
    //     .then(user => {
    //         // store user details and jwt token in local storage to keep user logged in between page refreshes
    //         localStorage.setItem('currentUser', JSON.stringify(user));
    //         currentUserSubject.next(user);

    //         return user;
    //     });
    //TODO Call backend
    localStorage.setItem('currentUser', JSON.stringify("tokenasdasd"));
    return true;
}

function logout() {
    localStorage.removeItem('currentUser');
}

export const authenticationService = {
    login,
    logout,
    token: JSON.parse(localStorage.getItem('currentUser') as string),
};
