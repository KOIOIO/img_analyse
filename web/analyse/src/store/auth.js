import { login, register } from '../api/users'

const state = {
    token: localStorage.getItem('token') || null,
    user: null
}

const mutations = {
    SET_TOKEN(state, token) {
        state.token = token
        localStorage.setItem('token', token)
    },
    SET_USER(state, user) {
        state.user = user
    },
    CLEAR_AUTH(state) {
        state.token = null
        state.user = null
        localStorage.removeItem('token')
    }
}

const actions = {
    async login({ commit }, credentials) {
        const res = await login(credentials)
        commit('SET_TOKEN', res.data.token)
        commit('SET_USER', {
            userName: res.data.userName,
            userId: res.data.userId
        })
        return res
    },

    async register(_, userData) {
        return await register(userData)
    },

    logout({ commit }) {
        commit('CLEAR_AUTH')
    }
}

const getters = {
    isAuthenticated: state => !!state.token,
    currentUser: state => state.user
}

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}