<template>
  <div class="page flex-center">
    <div class="sign-box">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <h3 class="title">系统注册</h3>
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
        <el-form-item label="确认密码" prop="cfpassword">
          <el-input
            v-model="form.cfpassword"
            type="password"
            placeholder="请确认密码"
          ></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-button type="primary" class="w100p" @click="register"
            >注册</el-button
          >
        </el-form-item>
        <div class="txt-r">
          <router-link to="/login">已有账号？去登录</router-link>
        </div>
      </el-form>
    </div>
  </div>
</template>
<script>
export default {
  name: "register",
  data() {
    const validateCfpassword = (rule, value, callback) => {
      if (value !== this.form.cfpassword) {
        callback(new Error("两次密码输入不一致"));
      } else {
        callback();
      }
    };
    return {
      form: {
        account: "",
        password: "",
        cfpassword: "",
      },
      rules: {
        account: [
          { required: true, message: "请输入用户名", trigger: ["change", "blur"] },
          { pattern: /^[a-zA-Z][a-zA-Z0-9_-]{3,31}$/, message: '用户名由英文字母开头的长度3-32位字母、_和-组成', trigger: ["change", "blur"] },
        ],
        password: {
          required: true,
          min: 6,
          message: "请输入至少6个字符的密码",
          trigger: ["change", "blur"],
        },
        cfpassword: [
          {
            required: true,
            message: "请确认密码",
            trigger: ["change", "blur"],
          },
          {
            validator: validateCfpassword,
            trigger: ["change", "blur"],
          },
        ],
      },
    };
  },
  methods: {
    register() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          // 模拟像后台请求数据
          const post = (params) => {
            // 模拟后台比对数据
            return new Promise((resolve, reject) => {
              const { account, password } = params
              let db_user = localStorage.getItem("db_user");
              if (db_user) {
                db_user = JSON.parse(db_user);
              } else {
                db_user = []
              }
              if (!db_user.find((v) => v.account === account)) {
                db_user.push({ account, password });
                localStorage.setItem("db_user", JSON.stringify(db_user));
                resolve ({ msg: '注册成功！'});
              } else {
                reject({msg: `用户名：${account}已存在！`});
              }
            })
          };
          post(this.form).then((res) => {
            alert(`${res.msg}用户名：${this.form.account},密码： ${this.form.password}` );
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