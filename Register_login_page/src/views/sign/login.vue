<template>
  <div class="page flex-center">
    <div class="sign-box">
      <el-form ref="form" :model="form" :rules="rules" label-width="60px">
        <h3 class="title">系统登录</h3>
        <el-form-item label="账号" prop="account">
          <el-input
            v-model="form.account"
            placeholder="请输入用户名"
          ></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            v-model="form.password"
            type="password"
            placeholder="请输入密码"
          ></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" class="w100p" @click="login"
            >登录</el-button
          >
        </el-form-item>
        <div class="txt-r">
          <router-link to="/register">没有账号？去注册</router-link>
        </div>
      </el-form>
    </div>
  </div>
</template>
<script>
export default {
  name: "login",
  data() {
    return {
      form: {
        account: "",
        password: "",
      },
      rules: {
        account: [
          {
            required: true,
            message: "请输入用户名",
            trigger: ["change", "blur"],
          },
          {
            pattern: /^[a-zA-Z][a-zA-Z0-9_-]{3,31}$/,
            message: "用户名由英文字母开头的长度4-32位字母、_和-组成",
            trigger: ["change", "blur"],
          },
        ],
        password: {
          required: true,
          min: 6,
          message: "请输入至少6个字符的密码",
          trigger: ["change", "blur"],
        },
      },
    };
  },
  methods: {
    login() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          // 模拟后台数据处理
          const post = ((params) => {
            // 模拟后台比对数据
            return new Promise((resolve, reject) => {
              let db_user = localStorage.getItem("db_user");
              if (db_user) {
                db_user = JSON.parse(db_user);
              }
              const { account, password } = params
              if (db_user.find((v) => v.account === account && v.password === password)) {
                resolve({msg: '登录成功！', token: Date.now()});
              } else {
                reject({msg: '用户名或密码错误！'})
              }
            })
          });
          post(this.form).then((res) => {
            alert(`${res.msg}` );
            localStorage.setItem("token", res.token);
            this.$router.push("/");
          }).catch(err => {
            alert(err.msg);
          })
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
  },
};
</script>

<style lang="css">
@import url(./index.css);
</style>