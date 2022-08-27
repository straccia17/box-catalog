import { client } from './http.service'

type Credential = {
    email: string,
    password: string
}

type LoginResponse = {
    token: string
}

export function login(credential: Credential): Promise<LoginResponse> {
    return client.post('/login', credential)
} 