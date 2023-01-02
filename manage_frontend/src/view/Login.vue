<template>

  <body id="poster">
    <el-form
      class="login-container"
      lable-position="left"
      label-width="0px"
    >
      <el-row
        type="flex"
        class="row-bg"
        justify="space-between"
      >
        <el-col :span="6">
          <div class="grid-content bg-purple"></div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple-light">
            <h3 class="login_title">
              系统登录
            </h3>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-link
              type="primary"
              @click="toRegister()"
            >点我注册</el-link>
          </div>
        </el-col>
      </el-row>

      <el-form-item label="">
        <el-input
          type="text"
          v-model="loginForm.loginName"
          autocomplete="off"
          placeholder="账号"
        ></el-input>
      </el-form-item>
      <el-form-item label="">
        <el-input
          type="password"
          v-model="loginForm.password"
          autocomplete="off"
          placeholder="密码"
        ></el-input>
      </el-form-item>

      <el-form-item style="width:100%;">
        <el-button
          type="primary"
          style="width:100%;"
          v-on:click="Login()"
        >登录</el-button>
      </el-form-item>
    </el-form>
  </body>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      loginForm: {
        loginName: '',
        password: '',
      }
    }
  },
  methods: {
    Login () {
      console.log('submit!', this.loginForm);
      this.$message({
        message: '登录成功，欢迎来到学生成绩管理系统！',
        type: 'success'
      });
      this.$router.push({ path: '/Home' })
    },
    toRegister () {
      this.$router.push({ path: '/Register' })
    }
  }
}
</script>

<style>
#poster {
  background-position: center;
  height: 100%;
  width: 100%;
  background-size: cover;
  position: fixed;
}
body {
  margin: 0px;
  padding: 0px;
}
.login-container {
  border-radius: 15px;
  background-clip: padding-box;
  margin: 90px auto;
  width: 350px;
  padding: 35px 35px 15px 35px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
}
.login_title {
  margin: 0px auto 40px auto;
  text-align: center;
  color: #505458;
}
</style>