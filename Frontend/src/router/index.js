import Vue from 'vue';
import VueRouter from 'vue-router';
// import HomePage from '../components/home page/HomePage.vue';
import Login from '../components/login/Login.vue'
import firebase from 'firebase';

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        name: 'Login Page',
        component: Login,
    },
    {
        path: '/homepage',
        name: 'Home Page',
        component: () =>
            import('../components/home page/HomePage.vue'),
        meta: {
            authRequired: true,
        },
    },
];

const router = new VueRouter({
    mode: 'history',
    // process.env.BASE_URL,
    base: process.env.BASE_URL, routes
    // routes: routes,
    // // routes,
});

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.authRequired)) {
        if (firebase.auth().currentUser) {
            next();
        } else {
            alert('You must be logged in to see this page');
            next({
                path: '/',
            });
        }
    } else {
        next();
    }
});
export default router;