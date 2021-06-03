<template>
<div class="w-100 h-100 container my-container" style="">
    <div class="row justify-content-center" style="min-height: 100vh;">
    <div 
    class="wrap-login d-flex flex-column h-100 justify-content-center"
    style="min-height: 100vh">
        <div class="wrap-login-form-title"><b>ACCOUNT LOGIN</b></div>
        <div class="wrap-login-form-modal border">
            <form class="w-100 h-100" @submit.prevent="login">
                <div class="wrap-input">
            <div class="input-group">
                <div class="input-group-prepend ml-3 mr-3" style="align-items: center">
                    <b-icon class="h4 mb-0" icon="person" style="color: gray"></b-icon>
                </div>

                <input
                type="email"
                class="form-control input" 
                placeholder="Email address..."
                aria-label="Username" 
                v-model="email"
                />
            </div>
                </div>
                <div class="wrap-input">
            <div class="input-group">
                <div class="input-group-prepend ml-3 mr-3" style="align-items: center">
                    <b-icon class="h4 mb-0" icon="lock" style="color: gray"></b-icon>
                </div>
                <input
                type="password"
                class="form-control input" 
                placeholder="password"
                aria-label="password" 
                v-model="password"
                />
            </div>
                </div>
            <div class="d-flex w-100 justify-content-center align-items-center" style="min-height: 130px">
                <button class="btn" type="submit">Login</button>    
            </div>
                
            </form>

        </div>
    </div>
    </div>
</div>
</template>

<script>
import firebase from 'firebase';
export default {
    name: 'Login',
    data() {
        return {
            email: '',
            password: '',
        };
    },
    methods: {
        login() {
            firebase
                .auth()
                .signInWithEmailAndPassword(this.email, this.password)
                .then(() => {
                    this.$router.push('/homepage');
                })
                .catch(error => {
                    alert(error.message);
                });
        },
    },
};
</script>

<style scoped>
.my-container {
    background-image: url('~@/assets/greenery-login-page.jpg');
    background-repeat: no-repeat;
    background-size: cover;

    margin-left: 0px;
    margin-right: 0px;
    width: 100%;
    min-width: 100%;
    }
.wrap-login {
  width: 390px;
  height: 390px;
  
}
.wrap-login-form-title {
  font-size: 28px;
  color: white;
  line-height: 1.2;
  text-align: center;
  text-transform: uppercase;

  display: block;
}
.wrap-login-form-modal {
  height: 350px;
  width: 390px;
  border-radius: 10px;
  overflow: hidden;
  margin-top: 30px;
  margin-bottom: 30px;
  background: white;
}
.wrap-input {
  width: 100%;
  position: relative;
  border-bottom: 1px solid #e6e6e6;
  padding: 29px 0;
}
.input {
  font-size: 20px;
  color: #555555;
  line-height: 1.2;
  border: none;
  display: block;
  width: 250px;
  height: 50px;
  background: transparent;
  margin-right: 50px;
}
.btn {
  font-size: 18px;
  background-color: #00BFA1;
  color: black;
  line-height: 1.2;
  text-transform: uppercase;

  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 20px;
  min-width: 160px;
  height: 42px;
  border-radius: 21px;
  position: relative;
  z-index: 1;

  -webkit-transition: all 0.4s;
  -o-transition: all 0.4s;
  -moz-transition: all 0.4s;
  transition: all 0.4s;
}
.btn:hover {
    opacity: 80%;
}
.form-control:focus {
    border-color: transparent;
    box-shadow:none;
} 

</style>