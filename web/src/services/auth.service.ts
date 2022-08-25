import axios from 'axios'

type Credential = {
    email: string,
    password: string
}

type LoginResponse = {
    token: string
}

export function login(credential: Credential): Promise<LoginResponse> {
    return axios.post(`${import.meta.env.VITE_API_BASE_URL}/login`, credential)
} 