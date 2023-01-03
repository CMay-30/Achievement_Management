<template>

  <div id="poster">
    <el-form
      :model="ruleForm"
      :rules="rules"
      lable-position="left"
      ref="ruleForm"
      label-width="0px"
      class="register-container"
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
            <h3 class="register_title">
              系统注册
            </h3>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="grid-content bg-purple">
            <el-link
              type="primary"
              @click="toLogin()"
            >去登录</el-link>
          </div>
        </el-col>
      </el-row>

      <el-form-item
        label=""
        prop="LoginName"
      >
        <el-input
          type="text"
          v-model="ruleForm.LoginName"
          autocomplete="off"
          placeholder="请输入用户名"
          prefix-icon="el-icon-user"
        ></el-input>
      </el-form-item>
      <el-form-item
        label=""
        prop="name"
      >
        <el-input
          type="text"
          v-model.number="ruleForm.name"
          autocomplete="off"
          placeholder="请输入姓名"
          prefix-icon="el-icon-user"
        ></el-input>
      </el-form-item>
      <el-form-item
        label=""
        prop="password"
      >
        <el-input
          type="password"
          v-model="ruleForm.password"
          autocomplete="off"
          placeholder="请输入密码"
          prefix-icon="el-icon-lock"
        ></el-input>
      </el-form-item>
      <el-form-item
        label=""
        prop="checkPass"
      >
        <el-input
          type="password"
          v-model="ruleForm.checkPass"
          autocomplete="off"
          placeholder="请确认密码"
          prefix-icon="el-icon-lock"
        ></el-input>
      </el-form-item>

      <el-form-item>
        <el-button
          type="primary"
          @click="submitForm(ruleForm)"
        >注册</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'Register',
  data () {
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass');
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.ruleForm.password) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        LoginName: '',
        name: '',
        password: '',
        checkPass: '',
      },
      rules: {
        LoginName: [
          { required: true, message: "请输入你的用户名", trigger: 'blur' },
          { min: 2, max: 9, message: "长度2到9个字符", trigger: 'blur' }
        ],
        password: [
          { validator: validatePass, trigger: 'blur' }
        ],
        checkPass: [
          { validator: validatePass2, trigger: 'blur' }
        ],
      }
    };
  },
  methods: {
    submitForm (ruleForm) {
      this.rulesForm = {};
      this.$message({
        message: '注册成功！',
        type: 'success',
      });
    },
    resetForm (formName) {
      this.$refs[formName].resetFields();
    },
    toLogin () {
      this.$router.push({ path: '/' })
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
  margin: 0px;
  padding: 0px;
}
.register-container {
  border-radius: 15px;
  background-clip: padding-box;
  margin: 90px auto;
  width: 350px;
  padding: 35px 35px 15px 35px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
}
.register_title {
  margin: 0px auto 40px auto;
  text-align: center;
  color: #505458;
}
</style>