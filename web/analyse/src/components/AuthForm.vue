<template>
  <form @submit.prevent="submit">
    <div class="form-group">
      <label>用户名</label>
      <input
          type="text"
          v-model="form.userName"
          required
          placeholder="请输入用户名"
      >
    </div>

    <div class="form-group">
      <label>密码</label>
      <input
          type="password"
          v-model="form.password"
          required
          placeholder="请输入密码"
      >
    </div>

    <div v-if="showConfirm" class="form-group">
      <label>确认密码</label>
      <input
          type="password"
          v-model="form.password2"
          required
          placeholder="请再次输入密码"
      >
    </div>

    <button type="submit">{{ buttonText }}</button>
  </form>
</template>

<script>
export default {
  props: {
    buttonText: {
      type: String,
      default: '提交'
    },
    showConfirm: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      form: {
        userName: '',
        password: '',
        password2: ''
      }
    }
  },
  methods: {
    submit() {
      this.$emit('submit', this.form)
    }
  }
}
</script>

<style scoped>
.form-group {
  margin-bottom: 15px;
}

input {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
}

button {
  padding: 8px 15px;
  background: #42b983;
  color: white;
  border: none;
  border-radius: 4px;
}
</style>