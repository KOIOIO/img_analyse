import axios from 'axios'

const api = axios.create({
    baseURL: '/api/users'
})

// 登录接口
export const login = (data) => api.post('/login', data)

// 注册接口
export const register = (data) => api.post('/register', {
    userName: data.userName,
    password: data.password1,
    password2: data.password2
})