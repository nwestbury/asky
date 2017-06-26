export const CREATE_SESSION = "CREATE_SESSION";

export const createSession = (username, token) => ({
    type: CREATE_SESSION,
    username,
    token
});