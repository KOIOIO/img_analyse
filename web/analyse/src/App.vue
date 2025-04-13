<template>
  <div id="app">
    <!-- 导航栏 (登录后显示) -->
    <nav v-if="$store.getters['auth/isAuthenticated']" class="app-nav">
      <div class="nav-links">
        <router-link to="/not-found" class="nav-link">404测试页</router-link>
      </div>
      <button @click="logout" class="logout-btn">退出登录</button>
    </nav>

    <!-- 页面内容 -->
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  methods: {
    logout() {
      this.$store.dispatch('auth/logout')
      this.$router.push('/login')
    }
  }
}
</script>

<style>
/* 基础样式 */
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

/* 导航栏样式 */
.app-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 20px;
  background-color: #f8f9fa;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-bottom: 30px;
}

.nav-links {
  display: flex;
  gap: 20px;
}

.nav-link {
  color: #2c3e50;
  text-decoration: none;
  font-weight: 500;
  padding: 8px 12px;
  border-radius: 4px;
  transition: all 0.3s ease;
}

.nav-link:hover {
  background-color: #e9ecef;
}

.nav-link.router-link-exact-active {
  color: #42b983;
  background-color: #e6f7ee;
}

/* 退出按钮样式 */
.logout-btn {
  padding: 8px 16px;
  background-color: #f8d7da;
  color: #721c24;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
}

.logout-btn:hover {
  background-color: #f1b0b7;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .app-nav {
    flex-direction: column;
    gap: 15px;
    padding: 15px;
  }

  .nav-links {
    width: 100%;
    justify-content: center;
  }
}
</style>