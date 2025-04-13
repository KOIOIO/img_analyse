<template>
  <div class="register-page">
    <h1>用户注册</h1>
    <AuthForm
        button-text="注册"
        :show-confirm="true"
        @submit="handleRegister"
    />
    <p v-if="error" class="error">{{ error }}</p>
    <p>已有账号？<router-link to="/login">立即登录</router-link></p>
  </div>
</template>

<script>
import AuthForm from '@/components/AuthForm.vue'

export default {
  name: 'RegisterPage',
  components: { AuthForm },
  data() {
    return {
      error: ''
    }
  },
  methods: {
    async handleRegister(formData) {
      if (formData.password !== formData.password2) {
        this.error = '两次密码输入不一致'
        return
      }

      try {
        await this.$store.dispatch('auth/register', {
          userName: formData.userName,
          password1: formData.password,
          password2: formData.password2
        })
        this.$router.push('/not-found')
      } catch (err) {
        this.error = err.response?.data?.message || '注册失败'
      }
    }
  }
}
</script>

<style scoped>
.register-page {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
}

.error {
  color: red;
}
</style>