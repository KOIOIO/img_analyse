<template>
  <div class="login-page">
    <h1>用户登录</h1>
    <AuthForm
        button-text="登录"
        @submit="handleLogin"
    />
    <p v-if="error" class="error">{{ error }}</p>
    <p>没有账号？<router-link to="/register">立即注册</router-link></p>
  </div>
</template>

<script>
import AuthForm from '@/components/AuthForm.vue'

export default {
  name: 'LoginPage',
  components: { AuthForm },
  data() {
    return {
      error: ''
    }
  },
  methods: {
    async handleLogin(formData) {
      try {
        await this.$store.dispatch('auth/login', {
          userName: formData.userName,
          password: formData.password
        })
        this.$router.push('/not-found')
      } catch (err) {
        this.error = err.response?.data?.message || '登录失败'
      }
    }
  }
}
</script>

<style scoped>
.login-page {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
}
.error {
  color: red;
}
</style>