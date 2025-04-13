<template>
  <div class="not-found-page">
    <img src="@/assets/Datum_404.png" alt="404 Not Found" class="not-found-image">
    <h1>页面未找到</h1>
    <p>抱歉，您访问的页面不存在或已被移除</p>
    <router-link to="/" class="home-link">返回首页</router-link>
  </div>
</template>

<script>
export default {
  name: 'NotFoundPage'
}
</script>

<style scoped>
.not-found-page {
  text-align: center;
  padding: 50px;
}

.not-found-image {
  max-width: 500px;
  margin-bottom: 30px;
}

.home-link {
  display: inline-block;
  margin-top: 20px;
  padding: 10px 20px;
  background-color: #42b983;
  color: white;
  text-decoration: none;
  border-radius: 4px;
}

.home-link:hover {
  background-color: #3aa876;
}
</style>